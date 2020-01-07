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
func (e *et) Matcher() M {
	return e.m
}

// GroupEvent ...
func (e *et) GroupEvent() GroupEvent {
	return e.e
}

func newEventTrigger(m M, e GroupEvent) ET {
	return &et{m: m, e: e}
}

// Added ...
func Added(keys ...uint) ET {
	return newEventTrigger(AllOf(keys...), GroupEventAdded)
}

// Updated ...
func Updated(keys ...uint) ET {
	return newEventTrigger(AllOf(keys...), GroupEventUpdated)
}

// Deleted ...
func Deleted(keys ...uint) ET {
	return newEventTrigger(AllOf(keys...), GroupEventDeleted)
}