package test

import (
	"testing"
	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/stretchr/testify/assert"
)

func Group(t *testing.T, e *zinc.EntityManager) {
	// Arrange, Act
	g := e.Group(zinc.AllOf(kit.LocalPosition2Key))
	// Assert
	assert.NotNil(t, g, "entity manager must not return nil group")
	assert.Equal(t, 1, e.GroupCount(), "group count must be equal to 1")
}

func GroupGlobal(t *testing.T) {
	// Arrange, Act
	g := zinc.Group(zinc.AllOf(kit.LocalPosition2Key))
	// Assert
	assert.NotNil(t, g, "entity manager must not return nil group")
	assert.Equal(t, 1, zinc.GroupCount(), "group count must be equal to 1")
}

func GroupCount(t *testing.T, e *zinc.EntityManager, groupCount int) {
	// Assert
	assert.Equalf(t, groupCount, e.GroupCount(), "group count does not match %d", groupCount)
}

func GroupCountGlobal(t *testing.T, groupCount int) {
	// Assert
	assert.Equalf(t, groupCount, zinc.GroupCount(), "group count does not match %d", groupCount)
}

func CreateEntities(t *testing.T, e *zinc.EntityManager, n []zinc.EntityID) {
	for _, nid := range n {
		// Act
		id := e.CreateEntity()
		// Assert
		assert.Equal(t, id, nid, "created entity id does match expected id")
	}
}

func CreateEntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	for _, nid := range n {
		// Act
		id := zinc.CreateEntity()
		// Assert
		assert.Equal(t, id, nid, "created entity id does match expected id")
	}
}

func DeleteEntity(t *testing.T, e *zinc.EntityManager, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		e.DeleteEntity(id)
		// Assert
		assert.NotContains(t, e.Entities(), id, "entity manager must not contain id")
	}
}

func DeleteEntityGlobal(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		zinc.DeleteEntity(id)
		// Assert
		assert.NotContains(t, zinc.Entities(), id, "entity manager must not contain id")
	}
}

func Entities(t *testing.T, e *zinc.EntityManager, n []zinc.EntityID) {
	// Assert
	assert.ElementsMatch(t, e.Entities(), n, "entity manager must contain ids")
}

func EntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	// Assert
	assert.ElementsMatch(t, zinc.Entities(), n, "entity manager must contain ids")
}

func HasEntities(t *testing.T, e *zinc.EntityManager, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := e.HasEntity(id)
		// Assert
		assert.True(t, has, "entity manager must have id")
	}
}

func HasEntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := zinc.HasEntity(id)
		// Assert
		assert.True(t, has, "entity manager must have id")
	}
}

func DoesNotHaveEntities(t *testing.T, e *zinc.EntityManager, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := e.HasEntity(id)
		// Assert
		assert.False(t, has, "entity manager must not have id")
	}
}

func DoesNotHaveEntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := zinc.HasEntity(id)
		// Assert
		assert.False(t, has, "entity manager must not have id")
	}
}