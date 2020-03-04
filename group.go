package zinc

// ZGroup ...
type ZGroup struct {
	matcher       *ZMatcher
	entityManager *ZEntityManager
	entityList    *el
	addedFunc     []EntityEventFunc
	updatedFunc   []EntityEventFunc
	deletedFunc   []EntityEventFunc
}

// newGroup ...
// TODO: Write TEST
func newGroup(e *ZEntityManager, m *ZMatcher) *ZGroup {
	return &ZGroup{
		entityManager: e,
		matcher:       m,
		entityList:    newEntityList(),
	}
}

func (g *ZGroup) addEntity(key uint, id ZEntityID) {
	if g.entityList.AddEntity(id) && len(g.addedFunc) > 0 {
		for _, h := range g.addedFunc {
			h(key, id)
		}
	}
}

func (g *ZGroup) deleteEntity(key uint, id ZEntityID) {
	if g.entityList.DeleteEntity(id) && len(g.deletedFunc) > 0 {
		for _, h := range g.deletedFunc {
			h(key, id)
		}
	}
}

// HandleEntitySilently ...
func (g *ZGroup) HandleEntitySilently(id ZEntityID) {
	if ok := g.matcher.Match(g.entityManager, id); ok {
		g.entityList.AddEntity(id)
	} else {
		g.entityList.DeleteEntity(id)
	}
}

// HandleEntity ...
func (g *ZGroup) HandleEntity(key uint, id ZEntityID) {
	if ok := g.matcher.Match(g.entityManager, id); ok {
		g.addEntity(key, id)
	} else {
		g.deleteEntity(key, id)
	}
}

// UpdateEntity ...
// TODO: Write TEST
func (g *ZGroup) UpdateEntity(key uint, id ZEntityID) {
	if g.entityList.HasEntity(id) && len(g.updatedFunc) > 0 {
		for _, h := range g.updatedFunc {
			h(key, id)
		}
	}
}

// HasEntity ...
// TODO: Write TEST
func (g *ZGroup) HasEntity(id ZEntityID) bool {
	return g.entityList.HasEntity(id)
}

// Entities ...
// TODO: Write TEST
func (g *ZGroup) Entities() []ZEntityID {
	return g.entityList.Entities()
}

// Hash ...
func (g *ZGroup) Hash() uint {
	return g.matcher.Hash()
}

// HandleEntityAdded ...
// TODO: Write TEST
func (g *ZGroup) HandleEntityAdded(f EntityEventFunc) {
	g.addedFunc = append(g.addedFunc, f)
}

// HandleEntityUpdated ...
// TODO: Write TEST
func (g *ZGroup) HandleEntityUpdated(f EntityEventFunc) {
	g.updatedFunc = append(g.updatedFunc, f)
}

// HandleEntityDeleted ...
// TODO: Write TEST
func (g *ZGroup) HandleEntityDeleted(f EntityEventFunc) {
	g.deletedFunc = append(g.deletedFunc, f)
}
