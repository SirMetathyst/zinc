package zinc

// EntityEventFunc ...
type EntityEventFunc func(key uint, id ZEntityID)

// ZContext ...
type ZContext struct {
	componentAddedFunc   EntityEventFunc
	componentDeletedFunc EntityEventFunc
	componentUpdatedFunc EntityEventFunc
	hasEntityFunc        func(id ZEntityID) bool
}

func newContext(
	// Component Event(s)
	componentAddedFunc EntityEventFunc,
	componentDeletedFunc EntityEventFunc,
	componentUpdatedFunc EntityEventFunc,
	// HasEntity
	hasEntityFunc func(id ZEntityID) bool) *ZContext {

	return &ZContext{
		componentAddedFunc:   componentAddedFunc,
		componentDeletedFunc: componentDeletedFunc,
		componentUpdatedFunc: componentUpdatedFunc,
		hasEntityFunc:        hasEntityFunc,
	}
}

// ComponentAdded ...
func (c ZContext) ComponentAdded(key uint, id ZEntityID) {
	c.componentAddedFunc(key, id)
}

// ComponentDeleted ...
func (c ZContext) ComponentDeleted(key uint, id ZEntityID) {
	c.componentDeletedFunc(key, id)
}

// ComponentUpdated ...
func (c ZContext) ComponentUpdated(key uint, id ZEntityID) {
	c.componentUpdatedFunc(key, id)
}

// HasEntity ...
func (c ZContext) HasEntity(id ZEntityID) bool {
	return c.hasEntityFunc(id)
}
