package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZIsPlaying ...
var ZIsPlaying uint = uint(4040087141)

// IsPlayingComponent ...
type IsPlayingComponent struct {
	ctx  *zinc.ZContext
	data map[zinc.ZEntityID]bool
}

// RegisterIsPlayingComponentWith ...
func RegisterIsPlayingComponentWith(e *zinc.ZEntityManager) {
	x := NewIsPlayingComponent()
	ctx := e.RegisterComponent(ZIsPlaying, x)
	x.SetContext(ctx)
}

// RegisterIsPlayingComponent ...
func RegisterIsPlayingComponent() {
	x := NewIsPlayingComponent()
	ctx := zinc.Default().RegisterComponent(ZIsPlaying, x)
	x.SetContext(ctx)
}

// NewIsPlayingComponent ...
func NewIsPlayingComponent() *IsPlayingComponent {
	return &IsPlayingComponent{data: make(map[zinc.ZEntityID]bool)}
}

func init() {
	RegisterIsPlayingComponent()
}

// SetContext ...
func (c *IsPlayingComponent) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// AddIsPlaying ...
func (c *IsPlayingComponent) AddIsPlaying(id zinc.ZEntityID, value bool) error {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = value
		c.ctx.ComponentAdded(ZIsPlaying, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// UpdateIsPlaying ...
func (c *IsPlayingComponent) UpdateIsPlaying(id zinc.ZEntityID, value bool, silent bool) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = value
		if !silent {
			c.ctx.ComponentUpdated(ZIsPlaying, id)
		}
		return nil
	}
	return zinc.ErrComponentNotFound
}

// HasEntity ...
func (c *IsPlayingComponent) HasEntity(id zinc.ZEntityID) bool {
	_, ok := c.data[id]
	return ok
}

// IsPlaying ...
func (c *IsPlayingComponent) IsPlaying(id zinc.ZEntityID) (bool, error) {
	data, ok := c.data[id]
	if ok {
		return data, nil
	}
	return data, zinc.ErrComponentNotFound
}

// DeleteEntity ...
func (c *IsPlayingComponent) DeleteEntity(id zinc.ZEntityID) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(ZIsPlaying, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// AddIsPlayingX ...
func AddIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID, value bool) error {
	v := e.Component(ZIsPlaying)
	c := v.(*IsPlayingComponent)
	return c.AddIsPlaying(id, value)
}

// MustAddIsPlayingX ...
func MustAddIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID, value bool) {
	err := AddIsPlayingX(e, id, value)
	if err != nil {
		panic(err)
	}
}

// AddIsPlaying ...
func AddIsPlaying(id zinc.ZEntityID, value bool) error {
	return AddIsPlayingX(zinc.Default(), id, value)
}

// MustAddIsPlaying ...
func MustAddIsPlaying(id zinc.ZEntityID, value bool) {
	err := AddIsPlayingX(zinc.Default(), id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlayingSilentlyX ...
func UpdateIsPlayingSilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, value bool) error {
	v := e.Component(ZIsPlaying)
	c := v.(*IsPlayingComponent)
	return c.UpdateIsPlaying(id, value, true)
}

// MustUpdateIsPlayingSilentlyX ...
func MustUpdateIsPlayingSilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, value bool) {
	err := UpdateIsPlayingSilentlyX(e, id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlayingSilently ...
func UpdateIsPlayingSilently(id zinc.ZEntityID, value bool) error {
	return UpdateIsPlayingSilentlyX(zinc.Default(), id, value)
}

// MustUpdateIsPlayingSilently ...
func MustUpdateIsPlayingSilently(id zinc.ZEntityID, value bool) {
	err := UpdateIsPlayingSilentlyX(zinc.Default(), id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlayingX ...
func UpdateIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID, value bool) error {
	v := e.Component(ZIsPlaying)
	c := v.(*IsPlayingComponent)
	return c.UpdateIsPlaying(id, value, false)
}

// MustUpdateIsPlayingX ...
func MustUpdateIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID, value bool) {
	err := UpdateIsPlayingX(e, id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlaying ...
func UpdateIsPlaying(id zinc.ZEntityID, value bool) error {
	return UpdateIsPlayingX(zinc.Default(), id, value)
}

// MustUpdateIsPlaying ...
func MustUpdateIsPlaying(id zinc.ZEntityID, value bool) {
	err := UpdateIsPlayingX(zinc.Default(), id, value)
	if err != nil {
		panic(err)
	}
}

// HasIsPlayingX ...
func HasIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	v := e.Component(ZIsPlaying)
	return v.HasEntity(id)
}

// HasIsPlaying ...
func HasIsPlaying(id zinc.ZEntityID) bool {
	return HasIsPlayingX(zinc.Default(), id)
}

// IsPlayingX ...
func IsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID) (bool, error) {
	v := e.Component(ZIsPlaying)
	c := v.(*IsPlayingComponent)
	return c.IsPlaying(id)
}

// MustIsPlayingX ...
func MustIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID) bool {
	data, err := IsPlayingX(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// IsPlaying ...
func IsPlaying(id zinc.ZEntityID) (bool, error) {
	return IsPlayingX(zinc.Default(), id)
}

// MustIsPlaying ...
func MustIsPlaying(id zinc.ZEntityID) bool {
	data, err := IsPlayingX(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteIsPlayingX ...
func DeleteIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID) error {
	v := e.Component(ZIsPlaying)
	return v.DeleteEntity(id)
}

// MustDeleteIsPlayingX ...
func MustDeleteIsPlayingX(e *zinc.ZEntityManager, id zinc.ZEntityID) {
	err := DeleteIsPlayingX(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteIsPlaying ...
func DeleteIsPlaying(id zinc.ZEntityID) error {
	return DeleteIsPlayingX(zinc.Default(), id)
}

// MustDeleteIsPlaying ...
func MustDeleteIsPlaying(id zinc.ZEntityID) {
	err := DeleteIsPlaying(id)
	if err != nil {
		panic(err)
	}
}
