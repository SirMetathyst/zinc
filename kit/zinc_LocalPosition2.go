package kit

import (
	"github.com/SirMetathyst/zinc"
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
func (c *LocalPosition2Component) AddLocalPosition2(id zinc.EntityID, data LocalPosition2Data) error {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = data
		c.ctx.ComponentAdded(LocalPosition2Key, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// UpdateLocalPosition2 ...
func (c *LocalPosition2Component) UpdateLocalPosition2(id zinc.EntityID, data LocalPosition2Data, silent bool) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = data
		if !silent {
			c.ctx.ComponentUpdated(LocalPosition2Key, id)
		}
		return nil
	}
	return zinc.ErrComponentNotFound
}

// HasEntity ...
func (c *LocalPosition2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalPosition2 ...
func (c *LocalPosition2Component) LocalPosition2(id zinc.EntityID) (LocalPosition2Data, error) {
	data, ok := c.data[id]
	if ok {
		return data, nil
	}
	return data, zinc.ErrComponentNotFound
}

// DeleteEntity ...
func (c *LocalPosition2Component) DeleteEntity(id zinc.EntityID) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(LocalPosition2Key, id)
		return nil
	} 
	return zinc.ErrComponentNotFound
}

// AddLocalPosition2X ...
func AddLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) error {
	v := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	return c.AddLocalPosition2(id, data)
}

// MustAddLocalPosition2X ...
func MustAddLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) {
	err := AddLocalPosition2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// AddLocalPosition2 ...
func AddLocalPosition2(id zinc.EntityID, data LocalPosition2Data) error {
	return AddLocalPosition2X(zinc.Default(), id, data)
}

// MustAddLocalPosition2 ...
func MustAddLocalPosition2(id zinc.EntityID, data LocalPosition2Data) {
	err := AddLocalPosition2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2SilentlyX ...
func UpdateLocalPosition2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) error {
	v := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	return c.UpdateLocalPosition2(id, data, true)
}

// MustUpdateLocalPosition2SilentlyX ...
func MustUpdateLocalPosition2SilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) {
	err := UpdateLocalPosition2SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2Silently ...
func UpdateLocalPosition2Silently(id zinc.EntityID, data LocalPosition2Data) error {
	return UpdateLocalPosition2SilentlyX(zinc.Default(), id, data)
}

// MustUpdateLocalPosition2Silently ...
func MustUpdateLocalPosition2Silently(id zinc.EntityID, data LocalPosition2Data) {
	err := UpdateLocalPosition2SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2X ...
func UpdateLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) error {
	v := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	return c.UpdateLocalPosition2(id, data, false)
}

// MustUpdateLocalPosition2X ...
func MustUpdateLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID, data LocalPosition2Data) {
	err := UpdateLocalPosition2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2 ...
func UpdateLocalPosition2(id zinc.EntityID, data LocalPosition2Data) error {
	return UpdateLocalPosition2X(zinc.Default(), id, data)
}

// MustUpdateLocalPosition2 ...
func MustUpdateLocalPosition2(id zinc.EntityID, data LocalPosition2Data) {
	err := UpdateLocalPosition2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// HasLocalPosition2X ...
func HasLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v := e.Component(LocalPosition2Key)
	return v.HasEntity(id)
}

// HasLocalPosition2 ...
func HasLocalPosition2(id zinc.EntityID) bool {
	return HasLocalPosition2X(zinc.Default(), id)
}

// LocalPosition2X ...
func LocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) (LocalPosition2Data, error) {
	v := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	return c.LocalPosition2(id)
}

// MustLocalPosition2X ...
func MustLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) LocalPosition2Data {
	data, err := LocalPosition2X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// LocalPosition2 ...
func LocalPosition2(id zinc.EntityID) (LocalPosition2Data, error) {
	return LocalPosition2X(zinc.Default(), id)
}

// MustLocalPosition2 ...
func MustLocalPosition2(id zinc.EntityID) LocalPosition2Data {
	data, err := LocalPosition2X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteLocalPosition2X ...
func DeleteLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) error {
	v := e.Component(LocalPosition2Key)
	return v.DeleteEntity(id)
}

// MustDeleteLocalPosition2X ...
func MustDeleteLocalPosition2X(e *zinc.ZEntityManager, id zinc.EntityID) {
	err := DeleteLocalPosition2X(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteLocalPosition2 ...
func DeleteLocalPosition2(id zinc.EntityID) error {
	return DeleteLocalPosition2X(zinc.Default(), id)
}

// MustDeleteLocalPosition2 ...
func MustDeleteLocalPosition2(id zinc.EntityID) {
	err := DeleteLocalPosition2(id)
	if err != nil {
		panic(err)
	}
}