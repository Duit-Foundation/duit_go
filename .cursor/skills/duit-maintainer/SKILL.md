---
name: duit-maintainer
description: Implements new Flutter widgets following Duit project patterns, conducts code reviews, writes tests, and enforces coding conventions. Use when implementing new widget functionality, reviewing Flutter widget code, or maintaining the Duit library.
---

# Duit Maintainer

## Quick Start

When implementing a new Flutter widget for the Duit library, follow this workflow:

1. Create attributes structure in `pkg/duit_attributes/`
2. Add widget type to `pkg/duit_core/element_type.go`
3. Add widget function in `pkg/duit_widget/`
4. Run code generation: `make gogen`
5. Write tests if required

## Widget Implementation Workflow

### Step 1: Create Attributes Structure

Create a new file in `pkg/duit_attributes/` named `{widget_name}_attributes.go`:

```go
package duit_attributes

import (
    "errors"
    "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// WidgetNameAttributes defines attributes for WidgetName widget.
// See: https://api.flutter.dev/flutter/widgets/WidgetName-class.html
type WidgetNameAttributes struct {
    *ValueReferenceHolder `gen:"wrap"`
    *ThemeConsumer        `gen:"wrap"`
    Property1             duit_utils.Tristate[bool] `json:"property1,omitempty"`
    Property2             string                     `json:"property2,omitempty"`
}

// NewWidgetNameAttributes creates a new WidgetNameAttributes instance.
func NewWidgetNameAttributes() *WidgetNameAttributes {
    return &WidgetNameAttributes{}
}

// SetProperty1 sets the property1 and returns WidgetNameAttributes for method chaining.
func (r *WidgetNameAttributes) SetProperty1(property1 bool) *WidgetNameAttributes {
    r.Property1 = duit_utils.BoolValue(property1)
    return r
}

// SetProperty2 sets the property2 and returns WidgetNameAttributes for method chaining.
func (r *WidgetNameAttributes) SetProperty2(property2 string) *WidgetNameAttributes {
    r.Property2 = property2
    return r
}

//bindgen:exclude
func (r *WidgetNameAttributes) Validate() error {
    if err := r.ValueReferenceHolder.Validate(); err != nil {
        return err
    }

    if err := r.ThemeConsumer.Validate(); err != nil {
        return err
    }

    if r.Property1 == nil {
        return errors.New("property1 is required")
    }

    return nil
}
```

**Rules:**

- Use `*ValueReferenceHolder` with `gen:"wrap"` tag
- Use `*ThemeConsumer` with `gen:"wrap"` tag
- Use `duit_utils.Tristate[T]` for optional boolean properties
- Add `json:"propertyName,omitempty"` tag for all fields
- Method receiver must be named `r`
- Add `//bindgen:exclude` comment to `Validate()` method
- Documentation must be in English

### Step 2: Add Widget Type

Add a constant to `pkg/duit_core/element_type.go`:

```go
const (
    // ... existing types
    WidgetName DuitElementType = "WidgetName"
)
```

Follow naming convention: Use PascalCase with no spaces or underscores.

### Step 3: Add Widget Function

Create a new file in `pkg/duit_widget/` named `{widget_name}.go`:

```go
package duit_widget

import (
    "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
    "github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
    WidgetName(
        duit_attributes.NewWidgetNameAttributes().
            SetProperty1(true).
            SetProperty2("value"),
        "id",
        false,
        nil,
    )
*/
func WidgetName(attributes *duit_attributes.WidgetNameAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
    checkAttributes(attributes)
    return new(duit_core.DuitElementModel).CreateElement(duit_core.WidgetName, id, "", attributes, nil, controlled, 1, nil, child)
}
```

**Function signature patterns:**

- **Single child**: `func WidgetName(attributes *Attrs, id string, controlled bool, child *DuitElementModel) *DuitElementModel`
- **Multiple children**: `func WidgetName(attributes *Attrs, id string, controlled bool, children []*DuitElementModel) *DuitElementModel`
- **Optional child**: `func WidgetName(attributes *Attrs, id string, child *DuitElementModel) *DuitElementModel`

**CreateElement parameters:**

1. Widget type constant
2. ID string
3. Tag string (empty for most widgets)
4. Attributes
5. Commands (nil for most widgets)
6. Controlled boolean
7. Children count (use 2 for widgets with optional child like Visibility)
8. Nil (reserved)
9. Children (variadic)

**Rules:**

- Always call `checkAttributes(attributes)` to validate
- Use `new(DuitElementModel).CreateElement(...)`
- Add example usage in comment

### Step 4: Run Code Generation

```bash
make gogen
```

This generates embedded struct delegates and other boilerplate code.

### Step 5: Write Tests

Only write tests when explicitly required. When required, create `{widget_name}_attributes_test.go` in `pkg/duit_attributes/`:

```go
package duit_attributes_test

import (
    "testing"

    "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestWidgetNameAttributes_Validate_ValidAttributes(t *testing.T) {
    attrs := &duit_attributes.WidgetNameAttributes{
        Property1: duit_utils.BoolValue(true),
    }

    err := attrs.Validate()
    if err != nil {
        t.Fatal("expected no error for valid attributes, got:", err)
    }
}

func TestWidgetNameAttributes_Validate_MissingRequiredProperty(t *testing.T) {
    attrs := &duit_attributes.WidgetNameAttributes{
        Property1: nil, // missing required property
    }

    err := attrs.Validate()
    if err == nil {
        t.Fatal("expected error for missing required property")
    }
}

func TestWidgetNameAttributes_SetProperty1(t *testing.T) {
    attrs := duit_attributes.NewWidgetNameAttributes()
    property1 := true

    result := attrs.SetProperty1(property1)

    if *attrs.Property1 != property1 {
        t.Fatalf("expected Property1 to be %v, got: %v", property1, attrs.Property1)
    }

    if result != attrs {
        t.Fatal("expected SetProperty1 to return same instance for method chaining")
    }
}
```

**Testing rules:**

- Test filename: `{target_file}_test.go`
- Only test properties explicitly used in structure's methods
- Always test `Tristate[T]` properties for JSON serialization
- Never check error text, only check if error exists
- Never cover embedded fields or embedded structs
- Test setter methods return same instance for chaining

## Code Review Checklist

When reviewing widget implementations, check:

- [ ] Receiver name is `r` for all structure methods
- [ ] Documentation is in English
- [ ] Method chaining returns same instance (`return r`)
- [ ] Required properties validated in `Validate()` method
- [ ] `Validate()` has `//bindgen:exclude` comment
- [ ] `checkAttributes()` called in widget function
- [ ] Widget type follows naming convention (PascalCase)
- [ ] Attributes use appropriate types (`Tristate[T]` for optionals)
- [ ] JSON tags use `omitempty` suffix
- [ ] Example usage comment present in widget file
- [ ] Nil checks for potentially nil receivers before method calls

## Coding Conventions

### Style Rules

From `pkg/.cursor/rules/style.mdc`:

1. Method receiver name must always be "r"
2. Documentation must be in English
3. Explicit nil checks for potentially nil receivers

### Testing Rules

From `pkg/.cursor/rules/testing.mdc`:

1. Never cover embedded fields or embedded structs with tests
2. Always add tests for `Tristate[T]` properties (JSON serialization)
3. Test filename: `{target_file}_test.go`
4. Only test properties explicitly used in structure's methods
5. Never check error text, only check if error exists

### Flutter Documentation Reference

Always reference official Flutter documentation in attributes structure:

```go
// See: https://api.flutter.dev/flutter/widgets/WidgetName-class.html
```

## Common Patterns

### Required Properties

```go
if r.RequiredProperty == nil {
    return errors.New("requiredProperty is required")
}
```

### Optional Properties (Tristate)

```go
type WidgetAttributes struct {
    OptionalProperty duit_utils.Tristate[bool] `json:"optionalProperty,omitempty"`
}

// SetOptionalProperty sets the optional property.
func (r *WidgetAttributes) SetOptionalProperty(optional bool) *WidgetAttributes {
    r.OptionalProperty = duit_utils.BoolValue(optional)
    return r
}
```

### String Properties

```go
type WidgetAttributes struct {
    TextProperty string `json:"textProperty,omitempty"`
}

// SetTextProperty sets the text property.
func (r *WidgetAttributes) SetTextProperty(text string) *WidgetAttributes {
    r.TextProperty = text
    return r
}
```

### Nil-Safe Method Calls

```go
func SomeFunction(widget *DuitElementModel) {
    if widget == nil {
        return
    }
    widget.SomeMethod()
}
```

## Makefile Commands

```bash
# Generate code
make gogen

# Run tests
make test

# Run vet
make vet
```

Always run `make gogen` after adding new widgets.

## Troubleshooting

### Code generation fails

- Check that `gen:"wrap"` tags are present for embedded structs
- Verify attribute struct has no syntax errors
- Ensure all referenced types are imported

### Tests fail

- Check that required properties are set in test
- Verify Tristate properties are tested for JSON serialization
- Ensure error checks don't compare error text

### Widget not recognized

- Verify type is added to `element_type.go`
- Check type name matches exactly in widget function
- Run `make gogen` to regenerate delegates
