package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZLocalScale2 ...
var ZLocalScale2 uint = uint(54082440)

// ZLocalScale2Data ...
type ZLocalScale2Data struct {
	X	float32
	Y	float32
	
}

// LocalScale2Component ...
type LocalScale2Component struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]ZLocalScale2Data
}

// RegisterLocalScale2ComponentWith ...
func RegisterLocalScale2ComponentWith(e *zinc.ZEntityManager) {
	x := NewLocalScale2Component()
	ctx := e.RegisterComponent(ZLocalScale2, x)
	x.SetContext(ctx)
}

// RegisterLocalScale2Component ...
func RegisterLocalScale2Component() {
	x := NewLocalScale2Component()
	ctx := zinc.Default().RegisterComponent(ZLocalScale2, x)
	x.SetContext(ctx)
}

// NewLocalScale2Component ...
func NewLocalScale2Component() *LocalScale2Component {
	return &LocalScale2Component{data: make(map[zinc.ZEntityID]ZLocalScale2Data)}
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
func (c *LocalScale2Component) AddLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = data
			c.ctx.ComponentAdded(ZLocalScale2, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// UpdateLocalScale2 ...
func (c *LocalScale2Component) UpdateLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data, silent bool) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = data
			if !silent {
				c.ctx.ComponentUpdated(ZLocalScale2, id)
			}
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *LocalScale2Component) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// LocalScale2 ...
func (c *LocalScale2Component) LocalScale2(id zinc.ZEntityID) (ZLocalScale2Data, error) {
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
func (c *LocalScale2Component) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(ZLocalScale2, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	} 
	return zinc.ErrEntityNotFound
}

// AddLocalScale2X ...
func AddLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) error {
	v := e.Component(ZLocalScale2)
	c := v.(*LocalScale2Component)
	return c.AddLocalScale2(id, data)
}


// MustAddLocalScale2X ...
func MustAddLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) {
	err := AddLocalScale2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// AddLocalScale2 ...
func AddLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) error {
	return AddLocalScale2X(zinc.Default(), id, data)
}


// MustAddLocalScale2 ...
func MustAddLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) {
	err := AddLocalScale2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalScale2SilentlyX ...
func UpdateLocalScale2SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) error {
	v := e.Component(ZLocalScale2)
	c := v.(*LocalScale2Component)
	return c.UpdateLocalScale2(id, data, true)
}


// MustUpdateLocalScale2SilentlyX ...
func MustUpdateLocalScale2SilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) {
	err := UpdateLocalScale2SilentlyX(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalScale2Silently ...
func UpdateLocalScale2Silently(id zinc.ZEntityID, data ZLocalScale2Data) error {
	return UpdateLocalScale2SilentlyX(zinc.Default(), id, data)
}


// MustUpdateLocalScale2Silently ...
func MustUpdateLocalScale2Silently(id zinc.ZEntityID, data ZLocalScale2Data) {
	err := UpdateLocalScale2SilentlyX(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalScale2X ...
func UpdateLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) error {
	v := e.Component(ZLocalScale2)
	c := v.(*LocalScale2Component)
	return c.UpdateLocalScale2(id, data, false)
}


// MustUpdateLocalScale2X ...
func MustUpdateLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) {
	err := UpdateLocalScale2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// UpdateLocalScale2 ...
func UpdateLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) error {
	return UpdateLocalScale2X(zinc.Default(), id, data)
}


// MustUpdateLocalScale2 ...
func MustUpdateLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) {
	err := UpdateLocalScale2X(zinc.Default(), id, data)
	if err != nil {
		panic(err)
	}
}

// SetLocalScale2X ...
func SetLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) error {
	v := e.Component(ZLocalScale2)
	c := v.(*LocalScale2Component)
	if c.HasEntity(id) {
		return c.UpdateLocalScale2(id, data, false)
	}
	return c.AddLocalScale2(id, data)
}


// MustSetLocalScale2X ...
func MustSetLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZLocalScale2Data) {
	err := SetLocalScale2X(e, id, data)
	if err != nil {
		panic(err)
	}
}

// SetLocalScale2 ...
func SetLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) error {
	return SetLocalScale2X(zinc.Default(), id, data)
}


// MustSetLocalScale2 ...
func MustSetLocalScale2(id zinc.ZEntityID, data ZLocalScale2Data) {
	err := SetLocalScale2(id, data)
	if err != nil {
		panic(err)
	}
}

// HasLocalScale2X ...
func HasLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(ZLocalScale2)
	return v.HasEntity(id)
}

// HasLocalScale2 ...
func HasLocalScale2(id zinc.ZEntityID) bool {
	return HasLocalScale2X(zinc.Default(), id)
}

// LocalScale2X ...
func LocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID) (ZLocalScale2Data, error) {
	v := e.Component(ZLocalScale2)
	c := v.(*LocalScale2Component)
	return c.LocalScale2(id)
}


// MustLocalScale2X ...
func MustLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID) ZLocalScale2Data {
	data, err := LocalScale2X(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// LocalScale2 ...
func LocalScale2(id zinc.ZEntityID) (ZLocalScale2Data, error) {
	return LocalScale2X(zinc.Default(), id)
}


// MustLocalScale2 ...
func MustLocalScale2(id zinc.ZEntityID) ZLocalScale2Data {
	data, err := LocalScale2X(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteLocalScale2X ...
func DeleteLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(ZLocalScale2)
	return v.DeleteEntity(id)
}


// MustDeleteLocalScale2X ...
func MustDeleteLocalScale2X(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := DeleteLocalScale2X(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteLocalScale2 ...
func DeleteLocalScale2(id zinc.ZEntityID) error {
	return DeleteLocalScale2X(zinc.Default(), id)
}


// MustDeleteLocalScale2 ...
func MustDeleteLocalScale2(id zinc.ZEntityID) {
	err := DeleteLocalScale2(id)
	if err != nil {
		panic(err)
	}
}