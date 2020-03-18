// Package kit ...
// generated by zinc/UniqueSingleComponentDataTemplate. DO NOT EDIT.
package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZLogService ...
var ZLogService uint = uint(727081216)

// LogServiceComponent ...
type LogServiceComponent struct {
	ctx  *zinc.ZContext
	data interface{}
}

// RegisterLogServiceComponentWith ...
func RegisterLogServiceComponentWith(e *zinc.ZEntityManager) {
	x := NewLogServiceComponent()
	ctx := e.RegisterComponent(ZLogService, x)
	x.SetContext(ctx)
}

// RegisterLogServiceComponent ...
func RegisterLogServiceComponent() {
	x := NewLogServiceComponent()
	ctx := zinc.Default().RegisterComponent(ZLogService, x)
	x.SetContext(ctx)
}

// NewLogServiceComponent ...
func NewLogServiceComponent() *LogServiceComponent {
	return &LogServiceComponent{}
}

func init() {
	RegisterLogServiceComponent()
}

// SetContext ...
func (c *LogServiceComponent) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// SetLogService ...
func (c *LogServiceComponent) SetLogService(v interface{}) {
	c.data = v
}

// LogService ...
func (c *LogServiceComponent) LogService() interface{} {
	return c.data
}

// HasEntity ...
func (c *LogServiceComponent) HasEntity(id zinc.ZEntityID) bool {
	return false
}

// DeleteEntity ...
func (c *LogServiceComponent) DeleteEntity(id zinc.ZEntityID) error {
	return nil
}

// SetLogServiceX ...
func SetLogServiceX(e *zinc.ZEntityManager, logger interface{}) {
	v := e.Component(ZLogService)
	c := v.(*LogServiceComponent)
	c.SetLogService(logger)
}

// SetLogService ...
func SetLogService(logger interface{}) {
	SetLogServiceX(zinc.Default(), logger)
}

// LogServiceX ...
func LogServiceX(e *zinc.ZEntityManager) interface{} {
	v := e.Component(ZLogService)
	c := v.(*LogServiceComponent)
	return c.LogService()
}

// LogService ...
func LogService() interface{} {
	return LogServiceX(zinc.Default())
}