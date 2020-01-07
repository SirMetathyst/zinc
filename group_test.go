package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc/kit"
	"github.com/stretchr/testify/assert"
	"github.com/SirMetathyst/zinc"
)

func TestGroupHandleEntitySilently(t *testing.T) {

	t.Run("handle entity silently", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2(id, kit.LocalPosition2Data{X: 10, Y: 10})

		// Act
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))

		// Assert
		assert.Equal(t, 0, len(g.Entities()), "group should not contain any entities")
	})

	t.Run("handle entity silently", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))

		// Assert
		assert.Equal(t, 1, len(g.Entities()), "group should contain 1 entity")
	})
}

func TestGroupHandleEntity(t *testing.T) {

	t.Run("handle entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2(id, kit.LocalPosition2Data{X: 10, Y: 10})
		
		// Assert
		assert.Equal(t, 0, len(g.Entities()), "group should not contain any entities")
	})

	t.Run("handle entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)
		
		// Arrange
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Assert
		assert.Equal(t, 1, len(g.Entities()), "group should contain 1 entity")
	})
}