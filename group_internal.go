package zinc

func newGroup(e *ZEntityManager, m *ZMatcher) *ZGroup {
	return &ZGroup{
		entityManager: e,
		matcher:       m,
		entityList:    newEntityList(),
	}
}

func (g *ZGroup) addEntity(key uint, id ZEntityID) {
	if g.entityList.AddEntity(id) {
		for _, h := range g.addedFunc {
			h(key, id)
		}
	}
}

func (g *ZGroup) updateEntity(key uint, id ZEntityID) {
	if g.entityList.HasEntity(id) {
		for _, h := range g.updatedFunc {
			h(key, id)
		}
	}
}

func (g *ZGroup) deleteEntity(key uint, id ZEntityID) {
	if g.entityList.DeleteEntity(id) {
		for _, h := range g.deletedFunc {
			h(key, id)
		}
	}
}

func (g *ZGroup) handleEntitySilently(id ZEntityID) {
	if ok := g.matcher.Match(g.entityManager, id); ok {
		g.entityList.AddEntity(id)
	} else {
		g.entityList.DeleteEntity(id)
	}
}

func (g *ZGroup) handleEntity(key uint, id ZEntityID) {
	if ok := g.matcher.Match(g.entityManager, id); ok {
		g.addEntity(key, id)
	} else {
		g.deleteEntity(key, id)
	}
}
