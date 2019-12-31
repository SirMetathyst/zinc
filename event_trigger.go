package atom

// GroupEvent ...
type GroupEvent uint

const (
	GroupEventAdded GroupEvent = iota
	GroupEventDeleted
	GroupEventUpdated
)

// EventTrigger ...
type EventTrigger struct {
	m *Matcher
	e GroupEvent
}

// Matcher ...
func (e *EventTrigger) Matcher() *Matcher {
	return e.m
}

// GroupEvent ...
func (e *EventTrigger) GroupEvent() GroupEvent {
	return e.e
}

// AddedX ...
func AddedX(e *EntityManager, keys ...uint) *EventTrigger {
	return &EventTrigger{m: AllOfX(e, keys...), e: GroupEventAdded}
}

// Added ...
func Added(keys ...uint) *EventTrigger {
	return AddedX(Default(), keys...)
}
