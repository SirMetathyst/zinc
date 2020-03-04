package zinc

// ZCollector ...
type ZCollector struct {
	collectorEvent []ZCollectorEvent
	entityList     *el
}

// ZCollectorEvent ...
type ZCollectorEvent struct {
	groupEvent GroupEvent
	group      *ZGroup
}

// Group ...
func (e *ZCollectorEvent) Group() *ZGroup {
	return e.group
}

// GroupEvent ...
func (e *ZCollectorEvent) GroupEvent() GroupEvent {
	return e.groupEvent
}

func newCollectorEvent(g *ZGroup, e GroupEvent) ZCollectorEvent {
	if g == nil {
		panic("group cannot be nil")
	}
	if !GroupEventValid(e) {
		panic("invalid group event")
	}
	return ZCollectorEvent{group: g, groupEvent: e}
}

// TODO: Write TEST
func newCollector(ce ...ZCollectorEvent) *ZCollector {
	if len(ce) == 0 || ce == nil {
		panic("collector event cannot be nil and must have at least one event")
	}
	collector := &ZCollector{
		collectorEvent: ce,
		entityList:     newEntityList(),
	}
	collector.activate()
	return collector
}

func (c *ZCollector) addEntity(key uint, id EntityID) {
	c.entityList.AddEntity(id)
}

// activate ...
func (c *ZCollector) activate() {
	for _, collectorEvent := range c.collectorEvent {
		group := collectorEvent.Group()
		groupEvent := collectorEvent.GroupEvent()
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
func (c *ZCollector) ClearCollectedEntities() {
	for _, id := range c.entityList.Entities() {
		c.entityList.DeleteEntity(id)
	}
}

// Entities ...
// TODO: Write TEST
func (c *ZCollector) Entities() []EntityID {
	return c.entityList.Entities()
}
