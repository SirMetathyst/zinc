package test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/stretchr/testify/assert"
)

// Group ...
func Group(t *testing.T, e *zinc.ZEntityManager) {
	// Arrange, Act
	g := e.Group(zinc.AllOf(kit.ZLocalPosition2))
	// Assert
	assert.NotNil(t, g, "entity manager must not return nil group")
	assert.Equal(t, 1, e.GroupCount(), "group count must be equal to 1")
}

// GroupGlobal ...
func GroupGlobal(t *testing.T) {
	// Arrange, Act
	g := zinc.Group(zinc.AllOf(kit.ZLocalPosition2))
	// Assert
	assert.NotNil(t, g, "entity manager must not return nil group")
	assert.Equal(t, 1, zinc.GroupCount(), "group count must be equal to 1")
}

// GroupCount ...
func GroupCount(t *testing.T, e *zinc.ZEntityManager, groupCount int) {
	// Assert
	assert.Equalf(t, groupCount, e.GroupCount(), "group count does not match %d", groupCount)
}

// GroupCountGlobal ...
func GroupCountGlobal(t *testing.T, expectedGroupCount int) {
	// Assert
	assert.Equalf(t, expectedGroupCount, zinc.GroupCount(), "group count does not match %d", expectedGroupCount)
}

// CreateEntities ...
func CreateEntities(t *testing.T, e *zinc.ZEntityManager, n []zinc.EntityID) {
	for _, expected := range n {
		// Act
		actual := e.CreateEntity()
		// Assert
		assert.Equal(t, expected, actual, "created entity id does match expected id")
	}
}

// CreateEntitiesGlobal ...
func CreateEntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	for _, expected := range n {
		// Act
		actual := zinc.CreateEntity()
		// Assert
		assert.Equal(t, expected, actual, "created entity id does match expected id")
	}
}

// DeleteEntity ...
func DeleteEntity(t *testing.T, e *zinc.ZEntityManager, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		e.DeleteEntity(id)
		// Assert
		assert.NotContains(t, e.Entities(), id, "entity manager must not contain id")
	}
}

// DeleteEntityGlobal ...
func DeleteEntityGlobal(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		zinc.DeleteEntity(id)
		// Assert
		assert.NotContains(t, zinc.Entities(), id, "entity manager must not contain id")
	}
}

// Entities ...
func Entities(t *testing.T, e *zinc.ZEntityManager, n []zinc.EntityID) {
	// Assert
	assert.ElementsMatch(t, e.Entities(), n, "entity manager must contain ids")
}

// EntitiesGlobal ...
func EntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	// Assert
	assert.ElementsMatch(t, zinc.Entities(), n, "entity manager must contain ids")
}

// HasEntities ...
func HasEntities(t *testing.T, e *zinc.ZEntityManager, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := e.HasEntity(id)
		// Assert
		assert.True(t, has, "entity manager must have id")
	}
}

// HasEntitiesGlobal ...
func HasEntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := zinc.HasEntity(id)
		// Assert
		assert.True(t, has, "entity manager must have id")
	}
}

// DoesNotHaveEntities ...
func DoesNotHaveEntities(t *testing.T, e *zinc.ZEntityManager, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := e.HasEntity(id)
		// Assert
		assert.False(t, has, "entity manager must not have id")
	}
}

// DoesNotHaveEntitiesGlobal ...
func DoesNotHaveEntitiesGlobal(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := zinc.HasEntity(id)
		// Assert
		assert.False(t, has, "entity manager must not have id")
	}
}
