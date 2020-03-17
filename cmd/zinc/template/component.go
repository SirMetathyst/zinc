package template

import (
	"bytes"
	"flag"
	"os"
	"text/template"
)

const (
	// MultiComponentDataTemplateSource ...
	MultiComponentDataTemplateSource = `package {{.PackageName}}
{{ $ImportLength := len .Imports }}{{if eq $ImportLength 0 }}
import (
	"github.com/SirMetathyst/zinc"
){{else if gt $ImportLength 0}}
import (
	zinc "github.com/SirMetathyst/zinc"
	{{range $Index, $Element := .Imports }}"{{$Element}}"
	{{end}}
){{end}}

// Z{{.UpperComponentName}} ...
var Z{{.UpperComponentName}} uint = uint({{.ComponentNameHash}})

// Z{{.UpperComponentName}}Data ...
type Z{{.UpperComponentName}}Data struct {
	{{range $Index, $Element := .ComponentVariables }}{{$Element.UpperIdentifier}}	{{$Element.Type}}
	{{end}}
}

// {{.UpperComponentName}}Component ...
type {{.UpperComponentName}}Component struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]Z{{.UpperComponentName}}Data
}

// Register{{.UpperComponentName}}ComponentWith ...
func Register{{.UpperComponentName}}ComponentWith(e *zinc.ZEntityManager) {
	x := New{{.UpperComponentName}}Component()
	ctx := e.RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// Register{{.UpperComponentName}}Component ...
func Register{{.UpperComponentName}}Component() {
	x := New{{.UpperComponentName}}Component()
	ctx := zinc.Default().RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// New{{.UpperComponentName}}Component ...
func New{{.UpperComponentName}}Component() *{{.UpperComponentName}}Component {
	return &{{.UpperComponentName}}Component{data: make(map[zinc.ZEntityID]Z{{.UpperComponentName}}Data)}
}

func init() {
	Register{{.UpperComponentName}}Component()
}

// SetContext ...
func (c *{{.UpperComponentName}}Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// Add{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Add{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = data
			c.ctx.ComponentAdded(Z{{.UpperComponentName}}, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// Update{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Update{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data, silent bool) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = data
			if !silent {
				c.ctx.ComponentUpdated(Z{{.UpperComponentName}}, id)
			}
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.ZEntityID) (Z{{.UpperComponentName}}Data, error) {
	data, ok := c.data[id]
	if c.ctx.HasEntity(id) {
		if ok {
			return data, nil
		}
		return data, zinc.ErrEntityComponentNotFound
	}
	return data, zinc.ErrEntityNotFound
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(Z{{.UpperComponentName}}, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	} 
	return zinc.ErrEntityNotFound
}

// Add{{.UpperComponentName}}X ...
func Add{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Add{{.UpperComponentName}}(id, data)
}

{{if .Extras }}
// MustAdd{{.UpperComponentName}}X ...
func MustAdd{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Add{{.UpperComponentName}}X(e, id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Add{{.UpperComponentName}} ...
func Add{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	return Add{{.UpperComponentName}}X(zinc.Default(), id, data)
}

{{if .Extras }}
// MustAdd{{.UpperComponentName}} ...
func MustAdd{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Add{{.UpperComponentName}}X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}}SilentlyX ...
func Update{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, data, true)
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}}SilentlyX ...
func MustUpdate{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}}Silently ...
func Update{{.UpperComponentName}}Silently(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	return Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, data)
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}}Silently ...
func MustUpdate{{.UpperComponentName}}Silently(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, data, false)
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}}X ...
func MustUpdate{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}X(e, id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}} ...
func Update{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	return Update{{.UpperComponentName}}X(zinc.Default(), id, data)
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}} ...
func MustUpdate{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Set{{.UpperComponentName}}X ...
func Set{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	if c.HasEntity(id) {
		return c.Update{{.UpperComponentName}}(id, data, false)
	}
	return c.Add{{.UpperComponentName}}(id, data)
}

{{if .Extras }}
// MustSet{{.UpperComponentName}}X ...
func MustSet{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Set{{.UpperComponentName}}X(e, id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Set{{.UpperComponentName}} ...
func Set{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) error {
	return Set{{.UpperComponentName}}X(zinc.Default(), id, data)
}

{{if .Extras }}
// MustSet{{.UpperComponentName}} ...
func MustSet{{.UpperComponentName}}(id zinc.ZEntityID, data Z{{.UpperComponentName}}Data) {
	err := Set{{.UpperComponentName}}(id, data)
	if err != nil {
		panic(err)
	}
}{{end}}

// Has{{.UpperComponentName}}X ...
func Has{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(Z{{.UpperComponentName}})
	return v.HasEntity(id)
}

// Has{{.UpperComponentName}} ...
func Has{{.UpperComponentName}}(id zinc.ZEntityID) bool {
	return Has{{.UpperComponentName}}X(zinc.Default(), id)
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) (Z{{.UpperComponentName}}Data, error) {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

{{if .Extras }}
// Must{{.UpperComponentName}}X ...
func Must{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) Z{{.UpperComponentName}}Data {
	data, err := {{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}{{end}}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.ZEntityID) (Z{{.UpperComponentName}}Data, error) {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

{{if .Extras }}
// Must{{.UpperComponentName}} ...
func Must{{.UpperComponentName}}(id zinc.ZEntityID) Z{{.UpperComponentName}}Data {
	data, err := {{.UpperComponentName}}X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}{{end}}

// Delete{{.UpperComponentName}}X ...
func Delete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(Z{{.UpperComponentName}})
	return v.DeleteEntity(id)
}

{{if .Extras }}
// MustDelete{{.UpperComponentName}}X ...
func MustDelete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := Delete{{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
}{{end}}

// Delete{{.UpperComponentName}} ...
func Delete{{.UpperComponentName}}(id zinc.ZEntityID) error {
	return Delete{{.UpperComponentName}}X(zinc.Default(), id)
}

{{if .Extras }}
// MustDelete{{.UpperComponentName}} ...
func MustDelete{{.UpperComponentName}}(id zinc.ZEntityID) {
	err := Delete{{.UpperComponentName}}(id)
	if err != nil {
		panic(err)
	}
}{{end}}`

	// SingleComponentDataTemplateSource ...
	SingleComponentDataTemplateSource = `package {{.PackageName}}
{{ $ImportLength := len .Imports }}{{if eq $ImportLength 0 }}
import (
	"github.com/SirMetathyst/zinc"
){{else if gt $ImportLength 0}}
import (
	zinc "github.com/SirMetathyst/zinc"
	{{range $Index, $Element := .Imports }}"{{$Element}}"
	{{end}}
){{end}}

// Z{{.UpperComponentName}} ...
var Z{{.UpperComponentName}} uint = uint({{.ComponentNameHash}})

// {{.UpperComponentName}}Component ...
type {{.UpperComponentName}}Component struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]{{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}
}

// Register{{.UpperComponentName}}ComponentWith ...
func Register{{.UpperComponentName}}ComponentWith(e *zinc.ZEntityManager) {
	x := New{{.UpperComponentName}}Component()
	ctx := e.RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// Register{{.UpperComponentName}}Component ...
func Register{{.UpperComponentName}}Component() {
	x := New{{.UpperComponentName}}Component()
	ctx := zinc.Default().RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// New{{.UpperComponentName}}Component ...
func New{{.UpperComponentName}}Component() *{{.UpperComponentName}}Component {
	return &{{.UpperComponentName}}Component{data: make(map[zinc.ZEntityID]{{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}})}
}

func init() {
	Register{{.UpperComponentName}}Component()
}

// SetContext ...
func (c *{{.UpperComponentName}}Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// Add{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Add{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
			c.ctx.ComponentAdded(Z{{.UpperComponentName}}, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// Update{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Update{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}, silent bool) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
			if !silent {
				c.ctx.ComponentUpdated(Z{{.UpperComponentName}}, id)
			}
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.ZEntityID) ({{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}, error) {
	data, ok := c.data[id]
	if c.ctx.HasEntity(id) {
		if ok {
			return data, nil
		}
		return data, zinc.ErrEntityComponentNotFound
	}
	return data, zinc.ErrEntityNotFound
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(Z{{.UpperComponentName}}, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// Add{{.UpperComponentName}}X ...
func Add{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Add{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

{{if .Extras }}
// MustAdd{{.UpperComponentName}}X ...
func MustAdd{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Add{{.UpperComponentName}}X(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Add{{.UpperComponentName}} ...
func Add{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Add{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

{{if .Extras }}
// MustAdd{{.UpperComponentName}} ...
func MustAdd{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Add{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}}SilentlyX ...
func Update{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, true)
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}}SilentlyX ...
func MustUpdate{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}SilentlyX(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}}Silently ...
func Update{{.UpperComponentName}}Silently(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}}Silently ...
func MustUpdate{{.UpperComponentName}}Silently(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, false)
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}}X ...
func MustUpdate{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}X(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Update{{.UpperComponentName}} ...
func Update{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Update{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

{{if .Extras }}
// MustUpdate{{.UpperComponentName}} ...
func MustUpdate{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Set{{.UpperComponentName}}X ...
func Set{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	if c.HasEntity(id) {
		return c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, false)
	}
	return c.Add{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

{{if .Extras }}
// MustSet{{.UpperComponentName}}X ...
func MustSet{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Set{{.UpperComponentName}}X(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Set{{.UpperComponentName}} ...
func Set{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Set{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

{{if .Extras }}
// MustSet{{.UpperComponentName}} ...
func MustSet{{.UpperComponentName}}(id zinc.ZEntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Set{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}{{end}}

// Has{{.UpperComponentName}}X ...
func Has{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(Z{{.UpperComponentName}})
	return v.HasEntity(id)
}

// Has{{.UpperComponentName}} ...
func Has{{.UpperComponentName}}(id zinc.ZEntityID) bool {
	return Has{{.UpperComponentName}}X(zinc.Default(), id)
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) ({{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}, error) {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

{{if .Extras }}
// Must{{.UpperComponentName}}X ...
func Must{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	data, err := {{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}{{end}}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.ZEntityID) ({{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}, error) {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

{{if .Extras }}
// Must{{.UpperComponentName}} ...
func Must{{.UpperComponentName}}(id zinc.ZEntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	data, err := {{.UpperComponentName}}X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}{{end}}

// Delete{{.UpperComponentName}}X ...
func Delete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(Z{{.UpperComponentName}})
	return v.DeleteEntity(id)
}

{{if .Extras }}
// MustDelete{{.UpperComponentName}}X ...
func MustDelete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := Delete{{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
}{{end}}

// Delete{{.UpperComponentName}} ...
func Delete{{.UpperComponentName}}(id zinc.ZEntityID) error {
	return Delete{{.UpperComponentName}}X(zinc.Default(), id)
}

{{if .Extras }}
// MustDelete{{.UpperComponentName}} ...
func MustDelete{{.UpperComponentName}}(id zinc.ZEntityID) {
	err := Delete{{.UpperComponentName}}(id)
	if err != nil {
		panic(err)
	}
}{{end}}`

	// FlagComponentDataTemplateSource ...
	FlagComponentDataTemplateSource = `package {{.PackageName}}
{{ $ImportLength := len .Imports }}{{if eq $ImportLength 0 }}
import (
	"github.com/SirMetathyst/zinc"
){{else if gt $ImportLength 0}}
import (
	zinc "github.com/SirMetathyst/zinc"
	{{range $Index, $Element := .Imports }}"{{$Element}}"
	{{end}}
){{end}}

// Z{{.UpperComponentName}} ...
var Z{{.UpperComponentName}} uint = uint({{.ComponentNameHash}})

// {{.UpperComponentName}}Component ...
type {{.UpperComponentName}}Component struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]bool
}

// Register{{.UpperComponentName}}ComponentWith ...
func Register{{.UpperComponentName}}ComponentWith(e *zinc.ZEntityManager) {
	x := New{{.UpperComponentName}}Component()
	ctx := e.RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// Register{{.UpperComponentName}}Component ...
func Register{{.UpperComponentName}}Component() {
	x := New{{.UpperComponentName}}Component()
	ctx := zinc.Default().RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// New{{.UpperComponentName}}Component ...
func New{{.UpperComponentName}}Component() *{{.UpperComponentName}}Component {
	return &{{.UpperComponentName}}Component{data: make(map[zinc.ZEntityID]bool)}
}

func init() {
	Register{{.UpperComponentName}}Component()
}

// SetContext ...
func (c *{{.UpperComponentName}}Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = true
			c.ctx.ComponentAdded(Z{{.UpperComponentName}}, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(Z{{.UpperComponentName}}, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

{{if .Extras }}
// Must{{.UpperComponentName}}X ...
func Must{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := {{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
}{{end}}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.ZEntityID) error {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

{{if .Extras }}
// Must{{.UpperComponentName}} ...
func Must{{.UpperComponentName}}(id zinc.ZEntityID) {
	err := {{.UpperComponentName}}X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
}{{end}}

// Is{{.UpperComponentName}}X ...
func Is{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(Z{{.UpperComponentName}})
	return v.HasEntity(id)
}

// Is{{.UpperComponentName}} ...
func Is{{.UpperComponentName}}(id zinc.ZEntityID) bool {
	return Is{{.UpperComponentName}}X(zinc.Default(), id)
}

// Not{{.UpperComponentName}}X ...
func Not{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(Z{{.UpperComponentName}})
	return v.DeleteEntity(id)
}

{{if .Extras }}
// MustNot{{.UpperComponentName}}X ...
func MustNot{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := Not{{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
}{{end}}

// Not{{.UpperComponentName}} ...
func Not{{.UpperComponentName}}(id zinc.ZEntityID) error {
	return Not{{.UpperComponentName}}X(zinc.Default(), id)
}

{{if .Extras }}
// MustNot{{.UpperComponentName}} ...
func MustNot{{.UpperComponentName}}(id zinc.ZEntityID) {
	err := Not{{.UpperComponentName}}(id)
	if err != nil {
		panic(err)
	}
}{{end}}`

	// UniqueFlagComponentDataTemplateSource ...
	UniqueFlagComponentDataTemplateSource = `package {{.PackageName}}
{{ $ImportLength := len .Imports }}{{if eq $ImportLength 0 }}
import (
	"github.com/SirMetathyst/zinc"
){{else if gt $ImportLength 0}}
import (
	zinc "github.com/SirMetathyst/zinc"
	{{range $Index, $Element := .Imports }}"{{$Element}}"
	{{end}}
){{end}}

// Z{{.UpperComponentName}} ...
var Z{{.UpperComponentName}} uint = uint({{.ComponentNameHash}})

// {{.UpperComponentName}}Component ...
type {{.UpperComponentName}}Component struct {
	ctx  *zinc.ZContext
	data bool
}

// Register{{.UpperComponentName}}ComponentWith ...
func Register{{.UpperComponentName}}ComponentWith(e *zinc.ZEntityManager) {
	x := New{{.UpperComponentName}}Component()
	ctx := e.RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// Register{{.UpperComponentName}}Component ...
func Register{{.UpperComponentName}}Component() {
	x := New{{.UpperComponentName}}Component()
	ctx := zinc.Default().RegisterComponent(Z{{.UpperComponentName}}, x)
	x.SetContext(ctx)
}

// New{{.UpperComponentName}}Component ...
func New{{.UpperComponentName}}Component() *{{.UpperComponentName}}Component {
	return &{{.UpperComponentName}}Component{data: false}
}

func init() {
	Register{{.UpperComponentName}}Component()
}

// SetContext ...
func (c *{{.UpperComponentName}}Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// Set{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Set{{.UpperComponentName}}(v bool) {
	c.data = v
}

// Is{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Is{{.UpperComponentName}}() bool {
	return c.data
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.ZEntityID) bool {
	return false
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.ZEntityID) error {
	return nil
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager) {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	c.Set{{.UpperComponentName}}(true)
}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}() {
	{{.UpperComponentName}}X(zinc.Default())
}

// Is{{.UpperComponentName}}X ...
func Is{{.UpperComponentName}}X(e *zinc.ZEntityManager) bool {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	return c.Is{{.UpperComponentName}}()
}

// Is{{.UpperComponentName}} ...
func Is{{.UpperComponentName}}() bool {
	return Is{{.UpperComponentName}}X(zinc.Default())
}

// Not{{.UpperComponentName}}X ...
func Not{{.UpperComponentName}}X(e *zinc.ZEntityManager) {
	v := e.Component(Z{{.UpperComponentName}})
	c := v.(*{{.UpperComponentName}}Component)
	c.Set{{.UpperComponentName}}(false)
}

// Not{{.UpperComponentName}} ...
func Not{{.UpperComponentName}}() {
	Not{{.UpperComponentName}}X(zinc.Default())
}`
)

var (
	// MultiComponentDataTemplate ...
	MultiComponentDataTemplate = template.New("MultiComponentData")
	// SingleComponentDataTemplate ...
	SingleComponentDataTemplate = template.New("SingleComponentData")
	// FlagComponentDataTemplate ...
	FlagComponentDataTemplate = template.New("FlagComponentData")
	// UniqueFlagComponentDataTemplate ...
	UniqueFlagComponentDataTemplate = template.New("UniqueFlagComponentData")
)

func init() {
	MultiComponentDataTemplate, _ = MultiComponentDataTemplate.Parse(MultiComponentDataTemplateSource)
	SingleComponentDataTemplate, _ = SingleComponentDataTemplate.Parse(SingleComponentDataTemplateSource)
	FlagComponentDataTemplate, _ = FlagComponentDataTemplate.Parse(FlagComponentDataTemplateSource)
	UniqueFlagComponentDataTemplate, _ = UniqueFlagComponentDataTemplate.Parse(UniqueFlagComponentDataTemplateSource)

	// Register Source Generator ...
	Register("component", ComponentTemplate)
}

// ParseComponentArgs ...
func ParseComponentArgs(Args []string) Data {

	// Vars
	var imports StringSliceFlag
	var componentData StringSliceFlag

	// Flags
	packageName := flag.String("package", "component", "")
	componentName := flag.String("name", "untitled", "")
	unique := flag.Bool("unique", false, "")
	extras := flag.Bool("extras", true, "")

	// Custom Flag Types
	flag.Var(&imports, "import", "")
	flag.Var(&componentData, "var", "")

	// Parse Flags
	flag.CommandLine.Parse(os.Args[2:])

	v := CliData{
		PackageName:   *packageName,
		Imports:       imports,
		ComponentName: *componentName,
		ComponentData: componentData,
		Unique:        *unique,
		Extras:        *extras,
	}
	return NewTemplateDataFrom(v)
}

// ComponentTemplate ...
func ComponentTemplate(args []string) (GeneratorOutput, error) {

	// Generate Template Data
	td := ParseComponentArgs(args)

	var b bytes.Buffer

	if td.Unique {

		// xxType
		if len(td.ComponentVariables) == 0 {
			UniqueFlagComponentDataTemplate.Execute(&b, td)
			return GeneratorOutput{Filename: td.ComponentName + ".go", Content: b.String()}, nil
		}

	} else {

		if len(td.ComponentVariables) == 0 {
			// xxType
			FlagComponentDataTemplate.Execute(&b, td)
			return GeneratorOutput{Filename: td.ComponentName + ".go", Content: b.String()}, nil
		} else if len(td.ComponentVariables) == 1 {
			// xxType
			SingleComponentDataTemplate.Execute(&b, td)
			return GeneratorOutput{Filename: td.ComponentName + ".go", Content: b.String()}, nil
		} else if len(td.ComponentVariables) > 1 {
			// type ComponentData struct {
			//	xxVariable1	xxType
			//	xxVariable2	xxType
			//	... >=2
			// }
			MultiComponentDataTemplate.Execute(&b, td)
			return GeneratorOutput{Filename: td.ComponentName + ".go", Content: b.String()}, nil
		}
	}

	// No Template ...
	return GeneratorOutput{}, ErrNoTemplateForTemplateData
}
