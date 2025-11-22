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
	"go/parser"
	"go/token"
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

	outPath := filepath.Join(pkgPath, *outFlag)
	
	// Parse existing file to find already generated methods
	existingMethods := parseExistingFile(outPath)

	gen, err := buildGenerationPlan(target, existingMethods)
	if err != nil {
		log.Fatalf("plan: %v", err)
	}
	if len(gen.Wrappers) == 0 {
		log.Printf("wrapgen: nothing to generate for package %s", target.PkgPath)
		return
	}

	code, err := render(target, gen, outPath)
	if err != nil {
		log.Fatalf("render: %v", err)
	}

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

// methodKey uniquely identifies a generated wrapper method
type methodKey struct {
	OuterName string
	MethodName string
}

// parseExistingFile parses the existing generated file and returns a set of already generated methods
func parseExistingFile(filePath string) map[methodKey]struct{} {
	existing := make(map[methodKey]struct{})
	
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return existing
	}

	// Read and parse the file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		// If parsing fails, assume file is corrupted or not a valid Go file
		// Return empty set to regenerate everything
		log.Printf("warning: failed to parse existing file %s: %v, will regenerate", filePath, err)
		return existing
	}

	// Walk AST to find all method declarations
	ast.Inspect(file, func(n ast.Node) bool {
		fd, ok := n.(*ast.FuncDecl)
		if !ok || fd.Recv == nil || fd.Name == nil {
			return true
		}

		// Check if this is a method with pointer receiver
		if len(fd.Recv.List) == 0 {
			return true
		}
		
		recvType := fd.Recv.List[0].Type
		star, ok := recvType.(*ast.StarExpr)
		if !ok {
			return true
		}
		
		ident, ok := star.X.(*ast.Ident)
		if !ok {
			return true
		}

		// This looks like a generated wrapper method: func (r *OuterName) MethodName(...) *OuterName
		// Check return type to confirm it's a fluent interface method
		if fd.Type.Results != nil && len(fd.Type.Results.List) == 1 {
			retType := fd.Type.Results.List[0].Type
			retStar, ok := retType.(*ast.StarExpr)
			if ok {
				retIdent, ok := retStar.X.(*ast.Ident)
				if ok && retIdent.Name == ident.Name {
					// This matches the pattern of generated wrapper methods
					key := methodKey{
						OuterName:  ident.Name,
						MethodName: fd.Name.Name,
					}
					existing[key] = struct{}{}
				}
			}
		}

		return true
	})

	return existing
}

func buildGenerationPlan(p *packages.Package, existingMethods map[methodKey]struct{}) (*generationPlan, error) {
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

				// Skip if this method already exists in the generated file
				key := methodKey{
					OuterName:  named.Name(),
					MethodName: m.Name(),
				}
				if _, exists := existingMethods[key]; exists {
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

func render(pkg *packages.Package, plan *generationPlan, outPath string) ([]byte, error) {
	// Check if file exists and read it
	var existingContent []byte
	var existingImports []string
	fileExists := false
	
	if _, err := os.Stat(outPath); err == nil {
		fileExists = true
		var err error
		existingContent, err = os.ReadFile(outPath)
		if err != nil {
			return nil, fmt.Errorf("read existing file: %w", err)
		}
		
		// Parse existing file to extract imports
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, outPath, existingContent, parser.ParseComments)
		if err == nil {
			for _, imp := range file.Imports {
				if imp.Path != nil {
					path := strings.Trim(imp.Path.Value, "\"")
					existingImports = append(existingImports, path)
				}
			}
		}
	}

	// If file doesn't exist or is empty, generate from scratch
	if !fileExists || len(existingContent) == 0 {
		return renderNew(pkg, plan)
	}

	// File exists, append new methods
	return renderAppend(pkg, plan, existingContent, existingImports)
}

func renderNew(pkg *packages.Package, plan *generationPlan) ([]byte, error) {
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

func renderAppend(pkg *packages.Package, plan *generationPlan, existingContent []byte, existingImports []string) ([]byte, error) {
	// Parse existing file to check if it's valid
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", existingContent, parser.ParseComments)
	if err != nil {
		// If we can't parse, regenerate from scratch
		log.Printf("warning: failed to parse existing file, regenerating from scratch: %v", err)
		return renderNew(pkg, plan)
	}

	// Group new wrappers by outer type
	byOuter := map[string][]wrapper{}
	for _, w := range plan.Wrappers {
		byOuter[w.OuterName] = append(byOuter[w.OuterName], w)
	}
	var outers []string
	for k := range byOuter {
		outers = append(outers, k)
	}
	sort.Strings(outers)

	// Build new methods content
	var newMethods bytes.Buffer
	for _, outer := range outers {
		// Check if section for this outer type already exists in the file
		sectionExists := strings.Contains(string(existingContent), fmt.Sprintf("// --- Wrappers for %s ---", outer))
		
		if !sectionExists {
			// Add new section header
			fmt.Fprintf(&newMethods, "// --- Wrappers for %s ---\n", outer)
			fmt.Fprintf(&newMethods, "\n")
		}

		for _, w := range byOuter[outer] {
			fmt.Fprintf(&newMethods, "%s\n\n", renderWrapper(w, pkg))
		}
	}

	// Check if we need to add new imports
	newImports := collectImports(pkg, plan)
	importsMap := make(map[string]struct{})
	for _, imp := range existingImports {
		importsMap[imp] = struct{}{}
	}
	
	var missingImports []string
	for _, imp := range newImports {
		if _, exists := importsMap[imp]; !exists {
			missingImports = append(missingImports, imp)
		}
	}

	// If no new methods and no new imports, return existing content
	if newMethods.Len() == 0 && len(missingImports) == 0 {
		return existingContent, nil
	}

	// Build the new file content by appending to existing
	var result bytes.Buffer
	
	// If we need to add imports, we need to rebuild the file
	if len(missingImports) > 0 {
		// Parse existing file to extract structure
		existingFile, err := parser.ParseFile(fset, "", existingContent, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("parse existing file: %w", err)
		}

		// Rebuild file with updated imports
		fmt.Fprintf(&result, "// Code generated by duit_go bindgen; DO NOT EDIT.\n")
		fmt.Fprintf(&result, "package %s\n\n", existingFile.Name.Name)

		// Merge imports
		allImports := append(existingImports, missingImports...)
		sort.Strings(allImports)
		if len(allImports) > 0 {
			fmt.Fprintf(&result, "import (\n")
			for _, imp := range allImports {
				fmt.Fprintf(&result, "\t\"%s\"\n", imp)
			}
			fmt.Fprintf(&result, ")\n\n")
		}

		// Extract existing methods (everything after package and imports)
		existingStr := string(existingContent)
		lines := strings.Split(existingStr, "\n")
		
		// Find where methods start (skip header, package, imports)
		methodStart := 0
		inImports := false
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "import (" {
				inImports = true
				continue
			}
			if inImports && trimmed == ")" {
				inImports = false
				continue
			}
			if !inImports && trimmed != "" && 
			   !strings.HasPrefix(trimmed, "// Code generated") &&
			   !strings.HasPrefix(trimmed, "package") &&
			   trimmed != "import (" && trimmed != ")" &&
			   !strings.HasPrefix(trimmed, "\"") {
				methodStart = i
				break
			}
		}

		// Write existing methods
		if methodStart < len(lines) {
			existingMethods := strings.Join(lines[methodStart:], "\n")
			result.WriteString(existingMethods)
			if !strings.HasSuffix(existingMethods, "\n") {
				result.WriteString("\n")
			}
		}
	} else {
		// No new imports, just append new methods to existing content
		result.Write(existingContent)
		// Ensure there's a newline before appending
		if !strings.HasSuffix(string(existingContent), "\n") {
			result.WriteString("\n")
		}
	}

	// Append new methods
	result.WriteString(newMethods.String())

	formatted, err := format.Source(result.Bytes())
	if err != nil {
		// Write unformatted to help debugging.
		return result.Bytes(), fmt.Errorf("gofmt failed: %w", err)
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
