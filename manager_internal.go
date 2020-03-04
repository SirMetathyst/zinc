package zinc

func (e *ZEntityManager) putID(id ZEntityID) {
	e.pool = append(e.pool, id)
}

func (e *ZEntityManager) getID() ZEntityID {
	l := len(e.pool)
	if l == 0 {
		e.id++
		return e.id
	}
	v := e.pool[l-1]
	e.pool = e.pool[:l-1]
	return v
}

func (e *ZEntityManager) deleteEntity(id ZEntityID) {
	if !e.entityList.DeleteEntity(id) {
		panic("what?!")
	}
	e.putID(id)
}

func (e *ZEntityManager) newGroup(m *ZMatcher) *ZGroup {
	g := newGroup(e, m)
	for _, id := range e.Entities() {
		g.handleEntitySilently(id)
	}
	return g
}

func (e *ZEntityManager) addGroup(g *ZGroup) {
	e.groups = append(e.groups, g)
	e.groupsMap[g.Hash()] = len(e.groups) - 1
}

func (e *ZEntityManager) group(m *ZMatcher) *ZGroup {
	if idx, ok := e.groupsMap[m.Hash()]; ok {
		return e.groups[idx]
	}
	return nil
}

func (e *ZEntityManager) groupHandleEntity(key uint, id ZEntityID) {
	for _, g := range e.groups {
		g.handleEntity(key, id)
	}
}

func (e *ZEntityManager) groupUpdateEntity(key uint, id ZEntityID) {
	for _, g := range e.groups {
		g.updateEntity(key, id)
	}
}

func (e *ZEntityManager) componentDeleteEntity(id ZEntityID) {
	for _, c := range e.componentMap {
		if c.HasEntity(id) {
			c.DeleteEntity(id)
		}
	}
}

func (e *ZEntityManager) componentAdded(key uint, id ZEntityID) {
	e.groupHandleEntity(key, id)
}

func (e *ZEntityManager) componentUpdated(key uint, id ZEntityID) {
	e.groupUpdateEntity(key, id)
}

func (e *ZEntityManager) componentDeleted(key uint, id ZEntityID) {
	e.groupHandleEntity(key, id)
}
