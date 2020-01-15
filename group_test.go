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
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		called := false

		// Act
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))
		g.HandleEntityAdded(func (key uint, id zinc.EntityID){
			called = true
		})

		// Assert
		assert.Equal(t, 0, len(g.Entities()), "group should not contain any entities")
		assert.False(t, called, "handle entity added should not be called")
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
		called := false

		// Act
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))
		g.HandleEntityAdded(func (key uint, id zinc.EntityID){
			called = true
		})

		// Assert
		assert.Equal(t, 1, len(g.Entities()), "group should contain 1 entity")
		assert.False(t, called, "handle entity added should not be called")
	})
}

func TestGroupHandleEntity(t *testing.T) {

	t.Run("handle entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		called := false
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))
		g.HandleEntityAdded(func (key uint, id zinc.EntityID){
			called = true
		})

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		
		// Assert
		assert.Equal(t, 0, len(g.Entities()), "group should not contain any entities")
		assert.False(t, called, "handle entity added should not be called")
	})

	t.Run("handle entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)
		
		// Arrange
		called := false
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key))
		g.HandleEntityAdded(func (key uint, id zinc.EntityID){
			called = true
		})

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Assert
		assert.Equal(t, 1, len(g.Entities()), "group should contain 1 entity")
		assert.True(t, called, "handle entity added should be called")
	})
}

func TestGroupUpdateEntity(t *testing.T) {

	t.Run("update entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		called := false
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key))
		g.HandleEntityUpdated(func(key uint, id zinc.EntityID){
			called = true
		})

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 20})

		// Assert
		assert.Equal(t, 1, len(g.Entities()), "group should contain 1 entity")
		assert.True(t, called, "handle entity updated should have been called")
	})
}

func TestGroupDeleteEntity(t *testing.T) {

	t.Run("delete entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		called := false
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key))
		g.HandleEntityDeleted(func(key uint, id zinc.EntityID){
			called = true
		})

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.DeleteLocalPosition2X(e, id)

		// Assert
		assert.Equal(t, 0, len(g.Entities()), "group should not contain any entities")
		assert.True(t, called, "handle entity deleted should have been called")
	})
}

func TestGroupHasEntity(t *testing.T) {

	t.Run("has entity", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		g := e.Group(zinc.AllOf(kit.LocalPosition2Key))

		// Act
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		has := g.HasEntity(id)

		// Assert
		assert.Equal(t, 1, len(g.Entities()), "group should contain 1 entity")
		assert.True(t, has, "should return true")
	})
}