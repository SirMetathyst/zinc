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
var Template1s string = `// Package ${Package} ...
// Generated by the zinc tool.  DO NOT EDIT!
// Source: zinc_${Component}
package ${Package}

import (
	"github.com/SirMetathyst/zinc"
${ImportData}
)

// ${Component}Key ...
const ${Component}Key uint = ${ComponentKey}

// ${Component}Data ...
type ${Component}Data struct {
${ComponentData}	
}

// ${Component}Component ...
type ${Component}Component struct {
	ctx zinc.CTX
	data map[zinc.EntityID]${Component}Data
}

// New${Component}Component ...
func New${Component}Component() *${Component}Component {
	return &${Component}Component{
		data: make(map[zinc.EntityID]${Component}Data),
	}
}

// SetContext ...
func (c *${Component}Component) SetContext(ctx zinc.CTX) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

func init() {
	x := New${Component}Component()
	ctx := zinc.Default().RegisterComponent(${Component}Key, x)
	x.SetContext(ctx)
}

// DeleteEntity ...
func (c *${Component}Component) DeleteEntity(id zinc.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *${Component}Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// Set${Component} ...
func (c *${Component}Component) Set${Component}(id zinc.EntityID, ${component} ${Component}Data) {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = ${component}
			c.ctx.ComponentUpdated(${Component}Key, id)
		} else {
			c.data[id] = ${component}
			c.ctx.ComponentAdded(${Component}Key, id)
		}
	}
}

// ${Component} ...
func (c *${Component}Component) ${Component}(id zinc.EntityID) ${Component}Data {
	return c.data[id]
}

// Delete${Component} ...
func (c *${Component}Component) Delete${Component}(id zinc.EntityID) {
	delete(c.data, id)
	c.ctx.ComponentDeleted(${Component}Key, id)
}

// Set${Component}X ...
func Set${Component}X(e *zinc.EntityManager, id zinc.EntityID, ${component} ${Component}Data) {
	v, _ := e.Component(${Component}Key)
	c := v.(*${Component}Component)
	c.Set${Component}(id, ${component})
}

// Set${Component} ...
func Set${Component}(id zinc.EntityID, ${component} ${Component}Data) {
	Set${Component}X(zinc.Default(), id, ${component})
}

// ${Component}X ...
func ${Component}X(e *zinc.EntityManager, id zinc.EntityID) ${Component}Data {
	v, _ := e.Component(${Component}Key)
	c := v.(*${Component}Component)
	return c.${Component}(id)
}

// ${Component} ...
func ${Component}(id zinc.EntityID) ${Component}Data {
	return ${Component}X(zinc.Default(), id)
}

// Delete${Component}X ...
func Delete${Component}X(e *zinc.EntityManager, id zinc.EntityID) {
	v, _ := e.Component(${Component}Key)
	c := v.(*${Component}Component)
	c.Delete${Component}(id)
}

// Delete${Component} ...
func Delete${Component}(id zinc.EntityID) {
	Delete${Component}X(zinc.Default(), id)
}

// Has${Component}X ...
func Has${Component}X(e *zinc.EntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(${Component}Key)
	return v.HasEntity(id)
}

// Has${Component} ...
func Has${Component}(id zinc.EntityID) bool {
	return Has${Component}X(zinc.Default(), id)
}`

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}

// ImportDataTemplate ...
func ImportDataTemplate(componentParams []string) string {
	bb := bytes.Buffer{}
	importMap := make(map[string]struct{}, 0)
	importSlice := make([]string, 0)
	for _, p := range componentParams {
		split := strings.Split(p, ":")
		if len(split) < 3 {
			continue
		}
		if _, ok := importMap[split[0]]; !ok {
			importSlice = append(importSlice, split[0])
			importMap[split[0]] = struct{}{}
		}
	}
	for i, x := range importSlice {
		bb.WriteString("\t\"")
		bb.WriteString(x)
		bb.WriteRune('"')
		if i != len(importSlice)-1 {
			bb.WriteRune('\n')
		}
	}
	return bb.String()
}

// StructDataTemplate ...
func StructDataTemplate(componentParams []string) string {
	bb := bytes.Buffer{}
	for i, p := range componentParams {
		split := strings.Split(p, ":")
		c := 0
		if len(split) < 2 {
			log.Fatal("err: template: data must be in format name:type or package:name:type")
		} else if len(split) >= 3 {
			c++
		}
		bb.WriteRune('\t')
		bb.WriteString(strings.Title(split[c]))
		bb.WriteRune(' ')
		bb.WriteString(split[c+1])
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
	tpl = strings.Replace(tpl, "${ImportData}", ImportDataTemplate(componentData), -1)
	err := ioutil.WriteFile(path.Join(outputPath, fmt.Sprint("zinc_", titledComponentName, ".go")), []byte(tpl), 0644)
	if err != nil {
		log.Fatalf("err: template: %s", err)
	}
}

func main() {
	app := &cli.App{
		Name:  "Zinc CLI",
		Usage: "manipulate your zinc project",
		Authors: []*cli.Author{
			{Name: "SirMetathyst"},
		},
		Version: "0.6.0",
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
							&cli.StringFlag{Name: "package", Aliases: []string{"p"}, Required: false, Value: "zinccomponents"},
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