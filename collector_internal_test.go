package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCollectorEventPanics(t *testing.T) {

	t.Run("new collector event with nil group panics", func(t *testing.T) {

		// Arrange, Act
		f := func() {
			newCollectorEvent(nil, 0)
		}

		// Assert
		assert.PanicsWithValue(t, ErrNilGroup, f, "newCollectorEvent must panic with ErrNilGroup when group given is nil")
	})

	t.Run("new collector event with invalid group event panics", func(t *testing.T) {

		// Arrange
		e := NewEntityManager()
		m := NewMatcher()
		g := newGroup(e, m)

		// Act
		f := func() {
			newCollectorEvent(g, 20)
		}

		// Assert
		assert.PanicsWithValue(t, ErrInvalidGroupEvent, f, "newCollectorEvent must panic with ErrInvalidGroupEvent when group event is invalid")
	})
}

func TestNewCollectorPanics(t *testing.T) {

	t.Run("new collector with nil collector event slice panics", func(t *testing.T) {

		// Arrange, Act
		f := func() {
			newCollector()
		}

		// Assert
		assert.PanicsWithValue(t, ErrNilCollectorEventSlice, f, "newCollector must panic with ErrNilCollectorEventSlice when collector event slice given is nil")
	})
}

func TestNewCollector(t *testing.T) {

	// Arrange, Act
	c := NewCollector(Added(0))

	// Assert
	assert.NotNil(t, c, "collector must not be nil")
}
