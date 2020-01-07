package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGroup(t *testing.T) {
	
	t.Run("new group", func(t *testing.T){
		// Arrange, Act
		g := newGroup(nil, nil)
		// Assert
		assert.NotNil(t, g, "new group must not return nil")
	})

	t.Run("new group", func(t *testing.T){
		// Arrange, Act
		g := newGroup(Default(), nil)
		// Assert
		assert.NotNil(t, g, "new group must not return nil")
	})

	t.Run("new group", func(t *testing.T){
		// Arrange, Act
		g := newGroup(nil, AllOf(0))
		// Assert
		assert.NotNil(t, g, "new group must not return nil")
	})

	t.Run("new group", func(t *testing.T){
		// Arrange, Act
		g := newGroup(Default(), AllOf(0))
		// Assert
		assert.NotNil(t, g, "new group must not return nil")
	})
}