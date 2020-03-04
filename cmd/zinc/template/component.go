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

// {{.UpperComponentName}}Key ...
var {{.UpperComponentName}}Key uint = uint({{.ComponentNameHash}})

//{{.UpperComponentName}}Data ...
type {{.UpperComponentName}}Data struct {
	{{range $Index, $Element := .ComponentVariables }}{{$Element.UpperIdentifier}}	{{$Element.Type}}
	{{end}}
}

// {{.UpperComponentName}}Component ...
type {{.UpperComponentName}}Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]{{.UpperComponentName}}Data
}

// Register{{.UpperComponentName}}ComponentWith ...
func Register{{.UpperComponentName}}ComponentWith(e *zinc.ZEntityManager) {
	x := New{{.UpperComponentName}}Component()
	ctx := e.RegisterComponent({{.UpperComponentName}}Key, x)
	x.SetContext(ctx)
}

// Register{{.UpperComponentName}}Component ...
func Register{{.UpperComponentName}}Component() {
	x := New{{.UpperComponentName}}Component()
	ctx := zinc.Default().RegisterComponent({{.UpperComponentName}}Key, x)
	x.SetContext(ctx)
}

// New{{.UpperComponentName}}Component ...
func New{{.UpperComponentName}}Component() *{{.UpperComponentName}}Component {
	return &{{.UpperComponentName}}Component{data: make(map[zinc.EntityID]{{.UpperComponentName}}Data)}
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
func (c *{{.UpperComponentName}}Component) Add{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded({{.UpperComponentName}}Key, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// Update{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Update{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data, silent bool) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated({{.UpperComponentName}}Key, id)
		}
		return nil
	}
	return zinc.ErrComponentNotFound
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.EntityID) ({{.UpperComponentName}}Data, error) {
	data, ok := c.data[id]
	if ok {
		return data, nil
	}
	return data, zinc.ErrComponentNotFound
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.EntityID) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted({{.UpperComponentName}}Key, id)
		return nil
	} 
	return zinc.ErrComponentNotFound
}

// Add{{.UpperComponentName}}X ...
func Add{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.Add{{.UpperComponentName}}(id, data)
}

// MustAdd{{.UpperComponentName}}X ...
func MustAdd{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) {
	err := Add{{.UpperComponentName}}X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// Add{{.UpperComponentName}} ...
func Add{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	return Add{{.UpperComponentName}}X(zinc.Default(), id, data)
}

// MustAdd{{.UpperComponentName}} ...
func MustAdd{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	err := Add{{.UpperComponentName}}X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}}SilentlyX ...
func Update{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, data, true)
}

// MustUpdate{{.UpperComponentName}}SilentlyX ...
func MustUpdate{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}}Silently ...
func Update{{.UpperComponentName}}Silently(id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	return Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, data)
}

// MustUpdate{{.UpperComponentName}}Silently ...
func MustUpdate{{.UpperComponentName}}Silently(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, data, false)
}

// MustUpdate{{.UpperComponentName}}X ...
func MustUpdate{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}} ...
func Update{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) error {
	return Update{{.UpperComponentName}}X(zinc.Default(), id, data)
}

// MustUpdate{{.UpperComponentName}} ...
func MustUpdate{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	err := Update{{.UpperComponentName}}X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// Has{{.UpperComponentName}}X ...
func Has{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v := e.Component({{.UpperComponentName}}Key)
	return v.HasEntity(id)
}

// Has{{.UpperComponentName}} ...
func Has{{.UpperComponentName}}(id zinc.EntityID) bool {
	return Has{{.UpperComponentName}}X(zinc.Default(), id)
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) ({{.UpperComponentName}}Data, error) {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

// Must{{.UpperComponentName}}X ...
func Must{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {{.UpperComponentName}}Data {
	data, err := {{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.EntityID) ({{.UpperComponentName}}Data, error) {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

// Must{{.UpperComponentName}} ...
func Must{{.UpperComponentName}}(id zinc.EntityID) {{.UpperComponentName}}Data {
	data, err := {{.UpperComponentName}}X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// Delete{{.UpperComponentName}}X ...
func Delete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) error {
	v := e.Component({{.UpperComponentName}}Key)
	return v.DeleteEntity(id)
}

// MustDelete{{.UpperComponentName}}X ...
func MustDelete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {
	err := Delete{{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
}

// Delete{{.UpperComponentName}} ...
func Delete{{.UpperComponentName}}(id zinc.EntityID) error {
	return Delete{{.UpperComponentName}}X(zinc.Default(), id)
}

// MustDelete{{.UpperComponentName}} ...
func MustDelete{{.UpperComponentName}}(id zinc.EntityID) {
	err := Delete{{.UpperComponentName}}(id)
	if err != nil {
		panic(err)
	}
}`

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

// {{.UpperComponentName}}Key ...
var {{.UpperComponentName}}Key uint = uint({{.ComponentNameHash}})

// {{.UpperComponentName}}Component ...
type {{.UpperComponentName}}Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]{{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}
}

// Register{{.UpperComponentName}}ComponentWith ...
func Register{{.UpperComponentName}}ComponentWith(e *zinc.ZEntityManager) {
	x := New{{.UpperComponentName}}Component()
	ctx := e.RegisterComponent({{.UpperComponentName}}Key, x)
	x.SetContext(ctx)
}

// Register{{.UpperComponentName}}Component ...
func Register{{.UpperComponentName}}Component() {
	x := New{{.UpperComponentName}}Component()
	ctx := zinc.Default().RegisterComponent({{.UpperComponentName}}Key, x)
	x.SetContext(ctx)
}

// New{{.UpperComponentName}}Component ...
func New{{.UpperComponentName}}Component() *{{.UpperComponentName}}Component {
	return &{{.UpperComponentName}}Component{data: make(map[zinc.EntityID]{{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}})}
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
func (c *{{.UpperComponentName}}Component) Add{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
		c.ctx.ComponentAdded({{.UpperComponentName}}Key, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// Update{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Update{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}, silent bool) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
		if !silent {
			c.ctx.ComponentUpdated({{.UpperComponentName}}Key, id)
		}
		return nil
	}
	return zinc.ErrComponentNotFound
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.EntityID) ({{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}, error) {
	data, ok := c.data[id]
	if ok {
		return data, nil
	}
	return data, zinc.ErrComponentNotFound
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.EntityID) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted({{.UpperComponentName}}Key, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// Add{{.UpperComponentName}}X ...
func Add{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.Add{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// MustAdd{{.UpperComponentName}}X ...
func MustAdd{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Add{{.UpperComponentName}}X(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}

// Add{{.UpperComponentName}} ...
func Add{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Add{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// MustAdd{{.UpperComponentName}} ...
func MustAdd{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Add{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}}SilentlyX ...
func Update{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, true)
}

// MustUpdate{{.UpperComponentName}}SilentlyX ...
func MustUpdate{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}SilentlyX(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}}Silently ...
func Update{{.UpperComponentName}}Silently(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// MustUpdate{{.UpperComponentName}}Silently ...
func MustUpdate{{.UpperComponentName}}Silently(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, false)
}

// MustUpdate{{.UpperComponentName}}X ...
func MustUpdate{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}X(e, id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}

// Update{{.UpperComponentName}} ...
func Update{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) error {
	return Update{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// MustUpdate{{.UpperComponentName}} ...
func MustUpdate{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	err := Update{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
	if err != nil {
		panic(err)
	}
}

// Has{{.UpperComponentName}}X ...
func Has{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v := e.Component({{.UpperComponentName}}Key)
	return v.HasEntity(id)
}

// Has{{.UpperComponentName}} ...
func Has{{.UpperComponentName}}(id zinc.EntityID) bool {
	return Has{{.UpperComponentName}}X(zinc.Default(), id)
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) ({{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}, error) {
	v := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

// Must{{.UpperComponentName}}X ...
func Must{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	data, err := {{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.EntityID) ({{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}}, error) {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

// Must{{.UpperComponentName}} ...
func Must{{.UpperComponentName}}(id zinc.EntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	data, err := {{.UpperComponentName}}X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// Delete{{.UpperComponentName}}X ...
func Delete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) error {
	v := e.Component({{.UpperComponentName}}Key)
	return v.DeleteEntity(id)
}

// MustDelete{{.UpperComponentName}}X ...
func MustDelete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {
	err := Delete{{.UpperComponentName}}X(e, id)
	if err != nil {
		panic(err)
	}
}

// Delete{{.UpperComponentName}} ...
func Delete{{.UpperComponentName}}(id zinc.EntityID) error {
	return Delete{{.UpperComponentName}}X(zinc.Default(), id)
}

// MustDelete{{.UpperComponentName}} ...
func MustDelete{{.UpperComponentName}}(id zinc.EntityID) {
	err := Delete{{.UpperComponentName}}(id)
	if err != nil {
		panic(err)
	}
}`
)

var (
	// MultiComponentDataTemplate ...
	MultiComponentDataTemplate = template.New("MultiComponentData")
	// SingleComponentDataTemplate ...
	SingleComponentDataTemplate = template.New("SingleComponentData")
)

func init() {
	MultiComponentDataTemplate, _ = MultiComponentDataTemplate.Parse(MultiComponentDataTemplateSource)
	SingleComponentDataTemplate, _ = SingleComponentDataTemplate.Parse(SingleComponentDataTemplateSource)

	// Register Source Generator ...
	Register("component", ComponentTemplate)
}

// ParseComponentArgs ...
func ParseComponentArgs(Args []string) Data {
	var imports StringSliceFlag
	var componentVariables StringSliceFlag
	packageName := flag.String("package", "component", "")
	componentName := flag.String("name", "untitled", "")
	flag.Var(&imports, "import", "")
	flag.Var(&componentVariables, "var", "")
	flag.CommandLine.Parse(os.Args[2:])
	return NewTemplateDataFrom(*packageName, imports, *componentName, componentVariables)
}

// ComponentTemplate ...
func ComponentTemplate(args []string) (GeneratorOutput, error) {

	// Generate Template Data
	td := ParseComponentArgs(args)

	var b bytes.Buffer

	// type ComponentData struct {
	//	xxVariable1	xxType
	//	xxVariable2	xxType
	//	... >=2
	// }
	if len(td.ComponentVariables) > 1 {
		MultiComponentDataTemplate.Execute(&b, td)
		return GeneratorOutput{Filename: td.ComponentName + ".go", Content: b.String()}, nil
	}

	// xxType
	if len(td.ComponentVariables) == 1 {
		SingleComponentDataTemplate.Execute(&b, td)
		return GeneratorOutput{Filename: td.ComponentName + ".go", Content: b.String()}, nil
	}

	// No Template ...
	return GeneratorOutput{}, ErrNoTemplateForTemplateData
}
