package zinc

import "errors"

var (
	// ErrInvalidGroupEvent ...
	ErrInvalidGroupEvent = errors.New("zinc: invalid group event")
)

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

// GroupEventValid returns true
// if the group event holds one of the
// three valid values and false if it does
// not.
func GroupEventValid(e GroupEvent) bool {
	if e >= 0 && e <= 2 {
		return true
	}
	return false
}

// ZEventTrigger ...
type ZEventTrigger struct {
	matcher    *ZMatcher
	groupEvent GroupEvent
}

// Matcher returns the matcher of this trigger.
func (e *ZEventTrigger) Matcher() *ZMatcher {
	return e.matcher
}

// GroupEvent returns the group event of this trigger.
func (e *ZEventTrigger) GroupEvent() GroupEvent {
	return e.groupEvent
}

// NewEventTrigger creates a new event trigger and returns it.
// Panics if the matcher is nil or the group event is not one of
// GroupEventAdded|GroupEventDeleted|GroupEventUpdated
func NewEventTrigger(m *ZMatcher, e GroupEvent) *ZEventTrigger {
	if m == nil {
		panic(ErrNilMatcher)
	}
	if !GroupEventValid(e) {
		panic(ErrInvalidGroupEvent)
	}
	return &ZEventTrigger{matcher: m, groupEvent: e}
}

// Added returns an event trigger with
// a matcher that includes the given
// component keys for matching an entity.
// It triggers when all of the given
// keys are added to the matcher.
func Added(keys ...uint) *ZEventTrigger {
	return NewEventTrigger(AllOf(keys...), GroupEventAdded)
}

// Updated returns an event trigger with
// a matcher that includes the given
// component keys for matching an entity.
// It triggers when any one of the given
// keys are updated in the matcher.
func Updated(keys ...uint) *ZEventTrigger {
	return NewEventTrigger(AllOf(keys...), GroupEventUpdated)
}

// Deleted returns an event trigger with
// a matcher that includes the given
// component keys for matching an entity.
// It triggers when any one of the given
// keys are deleted from the matcher.
func Deleted(keys ...uint) *ZEventTrigger {
	return NewEventTrigger(AllOf(keys...), GroupEventDeleted)
}
