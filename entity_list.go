package atom

// EntityList ...
type EntityList struct {
	entitiesMap map[EntityID]int
	entities    []EntityID
}

// NewEntityList ...
func NewEntityList() *EntityList {
	return &EntityList{entitiesMap: make(map[EntityID]int)}
}

// HasEntity ...
func (e *EntityList) HasEntity(id EntityID) bool {
	_, exist := e.entitiesMap[id]
	return exist
}

// AddEntity ...
func (e *EntityList) AddEntity(id EntityID) bool {
	if _, exist := e.entitiesMap[id]; !exist {
		e.entities = append(e.entities, id)
		e.entitiesMap[id] = len(e.entities) - 1
		return true
	}
	return false
}

// DeleteEntity ...
func (e *EntityList) DeleteEntity(id EntityID) bool {
	if idx, exist := e.entitiesMap[id]; exist {
		lid := e.entities[len(e.entities)-1]
		e.entities[idx] = lid
		e.entitiesMap[lid] = idx
		e.entities = e.entities[:len(e.entities)-1]
		delete(e.entitiesMap, id)
		return true
	}
	return false
}

// Entities ...
func (e *EntityList) Entities() []EntityID {
	return e.entities
}
