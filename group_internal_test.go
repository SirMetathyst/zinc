package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGroupPanics(t *testing.T) {

	t.Run("new group with nil entity manager panics", func(t *testing.T) {

		// Arrange, Act
		f := func() {
			newGroup(nil, nil)
		}

		// Assert
		assert.PanicsWithValue(t, ErrNilEntityManager, f, "newGroup must panic with ErrNilEntityManager when entity manager given is nil")
	})

	t.Run("new group with nil matcher panics", func(t *testing.T) {

		// Arrange, Act
		f := func() {
			newGroup(Default(), nil)
		}

		// Assert
		assert.PanicsWithValue(t, ErrNilMatcher, f, "newGroup must panic with ErrNilMatcher when matcher given is nil")

	})
}

func TestNewGroup(t *testing.T) {

	// Arrange, Act
	e := NewEntityManager()
	g := newGroup(e, AllOf(0))

	// Assert
	assert.NotNil(t, g, "new group must not return nil")
}
