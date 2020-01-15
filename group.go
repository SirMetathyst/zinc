package zinc

// G ...
type G interface {
	HandleEntitySilently(id EntityID)
	HandleEntity(key uint, id EntityID)
	UpdateEntity(key uint, id EntityID)
	HasEntity(id EntityID) bool
	Entities() []EntityID
	HandleEntityAdded(f EntityEventFunc)
	HandleEntityDeleted(f EntityEventFunc)
	HandleEntityUpdated(f EntityEventFunc)
}

// g ...
type g struct {
	matcher       M
	entityManager *EntityManager
	entityList    *el
	addedFunc     []EntityEventFunc
	updatedFunc   []EntityEventFunc
	deletedFunc   []EntityEventFunc
}

// newGroup ...
// TODO: Write TEST
func newGroup(e *EntityManager, m M) *g {
	return &g{
		entityManager: e,
		matcher:       m,
		entityList:    newEntityList(),
	}
}

func (g *g) addEntity(key uint, id EntityID) {
	if g.entityList.AddEntity(id) && len(g.addedFunc) > 0 {
		for _, h := range g.addedFunc {
			h(key, id)
		}
	}
}

func (g *g) deleteEntity(key uint, id EntityID) {
	if g.entityList.DeleteEntity(id) && len(g.deletedFunc) > 0 {
		for _, h := range g.deletedFunc {
			h(key, id)
		}
	}
}

// HandleEntitySilently ...
func (g *g) HandleEntitySilently(id EntityID) {
	if ok := g.matcher.Match(g.entityManager, id); ok {
		g.entityList.AddEntity(id)
	} else {
		g.entityList.DeleteEntity(id)
	}
}

// HandleEntity ...
func (g *g) HandleEntity(key uint, id EntityID) {
	if ok := g.matcher.Match(g.entityManager, id); ok {
		g.addEntity(key, id)
	} else {
		g.deleteEntity(key, id)
	}
}

// UpdateEntity ...
// TODO: Write TEST
func (g *g) UpdateEntity(key uint, id EntityID) {
	if g.entityList.HasEntity(id) && len(g.updatedFunc) > 0 {
		for _, h := range g.updatedFunc {
			h(key, id)
		}
	}
}

// HasEntity ...
// TODO: Write TEST
func (g *g) HasEntity(id EntityID) bool {
	return g.entityList.HasEntity(id)
}

// Entities ...
// TODO: Write TEST
func (g *g) Entities() []EntityID {
	return g.entityList.Entities()
}

// HandleEntityAdded ...
// TODO: Write TEST
func (g *g) HandleEntityAdded(f EntityEventFunc) {
	g.addedFunc = append(g.addedFunc, f)
}

// HandleEntityUpdated ...
// TODO: Write TEST
func (g *g) HandleEntityUpdated(f EntityEventFunc) {
	g.updatedFunc = append(g.updatedFunc, f)
}

// HandleEntityDeleted ...
// TODO: Write TEST
func (g *g) HandleEntityDeleted(f EntityEventFunc) {
	g.deletedFunc = append(g.deletedFunc, f)
}
