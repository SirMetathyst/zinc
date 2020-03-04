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

// Group returns the group of this collector
// event.
func (e *ZCollectorEvent) Group() *ZGroup {
	return e.group
}

// GroupEvent returns the group event of this collector
// event.
func (e *ZCollectorEvent) GroupEvent() GroupEvent {
	return e.groupEvent
}

// ClearCollectedEntities deletes all the entity ids
// from this collector.
func (c *ZCollector) ClearCollectedEntities() {
	for _, id := range c.entityList.Entities() {
		c.entityList.DeleteEntity(id)
	}
}

// Entities returns a slice of entity ids
// associated with this collector.
func (c *ZCollector) Entities() []ZEntityID {
	return c.entityList.Entities()
}
