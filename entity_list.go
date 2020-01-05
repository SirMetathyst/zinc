package zinc

type el struct {
	entitiesMap map[EntityID]int
	entities    []EntityID
}

func newEntityList() *el {
	return &el{entitiesMap: make(map[EntityID]int)}
}

// AddEntity ...
func (e *el) AddEntity(id EntityID) bool {
	if _, exist := e.entitiesMap[id]; !exist {
		e.entities = append(e.entities, id)
		e.entitiesMap[id] = len(e.entities) - 1
		return true
	}
	return false
}

// DeleteEntity ...
func (e *el) DeleteEntity(id EntityID) bool {
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

// HasEntity ...
func (e *el) HasEntity(id EntityID) bool {
	_, exist := e.entitiesMap[id]
	return exist
}

// Entities ...
func (e *el) Entities() []EntityID {
	return e.entities
}
