package template

import (
	"bytes"
	"errors"
	"hash/fnv"
	"log"
	"strings"
)

var (
	// ErrNoTemplateForTemplateData ...
	ErrNoTemplateForTemplateData = errors.New("template: No template for template data")
)

// GeneratorOutput ...
type GeneratorOutput struct {
	Filename string
	Content  string
}

// GeneratorFunc ...
type GeneratorFunc func(Args []string) (GeneratorOutput, error)

var (
	// Generator ...
	Generator = map[string]GeneratorFunc{}
)

// ComponentVariable ...
type ComponentVariable struct {
	Identifier string
	Type       string
}

// UpperIdentifier ...
func (c *ComponentVariable) UpperIdentifier() string {
	return strings.Title(c.Identifier)
}

// Data ...
type Data struct {
	PackageName        string
	Imports            []string
	ComponentName      string
	ComponentVariables []ComponentVariable
}

// UpperComponentName ...
func (td Data) UpperComponentName() string {
	return strings.Title(td.ComponentName)
}

// ComponentNameHash ...
func (td Data) ComponentNameHash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(strings.ToLower(td.ComponentName)))
	return h.Sum32()
}

// Register ...
func Register(command string, f GeneratorFunc) {
	Generator[command] = f
}

// NewComponentVariablesFrom ...
func NewComponentVariablesFrom(variables []string, delimeter string) []ComponentVariable {
	var s []ComponentVariable
	for _, variable := range variables {
		param := strings.Split(variable, delimeter)
		if len(param) == 1 {
			s = append(s, ComponentVariable{Identifier: "value", Type: param[0]})
		} else if len(param) == 2 {
			s = append(s, ComponentVariable{Identifier: param[0], Type: param[1]})
		} else {
			log.Fatalf("Format Usage: name%stype", delimeter)
		}
	}
	return s
}

// NewTemplateDataFrom ...
func NewTemplateDataFrom(packageName string, imports []string, componentName string, componentData []string) Data {
	componentVariables := NewComponentVariablesFrom(componentData, ":")
	return Data{
		PackageName:        packageName,
		ComponentName:      componentName,
		ComponentVariables: componentVariables,
		Imports:            imports,
	}
}

// StringSliceFlag ...
type StringSliceFlag []string

func (i *StringSliceFlag) String() string {
	return "my string representation"
}

// Set ...
func (i *StringSliceFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// PrintCommands ...
func PrintCommands() string {
	var b bytes.Buffer
	i := 0
	for k := range Generator {
		b.WriteString(k)
		if i != len(Generator)-1 {
			b.WriteRune('|')
		}
		i++
	}
	return b.String()
}
