package zinc

var entityManager = NewEntityManager()

// Default returns the default entity manager.
func Default() *ZEntityManager {
	return entityManager
}

// ResetAll deletes all entities in the
// current entity manager, resets the
// id pool to zero and unregisters all components
// and groups.
func ResetAll() {
	Default().ResetAll()
}

// Reset deletes all entities in the
// current entity manager and resets the
// id pool to zero.
func Reset() {
	Default().Reset()
}

// DeleteEntities deletes all active entities in
// the current entity manager. Each entity and all
// components associated with that entity are deleted.
func DeleteEntities() {
	Default().DeleteEntities()
}

// CreateEntity creates a new entity and returns the entity id.
// If an entity was previously deleted it will use the cached
// entity id and use that one without incrementing the current id.
func CreateEntity() ZEntityID {
	return Default().CreateEntity()
}

// HasEntity returns true if the entity id exists or
// false if it does not.
func HasEntity(id ZEntityID) bool {
	return Default().HasEntity(id)
}

// DeleteEntity deletes all components associated with
// the entity id and then deletes the entity from the current
// entity manager.
func DeleteEntity(id ZEntityID) {
	Default().DeleteEntity(id)
}

// Entities returns a slice of currently active entity ids.
func Entities() []ZEntityID {
	return Default().Entities()
}

// NewCollector returns a new collector for the given event triggers.
//
// You can use the built-in Added(Component...), Updated(Component...), Deleted(Component...) triggers in most cases,
// but you can also build custom triggers using the NewEventTrigger(Matcher, GroupEvent) function.
//
// As an example, this is useful for when you need to collect an entity that has ComponentA but not ComponentB and
// you want to collect the entity when it has been added to the group.
func NewCollector(et ...*ZEventTrigger) *ZCollector {
	return Default().NewCollector(et...)
}

// GroupCount returns the number of groups used by the current entity manager.
func GroupCount() int {
	return Default().GroupCount()
}

// Group returns a new group from a matcher or returns an existing one using the matcher hash.
func Group(m *ZMatcher) *ZGroup {
	return Default().Group(m)
}
