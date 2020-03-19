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
