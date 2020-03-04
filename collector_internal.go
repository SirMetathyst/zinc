package zinc

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

func (c *ZCollector) addEntity(key uint, id ZEntityID) {
	c.entityList.AddEntity(id)
}

// activate ...
func (c *ZCollector) activate() {
	for _, collectorEvent := range c.collectorEvent {
		group := collectorEvent.Group()
		groupEvent := collectorEvent.GroupEvent()
		switch groupEvent {
		case GroupEventAdded:
			group.RegisterEntityAddedFunc(c.addEntity)
			break
		case GroupEventUpdated:
			group.RegisterEntityUpdatedFunc(c.addEntity)
			break
		case GroupEventDeleted:
			group.RegisterEntityDeletedFunc(c.addEntity)
			break
		}
	}
}
