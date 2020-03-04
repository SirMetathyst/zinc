package kit

import (
	"github.com/SirMetathyst/zinc"
	"fmt"
)

// Velocity2Key ...
var Velocity2Key uint = uint(1825051648)

//Velocity2Data ...
type Velocity2Data struct {
	X	float32
	Y	float32
	
}

// Velocity2Component ...
type Velocity2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]Velocity2Data
}

// RegisterVelocity2ComponentWith ...
func RegisterVelocity2ComponentWith(e *zinc.ZEntityManager) {
	x := NewVelocity2Component()
	ctx := e.RegisterComponent(Velocity2Key, x)
	x.SetContext(ctx)
}

// RegisterVelocity2Component ...
func RegisterVelocity2Component() {
	x := NewVelocity2Component()
	ctx := zinc.Default().RegisterComponent(Velocity2Key, x)
	x.SetContext(ctx)
}

// NewVelocity2Component ...
func NewVelocity2Component() *Velocity2Component {
	return &Velocity2Component{data: make(map[zinc.EntityID]Velocity2Data)}
}

func init() {
	RegisterVelocity2Component()
}

// SetContext ...
func (c *Velocity2Component) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// AddVelocity2 ...
func (c *Velocity2Component) AddVelocity2(id zinc.EntityID, data Velocity2Data) {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded(Velocity2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component Velocity2 already exists on entity %d", id))
	}
}

// UpdateVelocity2 ...
func (c *Velocity2Component) UpdateVelocity2(id zinc.EntityID, data Velocity2Data, silent bool) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated(Velocity2Key, id)
		}
	} else {
		panic(fmt.Sprintf("zinc: component Velocity2 does not exist on entity %d", id))
	}
}

// HasEntity ...
func (c *Velocity2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// Velocity2 ...
func (c *Velocity2Component) Velocity2(id zinc.EntityID) Velocity2Data {
	if data, ok := c.data[id]; ok {
		return data
	}
	panic(fmt.Sprintf("zinc: component Velocity2 does not exist on entity %d", id))
}

// DeleteEntity ...
func (c *Velocity2Component) DeleteEntity(id zinc.EntityID) {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(Velocity2Key, id)
	} else {
		panic(fmt.Sprintf("zinc: component Velocity2 does not exist on entity %d", id))
	}
}

// AddVelocity2X ...
func AddVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID, data Velocity2Data) {
	v, _ := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	c.AddVelocity2(id, data)
}

// AddVelocity2 ...
func AddVelocity2(id zinc.EntityID, data Velocity2Data) {
	AddVelocity2X(zinc.Default(), id, data)
}

// UpdateVelocity2X ...
func UpdateVelocity2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data Velocity2Data) {
	v, _ := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	c.UpdateVelocity2(id, data, true)
}

// UpdateVelocity2Silently ...
func UpdateVelocity2Silently(id zinc.EntityID, data Velocity2Data) {
	UpdateVelocity2SilentlyX(zinc.Default(), id, data)
}

// UpdateVelocity2X ...
func UpdateVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID, data Velocity2Data) {
	v, _ := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	c.UpdateVelocity2(id, data, false)
}

// UpdateVelocity2 ...
func UpdateVelocity2(id zinc.EntityID, data Velocity2Data) {
	UpdateVelocity2X(zinc.Default(), id, data)
}

// HasVelocity2X ...
func HasVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(Velocity2Key)
	return v.HasEntity(id)
}

// HasVelocity2 ...
func HasVelocity2(id zinc.EntityID) bool {
	return HasVelocity2X(zinc.Default(), id)
}

// Velocity2X ...
func Velocity2X(e *zinc.ZEntityManager, id zinc.EntityID) Velocity2Data {
	v, _ := e.Component(Velocity2Key)
	c := v.(*Velocity2Component)
	return c.Velocity2(id)
}

// Velocity2 ...
func Velocity2(id zinc.EntityID) Velocity2Data {
	return Velocity2X(zinc.Default(), id)
}

// DeleteVelocity2X ...
func DeleteVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID) {
	v, _ := e.Component(Velocity2Key)
	v.DeleteEntity(id)
}

// DeleteVelocity2 ...
func DeleteVelocity2(id zinc.EntityID) {
	DeleteVelocity2X(zinc.Default(), id)
}