package kit

import (
	"github.com/SirMetathyst/zinc"
)

// IsPlayingKey ...
var IsPlayingKey uint = uint(4040087141)

// IsPlayingComponent ...
type IsPlayingComponent struct {
	ctx  *zinc.ZContext
	data map[zinc.EntityID]bool
}

// RegisterIsPlayingComponentWith ...
func RegisterIsPlayingComponentWith(e *zinc.ZEntityManager) {
	x := NewIsPlayingComponent()
	ctx := e.RegisterComponent(IsPlayingKey, x)
	x.SetContext(ctx)
}

// RegisterIsPlayingComponent ...
func RegisterIsPlayingComponent() {
	x := NewIsPlayingComponent()
	ctx := zinc.Default().RegisterComponent(IsPlayingKey, x)
	x.SetContext(ctx)
}

// NewIsPlayingComponent ...
func NewIsPlayingComponent() *IsPlayingComponent {
	return &IsPlayingComponent{data: make(map[zinc.EntityID]bool)}
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
func (c *IsPlayingComponent) AddIsPlaying(id zinc.EntityID, value bool) error {
	if c.ctx.HasEntity(id) && !c.HasEntity(id) {
		c.data[id] = value
		c.ctx.ComponentAdded(IsPlayingKey, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// UpdateIsPlaying ...
func (c *IsPlayingComponent) UpdateIsPlaying(id zinc.EntityID, value bool, silent bool) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		c.data[id] = value
		if !silent {
			c.ctx.ComponentUpdated(IsPlayingKey, id)
		}
		return nil
	}
	return zinc.ErrComponentNotFound
}

// HasEntity ...
func (c *IsPlayingComponent) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// IsPlaying ...
func (c *IsPlayingComponent) IsPlaying(id zinc.EntityID) (bool, error) {
	data, ok := c.data[id]
	if ok {
		return data, nil
	}
	return data, zinc.ErrComponentNotFound
}

// DeleteEntity ...
func (c *IsPlayingComponent) DeleteEntity(id zinc.EntityID) error {
	if c.ctx.HasEntity(id) && c.HasEntity(id) {
		delete(c.data, id)
		c.ctx.ComponentDeleted(IsPlayingKey, id)
		return nil
	}
	return zinc.ErrComponentNotFound
}

// AddIsPlayingX ...
func AddIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID, value bool) error {
	v := e.Component(IsPlayingKey)
	c := v.(*IsPlayingComponent)
	return c.AddIsPlaying(id, value)
}

// MustAddIsPlayingX ...
func MustAddIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID, value bool) {
	err := AddIsPlayingX(e, id, value)
	if err != nil {
		panic(err)
	}
}

// AddIsPlaying ...
func AddIsPlaying(id zinc.EntityID, value bool) error {
	return AddIsPlayingX(zinc.Default(), id, value)
}

// MustAddIsPlaying ...
func MustAddIsPlaying(id zinc.EntityID, value bool) {
	err := AddIsPlayingX(zinc.Default(), id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlayingSilentlyX ...
func UpdateIsPlayingSilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, value bool) error {
	v := e.Component(IsPlayingKey)
	c := v.(*IsPlayingComponent)
	return c.UpdateIsPlaying(id, value, true)
}

// MustUpdateIsPlayingSilentlyX ...
func MustUpdateIsPlayingSilentlyX(e *zinc.ZEntityManager, id zinc.EntityID, value bool) {
	err := UpdateIsPlayingSilentlyX(e, id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlayingSilently ...
func UpdateIsPlayingSilently(id zinc.EntityID, value bool) error {
	return UpdateIsPlayingSilentlyX(zinc.Default(), id, value)
}

// MustUpdateIsPlayingSilently ...
func MustUpdateIsPlayingSilently(id zinc.EntityID, value bool) {
	err := UpdateIsPlayingSilentlyX(zinc.Default(), id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlayingX ...
func UpdateIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID, value bool) error {
	v := e.Component(IsPlayingKey)
	c := v.(*IsPlayingComponent)
	return c.UpdateIsPlaying(id, value, false)
}

// MustUpdateIsPlayingX ...
func MustUpdateIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID, value bool) {
	err := UpdateIsPlayingX(e, id, value)
	if err != nil {
		panic(err)
	}
}

// UpdateIsPlaying ...
func UpdateIsPlaying(id zinc.EntityID, value bool) error {
	return UpdateIsPlayingX(zinc.Default(), id, value)
}

// MustUpdateIsPlaying ...
func MustUpdateIsPlaying(id zinc.EntityID, value bool) {
	err := UpdateIsPlayingX(zinc.Default(), id, value)
	if err != nil {
		panic(err)
	}
}

// HasIsPlayingX ...
func HasIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	v := e.Component(IsPlayingKey)
	return v.HasEntity(id)
}

// HasIsPlaying ...
func HasIsPlaying(id zinc.EntityID) bool {
	return HasIsPlayingX(zinc.Default(), id)
}

// IsPlayingX ...
func IsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID) (bool, error) {
	v := e.Component(IsPlayingKey)
	c := v.(*IsPlayingComponent)
	return c.IsPlaying(id)
}

// MustIsPlayingX ...
func MustIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID) bool {
	data, err := IsPlayingX(e, id)
	if err != nil {
		panic(err)
	}
	return data
}

// IsPlaying ...
func IsPlaying(id zinc.EntityID) (bool, error) {
	return IsPlayingX(zinc.Default(), id)
}

// MustIsPlaying ...
func MustIsPlaying(id zinc.EntityID) bool {
	data, err := IsPlayingX(zinc.Default(), id)
	if err != nil {
		panic(err)
	}
	return data
}

// DeleteIsPlayingX ...
func DeleteIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID) error {
	v := e.Component(IsPlayingKey)
	return v.DeleteEntity(id)
}

// MustDeleteIsPlayingX ...
func MustDeleteIsPlayingX(e *zinc.ZEntityManager, id zinc.EntityID) {
	err := DeleteIsPlayingX(e, id)
	if err != nil {
		panic(err)
	}
}

// DeleteIsPlaying ...
func DeleteIsPlaying(id zinc.EntityID) error {
	return DeleteIsPlayingX(zinc.Default(), id)
}

// MustDeleteIsPlaying ...
func MustDeleteIsPlaying(id zinc.EntityID) {
	err := DeleteIsPlaying(id)
	if err != nil {
		panic(err)
	}
}