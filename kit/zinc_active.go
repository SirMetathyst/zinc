package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZActive ...
var ZActive uint = uint(3648362799)

// ActiveComponent ...
type ActiveComponent struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]bool
}

// RegisterActiveComponentWith ...
func RegisterActiveComponentWith(e *zinc.ZEntityManager) {
	x := NewActiveComponent()
	ctx := e.RegisterComponent(ZActive, x)
	x.SetContext(ctx)
}

// RegisterActiveComponent ...
func RegisterActiveComponent() {
	x := NewActiveComponent()
	ctx := zinc.Default().RegisterComponent(ZActive, x)
	x.SetContext(ctx)
}

// NewActiveComponent ...
func NewActiveComponent() *ActiveComponent {
	return &ActiveComponent{data: make(map[zinc.ZEntityID]bool)}
}

func init() {
	RegisterActiveComponent()
}

// SetContext ...
func (c *ActiveComponent) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// Active ...
func (c *ActiveComponent) Active(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if !c.HasEntity(id) {
			c.data[id] = true
			c.ctx.ComponentAdded(ZActive, id)
			return nil
		}
		return zinc.ErrEntityComponentAlreadyExists
	}
	return zinc.ErrEntityNotFound
}

// HasEntity ...
func (c *ActiveComponent) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// DeleteEntity ...
func (c *ActiveComponent) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			delete(c.data, id)
			c.ctx.ComponentDeleted(ZActive, id)
			return nil
		}
		return zinc.ErrEntityComponentNotFound
	}
	return zinc.ErrEntityNotFound
}

// ActiveX ...
func ActiveX(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(ZActive)
	c := v.(*ActiveComponent)
	return c.Active(id)
}


// MustActiveX ...
func MustActiveX(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := ActiveX(e, id)
	if err != nil {
		panic(err)
	}
}

// Active ...
func Active(id zinc.ZEntityID) error {
	return ActiveX(zinc.Default(), id)
}


// MustActive ...
func MustActive(id zinc.ZEntityID) {
	err := ActiveX(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
}

// IsActiveX ...
func IsActiveX(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(ZActive)
	return v.HasEntity(id)
}

// IsActive ...
func IsActive(id zinc.ZEntityID) bool {
	return IsActiveX(zinc.Default(), id)
}

// NotActiveX ...
func NotActiveX(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(ZActive)
	return v.DeleteEntity(id)
}


// MustNotActiveX ...
func MustNotActiveX(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := NotActiveX(e, id)
	if err != nil {
		panic(err)
	}
}

// NotActive ...
func NotActive(id zinc.ZEntityID) error {
	return NotActiveX(zinc.Default(), id)
}


// MustNotActive ...
func MustNotActive(id zinc.ZEntityID) {
	err := NotActive(id)
	if err != nil {
		panic(err)
	}
}