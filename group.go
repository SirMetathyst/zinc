package atom

// Grouper ...
type Grouper interface {
	HasEntity(id EntityID) bool
	Entities() []EntityID
}

// BasicGroup ...
type BasicGroup struct {
	matcher       *Matcher
	entityManager *EntityManager
	entitiesMap   map[EntityID]int
	entities      []EntityID
}

// NewBasicGroup ...
func NewBasicGroup(e *EntityManager, m *Matcher) *BasicGroup {
	return &BasicGroup{
		entityManager: e,
		matcher:       m,
		entitiesMap:   make(map[EntityID]int, 0),
	}
}

func (g *BasicGroup) indexOf(id EntityID) int {
	idx, ok := g.entitiesMap[id]
	if !ok {
		return -1
	}
	return idx
}

func (g *BasicGroup) deleteEntity(id EntityID) {
	idx := g.indexOf(id)
	lastv := g.entities[len(g.entities)-1]
	g.entities[idx] = lastv
	g.entitiesMap[lastv] = idx
	g.entities = g.entities[:len(g.entities)-1]
	delete(g.entitiesMap, id)
}

func (g *BasicGroup) addEntity(id EntityID) {
	g.entities = append(g.entities, id)
	g.entitiesMap[id] = len(g.entities) - 1
}

func (g *BasicGroup) entityDeleted(id EntityID) {
	g.deleteEntity(id)
}

func (g *BasicGroup) handleEntity(key uint, id EntityID) {
	ok := g.matcher.Match(id)
	if ok && !g.HasEntity(id) {
		g.addEntity(id)
	} else if !ok && g.HasEntity(id) {
		g.deleteEntity(id)
	}
}

// HasEntity ...
func (g *BasicGroup) HasEntity(id EntityID) bool {
	if idx := g.indexOf(id); idx == -1 {
		return false
	}
	return true
}

// Entities ...
func (g *BasicGroup) Entities() []EntityID {
	return g.entities
}
