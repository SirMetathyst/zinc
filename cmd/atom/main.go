package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

// Template1s ...
var Template1s string = `package ${Package}

import "github.com/SirMetathyst/atom"

// ${Component}Key ...
const ${Component}Key uint = ${ComponentKey}

// ${Component}Data ...
type ${Component}Data struct {
${ComponentData}	
}

// ${Component}Component ...
type ${Component}Component struct {
	context atom.Context
	data map[atom.EntityID]${Component}Data
}

// New${Component}Component ...
func New${Component}Component() *${Component}Component {
	return &${Component}Component{
		data: make(map[atom.EntityID]${Component}Data),
	}
}

func init() {
	x := New${Component}Component()
	context := atom.Default().RegisterComponent(${Component}Key, x)
	x.context = context 
}

// EntityDeleted ...
func (c *${Component}Component) EntityDeleted(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *${Component}Component) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// Set${Component} ...
func (c *${Component}Component) Set${Component}(id atom.EntityID, ${component} ${Component}Data) {
	if c.context.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = ${component}
			c.context.ComponentUpdated(${Component}Key, id)
		} else {
			c.data[id] = ${component}
			c.context.ComponentAdded(${Component}Key, id)
		}
	}
}

// ${Component} ...
func (c *${Component}Component) ${Component}(id atom.EntityID) ${Component}Data {
	return c.data[id]
}

// Delete${Component} ...
func (c *${Component}Component) Delete${Component}(id atom.EntityID) {
	delete(c.data, id)
	c.context.ComponentDeleted(${Component}Key, id)
}

// Set${Component}X ...
func Set${Component}X(e *atom.EntityManager, id atom.EntityID, ${component} ${Component}Data) {
	v := e.Component(${Component}Key)
	c := v.(*${Component}Component)
	c.Set${Component}(id, ${component})
}

// Set${Component} ...
func Set${Component}(id atom.EntityID, ${component} ${Component}Data) {
	Set${Component}X(atom.Default(), id, ${component})
}

// ${Component}X ...
func ${Component}X(e *atom.EntityManager, id atom.EntityID) ${Component}Data {
	v := e.Component(${Component}Key)
	c := v.(*${Component}Component)
	return c.${Component}(id)
}

// ${Component} ...
func ${Component}(id atom.EntityID) ${Component}Data {
	return ${Component}X(atom.Default(), id)
}

// Delete${Component}X ...
func Delete${Component}X(e *atom.EntityManager, id atom.EntityID) {
	v := e.Component(${Component}Key)
	c := v.(*${Component}Component)
	c.Delete${Component}(id)
}

// Delete${Component} ...
func Delete${Component}(id atom.EntityID) {
	Delete${Component}X(atom.Default(), id)
}

// Has${Component}X ...
func Has${Component}X(e *atom.EntityManager, id atom.EntityID) bool {
	v := e.Component(${Component}Key)
	return v.HasEntity(id)
}

// Has${Component} ...
func Has${Component}(id atom.EntityID) bool {
	return Has${Component}X(atom.Default(), id)
}`

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}

// StructDataTemplate ...
func StructDataTemplate(componentParams []string) string {
	bb := bytes.Buffer{}
	for i, p := range componentParams {
		split := strings.Split(p, ":")
		if len(split) != 2 {
			log.Fatal("err: template: data must be in format name:type")
		}
		bb.WriteRune('\t')
		bb.WriteString(strings.Title(split[0]))
		bb.WriteRune(' ')
		bb.WriteString(split[1])
		if i != len(componentParams)-1 {
			bb.WriteRune('\n')
		}
	}
	return bb.String()
}

// Template1 ...
func Template1(packageName string, componentName string, componentData []string, outputPath string) {
	titledComponentName := strings.Title(componentName)
	tpl := strings.Replace(Template1s, "${Package}", packageName, -1)
	tpl = strings.Replace(tpl, "${Component}", titledComponentName, -1)
	tpl = strings.Replace(tpl, "${ComponentKey}", hash(titledComponentName), -1)
	tpl = strings.Replace(tpl, "${component}", strings.ToLower(componentName), -1)
	tpl = strings.Replace(tpl, "${ComponentData}", StructDataTemplate(componentData), -1)
	err := ioutil.WriteFile(path.Join(outputPath, fmt.Sprint("atom_", titledComponentName, ".go")), []byte(tpl), 0644)
	if err != nil {
		log.Fatalf("err: template: %s", err)
	}
}

func main() {
	app := &cli.App{
		Name:  "Atom CLI",
		Usage: "manipulate your atom project",
		Authors: []*cli.Author{
			{Name: "SirMetathyst"},
		},
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:    "component",
				Aliases: []string{"c"},
				Usage:   "Shows a list of commands about components",
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "add a component",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "package", Aliases: []string{"p"}, Required: false, Value: "atomcomponents"},
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
							&cli.StringSliceFlag{Name: "data", Aliases: []string{"d"}, Required: true},
							&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Required: false, Value: "./"},
						},
						Action: func(c *cli.Context) error {
							Template1(c.String("package"), c.String("name"), c.StringSlice("data"), c.String("output"))
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
