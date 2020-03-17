package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZLocalPosition2 ...
var ZLocalPosition2 uint = uint(1950915755)

// ZLocalPosition2Data ...
type ZLocalPosition2Data struct {
	X	float32
	Y	float32
	
}

// LocalPosition2Component ...
type LocalPosition2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]ZLocalPosition2Data
}

// RegisterLocalPosition2ComponentWith ...
func RegisterLocalPosition2ComponentWith(e *zinc.ZEntityManager) {
	x := NewLocalPosition2Component()
	ctx := e.RegisterComponent(ZLocalPosition2, x)
	x.SetContext(ctx)
}

// RegisterLocalPosition2Component ...
func RegisterLocalPosition2Component() {
	x := NewLocalPosition2Component()
	ctx := zinc.Default().RegisterComponent(ZLocalPosition2, x)
	x.SetContext(ctx)
}

// NewLocalPosition2Component ...
func NewLocalPosition2Component() *LocalPosition2Component {
	return &LocalPosition2Component{data: make(map[zinc.ZEntityID]ZLocalPosition2Data)}
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
func (c *LocalPosition2Component) AddLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = data
			c.ctx.ComponentAdded(ZLocalPosition2, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// UpdateLocalPosition2 ...
func (c *LocalPosition2Component) UpdateLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data, silent bool) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = data
			if !silent {
				c.ctx.ComponentUpdated(ZLocalPosition2, id)
			}
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *LocalPosition2Component) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalPosition2 ...
func (c *LocalPosition2Component) LocalPosition2(id zinc.ZEntityID) (ZLocalPosition2Data, error) {
	data, ok := c.data[id]
	if c.ctx.HasEntity(id) {
		if ok {
			return data, nil
		}
		return data, zinc.ErrEntityComponentNotFound
	}
	return data, zinc.ErrEntityNotFound
}

// DeleteEntity ...
func (c *LocalPosition2Component) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(ZLocalPosition2, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	} 
	return zinc.ErrEntityNotFound
}

// AddLocalPosition2X ...
func AddLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) error {
	v := e.Component(ZLocalPosition2)
	c := v.(*LocalPosition2Component)
	return c.AddLocalPosition2(id, data)
}


// MustAddLocalPosition2X ...
func MustAddLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := AddLocalPosition2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// AddLocalPosition2 ...
func AddLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) error {
	return AddLocalPosition2X(zinc.Default(), id, data)
}


// MustAddLocalPosition2 ...
func MustAddLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := AddLocalPosition2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2SilentlyX ...
func UpdateLocalPosition2SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) error {
	v := e.Component(ZLocalPosition2)
	c := v.(*LocalPosition2Component)
	return c.UpdateLocalPosition2(id, data, true)
}


// MustUpdateLocalPosition2SilentlyX ...
func MustUpdateLocalPosition2SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := UpdateLocalPosition2SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2Silently ...
func UpdateLocalPosition2Silently(id zinc.ZEntityID, data ZLocalPosition2Data) error {
	return UpdateLocalPosition2SilentlyX(zinc.Default(), id, data)
}


// MustUpdateLocalPosition2Silently ...
func MustUpdateLocalPosition2Silently(id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := UpdateLocalPosition2SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2X ...
func UpdateLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) error {
	v := e.Component(ZLocalPosition2)
	c := v.(*LocalPosition2Component)
	return c.UpdateLocalPosition2(id, data, false)
}


// MustUpdateLocalPosition2X ...
func MustUpdateLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := UpdateLocalPosition2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalPosition2 ...
func UpdateLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) error {
	return UpdateLocalPosition2X(zinc.Default(), id, data)
}


// MustUpdateLocalPosition2 ...
func MustUpdateLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := UpdateLocalPosition2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// SetLocalPosition2X ...
func SetLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) error {
	v := e.Component(ZLocalPosition2)
	c := v.(*LocalPosition2Component)
	if c.HasEntity(id) {
		return c.UpdateLocalPosition2(id, data, false)
	}
	return c.AddLocalPosition2(id, data)
}


// MustSetLocalPosition2X ...
func MustSetLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := SetLocalPosition2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// SetLocalPosition2 ...
func SetLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) error {
	return SetLocalPosition2X(zinc.Default(), id, data)
}


// MustSetLocalPosition2 ...
func MustSetLocalPosition2(id zinc.ZEntityID, data ZLocalPosition2Data) {
	err := SetLocalPosition2(id, data)
	if err != nil {
		panic(err)
	}
}

// HasLocalPosition2X ...
func HasLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(ZLocalPosition2)
	return v.HasEntity(id)
}

// HasLocalPosition2 ...
func HasLocalPosition2(id zinc.ZEntityID) bool {
	return HasLocalPosition2X(zinc.Default(), id)
}

// LocalPosition2X ...
func LocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID) (ZLocalPosition2Data, error) {
	v := e.Component(ZLocalPosition2)
	c := v.(*LocalPosition2Component)
	return c.LocalPosition2(id)
}


// MustLocalPosition2X ...
func MustLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID) ZLocalPosition2Data {
	data, err := LocalPosition2X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// LocalPosition2 ...
func LocalPosition2(id zinc.ZEntityID) (ZLocalPosition2Data, error) {
	return LocalPosition2X(zinc.Default(), id)
}


// MustLocalPosition2 ...
func MustLocalPosition2(id zinc.ZEntityID) ZLocalPosition2Data {
	data, err := LocalPosition2X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteLocalPosition2X ...
func DeleteLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(ZLocalPosition2)
	return v.DeleteEntity(id)
}


// MustDeleteLocalPosition2X ...
func MustDeleteLocalPosition2X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := DeleteLocalPosition2X(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteLocalPosition2 ...
func DeleteLocalPosition2(id zinc.ZEntityID) error {
	return DeleteLocalPosition2X(zinc.Default(), id)
}


// MustDeleteLocalPosition2 ...
func MustDeleteLocalPosition2(id zinc.ZEntityID) {
	err := DeleteLocalPosition2(id)
	if err != nil {
		panic(err)
	}
}