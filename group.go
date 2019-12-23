package atom

// G ...
type G struct {
	matcher       *Matcher
	entityManager *EntityManager
	entitiesMap   map[EntityID]int
	entities      []EntityID
}

// NewGroup ...
func NewGroup(e *EntityManager, m *Matcher) *G {
	return &G{
		entityManager: e,
		matcher:       m,
		entitiesMap:   make(map[EntityID]int, 0),
	}
}

func (g *G) indexOf(id EntityID) int {
	idx, ok := g.entitiesMap[id]
	if !ok {
		return -1
	}
	return idx
}

func (g *G) deleteEntity(id EntityID) {
	idx := g.indexOf(id)
	g.entities = append(g.entities[:idx], g.entities[idx+1:]...)
	delete(g.entitiesMap, id)
}

func (g *G) addEntity(id EntityID) {
	g.entities = append(g.entities, id)
	g.entitiesMap[id] = len(g.entities) - 1
}

func (g *G) entityDeleted(id EntityID) {
	g.deleteEntity(id)
}

func (g *G) handleEntity(key uint, id EntityID) {
	ok := g.matcher.Match(id)
	if ok && !g.HasEntity(id) {
		g.addEntity(id)
	} else if !ok && g.HasEntity(id) {
		g.deleteEntity(id)
	}
}

// HasEntity ...
func (g *G) HasEntity(id EntityID) bool {
	idx := g.indexOf(id)
	if idx == -1 {
		return false
	}
	return true
}

// Entities ...
func (g *G) Entities() []EntityID {
	return g.entities
}
