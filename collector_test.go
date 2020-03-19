package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/stretchr/testify/assert"
)

func TestCollectorAdded(t *testing.T) {

	// Arrange
	e := zinc.NewEntityManager()
	kit.RegisterLocalPosition2ComponentWith(e)
	c := e.NewCollector(zinc.Added(kit.ZLocalPosition2))

	// Act
	id := e.CreateEntity()
	kit.AddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})

	// Assert
	assert.Contains(t, c.Entities(), id, "collector must contain id")
}

func TestCollectorUpdated(t *testing.T) {

	// Arrange
	e := zinc.NewEntityManager()
	kit.RegisterLocalPosition2ComponentWith(e)
	c := e.NewCollector(zinc.Updated(kit.ZLocalPosition2))

	// Act
	id := e.CreateEntity()
	kit.MustAddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})

	// Assert
	assert.NotContains(t, c.Entities(), id, "collector must not contain id because the component wasn't updated, but added")

	// Act
	kit.MustUpdateLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 12, Y: 12})

	// Assert
	assert.Contains(t, c.Entities(), id, "collector must contain id because the component was updated")
}

func TestCollectorDeleted(t *testing.T) {

	// Arrange
	e := zinc.NewEntityManager()
	kit.RegisterLocalPosition2ComponentWith(e)
	c := e.NewCollector(zinc.Deleted(kit.ZLocalPosition2))

	// Act
	id := e.CreateEntity()
	kit.MustAddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})

	// Assert
	assert.NotContains(t, c.Entities(), id, "collector must not contain id because the component wasn't deleted, but added")

	// Act
	kit.MustDeleteLocalPosition2X(e, id)

	// Assert
	assert.Contains(t, c.Entities(), id, "collector must contain id because the component was deleted")
}

func TestCollectorClearCollectedEntities(t *testing.T) {

	// Arrange
	e := zinc.NewEntityManager()
	kit.RegisterLocalPosition2ComponentWith(e)
	c := e.NewCollector(zinc.Added(kit.ZLocalPosition2))

	// Act
	id1 := e.CreateEntity()
	kit.MustAddLocalPosition2X(e, id1, kit.ZLocalPosition2Data{X: 10, Y: 10})

	id2 := e.CreateEntity()
	kit.MustAddLocalPosition2X(e, id2, kit.ZLocalPosition2Data{X: 10, Y: 10})

	// Assert
	assert.Equal(t, 2, len(c.Entities()), "collector must have 2 entities in entities slice")

	// Act
	c.ClearCollectedEntities()

	// Assert
	assert.Equal(t, 0, len(c.Entities()), "collector must not have any entities in entities slice because ClearCollectedEntities was called")
}
