package zinc

var entityManager = NewEntityManager()

// Default ...
func Default() *ZEntityManager {
	return entityManager
}

// ResetAll ...
func ResetAll() {
	Default().ResetAll()
}

// Reset ...
func Reset() {
	Default().Reset()
}

// DeleteEntities ...
func DeleteEntities() {
	Default().DeleteEntities()
}

// CreateEntity ...
func CreateEntity() EntityID {
	return Default().CreateEntity()
}

// HasEntity ...
func HasEntity(id EntityID) bool {
	return Default().HasEntity(id)
}

// DeleteEntity ...
func DeleteEntity(id EntityID) {
	Default().DeleteEntity(id)
}

// Entities ...
func Entities() []EntityID {
	return Default().Entities()
}

// NewCollector ...
func NewCollector(et ...*ZEventTrigger) *ZCollector {
	return Default().NewCollector(et...)
}

// GroupCount ...
func GroupCount() int {
	return Default().GroupCount()
}

// Group ...
func Group(m *ZMatcher) *ZGroup {
	return Default().Group(m)
}
