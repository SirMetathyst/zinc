package zinc

// EntityEventFunc ...
type EntityEventFunc func(key uint, id EntityID)

// ZContext ...
type ZContext struct {
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
	hasEntityFunc func(id EntityID) bool) *ZContext {

	return &ZContext{
		componentAddedFunc:   componentAddedFunc,
		componentDeletedFunc: componentDeletedFunc,
		componentUpdatedFunc: componentUpdatedFunc,
		hasEntityFunc:        hasEntityFunc,
	}
}

// ComponentAdded ...
func (c ZContext) ComponentAdded(key uint, id EntityID) {
	c.componentAddedFunc(key, id)
}

// ComponentDeleted ...
func (c ZContext) ComponentDeleted(key uint, id EntityID) {
	c.componentDeletedFunc(key, id)
}

// ComponentUpdated ...
func (c ZContext) ComponentUpdated(key uint, id EntityID) {
	c.componentUpdatedFunc(key, id)
}

// HasEntity ...
func (c ZContext) HasEntity(id EntityID) bool {
	return c.hasEntityFunc(id)
}
