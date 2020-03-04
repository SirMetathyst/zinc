package kit

import (
	"github.com/SirMetathyst/zinc"
	"fmt"
)

// LocalRotation2Key ...
var LocalRotation2Key uint = uint(4108562484)

//LocalRotation2Data ...
type LocalRotation2Data struct {
	X	float32
	Y	float32
	
}

// LocalRotation2Component ...
type LocalRotation2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]LocalRotation2Data
}

// RegisterLocalRotation2ComponentWith ...
func RegisterLocalRotation2ComponentWith(e *zinc.ZEntityManager) {
	x := NewLocalRotation2Component()
	ctx := e.RegisterComponent(LocalRotation2Key, x)
	x.SetContext(ctx)
}

// RegisterLocalRotation2Component ...
func RegisterLocalRotation2Component() {
	x := NewLocalRotation2Component()
	ctx := zinc.Default().RegisterComponent(LocalRotation2Key, x)
	x.SetContext(ctx)
}

// NewLocalRotation2Component ...
func NewLocalRotation2Component() *LocalRotation2Component {
	return &LocalRotation2Component{data: make(map[zinc.EntityID]LocalRotation2Data)}
}

func init() {
	RegisterLocalRotation2Component()
}

// SetContext ...
func (c *LocalRotation2Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// AddLocalRotation2 ...
func (c *LocalRotation2Component) AddLocalRotation2(id zinc.EntityID, data LocalRotation2Data) {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded(LocalRotation2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component LocalRotation2 already exists on entity %d", id))
	}
}

// UpdateLocalRotation2 ...
func (c *LocalRotation2Component) UpdateLocalRotation2(id zinc.EntityID, data LocalRotation2Data, silent bool) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated(LocalRotation2Key, id)
		}
	} else {
		panic(fmt.Sprintf("zinc: component LocalRotation2 does not exist on entity %d", id))
	}
}

// HasEntity ...
func (c *LocalRotation2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalRotation2 ...
func (c *LocalRotation2Component) LocalRotation2(id zinc.EntityID) LocalRotation2Data {
	if data, ok := c.data[id]; ok {
		return data
	}
	panic(fmt.Sprintf("zinc: component LocalRotation2 does not exist on entity %d", id))
}

// DeleteEntity ...
func (c *LocalRotation2Component) DeleteEntity(id zinc.EntityID) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(LocalRotation2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component LocalRotation2 does not exist on entity %d", id))
	}
}

// AddLocalRotation2X ...
func AddLocalRotation2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalRotation2Data) {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	c.AddLocalRotation2(id, data)
}

// AddLocalRotation2 ...
func AddLocalRotation2(id zinc.EntityID, data LocalRotation2Data) {
	AddLocalRotation2X(zinc.Default(), id, data)
}

// UpdateLocalRotation2X ...
func UpdateLocalRotation2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data LocalRotation2Data) {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	c.UpdateLocalRotation2(id, data, true)
}

// UpdateLocalRotation2Silently ...
func UpdateLocalRotation2Silently(id zinc.EntityID, data LocalRotation2Data) {
	UpdateLocalRotation2SilentlyX(zinc.Default(), id, data)
}

// UpdateLocalRotation2X ...
func UpdateLocalRotation2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalRotation2Data) {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	c.UpdateLocalRotation2(id, data, false)
}

// UpdateLocalRotation2 ...
func UpdateLocalRotation2(id zinc.EntityID, data LocalRotation2Data) {
	UpdateLocalRotation2X(zinc.Default(), id, data)
}

// HasLocalRotation2X ...
func HasLocalRotation2X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(LocalRotation2Key)
	return v.HasEntity(id)
}

// HasLocalRotation2 ...
func HasLocalRotation2(id zinc.EntityID) bool {
	return HasLocalRotation2X(zinc.Default(), id)
}

// LocalRotation2X ...
func LocalRotation2X(e *zinc.ZEntityManager, id zinc.EntityID) LocalRotation2Data {
	v, _ := e.Component(LocalRotation2Key)
	c := v.(*LocalRotation2Component)
	return c.LocalRotation2(id)
}

// LocalRotation2 ...
func LocalRotation2(id zinc.EntityID) LocalRotation2Data {
	return LocalRotation2X(zinc.Default(), id)
}

// DeleteLocalRotation2X ...
func DeleteLocalRotation2X(e *zinc.ZEntityManager, id zinc.EntityID) {
	v, _ := e.Component(LocalRotation2Key)
	v.DeleteEntity(id)
}

// DeleteLocalRotation2 ...
func DeleteLocalRotation2(id zinc.EntityID) {
	DeleteLocalRotation2X(zinc.Default(), id)
}