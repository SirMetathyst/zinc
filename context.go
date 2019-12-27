package atom

// ComponentEventFunc ...
type ComponentEventFunc func(key uint, id EntityID)

// Context ...
type Context interface {
	ComponentAdded(key uint, id EntityID)
	ComponentDeleted(key uint, id EntityID)
	ComponentUpdated(key uint, id EntityID)
	HasEntity(id EntityID) bool
}

type ctx struct {
	componentAddedFunc   ComponentEventFunc
	componentDeletedFunc ComponentEventFunc
	componentUpdatedFunc ComponentEventFunc
	hasEntityFunc        func(id EntityID) bool
}

// NewContext ...
func NewContext(
	// Component Event(s)
	componentAddedFunc ComponentEventFunc,
	componentDeletedFunc ComponentEventFunc,
	componentUpdatedFunc ComponentEventFunc,
	// HasEntity
	hasEntityFunc func(id EntityID) bool) Context {

	return ctx{
		componentAddedFunc:   componentAddedFunc,
		componentDeletedFunc: componentDeletedFunc,
		componentUpdatedFunc: componentUpdatedFunc,
		hasEntityFunc:        hasEntityFunc,
	}
}

// ComponentAdded ...
func (c ctx) ComponentAdded(key uint, id EntityID) {
	c.componentAddedFunc(key, id)
}

// ComponentDeleted ...
func (c ctx) ComponentDeleted(key uint, id EntityID) {
	c.componentDeletedFunc(key, id)
}

// ComponentUpdated ...
func (c ctx) ComponentUpdated(key uint, id EntityID) {
	c.componentUpdatedFunc(key, id)
}

// HasEntity ...
func (c ctx) HasEntity(id EntityID) bool {
	return c.hasEntityFunc(id)
}
