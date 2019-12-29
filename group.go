package atom

// G ...
type G interface {
	HandleEntitySilently(id EntityID)
	HandleEntity(key uint, id EntityID)
	UpdateEntity(key uint, id EntityID)
	DeleteEntity(id EntityID)
	HasEntity(id EntityID) bool
	Entities() []EntityID
	HandleEntityAdded(f EntityEventFunc)
	HandleEntityDeleted(f EntityEventFunc)
	HandleEntityUpdated(f EntityEventFunc)
}

// g ...
type g struct {
	matcher       *Matcher
	entityManager *EntityManager
	entityList    *EntityList
	addedFunc     []EntityEventFunc
	updatedFunc   []EntityEventFunc
	deletedFunc   []EntityEventFunc
}

// newGroup ...
func newGroup(e *EntityManager, m *Matcher) *g {
	return &g{
		entityManager: e,
		matcher:       m,
		entityList:    NewEntityList(),
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
	if ok := g.matcher.Match(id); ok {
		g.entityList.AddEntity(id)
	} else {
		g.entityList.DeleteEntity(id)
	}
}

// HandleEntity ...
func (g *g) HandleEntity(key uint, id EntityID) {
	if ok := g.matcher.Match(id); ok {
		g.addEntity(key, id)
	} else {
		g.deleteEntity(key, id)
	}
}

// UpdateEntity ...
func (g *g) UpdateEntity(key uint, id EntityID) {
	if g.entityList.HasEntity(id) && len(g.updatedFunc) > 0 {
		for _, h := range g.updatedFunc {
			h(key, id)
		}
	}
}

// DeleteEntity ...
func (g *g) DeleteEntity(id EntityID) {
	g.entityList.DeleteEntity(id)
}

// HasEntity ...
func (g *g) HasEntity(id EntityID) bool {
	return g.entityList.HasEntity(id)
}

// Entities ...
func (g *g) Entities() []EntityID {
	return g.entityList.Entities()
}

// HandleEntityAdded ...
func (g *g) HandleEntityAdded(f EntityEventFunc) {
	g.addedFunc = append(g.addedFunc, f)
}

// HandleEntityUpdated ...
func (g *g) HandleEntityUpdated(f EntityEventFunc) {
	g.updatedFunc = append(g.updatedFunc, f)
}

// HandleEntityDeleted ...
func (g *g) HandleEntityDeleted(f EntityEventFunc) {
	g.deletedFunc = append(g.deletedFunc, f)
}
