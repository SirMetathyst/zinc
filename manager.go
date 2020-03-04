// Licensed to Alexander Hunt under one or more agreements.
// Alexander Hunt licenses this file to you under the MIT license.
// See the LICENSE file in the project root for more information.

package zinc

import (
	"errors"
)

var (
	// ErrComponentNil ...
	ErrComponentNil = errors.New("zinc: component is nil")
	// ErrComponentAlreadyRegistered ...
	ErrComponentAlreadyRegistered = errors.New("zinc: component has already been registered. Perhaps RegisterComponent was called more than once or there was a hash collision?")
	// ErrComponentNotFound ...
	ErrComponentNotFound = errors.New("zinc: component was not found. Are you sure the entity id is correct?")
	// ErrUnregisteredComponent ...
	ErrUnregisteredComponent = errors.New("zinc: component was not found. Did you forget to call RegisterComponent")
)

// ZEntityID ...
type ZEntityID uint

// Component ...
type Component interface {
	DeleteEntity(id ZEntityID) error
	HasEntity(id ZEntityID) bool
}

// ZEntityManager ...
type ZEntityManager struct {
	entityList   *el
	groupsMap    map[uint]int
	pool         []ZEntityID
	id           ZEntityID
	groups       []*ZGroup
	context      *ZContext
	componentMap map[uint]Component
}

// NewEntityManager creates a new entity manager and
// returns it.
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

// ResetAll deletes all entities in the
// current entity manager, resets the
// id pool to zero and unregisters all components
// and groups.
func (e *ZEntityManager) ResetAll() {
	e.Reset()
	e.groupsMap = make(map[uint]int, 0)
	e.groups = nil
	e.componentMap = make(map[uint]Component)
}

// Reset deletes all entities in the
// current entity manager and resets the
// id pool to zero.
func (e *ZEntityManager) Reset() {
	e.DeleteEntities()
	e.pool = e.pool[:0]
	e.id = 0
}

// DeleteEntities deletes all active entities in
// the current entity manager. Each entity and all
// components associated with that entity are deleted.
func (e *ZEntityManager) DeleteEntities() {
	for _, id := range e.Entities() {
		e.DeleteEntity(id)
	}
}

// CreateEntity creates a new entity id and returns it.
// If an entity was previously deleted it will use the cached
// entity id and use that one without incrementing the current id.
func (e *ZEntityManager) CreateEntity() ZEntityID {
	id := e.getID()
	if !e.entityList.AddEntity(id) {
		panic("what?!")
	}
	return id
}

// HasEntity returns true if the entity id exists or
// false if it does not.
func (e *ZEntityManager) HasEntity(id ZEntityID) bool {
	return e.entityList.HasEntity(id)
}

// DeleteEntity deletes all components associated with
// the entity id and then deletes the entity from the current
// entity manager.
func (e *ZEntityManager) DeleteEntity(id ZEntityID) {
	if e.HasEntity(id) {
		e.componentDeleteEntity(id)
		e.deleteEntity(id)
	}
}

// Entities returns a slice of currently active entity ids.
func (e *ZEntityManager) Entities() []ZEntityID {
	return e.entityList.Entities()
}

// RegisterComponent registers an implementation of a component. This
// component, in most cases unless it is written by
// a third pary; will store of all the data of a certain type
// for all entities you have attached data of that type to.
//
// If the given key retrieves a component, this method will
// panic with ErrComponentAlreadyRegistered. If the given key could
// not retrieve a component it will be added to the current entity manager as
// long as the component is not nil. Otherwise, it will panic with
// ErrComponentNil.
func (e *ZEntityManager) RegisterComponent(key uint, c Component) *ZContext {
	if _, exist := e.componentMap[key]; !exist {
		if c != nil {
			e.componentMap[key] = c
		} else {
			panic(ErrComponentNil)
		}
	} else {
		panic(ErrComponentAlreadyRegistered)
	}
	return e.context
}

// Component returns the component interface which
// holds an implementation of a component. This
// component, in most cases unless it is written by
// a third pary; will store of all the data of a certain type
// for all entities you have attached data of that type to.
//
// If the component with the given key could not be
// retrieved then this method will panic with
// ErrUnregisteredComponent.
func (e *ZEntityManager) Component(key uint) Component {
	c, ok := e.componentMap[key]
	if !ok {
		panic(ErrUnregisteredComponent)
	}
	return c
}

// NewCollector returns a new collector for the given event triggers.
//
// You can use the built-in Added(Component...), Updated(Component...), Deleted(Component...) triggers in most cases,
// but you can also build custom triggers using the NewEventTrigger(Matcher, GroupEvent) function.
//
// As an example, this is useful for when you need to collect an entity that has ComponentA but not ComponentB and
// you want to collect the entity when it has been added to the group.
func (e *ZEntityManager) NewCollector(et ...*ZEventTrigger) *ZCollector {
	var collectorEvent []ZCollectorEvent
	for _, v := range et {
		collectorEvent = append(collectorEvent,
			newCollectorEvent(
				e.Group(v.Matcher()),
				v.GroupEvent()))
	}
	return newCollector(collectorEvent...)
}

// GroupCount returns the number of groups used by the current entity manager.
func (e *ZEntityManager) GroupCount() int {
	return len(e.groups)
}

// Group returns a new group from a matcher or returns an existing one using the matcher hash.
func (e *ZEntityManager) Group(m *ZMatcher) *ZGroup {

	// Fetch existing group if it exists
	if g := e.group(m); g != nil {
		return g
	}

	// Create new group
	ng := e.newGroup(m)
	e.addGroup(ng)

	return ng
}
