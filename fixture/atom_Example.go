package fixture

import "github.com/SirMetathyst/atom"

// ExampleKey ...
const ExampleKey uint = 311628609

// ExampleData ...
type ExampleData struct {
	Value int	
}

// ExampleComponent ...
type ExampleComponent struct {
	context atom.Context
	data map[atom.EntityID]ExampleData
}

// NewExampleComponent ...
func NewExampleComponent() *ExampleComponent {
	return &ExampleComponent{
		data: make(map[atom.EntityID]ExampleData),
	}
}

func init() {
	x := NewExampleComponent()
	context := atom.Default().RegisterComponent(ExampleKey, x)
	x.context = context 
}

// EntityDeleted ...
func (c *ExampleComponent) EntityDeleted(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *ExampleComponent) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetExample ...
func (c *ExampleComponent) SetExample(id atom.EntityID, example ExampleData) {
	if c.context.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = example
			c.context.ComponentUpdated(ExampleKey, id)
		} else {
			c.data[id] = example
			c.context.ComponentAdded(ExampleKey, id)
		}
	}
}

// Example ...
func (c *ExampleComponent) Example(id atom.EntityID) ExampleData {
	return c.data[id]
}

// DeleteExample ...
func (c *ExampleComponent) DeleteExample(id atom.EntityID) {
	delete(c.data, id)
	c.context.ComponentDeleted(ExampleKey, id)
}

// SetExampleX ...
func SetExampleX(e *atom.EntityManager, id atom.EntityID, example ExampleData) {
	v := e.Component(ExampleKey)
	c := v.(*ExampleComponent)
	c.SetExample(id, example)
}

// SetExample ...
func SetExample(id atom.EntityID, example ExampleData) {
	SetExampleX(atom.Default(), id, example)
}

// ExampleX ...
func ExampleX(e *atom.EntityManager, id atom.EntityID) ExampleData {
	v := e.Component(ExampleKey)
	c := v.(*ExampleComponent)
	return c.Example(id)
}

// Example ...
func Example(id atom.EntityID) ExampleData {
	return ExampleX(atom.Default(), id)
}

// DeleteExampleX ...
func DeleteExampleX(e *atom.EntityManager, id atom.EntityID) {
	v := e.Component(ExampleKey)
	c := v.(*ExampleComponent)
	c.DeleteExample(id)
}

// DeleteExample ...
func DeleteExample(id atom.EntityID) {
	DeleteExampleX(atom.Default(), id)
}

// HasExampleX ...
func HasExampleX(e *atom.EntityManager, id atom.EntityID) bool {
	v := e.Component(ExampleKey)
	return v.HasEntity(id)
}

// HasExample ...
func HasExample(id atom.EntityID) bool {
	return HasExampleX(atom.Default(), id)
}