/*
TODO: instead of registering component with entity manager on init func, register
with a component registry so that you can call a single method e.g. EntityManager.RegisterComponentsWithRegistry...
instead of having call myKit.RegisterMyComponent for every single component for multiple entity managers...
*/

package zinc

import "fmt"

var (
	
)

// EntityID ...
type EntityID uint

// ZEntityManager ...
type ZEntityManager struct {
	entityList   *el
	groupsMap    map[uint]int
	pool         []EntityID
	id           EntityID
	groups       []*ZGroup
	context      *ZContext
	componentMap map[uint]Component
}

// NewEntityManager ...
func NewEntityManager() *ZEntityManager {
	e := &ZEntityManager{
		entityList:   newEntityList(),
		groupsMap:    make(map[uint]int, 0),
		componentMap: make(map[uint]Component),
	}

	e.context = newContext(
		e.componentAdded,
		e.componentDeleted,
		e.componentUpdated,
		e.HasEntity)

	return e
}

func (e *ZEntityManager) addGroup(g *ZGroup) {
	e.groups = append(e.groups, g)
	e.groupsMap[g.Hash()] = len(e.groups) - 1
}

func (e *ZEntityManager) deleteEntity(id EntityID) {
	if !e.entityList.DeleteEntity(id) {
		panic("what?!")
	}
	e.putID(id)
}

func (e *ZEntityManager) componentDeleteEntity(id EntityID) {
	for _, c := range e.componentMap {
		if c.HasEntity(id) {
			c.DeleteEntity(id)
		}
	}
}

func (e *ZEntityManager) groupHandleEntity(key uint, id EntityID) {
	for _, g := range e.groups {
		g.HandleEntity(key, id)
	}
}

func (e *ZEntityManager) groupUpdateEntity(key uint, id EntityID) {
	for _, g := range e.groups {
		g.UpdateEntity(key, id)
	}
}

func (e *ZEntityManager) componentAdded(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

func (e *ZEntityManager) componentUpdated(key uint, id EntityID) {
	e.groupUpdateEntity(key, id)
}

func (e *ZEntityManager) componentDeleted(key uint, id EntityID) {
	e.groupHandleEntity(key, id)
}

// ResetAll ...
func (e *ZEntityManager) ResetAll() {
	e.Reset()
	e.groupsMap = make(map[uint]int, 0)
	e.groups = nil
	e.componentMap = make(map[uint]Component)
}

// Reset ...
func (e *ZEntityManager) Reset() {
	e.DeleteEntities()
	e.pool = e.pool[:0]
	e.id = 0
}

// DeleteEntities ...
func (e *ZEntityManager) DeleteEntities() {
	for _, id := range e.Entities() {
		e.DeleteEntity(id)
	}
}

func (e *ZEntityManager) putID(id EntityID) {
	e.pool = append(e.pool, id)
}

func (e *ZEntityManager) getID() EntityID {
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
func (e *ZEntityManager) CreateEntity() EntityID {
	id := e.getID()
	if !e.entityList.AddEntity(id) {
		panic("what?!")
	}
	return id
}

// HasEntity ...
func (e *ZEntityManager) HasEntity(id EntityID) bool {
	return e.entityList.HasEntity(id)
}

// DeleteEntity ...
func (e *ZEntityManager) DeleteEntity(id EntityID) {
	if e.HasEntity(id) {
		e.componentDeleteEntity(id)
		e.deleteEntity(id)
	}
}

// Entities ...
func (e *ZEntityManager) Entities() []EntityID {
	return e.entityList.Entities()
}

// RegisterComponent ...
func (e *ZEntityManager) RegisterComponent(key uint, c Component) *ZContext {
	if _, exist := e.componentMap[key]; !exist {
		if c != nil {
			e.componentMap[key] = c
		} else {
			panic("component must not be nil")
		}
	} else {
		panic("component key already registered")
	}
	return e.context
}

// Component ...
func (e *ZEntityManager) Component(key uint) (c Component, ok bool) {
	c, ok = e.componentMap[key]
	if !ok {
		panic(fmt.Sprintf("Failed to get component for key %v. Did you forgot to register the component?", key))
	}
	return
}

// NewCollector ...
func (e *ZEntityManager) NewCollector(et ...*ZEventTrigger) *ZCollector {
	var collectorEvent []ZCollectorEvent
	for _, v := range et {
		groupEvent := v.GroupEvent()
		group := e.Group(v.Matcher())
		collectorEvent = append(collectorEvent, newCollectorEvent(group, groupEvent))
	}
	return newCollector(collectorEvent...)
}

// GroupCount ...
func (e *ZEntityManager) GroupCount() int {
	return len(e.groups)
}

// Group ...
func (e *ZEntityManager) Group(m *ZMatcher) *ZGroup {
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

// Component ...
type Component interface {
	DeleteEntity(id EntityID)
	HasEntity(id EntityID) bool
}
