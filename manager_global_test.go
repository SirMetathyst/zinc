package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/SirMetathyst/zinc/test"
	"github.com/stretchr/testify/assert"
)

func TestEntityManagerCreateEntityGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerHasEntityGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerDeleteEntityGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.DeleteEntityGlobal(t, []zinc.EntityID{1, 2, 3})
	test.DoesNotHaveEntitiesGlobal(t, []zinc.EntityID{1, 2, 3})
	test.HasEntitiesGlobal(t, []zinc.EntityID{4, 5})
	test.CreateEntitiesGlobal(t, []zinc.EntityID{3, 2, 1})
}

func TestEntityManagerDeleteEntitiesGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	zinc.DeleteEntities()
	test.DoesNotHaveEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerEntitiesGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.EntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerResetGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	zinc.Reset()
	test.DoesNotHaveEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerRegisterComponentGlobal(t *testing.T) {

	t.Run("register component type once", func(t *testing.T) {

		// Setup
		zinc.ResetAll()
		kit.RegisterLocalPosition2Component()

		// Arrange
		id := zinc.CreateEntity()

		// Act
		do := func() {
			kit.AddLocalPosition2(id, kit.LocalPosition2Data{X: 10, Y: 20})
		}

		// Assert
		assert.NotPanics(t, do, "should not panic because component type has been registered")
	})

	t.Run("register component type twice", func(t *testing.T) {

		// Setup
		zinc.ResetAll()

		// Act
		do := func() { kit.RegisterLocalPosition2Component() }
		do()

		// Assert
		assert.Panics(t, do, "must panic if you try to register a component with the same key again")
	})
}

func TestEntityManagerResetAllGlobal(t *testing.T) {

	// Setup
	zinc.ResetAll()
	kit.RegisterLocalPosition2Component()

	// Arrange
	id := zinc.CreateEntity()

	// Act
	zinc.ResetAll()
	do := func() { kit.AddLocalPosition2(id, kit.LocalPosition2Data{X: 10, Y: 20}) }

	// Assert
	assert.Panics(t, do, "should panic because component type has been deleted due to reset all")
}

func TestEntityManagerGroupGlobal(t *testing.T) {
	// Setup
	zinc.ResetAll()
	kit.RegisterLocalPosition2Component()
	// Assert
	test.GroupGlobal(t)
	test.GroupCountGlobal(t, 1)
	test.GroupGlobal(t)
	test.GroupCountGlobal(t, 1)
}

func TestEntityManagerCollectorGlobal(t *testing.T) {
	// Setup
	zinc.ResetAll()
	// Arrange, Act
	c := zinc.NewCollector(zinc.Added(kit.LocalPosition2Key))
	// Assert
	assert.NotNil(t, c, "entity manager must not return nil collector")
}
