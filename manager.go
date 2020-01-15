/*
TODO: instead of registering component with entity manager on init func, register
with a component registry so that you can call a single method e.g. EntityManager.RegisterComponentsWithRegistry...
instead of having call myKit.RegisterMyComponent for every single component for multiple entity managers...
*/

package zinc

// EntityID ...
type EntityID uint

// EntityManager ...
type EntityManager struct {
	entityList   *el
	groupsMap    map[uint]int
	pool         []EntityID
	id           EntityID
	groups       []*g
	context      CTX
	componentMap map[uint]CMP
}

// NewEntityManager ...
func NewEntityManager() *EntityManager {
	e := &EntityManager{
		entityList:   newEntityList(),
		groupsMap:    make(map[uint]int, 0),
		componentMap: make(map[uint]CMP),
	}

	e.context = newContext(
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
	if !e.entityList.DeleteEntity(id) {
		panic("what?!")
	}
	e.putID(id)
}

func (e *EntityManager) componentDeleteEntity(id EntityID) {
	for _, c := range e.componentMap {
		c.DeleteEntity(id)
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

// ResetAll ...
func (e *EntityManager) ResetAll() {
	e.Reset()
	e.groupsMap = make(map[uint]int, 0)
	e.groups = nil
	e.componentMap = make(map[uint]CMP)
}

// Reset ...
func (e *EntityManager) Reset() {
	e.DeleteEntities()
	e.pool = e.pool[:0]
	e.id = 0
}

// DeleteEntities ...
func (e *EntityManager) DeleteEntities() {
	for _, id := range e.Entities() {
		e.DeleteEntity(id)
	}
}

func (e *EntityManager) putID(id EntityID) {
	e.pool = append(e.pool, id)
}

func (e *EntityManager) getID() EntityID {
	l := len(e.pool)
	if l == 0 {
		e.id++
		return e.id
	}
	i := l - 1
	v := e.pool[i]
	e.pool = e.pool[:i]
	return v
}

// CreateEntity ...
func (e *EntityManager) CreateEntity() EntityID {
	id := e.getID()
	if !e.entityList.AddEntity(id) {
		panic("what?!")
	}
	return id
}

// HasEntity ...
func (e *EntityManager) HasEntity(id EntityID) bool {
	return e.entityList.HasEntity(id)
}

// DeleteEntity ...
func (e *EntityManager) DeleteEntity(id EntityID) {
	if e.HasEntity(id) {
		e.deleteEntity(id)
		e.componentDeleteEntity(id)
	}
}

// Entities ...
func (e *EntityManager) Entities() []EntityID {
	return e.entityList.Entities()
}

// RegisterComponent ...
func (e *EntityManager) RegisterComponent(key uint, c CMP) CTX {
	if _, exist := e.componentMap[key]; !exist {
		e.componentMap[key] = c
	} else {
		panic("component key already registered")
	}
	return e.context
}

// Component ...
func (e *EntityManager) Component(key uint) (c CMP, ok bool) {
	c, ok = e.componentMap[key]
	return
}

// CreateCollector ...
func (e *EntityManager) CreateCollector(et ...ET) C {
	groupEvent := make([]GroupEvent, len(et))
	group := make([]G, len(et))
	for i, v := range et {
		groupEvent[i] = v.GroupEvent()
		group[i] = e.Group(v.Matcher())
	}
	return NewCollector(group, groupEvent)
}

// GroupCount ...
func (e *EntityManager) GroupCount() int {
	return len(e.groups)
}

// Group ...
func (e *EntityManager) Group(m M) G {
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

// CMP ...
type CMP interface {
	DeleteEntity(id EntityID)
	HasEntity(id EntityID) bool
}
