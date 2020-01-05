package zinc

// EntityEventFunc ...
type EntityEventFunc func(key uint, id EntityID)

// CTX ...
type CTX interface {
	ComponentAdded(key uint, id EntityID)
	ComponentDeleted(key uint, id EntityID)
	ComponentUpdated(key uint, id EntityID)
	HasEntity(id EntityID) bool
}

type ctx struct {
	componentAddedFunc   EntityEventFunc
	componentDeletedFunc EntityEventFunc
	componentUpdatedFunc EntityEventFunc
	hasEntityFunc        func(id EntityID) bool
}

func newContext(
	// Component Event(s)
	componentAddedFunc EntityEventFunc,
	componentDeletedFunc EntityEventFunc,
	componentUpdatedFunc EntityEventFunc,
	// HasEntity
	hasEntityFunc func(id EntityID) bool) CTX {

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
