package atom

// C ...
type C interface {
	Entities() []EntityID
	ClearCollectedEntities()
}

type c struct {
	group      []G
	groupEvent []GroupEvent
	entityList *el
}

// NewCollector ...
// TODO: Write TEST
func NewCollector(group []G, groupEvent []GroupEvent) C {
	collector := &c{
		group:      group,
		groupEvent: groupEvent,
		entityList: newEntityList(),
	}
	collector.activate()
	return collector
}

func (c *c) addEntity(key uint, id EntityID) {
	c.entityList.AddEntity(id)
}

// activate ...
func (c *c) activate() {
	for i, group := range c.group {
		groupEvent := c.groupEvent[i]
		switch groupEvent {
		case GroupEventAdded:
			group.HandleEntityAdded(c.addEntity)
			break
		case GroupEventUpdated:
			group.HandleEntityUpdated(c.addEntity)
			break
		case GroupEventDeleted:
			group.HandleEntityDeleted(c.addEntity)
			break
		}
	}
}

// ClearCollectedEntities ...
// TODO: Write TEST
func (c *c) ClearCollectedEntities() {
	for _, id := range c.entityList.Entities() {
		c.entityList.DeleteEntity(id)
	}
}

// Entities ...
// TODO: Write TEST
func (c *c) Entities() []EntityID {
	return c.entityList.Entities()
}
