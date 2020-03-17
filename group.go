package zinc

import "errors"

var (
	// ErrNilGroup ...
	ErrNilGroup = errors.New("zinc: group is nil")
)

// ZGroup ...
type ZGroup struct {
	matcher       *ZMatcher
	entityManager *ZEntityManager
	entityList    *el
	addedFunc     []EntityEventFunc
	updatedFunc   []EntityEventFunc
	deletedFunc   []EntityEventFunc
}

// HasEntity returns true if the entity id is present
// in the group or false if it does not.
func (g *ZGroup) HasEntity(id ZEntityID) bool {
	return g.entityList.HasEntity(id)
}

// Entities returns a slice of entity ids
// associated with the current group.
func (g *ZGroup) Entities() []ZEntityID {
	return g.entityList.Entities()
}

// Hash returns a hash of the group's matcher.
func (g *ZGroup) Hash() uint {
	return g.matcher.Hash()
}

// RegisterEntityAddedFunc adds a callback function which
// is called when an entity has been added to the group.
func (g *ZGroup) RegisterEntityAddedFunc(f EntityEventFunc) {
	g.addedFunc = append(g.addedFunc, f)
}

// RegisterEntityUpdatedFunc adds a callback function which
// is called when an entity in the group has been updated.
func (g *ZGroup) RegisterEntityUpdatedFunc(f EntityEventFunc) {
	g.updatedFunc = append(g.updatedFunc, f)
}

// RegisterEntityDeletedFunc adds a callback function which
// is called when an entity has been deleted from the group.
func (g *ZGroup) RegisterEntityDeletedFunc(f EntityEventFunc) {
	g.deletedFunc = append(g.deletedFunc, f)
}
