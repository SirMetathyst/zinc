package kit

import (
	"github.com/SirMetathyst/zinc"
	"fmt"
)

// LocalPosition2Key ...
var LocalPosition2Key uint = uint(1950915755)

//LocalPosition2Data ...
type LocalPosition2Data struct {
	X	float32
	Y	float32
	
}

// LocalPosition2Component ...
type LocalPosition2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]LocalPosition2Data
}

// RegisterLocalPosition2ComponentWith ...
func RegisterLocalPosition2ComponentWith(e *zinc.ZEntityManager) {
	x := NewLocalPosition2Component()
	ctx := e.RegisterComponent(LocalPosition2Key, x)
	x.SetContext(ctx)
}

// RegisterLocalPosition2Component ...
func RegisterLocalPosition2Component() {
	x := NewLocalPosition2Component()
	ctx := zinc.Default().RegisterComponent(LocalPosition2Key, x)
	x.SetContext(ctx)
}

// NewLocalPosition2Component ...
func NewLocalPosition2Component() *LocalPosition2Component {
	return &LocalPosition2Component{data: make(map[zinc.EntityID]LocalPosition2Data)}
}

func init() {
	RegisterLocalPosition2Component()
}

// SetContext ...
func (c *LocalPosition2Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// AddLocalPosition2 ...
func (c *LocalPosition2Component) AddLocalPosition2(id zinc.EntityID, data LocalPosition2Data) {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded(LocalPosition2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component LocalPosition2 already exists on entity %d", id))
	}
}

// UpdateLocalPosition2 ...
func (c *LocalPosition2Component) UpdateLocalPosition2(id zinc.EntityID, data LocalPosition2Data, silent bool) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated(LocalPosition2Key, id)
		}
	} else {
		panic(fmt.Sprintf("zinc: component LocalPosition2 does not exist on entity %d", id))
	}
}

// HasEntity ...
func (c *LocalPosition2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalPosition2 ...
func (c *LocalPosition2Component) LocalPosition2(id zinc.EntityID) LocalPosition2Data {
	if data, ok := c.data[id]; ok {
		return data
	}
	panic(fmt.Sprintf("zinc: component LocalPosition2 does not exist on entity %d", id))
}

// DeleteEntity ...
func (c *LocalPosition2Component) DeleteEntity(id zinc.EntityID) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(LocalPosition2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component LocalPosition2 does not exist on entity %d", id))
	}
}

// AddLocalPosition2X ...
func AddLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	c.AddLocalPosition2(id, data)
}

// AddLocalPosition2 ...
func AddLocalPosition2(id zinc.EntityID, data LocalPosition2Data) {
	AddLocalPosition2X(zinc.Default(), id, data)
}

// UpdateLocalPosition2X ...
func UpdateLocalPosition2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	c.UpdateLocalPosition2(id, data, true)
}

// UpdateLocalPosition2Silently ...
func UpdateLocalPosition2Silently(id zinc.EntityID, data LocalPosition2Data) {
	UpdateLocalPosition2SilentlyX(zinc.Default(), id, data)
}

// UpdateLocalPosition2X ...
func UpdateLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	c.UpdateLocalPosition2(id, data, false)
}

// UpdateLocalPosition2 ...
func UpdateLocalPosition2(id zinc.EntityID, data LocalPosition2Data) {
	UpdateLocalPosition2X(zinc.Default(), id, data)
}

// HasLocalPosition2X ...
func HasLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(LocalPosition2Key)
	return v.HasEntity(id)
}

// HasLocalPosition2 ...
func HasLocalPosition2(id zinc.EntityID) bool {
	return HasLocalPosition2X(zinc.Default(), id)
}

// LocalPosition2X ...
func LocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) LocalPosition2Data {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	return c.LocalPosition2(id)
}

// LocalPosition2 ...
func LocalPosition2(id zinc.EntityID) LocalPosition2Data {
	return LocalPosition2X(zinc.Default(), id)
}

// DeleteLocalPosition2X ...
func DeleteLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) {
	v, _ := e.Component(LocalPosition2Key)
	v.DeleteEntity(id)
}

// DeleteLocalPosition2 ...
func DeleteLocalPosition2(id zinc.EntityID) {
	DeleteLocalPosition2X(zinc.Default(), id)
}