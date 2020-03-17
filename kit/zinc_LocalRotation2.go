package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZLocalRotation2 ...
var ZLocalRotation2 uint = uint(4108562484)

// ZLocalRotation2Data ...
type ZLocalRotation2Data struct {
	X	float32
	Y	float32
	
}

// LocalRotation2Component ...
type LocalRotation2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]ZLocalRotation2Data
}

// RegisterLocalRotation2ComponentWith ...
func RegisterLocalRotation2ComponentWith(e *zinc.ZEntityManager) {
	x := NewLocalRotation2Component()
	ctx := e.RegisterComponent(ZLocalRotation2, x)
	x.SetContext(ctx)
}

// RegisterLocalRotation2Component ...
func RegisterLocalRotation2Component() {
	x := NewLocalRotation2Component()
	ctx := zinc.Default().RegisterComponent(ZLocalRotation2, x)
	x.SetContext(ctx)
}

// NewLocalRotation2Component ...
func NewLocalRotation2Component() *LocalRotation2Component {
	return &LocalRotation2Component{data: make(map[zinc.ZEntityID]ZLocalRotation2Data)}
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
func (c *LocalRotation2Component) AddLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = data
			c.ctx.ComponentAdded(ZLocalRotation2, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// UpdateLocalRotation2 ...
func (c *LocalRotation2Component) UpdateLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data, silent bool) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = data
			if !silent {
				c.ctx.ComponentUpdated(ZLocalRotation2, id)
			}
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *LocalRotation2Component) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalRotation2 ...
func (c *LocalRotation2Component) LocalRotation2(id zinc.ZEntityID) (ZLocalRotation2Data, error) {
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
func (c *LocalRotation2Component) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(ZLocalRotation2, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	} 
	return zinc.ErrEntityNotFound
}

// AddLocalRotation2X ...
func AddLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) error {
	v := e.Component(ZLocalRotation2)
	c := v.(*LocalRotation2Component)
	return c.AddLocalRotation2(id, data)
}


// MustAddLocalRotation2X ...
func MustAddLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := AddLocalRotation2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// AddLocalRotation2 ...
func AddLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) error {
	return AddLocalRotation2X(zinc.Default(), id, data)
}


// MustAddLocalRotation2 ...
func MustAddLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := AddLocalRotation2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalRotation2SilentlyX ...
func UpdateLocalRotation2SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) error {
	v := e.Component(ZLocalRotation2)
	c := v.(*LocalRotation2Component)
	return c.UpdateLocalRotation2(id, data, true)
}


// MustUpdateLocalRotation2SilentlyX ...
func MustUpdateLocalRotation2SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := UpdateLocalRotation2SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalRotation2Silently ...
func UpdateLocalRotation2Silently(id zinc.ZEntityID, data ZLocalRotation2Data) error {
	return UpdateLocalRotation2SilentlyX(zinc.Default(), id, data)
}


// MustUpdateLocalRotation2Silently ...
func MustUpdateLocalRotation2Silently(id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := UpdateLocalRotation2SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalRotation2X ...
func UpdateLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) error {
	v := e.Component(ZLocalRotation2)
	c := v.(*LocalRotation2Component)
	return c.UpdateLocalRotation2(id, data, false)
}


// MustUpdateLocalRotation2X ...
func MustUpdateLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := UpdateLocalRotation2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalRotation2 ...
func UpdateLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) error {
	return UpdateLocalRotation2X(zinc.Default(), id, data)
}


// MustUpdateLocalRotation2 ...
func MustUpdateLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := UpdateLocalRotation2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// SetLocalRotation2X ...
func SetLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) error {
	v := e.Component(ZLocalRotation2)
	c := v.(*LocalRotation2Component)
	if c.HasEntity(id) {
		return c.UpdateLocalRotation2(id, data, false)
	}
	return c.AddLocalRotation2(id, data)
}


// MustSetLocalRotation2X ...
func MustSetLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := SetLocalRotation2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// SetLocalRotation2 ...
func SetLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) error {
	return SetLocalRotation2X(zinc.Default(), id, data)
}


// MustSetLocalRotation2 ...
func MustSetLocalRotation2(id zinc.ZEntityID, data ZLocalRotation2Data) {
	err := SetLocalRotation2(id, data)
	if err != nil {
		panic(err)
	}
}

// HasLocalRotation2X ...
func HasLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(ZLocalRotation2)
	return v.HasEntity(id)
}

// HasLocalRotation2 ...
func HasLocalRotation2(id zinc.ZEntityID) bool {
	return HasLocalRotation2X(zinc.Default(), id)
}

// LocalRotation2X ...
func LocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID) (ZLocalRotation2Data, error) {
	v := e.Component(ZLocalRotation2)
	c := v.(*LocalRotation2Component)
	return c.LocalRotation2(id)
}


// MustLocalRotation2X ...
func MustLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID) ZLocalRotation2Data {
	data, err := LocalRotation2X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// LocalRotation2 ...
func LocalRotation2(id zinc.ZEntityID) (ZLocalRotation2Data, error) {
	return LocalRotation2X(zinc.Default(), id)
}


// MustLocalRotation2 ...
func MustLocalRotation2(id zinc.ZEntityID) ZLocalRotation2Data {
	data, err := LocalRotation2X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteLocalRotation2X ...
func DeleteLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(ZLocalRotation2)
	return v.DeleteEntity(id)
}


// MustDeleteLocalRotation2X ...
func MustDeleteLocalRotation2X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := DeleteLocalRotation2X(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteLocalRotation2 ...
func DeleteLocalRotation2(id zinc.ZEntityID) error {
	return DeleteLocalRotation2X(zinc.Default(), id)
}


// MustDeleteLocalRotation2 ...
func MustDeleteLocalRotation2(id zinc.ZEntityID) {
	err := DeleteLocalRotation2(id)
	if err != nil {
		panic(err)
	}
}