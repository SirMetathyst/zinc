package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZVelocity2 ...
var ZVelocity2 uint = uint(1825051648)

// ZVelocity2Data ...
type ZVelocity2Data struct {
	X	float32
	Y	float32
	
}

// Velocity2Component ...
type Velocity2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]ZVelocity2Data
}

// RegisterVelocity2ComponentWith ...
func RegisterVelocity2ComponentWith(e *zinc.ZEntityManager) {
	x := NewVelocity2Component()
	ctx := e.RegisterComponent(ZVelocity2, x)
	x.SetContext(ctx)
}

// RegisterVelocity2Component ...
func RegisterVelocity2Component() {
	x := NewVelocity2Component()
	ctx := zinc.Default().RegisterComponent(ZVelocity2, x)
	x.SetContext(ctx)
}

// NewVelocity2Component ...
func NewVelocity2Component() *Velocity2Component {
	return &Velocity2Component{data: make(map[zinc.EntityID]ZVelocity2Data)}
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
func (c *Velocity2Component) AddVelocity2(id zinc.EntityID, data ZVelocity2Data) error {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded(ZVelocity2, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// UpdateVelocity2 ...
func (c *Velocity2Component) UpdateVelocity2(id zinc.EntityID, data ZVelocity2Data, silent bool) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated(ZVelocity2, id)
		}
		return nil
	}
	return zinc.ErrComponentNotFound
}

// HasEntity ...
func (c *Velocity2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// Velocity2 ...
func (c *Velocity2Component) Velocity2(id zinc.EntityID) (ZVelocity2Data, error) {
	data, ok := c.data[id]
	if ok {
		return data, nil
	}
	return data, zinc.ErrComponentNotFound
}

// DeleteEntity ...
func (c *Velocity2Component) DeleteEntity(id zinc.EntityID) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(ZVelocity2, id)
		return nil
	} 
	return zinc.ErrComponentNotFound
}

// AddVelocity2X ...
func AddVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID, data ZVelocity2Data) error {
	v := e.Component(ZVelocity2)
	c := v.(*Velocity2Component)
	return c.AddVelocity2(id, data)
}

// MustAddVelocity2X ...
func MustAddVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID, data ZVelocity2Data) {
	err := AddVelocity2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// AddVelocity2 ...
func AddVelocity2(id zinc.EntityID, data ZVelocity2Data) error {
	return AddVelocity2X(zinc.Default(), id, data)
}

// MustAddVelocity2 ...
func MustAddVelocity2(id zinc.EntityID, data ZVelocity2Data) {
	err := AddVelocity2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateVelocity2SilentlyX ...
func UpdateVelocity2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data ZVelocity2Data) error {
	v := e.Component(ZVelocity2)
	c := v.(*Velocity2Component)
	return c.UpdateVelocity2(id, data, true)
}

// MustUpdateVelocity2SilentlyX ...
func MustUpdateVelocity2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data ZVelocity2Data) {
	err := UpdateVelocity2SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateVelocity2Silently ...
func UpdateVelocity2Silently(id zinc.EntityID, data ZVelocity2Data) error {
	return UpdateVelocity2SilentlyX(zinc.Default(), id, data)
}

// MustUpdateVelocity2Silently ...
func MustUpdateVelocity2Silently(id zinc.EntityID, data ZVelocity2Data) {
	err := UpdateVelocity2SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateVelocity2X ...
func UpdateVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID, data ZVelocity2Data) error {
	v := e.Component(ZVelocity2)
	c := v.(*Velocity2Component)
	return c.UpdateVelocity2(id, data, false)
}

// MustUpdateVelocity2X ...
func MustUpdateVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID, data ZVelocity2Data) {
	err := UpdateVelocity2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateVelocity2 ...
func UpdateVelocity2(id zinc.EntityID, data ZVelocity2Data) error {
	return UpdateVelocity2X(zinc.Default(), id, data)
}

// MustUpdateVelocity2 ...
func MustUpdateVelocity2(id zinc.EntityID, data ZVelocity2Data) {
	err := UpdateVelocity2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// HasVelocity2X ...
func HasVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v := e.Component(ZVelocity2)
	return v.HasEntity(id)
}

// HasVelocity2 ...
func HasVelocity2(id zinc.EntityID) bool {
	return HasVelocity2X(zinc.Default(), id)
}

// Velocity2X ...
func Velocity2X(e *zinc.ZEntityManager, id zinc.EntityID) (ZVelocity2Data, error) {
	v := e.Component(ZVelocity2)
	c := v.(*Velocity2Component)
	return c.Velocity2(id)
}

// MustVelocity2X ...
func MustVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID) ZVelocity2Data {
	data, err := Velocity2X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// Velocity2 ...
func Velocity2(id zinc.EntityID) (ZVelocity2Data, error) {
	return Velocity2X(zinc.Default(), id)
}

// MustVelocity2 ...
func MustVelocity2(id zinc.EntityID) ZVelocity2Data {
	data, err := Velocity2X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteVelocity2X ...
func DeleteVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID) error {
	v := e.Component(ZVelocity2)
	return v.DeleteEntity(id)
}

// MustDeleteVelocity2X ...
func MustDeleteVelocity2X(e *zinc.ZEntityManager, id zinc.EntityID) {
	err := DeleteVelocity2X(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteVelocity2 ...
func DeleteVelocity2(id zinc.EntityID) error {
	return DeleteVelocity2X(zinc.Default(), id)
}

// MustDeleteVelocity2 ...
func MustDeleteVelocity2(id zinc.EntityID) {
	err := DeleteVelocity2(id)
	if err != nil {
		panic(err)
	}
}