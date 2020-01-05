package zinc

// GroupEvent ...
type GroupEvent uint

const (
	// GroupEventAdded ...
	GroupEventAdded GroupEvent = iota
	// GroupEventDeleted ...
	GroupEventDeleted
	// GroupEventUpdated ...
	GroupEventUpdated
)

// ET ...
type ET interface {
	Matcher() M
	GroupEvent() GroupEvent
}

type et struct {
	m M
	e GroupEvent
}

// Matcher ...
// TODO: Write TEST
func (e *et) Matcher() M {
	return e.m
}

// GroupEvent ...
// TODO: Write TEST
func (e *et) GroupEvent() GroupEvent {
	return e.e
}

// TODO: Write TEST
func newEventTrigger(m M, e GroupEvent) ET {
	return &et{m: m, e: e}
}

// Added ...
// TODO: Write TEST
func Added(keys ...uint) ET {
	return newEventTrigger(AllOf(keys...), GroupEventAdded)
}

// Updated ...
// TODO: Write TEST
func Updated(keys ...uint) ET {
	return newEventTrigger(AllOf(keys...), GroupEventUpdated)
}

// Deleted ...
// TODO: Write TEST
func Deleted(keys ...uint) ET {
	return newEventTrigger(AllOf(keys...), GroupEventDeleted)
}