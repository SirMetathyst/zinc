package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTrigger(t *testing.T) {

	t.Run("new event trigger", func(t *testing.T) {
		et := NewEventTrigger(AllOf(0), GroupEventAdded)
		assert.NotNil(t, et, "event trigger must not be nil")
	})

	t.Run("new event trigger", func(t *testing.T) {
		et := NewEventTrigger(nil, GroupEventAdded)
		assert.NotNil(t, et, "event trigger must not be nil")
	})

	t.Run("new event trigger", func(t *testing.T) {
		et := NewEventTrigger(nil, 20)
		assert.NotNil(t, et, "event trigger must not be nil")
	})

	t.Run("new event trigger", func(t *testing.T) {
		et := NewEventTrigger(AllOf(0), 20)
		assert.NotNil(t, et, "event trigger must not be nil")
	})
}
