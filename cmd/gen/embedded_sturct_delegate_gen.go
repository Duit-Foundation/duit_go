// Command wrapgen generates wrapper methods for embedded struct fields tagged with `gen:"wrap"`.
// It uses go/types (through golang.org/x/tools/go/packages) and proxies any method
// signature (params, variadics, and multiple return values). It skips methods
// that have a leading comment line containing [excludeTag].
//
// How it decides which package to process:
//   - By default it processes the package in the current directory (flag -dir).
//   - The generator automatically processes any package that contains structs
//     with embedded fields tagged with `gen:"wrap"`.
//
// How it decides what to wrap:
//   - It scans *struct types* in the package.
//   - For each embedded field whose struct tag contains `gen:"wrap"`, it finds
//     methods of the embedded type (pointer or value receiver) defined in the
//     *same package* and generates wrappers on the outer struct that:
//   - Lazily initialize the embedded field if it's a pointer and nil.
//   - Forward all parameters and return a pointer to the outer struct (fluent interface).
//   - If the original method's leading comment includes [excludeTag],
//     that method is skipped.
//   - If the outer struct already defines a method with the same name,
//     the wrapper is not generated (no clobbering).
//
// Example usage in your code:
//
//	type ValueReferenceHolder struct { /* ... */ }
//
//	// Some is eligible for wrapping (unless you add [excludeTag] above).
//	func (r *ValueReferenceHolder) Some() { /* ... */ }
//
//	type IgnorePointerAttributes struct {
//	    *ValueReferenceHolder `gen:"wrap"` // <- tag tells wrapgen to proxy methods
//	    *ThemeConsumer                      // <- not tagged, ignored
//	    Ignoring duit_utils.Tristate[bool]  `json:"ignoring,omitempty"`
//	}
//
//	// Generated wrapper will return *IgnorePointerAttributes for fluent interface:
//	attrs := NewIgnorePointerAttributes().
//	    Some().  // returns *IgnorePointerAttributes
//	    Some2()  // returns *IgnorePointerAttributes
package gen

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	excludeTag  = "//bindgen:exclude"
	outFileName = "zz_embedded_structs_delegates.go"
)

var (
	dirFlag = flag.String("dir", ".", "directory of the target package (default: current directory)")
	outFlag = flag.String("out", outFileName, "output filename inside the package directory")
)

func EmbeddedStructDelegateGenerator() {
	flag.Parse()

	pkgPath, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatalf("resolve dir: %v", err)
	}

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
		Dir: pkgPath,
	}
	pkgs, err := packages.Load(cfg, "./...") // allow local replace; we'll filter by Dir
	if err != nil {
		log.Fatalf("load packages: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Find the package exactly at dirFlag.
	var target *packages.Package
	for _, p := range pkgs {
		if sameDir(filepath.Dir(firstNonEmpty(p.CompiledGoFiles)), pkgPath) {
			target = p
			break
		}
	}
	if target == nil {
		log.Fatalf("no package found in %s", pkgPath)
	}

	gen, err := buildGenerationPlan(target)
	if err != nil {
		log.Fatalf("plan: %v", err)
	}
	if len(gen.Wrappers) == 0 {
		log.Printf("wrapgen: nothing to generate for package %s", target.PkgPath)
		return
	}

	code, err := render(target, gen)
	if err != nil {
		log.Fatalf("render: %v", err)
	}

	outPath := filepath.Join(pkgPath, *outFlag)
	if err := os.WriteFile(outPath, code, 0o644); err != nil {
		log.Fatalf("write output: %v", err)
	}
	log.Printf("wrapgen: wrote %s", outPath)
}

func sameDir(a, b string) bool {
	ap, _ := filepath.Abs(a)
	bp, _ := filepath.Abs(b)
	return ap == bp
}

func firstNonEmpty(ss []string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

// generationPlan collects all wrapper methods to emit.

type generationPlan struct {
	PkgName  string
	Wrappers []wrapper
}

type wrapper struct {
	OuterName string       // e.g., IgnorePointerAttributes
	FieldName string       // e.g., ValueReferenceHolder
	FieldType types.Type   // full field type (may be *T or T)
	ElemNamed *types.Named // named type for method set (underlying element if pointer)
	Method    *types.Func  // method on the embedded type to proxy
	Sig       *types.Signature
}

func buildGenerationPlan(p *packages.Package) (*generationPlan, error) {
	info := p.TypesInfo
	plan := &generationPlan{PkgName: p.Name}

	// Map: method object -> its leading comments (to check excludeTag)
	methodComments := map[*types.Func]string{}
	// Also build a set of methods declared on each outer type to avoid duplicates.
	outerHasMethod := map[string]map[string]struct{}{}

	// Walk files to collect AST-based metadata.
	for _, file := range p.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			fd, ok := n.(*ast.FuncDecl)
			if !ok || fd.Recv == nil || fd.Name == nil {
				return true
			}
			obj := info.Defs[fd.Name]
			fn, _ := obj.(*types.Func)
			if fn == nil {
				return true
			}
			// Save leading comments for methods.
			if fd.Doc != nil {
				var b strings.Builder
				for _, c := range fd.Doc.List {
					b.WriteString(strings.TrimSpace(c.Text))
					b.WriteRune('\n')
				}
				methodComments[fn] = b.String()
			}

			// Track methods existing on outer types.
			recv := fn.Type().(*types.Signature).Recv().Type()
			outer := derefNamed(recv)
			if outer != nil {
				if _, ok := outerHasMethod[outer.Obj().Name()]; !ok {
					outerHasMethod[outer.Obj().Name()] = map[string]struct{}{}
				}
				outerHasMethod[outer.Obj().Name()][fn.Name()] = struct{}{}
			}
			return true
		})
	}

	// Now iterate named struct types to find embedded fields with tag gen:"wrap".
	for _, name := range p.Types.Scope().Names() {
		obj := p.Types.Scope().Lookup(name)
		named, _ := obj.(*types.TypeName)
		if named == nil {
			continue
		}
		st, ok := named.Type().Underlying().(*types.Struct)
		if !ok {
			continue
		}

		for i := 0; i < st.NumFields(); i++ {
			fld := st.Field(i)
			if !fld.Embedded() {
				continue
			}
			tag := st.Tag(i)
			if !hasWrapTag(tag) {
				continue
			}
			fldTyp := fld.Type()
			elNamed := derefNamed(fldTyp)
			if elNamed == nil {
				continue // we only handle named types
			}

			// Limit to methods defined in the same package as elNamed, so we can read comments.
			if elNamed.Obj().Pkg() == nil || elNamed.Obj().Pkg().Path() != p.PkgPath {
				continue
			}

			// Compute method set of the element type's *pointer* to capture both value and pointer methods.
			ptrTo := types.NewPointer(elNamed)
			mset := types.NewMethodSet(ptrTo)

			for j := 0; j < mset.Len(); j++ {
				m := mset.At(j).Obj().(*types.Func)

				// Skip if method is marked excludeTag in its decl comments.
				if comments, ok := methodComments[m]; ok {
					if containsGenExclude(comments) {
						continue
					}
				}

				// Skip if outer already has a method with this name.
				if has := outerHasMethod[named.Name()]; has != nil {
					if _, exists := has[m.Name()]; exists {
						continue
					}
				}

				// Only generate for methods whose receiver base is the element type (not promoted ones from deeper embedding).
				recv := m.Type().(*types.Signature).Recv()
				if recv == nil {
					continue
				}
				base := derefNamed(recv.Type())
				if base == nil || base.Obj() != elNamed.Obj() {
					continue
				}

				plan.Wrappers = append(plan.Wrappers, wrapper{
					OuterName: named.Name(),
					FieldName: fld.Name(),
					FieldType: fldTyp,
					ElemNamed: elNamed,
					Method:    m,
					Sig:       m.Type().(*types.Signature),
				})
			}
		}
	}

	// Make output deterministic.
	sort.Slice(plan.Wrappers, func(i, j int) bool {
		if plan.Wrappers[i].OuterName != plan.Wrappers[j].OuterName {
			return plan.Wrappers[i].OuterName < plan.Wrappers[j].OuterName
		}
		if plan.Wrappers[i].FieldName != plan.Wrappers[j].FieldName {
			return plan.Wrappers[i].FieldName < plan.Wrappers[j].FieldName
		}
		return plan.Wrappers[i].Method.Name() < plan.Wrappers[j].Method.Name()
	})

	return plan, nil
}

func hasWrapTag(tag string) bool {
	if tag == "" {
		return false
	}
	st := reflect.StructTag(tag)
	return st.Get("gen") == "wrap"
}

func derefNamed(t types.Type) *types.Named {
	switch tt := t.(type) {
	case *types.Pointer:
		if n, ok := tt.Elem().(*types.Named); ok {
			return n
		}
	case *types.Named:
		return tt
	}
	return nil
}

func containsGenExclude(comment string) bool {
	for _, line := range strings.Split(comment, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), excludeTag) {
			return true
		}
	}
	return false
}

func collectImports(pkg *packages.Package, plan *generationPlan) []string {
	importsSet := make(map[string]struct{})

	// Create a qualifier that collects external package paths
	collectingQualifier := func(p *types.Package) string {
		if p.Path() != pkg.PkgPath {
			importsSet[p.Path()] = struct{}{}
		}
		return p.Name()
	}

	for _, w := range plan.Wrappers {
		sig := w.Sig

		// Check parameter types
		params := sig.Params()
		for i := 0; i < params.Len(); i++ {
			types.TypeString(params.At(i).Type(), collectingQualifier)
		}

		// Check result types
		results := sig.Results()
		for i := 0; i < results.Len(); i++ {
			types.TypeString(results.At(i).Type(), collectingQualifier)
		}
	}

	// Convert set to sorted slice
	var imports []string
	for imp := range importsSet {
		imports = append(imports, imp)
	}
	sort.Strings(imports)

	return imports
}

func render(pkg *packages.Package, plan *generationPlan) ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by duit_go bindgen; DO NOT EDIT.\n")
	fmt.Fprintf(&b, "package %s\n\n", plan.PkgName)

	// Collect all external packages used in method signatures
	imports := collectImports(pkg, plan)
	if len(imports) > 0 {
		fmt.Fprintf(&b, "import (\n")
		for _, imp := range imports {
			fmt.Fprintf(&b, "\t\"%s\"\n", imp)
		}
		fmt.Fprintf(&b, ")\n\n")
	}

	// Group wrappers by outer type for nicer layout.
	byOuter := map[string][]wrapper{}
	for _, w := range plan.Wrappers {
		byOuter[w.OuterName] = append(byOuter[w.OuterName], w)
	}
	var outers []string
	for k := range byOuter {
		outers = append(outers, k)
	}
	sort.Strings(outers)

	for _, outer := range outers {
		fmt.Fprintf(&b, "// --- Wrappers for %s ---\n", outer)
		fmt.Fprintf(&b, "\n")
		for _, w := range byOuter[outer] {
			fmt.Fprintf(&b, "%s\n\n", renderWrapper(w, pkg))
		}
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		// Write unformatted to help debugging.
		return b.Bytes(), fmt.Errorf("gofmt failed: %w", err)
	}
	return formatted, nil
}

func renderWrapper(w wrapper, pkg *packages.Package) string {
	// Create qualifier that omits package name for current package
	qualifier := func(p *types.Package) string {
		if p.Path() == pkg.PkgPath {
			return ""
		}
		return p.Name()
	}

	// Build parameter and result lists.
	paramsSig := w.Sig.Params()

	var (
		paramDecls []string // with names and types
		argNames   []string // just names
	)
	for i := 0; i < paramsSig.Len(); i++ {
		v := paramsSig.At(i)
		name := v.Name()
		if name == "" {
			name = fmt.Sprintf("p%d", i)
		}
		// If last param is variadic, show ...T in decl.
		t := v.Type()
		if w.Sig.Variadic() && i == paramsSig.Len()-1 {
			if sl, ok := t.(*types.Slice); ok {
				paramDecls = append(paramDecls, fmt.Sprintf("%s ...%s", name, types.TypeString(sl.Elem(), qualifier)))
				argNames = append(argNames, name+"...")
				continue
			}
		}
		paramDecls = append(paramDecls, fmt.Sprintf("%s %s", name, types.TypeString(t, qualifier)))
		argNames = append(argNames, name)
	}

	// For fluent interface, always return *OuterName
	resultDecl := fmt.Sprintf("*%s", w.OuterName)

	decl := fmt.Sprintf("func (r *%s) %s(%s) %s {", w.OuterName, w.Method.Name(), strings.Join(paramDecls, ", "), resultDecl)

	var body []string
	// Initialization if field is a pointer type.
	if _, ok := w.FieldType.(*types.Pointer); ok {
		body = append(body, fmt.Sprintf("\tif r.%s == nil {", w.FieldName))
		body = append(body, fmt.Sprintf("\t\tr.%s = &%s{}", w.FieldName, w.ElemNamed.Obj().Name()))
		body = append(body, "\t}")
	}

	// Call the embedded field's method
	call := fmt.Sprintf("r.%s.%s(%s)", w.FieldName, w.Method.Name(), strings.Join(argNames, ", "))

	// For fluent interface, always return the receiver
	body = append(body, "\t"+call)
	body = append(body, "\treturn r")

	return decl + "\n" + strings.Join(body, "\n") + "\n}"
}
