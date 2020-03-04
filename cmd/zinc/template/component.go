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
	"fmt"
){{else if gt $ImportLength 0}}
import (
	zinc "github.com/SirMetathyst/zinc"
	"fmt"
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
func (c *{{.UpperComponentName}}Component) Add{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded({{.UpperComponentName}}Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} already exists on entity %d", id))
	}
}

// Update{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Update{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data, silent bool) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated({{.UpperComponentName}}Key, id)
		}
	} else {
		panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} does not exist on entity %d", id))
	}
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.EntityID) {{.UpperComponentName}}Data {
	if data, ok := c.data[id]; ok {
		return data
	}
	panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} does not exist on entity %d", id))
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.EntityID) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted({{.UpperComponentName}}Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} does not exist on entity %d", id))
	}
}

// Add{{.UpperComponentName}}X ...
func Add{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	c.Add{{.UpperComponentName}}(id, data)
}

// Add{{.UpperComponentName}} ...
func Add{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	Add{{.UpperComponentName}}X(zinc.Default(), id, data)
}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	c.Update{{.UpperComponentName}}(id, data, true)
}

// Update{{.UpperComponentName}}Silently ...
func Update{{.UpperComponentName}}Silently(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, data)
}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, data {{.UpperComponentName}}Data) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	c.Update{{.UpperComponentName}}(id, data, false)
}

// Update{{.UpperComponentName}} ...
func Update{{.UpperComponentName}}(id zinc.EntityID, data {{.UpperComponentName}}Data) {
	Update{{.UpperComponentName}}X(zinc.Default(), id, data)
}

// Has{{.UpperComponentName}}X ...
func Has{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v, _ := e.Component({{.UpperComponentName}}Key)
	return v.HasEntity(id)
}

// Has{{.UpperComponentName}} ...
func Has{{.UpperComponentName}}(id zinc.EntityID) bool {
	return Has{{.UpperComponentName}}X(zinc.Default(), id)
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {{.UpperComponentName}}Data {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.EntityID) {{.UpperComponentName}}Data {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

// Delete{{.UpperComponentName}}X ...
func Delete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	v.DeleteEntity(id)
}

// Delete{{.UpperComponentName}} ...
func Delete{{.UpperComponentName}}(id zinc.EntityID) {
	Delete{{.UpperComponentName}}X(zinc.Default(), id)
}`

	// SingleComponentDataTemplateSource ...
	SingleComponentDataTemplateSource = `package {{.PackageName}}
{{ $ImportLength := len .Imports }}{{if eq $ImportLength 0 }}
import (
	"github.com/SirMetathyst/zinc"
	"fmt"
){{else if gt $ImportLength 0}}
import (
	zinc "github.com/SirMetathyst/zinc"
	"fmt"
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
func (c *{{.UpperComponentName}}Component) Add{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
		c.ctx.ComponentAdded({{.UpperComponentName}}Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} already exists on entity %d", id))
	}
}

// Update{{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) Update{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}, silent bool) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
		if !silent {
			c.ctx.ComponentUpdated({{.UpperComponentName}}Key, id)
		}
	} else {
		panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} does not exist on entity %d", id))
	}
}

// HasEntity ...
func (c *{{.UpperComponentName}}Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// {{.UpperComponentName}} ...
func (c *{{.UpperComponentName}}Component) {{.UpperComponentName}}(id zinc.EntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	if {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, ok := c.data[id]; ok {
		return {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}
	}
	panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} does not exist on entity %d", id))
}

// DeleteEntity ...
func (c *{{.UpperComponentName}}Component) DeleteEntity(id zinc.EntityID) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted({{.UpperComponentName}}Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component {{.UpperComponentName}} does not exist on entity %d", id))
	}
}

// Add{{.UpperComponentName}}X ...
func Add{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	c.Add{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// Add{{.UpperComponentName}} ...
func Add{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	Add{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, true)
}

// Update{{.UpperComponentName}}Silently ...
func Update{{.UpperComponentName}}Silently(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	Update{{.UpperComponentName}}SilentlyX(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// Update{{.UpperComponentName}}X ...
func Update{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	c.Update{{.UpperComponentName}}(id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}}, false)
}

// Update{{.UpperComponentName}} ...
func Update{{.UpperComponentName}}(id zinc.EntityID, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}} {{$Element.Type}}{{end}}) {
	Update{{.UpperComponentName}}X(zinc.Default(), id, {{range $Index, $Element := .ComponentVariables }}{{$Element.Identifier}}{{end}})
}

// Has{{.UpperComponentName}}X ...
func Has{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v, _ := e.Component({{.UpperComponentName}}Key)
	return v.HasEntity(id)
}

// Has{{.UpperComponentName}} ...
func Has{{.UpperComponentName}}(id zinc.EntityID) bool {
	return Has{{.UpperComponentName}}X(zinc.Default(), id)
}

// {{.UpperComponentName}}X ...
func {{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	v, _ := e.Component({{.UpperComponentName}}Key)
	c := v.(*{{.UpperComponentName}}Component)
	return c.{{.UpperComponentName}}(id)
}

// {{.UpperComponentName}} ...
func {{.UpperComponentName}}(id zinc.EntityID) {{range $Index, $Element := .ComponentVariables }}{{$Element.Type}}{{end}} {
	return {{.UpperComponentName}}X(zinc.Default(), id)
}

// Delete{{.UpperComponentName}}X ...
func Delete{{.UpperComponentName}}X(e *zinc.ZEntityManager, id zinc.EntityID) {
	v, _ := e.Component({{.UpperComponentName}}Key)
	v.DeleteEntity(id)
}

// Delete{{.UpperComponentName}} ...
func Delete{{.UpperComponentName}}(id zinc.EntityID) {
	Delete{{.UpperComponentName}}X(zinc.Default(), id)
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
