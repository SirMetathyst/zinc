package atom

// Context ...
type Context interface {
	ComponentAdded(key uint, id EntityID)
	ComponentDeleted(key uint, id EntityID)
	ComponentUpdated(key uint, id EntityID)
	HasEntity(id EntityID) bool
}

type ctx struct {
	entityManager *EntityManager
}

// NewContext ...
func NewContext(e *EntityManager) Context {
	return ctx{entityManager: e}
}

// ComponentAdded ...
func (c ctx) ComponentAdded(key uint, id EntityID) {
	c.entityManager.componentAdded(key, id)
}

// ComponentDeleted ...
func (c ctx) ComponentDeleted(key uint, id EntityID) {
	c.entityManager.componentDeleted(key, id)
}

// ComponentUpdated ...
func (c ctx) ComponentUpdated(key uint, id EntityID) {
	c.entityManager.componentUpdated(key, id)
}

// HasEntity ...
func (c ctx) HasEntity(id EntityID) bool {
	return c.entityManager.HasEntity(id)
}
