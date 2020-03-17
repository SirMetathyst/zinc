package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventTriggerPanics(t *testing.T) {

	t.Run("new event trigger with nil matcher panics", func(t *testing.T) {

		// Arrange, Act
		f := func() {
			NewEventTrigger(nil, GroupEventAdded)
		}

		// Assert
		assert.PanicsWithValue(t, ErrNilMatcher, f, "NewEventTrigger must panic with ErrNilMatcher when matcher given is nil")
	})

	t.Run("new event trigger with invalid group event panics", func(t *testing.T) {

		// Arrange, Act
		f := func() {
			NewEventTrigger(AllOf(0), 20)
		}

		// Assert
		assert.PanicsWithValue(t, ErrInvalidGroupEvent, f, "NewEventTrigger must panic with ErrInvalidGroupEvent when group event given is invalid")
	})
}

func TestNewEventTrigger(t *testing.T) {

	// Arrange, Act
	et := NewEventTrigger(AllOf(0), GroupEventAdded)

	// Assert
	assert.NotNil(t, et, "event trigger must not be nil")
}
