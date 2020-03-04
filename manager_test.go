package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/SirMetathyst/zinc/test"
	"github.com/stretchr/testify/assert"
)

func TestNewEntityManager(t *testing.T) {
	// Arrange, Act
	e := zinc.NewEntityManager()
	// Assert
	assert.NotNil(t, e, "must not return nil")
}

func TestEntityManagerCreateEntity(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerHasEntity(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerDeleteEntity(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.DeleteEntity(t, e, []zinc.EntityID{1, 2, 3})
	test.DoesNotHaveEntities(t, e, []zinc.EntityID{1, 2, 3})
	test.HasEntities(t, e, []zinc.EntityID{4, 5})
	test.CreateEntities(t, e, []zinc.EntityID{3, 2, 1})
}

func TestEntityManagerDeleteEntities(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	e.DeleteEntities()
	test.DoesNotHaveEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.CreateEntities(t, e, []zinc.EntityID{5, 4, 3})
}

func TestEntityManagerEntities(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.Entities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerReset(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	e.Reset()
	test.DoesNotHaveEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerRegisterComponent(t *testing.T) {

	t.Run("register component type once", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()

		// Act
		do := func() {
			kit.AddLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 20})
		}

		// Assert
		assert.NotPanics(t, do, "should not panic because component type has been registered")
	})

	t.Run("register component type twice", func(t *testing.T) {

		// Arrange
		e := zinc.NewEntityManager()

		// Act
		do := func() { kit.RegisterLocalPosition2ComponentWith(e) }
		do()

		// Assert
		assert.Panics(t, do, "must panic if you try to register a component with the same key again")
	})
}

func TestEntityManagerResetAll(t *testing.T) {

	// Setup
	e := zinc.NewEntityManager()
	kit.RegisterLocalPosition2ComponentWith(e)

	// Arrange
	id := e.CreateEntity()

	// Act
	e.ResetAll()
	do := func() { kit.AddLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 20}) }

	// Assert
	assert.Panics(t, do, "should panic because component type has been deleted due to reset all")
}

func TestEntityManagerGroup(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	kit.RegisterLocalPosition2ComponentWith(e)
	// Assert
	test.Group(t, e)
	test.GroupCount(t, e, 1)
	test.Group(t, e)
	test.GroupCount(t, e, 1)
}

func TestEntityManagerCollector(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Arrange, Act
	c := e.NewCollector(zinc.Added(kit.LocalPosition2Key))
	// Assert
	assert.NotNil(t, c, "entity manager must not return nil collector")
}
