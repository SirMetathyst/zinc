package atom

// C ...
type C interface {
	Entities() []EntityID
	ClearCollectedEntities()
}

type c struct {
	group      []G
	groupEvent []GroupEvent
	entityList *EntityList
}

func newCollector(group []G, groupEvent []GroupEvent) *c {
	collector := &c{
		group:      group,
		groupEvent: groupEvent,
		entityList: NewEntityList(),
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
func (c *c) ClearCollectedEntities() {
	for _, id := range c.entityList.Entities() {
		c.entityList.DeleteEntity(id)
	}
}

// Entities ...
func (c *c) Entities() []EntityID {
	return c.entityList.Entities()
}
