package kit

import (
	"github.com/SirMetathyst/zinc"
	"fmt"
)

// LocalScale2Key ...
var LocalScale2Key uint = uint(54082440)

//LocalScale2Data ...
type LocalScale2Data struct {
	X	float32
	Y	float32
	
}

// LocalScale2Component ...
type LocalScale2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]LocalScale2Data
}

// RegisterLocalScale2ComponentWith ...
func RegisterLocalScale2ComponentWith(e *zinc.ZEntityManager) {
	x := NewLocalScale2Component()
	ctx := e.RegisterComponent(LocalScale2Key, x)
	x.SetContext(ctx)
}

// RegisterLocalScale2Component ...
func RegisterLocalScale2Component() {
	x := NewLocalScale2Component()
	ctx := zinc.Default().RegisterComponent(LocalScale2Key, x)
	x.SetContext(ctx)
}

// NewLocalScale2Component ...
func NewLocalScale2Component() *LocalScale2Component {
	return &LocalScale2Component{data: make(map[zinc.EntityID]LocalScale2Data)}
}

func init() {
	RegisterLocalScale2Component()
}

// SetContext ...
func (c *LocalScale2Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// AddLocalScale2 ...
func (c *LocalScale2Component) AddLocalScale2(id zinc.EntityID, data LocalScale2Data) {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded(LocalScale2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component LocalScale2 already exists on entity %d", id))
	}
}

// UpdateLocalScale2 ...
func (c *LocalScale2Component) UpdateLocalScale2(id zinc.EntityID, data LocalScale2Data, silent bool) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated(LocalScale2Key, id)
		}
	} else {
		panic(fmt.Sprintf("zinc: component LocalScale2 does not exist on entity %d", id))
	}
}

// HasEntity ...
func (c *LocalScale2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalScale2 ...
func (c *LocalScale2Component) LocalScale2(id zinc.EntityID) LocalScale2Data {
	if data, ok := c.data[id]; ok {
		return data
	}
	panic(fmt.Sprintf("zinc: component LocalScale2 does not exist on entity %d", id))
}

// DeleteEntity ...
func (c *LocalScale2Component) DeleteEntity(id zinc.EntityID) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(LocalScale2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component LocalScale2 does not exist on entity %d", id))
	}
}

// AddLocalScale2X ...
func AddLocalScale2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalScale2Data) {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	c.AddLocalScale2(id, data)
}

// AddLocalScale2 ...
func AddLocalScale2(id zinc.EntityID, data LocalScale2Data) {
	AddLocalScale2X(zinc.Default(), id, data)
}

// UpdateLocalScale2X ...
func UpdateLocalScale2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data LocalScale2Data) {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	c.UpdateLocalScale2(id, data, true)
}

// UpdateLocalScale2Silently ...
func UpdateLocalScale2Silently(id zinc.EntityID, data LocalScale2Data) {
	UpdateLocalScale2SilentlyX(zinc.Default(), id, data)
}

// UpdateLocalScale2X ...
func UpdateLocalScale2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalScale2Data) {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	c.UpdateLocalScale2(id, data, false)
}

// UpdateLocalScale2 ...
func UpdateLocalScale2(id zinc.EntityID, data LocalScale2Data) {
	UpdateLocalScale2X(zinc.Default(), id, data)
}

// HasLocalScale2X ...
func HasLocalScale2X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(LocalScale2Key)
	return v.HasEntity(id)
}

// HasLocalScale2 ...
func HasLocalScale2(id zinc.EntityID) bool {
	return HasLocalScale2X(zinc.Default(), id)
}

// LocalScale2X ...
func LocalScale2X(e *zinc.ZEntityManager, id zinc.EntityID) LocalScale2Data {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	return c.LocalScale2(id)
}

// LocalScale2 ...
func LocalScale2(id zinc.EntityID) LocalScale2Data {
	return LocalScale2X(zinc.Default(), id)
}

// DeleteLocalScale2X ...
func DeleteLocalScale2X(e *zinc.ZEntityManager, id zinc.EntityID) {
	v, _ := e.Component(LocalScale2Key)
	v.DeleteEntity(id)
}

// DeleteLocalScale2 ...
func DeleteLocalScale2(id zinc.EntityID) {
	DeleteLocalScale2X(zinc.Default(), id)
}