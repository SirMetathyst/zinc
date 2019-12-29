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
	entityList   *EntityList
	groupsMap    map[uint]int
	groups       []*g
	pool         *Pool
	context      Context
	componentMap map[uint]Component
}

// NewEntityManager ...
func NewEntityManager() *EntityManager {
	e := &EntityManager{
		entityList:   NewEntityList(),
		groupsMap:    make(map[uint]int, 0),
		pool:         NewPool(NewEntityIDFactory()),
		componentMap: make(map[uint]Component),
	}

	e.context = NewContext(
		e.componentAdded,
		e.componentDeleted,
		e.componentUpdated,
		e.HasEntity)

	return e
}

func (e *EntityManager) addGroup(g *g) {
	e.groups = append(e.groups, g)
	e.groupsMap[g.matcher.Hash()] = len(e.groups) - 1
}

func (e *EntityManager) deleteEntity(id EntityID) {
	e.entityList.DeleteEntity(id)
	e.pool.Put(id)
}

func (e *EntityManager) componentDeleteEntity(id EntityID) {
	for _, c := range e.componentMap {
		c.DeleteEntity(id)
	}
}

func (e *EntityManager) groupDeleteEntity(id EntityID) {
	for _, g := range e.groups {
		g.DeleteEntity(id)
	}
}

func (e *EntityManager) groupHandleEntity(key uint, id EntityID) {
	for _, g := range e.groups {
		g.HandleEntity(key, id)
	}
}

func (e *EntityManager) groupUpdateEntity(key uint, id EntityID) {
	for _, g := range e.groups {
		g.UpdateEntity(key, id)
	}
}

func (e *EntityManager) componentAdded(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

func (e *EntityManager) componentUpdated(key uint, id EntityID) {
	e.groupUpdateEntity(key, id)
}

func (e *EntityManager) componentDeleted(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

// DeleteEntities ...
func (e *EntityManager) DeleteEntities() {
	for _, id := range e.Entities() {
		e.DeleteEntity(id)
	}
}

// DeleteEntities ...
func DeleteEntities() {
	Default().DeleteEntities()
}

// CreateEntity ...
func (e *EntityManager) CreateEntity() EntityID {
	v := e.pool.Get()
	id := v.(EntityID)
	e.entityList.AddEntity(id)
	return id
}

// CreateEntity ...
func CreateEntity() EntityID {
	return Default().CreateEntity()
}

// HasEntity ...
func (e *EntityManager) HasEntity(id EntityID) bool {
	return e.entityList.HasEntity(id)
}

// HasEntity ...
func HasEntity(id EntityID) bool {
	return Default().HasEntity(id)
}

// DeleteEntity ...
func (e *EntityManager) DeleteEntity(id EntityID) {
	if e.HasEntity(id) {
		e.deleteEntity(id)
		e.componentDeleteEntity(id)
		e.groupDeleteEntity(id)
	}
}

// DeleteEntity ...
func DeleteEntity(id EntityID) {
	Default().DeleteEntity(id)
}

// Entities ...
func (e *EntityManager) Entities() []EntityID {
	return e.entityList.Entities()
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
func (e *EntityManager) Component(key uint) (c Component, ok bool) {
	c, ok = e.componentMap[key]
	return
}

// CreateCollector ...
func (e *EntityManager) CreateCollector(et ...*EventTrigger) C {
	groupEvent := make([]GroupEvent, len(et))
	group := make([]G, len(et))
	for i, v := range et {
		groupEvent[i] = v.GroupEvent()
		group[i] = e.Group(v.Matcher())
	}
	return newCollector(group, groupEvent)
}

// Group ...
func (e *EntityManager) Group(m *Matcher) G {
	if idx, ok := e.groupsMap[m.Hash()]; ok {
		return e.groups[idx]
	}
	g := newGroup(e, m)
	for _, id := range e.Entities() {
		g.HandleEntitySilently(id)
	}
	e.addGroup(g)
	return g
}

// Group ...
func Group(m *Matcher) G {
	return Default().Group(m)
}

// Component ...
type Component interface {
	DeleteEntity(id EntityID)
	HasEntity(id EntityID) bool
}
