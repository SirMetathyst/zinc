package atom

var entityManager = NewEntityManager()

// Default ...
func Default() *EntityManager {
	return entityManager
}

// EntityID ...
type EntityID uint

// NewEntityIDFactory ...
func NewEntityIDFactory() FactoryFunc {
	id := EntityID(0)
	return func() interface{} {
		id++
		return id
	}
}

// EntityManager ...
type EntityManager struct {
	entitiesMap  map[EntityID]int
	entities     []EntityID
	groupsMap    map[string]int
	groups       []*G
	pool         *Pool
	context      Context
	componentMap map[uint]Component
}

// NewEntityManager ...
func NewEntityManager() *EntityManager {
	e := &EntityManager{
		entitiesMap:  make(map[EntityID]int, 0),
		groupsMap:    make(map[string]int, 0),
		pool:         NewPool(NewEntityIDFactory()),
		componentMap: make(map[uint]Component),
	}
	e.context = NewContext(e)
	return e
}

func (e *EntityManager) indexOf(id EntityID) int {
	idx, ok := e.entitiesMap[id]
	if !ok {
		return -1
	}
	return idx
}

func (e *EntityManager) addEntity(id EntityID) {
	e.entities = append(e.entities, id)
	e.entitiesMap[id] = len(e.entities) - 1
}

func (e *EntityManager) addGroup(g *G) {
	e.groups = append(e.groups, g)
	e.groupsMap[g.matcher.Hash()] = len(e.groups) - 1
}

func (e *EntityManager) deleteEntity(id EntityID) {
	idx := e.indexOf(id)
	e.entities = append(e.entities[:idx], e.entities[idx+1:]...)
	delete(e.entitiesMap, id)
	e.pool.Put(id)
}

func (e *EntityManager) entityDeleted(id EntityID) {
	e.componentEntityDeleted(id)
	e.groupEntityDeleted(id)
}

func (e *EntityManager) componentEntityDeleted(id EntityID) {
	for _, v := range e.componentMap {
		v.EntityDeleted(id)
	}
}

func (e *EntityManager) groupEntityDeleted(id EntityID) {
	for _, g := range e.groups {
		g.entityDeleted(id)
	}
}

func (e *EntityManager) groupHandleEntity(key uint, id EntityID) {
	for _, g := range e.groups {
		g.handleEntity(key, id)
	}
}

func (e *EntityManager) componentAdded(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

func (e *EntityManager) componentUpdated(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

func (e *EntityManager) componentDeleted(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

// Reset ...
func (e *EntityManager) Reset() {
	for _, x := range e.Entities() {
		e.DeleteEntity(x)
	}
	e.pool = NewPool(NewEntityIDFactory())
}

// Reset ...
func Reset() {
	Default().Reset()
}

// CreateEntity ...
func (e *EntityManager) CreateEntity() EntityID {
	v := e.pool.Get()
	id := v.(EntityID)
	e.addEntity(id)
	return id
}

// CreateEntity ...
func CreateEntity() EntityID {
	return Default().CreateEntity()
}

// HasEntity ...
func (e *EntityManager) HasEntity(id EntityID) bool {
	idx := e.indexOf(id)
	if idx == -1 {
		return false
	}
	return true
}

// HasEntity ...
func HasEntity(id EntityID) bool {
	return Default().HasEntity(id)
}

// DeleteEntity ...
func (e *EntityManager) DeleteEntity(id EntityID) {
	if e.HasEntity(id) {
		e.deleteEntity(id)
		e.entityDeleted(id)
	}
}

// DeleteEntity ...
func DeleteEntity(id EntityID) {
	Default().DeleteEntity(id)
}

// Entities ...
func (e *EntityManager) Entities() []EntityID {
	return e.entities
}

// Entities ...
func Entities() []EntityID {
	return Default().Entities()
}

// RegisterComponent ...
func (e *EntityManager) RegisterComponent(key uint, c Component) Context {
	e.componentMap[key] = c
	return e.context
}

// Component ...
func (e *EntityManager) Component(key uint) Component {
	v, _ := e.componentMap[key]
	return v
}

// Group ...
func (e *EntityManager) Group(m *Matcher) *G {
	h := m.Hash()
	if idx, ok := e.groupsMap[h]; ok {
		return e.groups[idx]
	}
	g := NewGroup(e, m)
	e.addGroup(g)
	return g
}

// Group ...
func Group(m *Matcher) *G {
	g := Default().Group(m)
	return g
}

// Component ...
type Component interface {
	EntityDeleted(id EntityID)
	HasEntity(id EntityID) bool
}