package zinc

type el struct {
	entitiesMap map[ZEntityID]int
	entities    []ZEntityID
}

func newEntityList() *el {
	return &el{entitiesMap: make(map[ZEntityID]int)}
}

// AddEntity ...
func (e *el) AddEntity(id ZEntityID) bool {
	if _, exist := e.entitiesMap[id]; !exist {
		e.entities = append(e.entities, id)
		e.entitiesMap[id] = len(e.entities) - 1
		return true
	}
	return false
}

// DeleteEntity ...
func (e *el) DeleteEntity(id ZEntityID) bool {
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
func (e *el) HasEntity(id ZEntityID) bool {
	_, exist := e.entitiesMap[id]
	return exist
}

// Entities ...
func (e *el) Entities() []ZEntityID {
	return e.entities
}
