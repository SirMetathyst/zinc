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

// GroupEventValid ...
func GroupEventValid(e GroupEvent) bool {
	if e >= 0 || e <= 2 {
		return true
	}
	return false
}

// ZEventTrigger ...
type ZEventTrigger struct {
	matcher    *ZMatcher
	groupEvent GroupEvent
}

// Matcher ...
func (e *ZEventTrigger) Matcher() *ZMatcher {
	return e.matcher
}

// GroupEvent ...
func (e *ZEventTrigger) GroupEvent() GroupEvent {
	return e.groupEvent
}

// NewEventTrigger ...
func NewEventTrigger(m *ZMatcher, e GroupEvent) *ZEventTrigger {
	return &ZEventTrigger{matcher: m, groupEvent: e}
}

// Added ...
func Added(keys ...uint) *ZEventTrigger {
	return NewEventTrigger(AllOf(keys...), GroupEventAdded)
}

// Updated ...
func Updated(keys ...uint) *ZEventTrigger {
	return NewEventTrigger(AllOf(keys...), GroupEventUpdated)
}

// Deleted ...
func Deleted(keys ...uint) *ZEventTrigger {
	return NewEventTrigger(AllOf(keys...), GroupEventDeleted)
}
