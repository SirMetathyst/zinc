package zinc_test

import (
	"testing"
	"github.com/SirMetathyst/zinc/test"
	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
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

func TestEntityManagerCreateEntityGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerHasEntity(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerHasEntityGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerDeleteEntity(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.DeleteEntity(t, e, []zinc.EntityID{1, 2, 3})
	test.DoesNotHaveEntities(t, e, []zinc.EntityID{1, 2, 3})
	test.HasEntities(t, e, []zinc.EntityID{4, 5})
}

func TestEntityManagerDeleteEntityGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.DeleteEntityGlobal(t, []zinc.EntityID{1, 2, 3})
	test.DoesNotHaveEntitiesGlobal(t, []zinc.EntityID{1, 2, 3})
	test.HasEntitiesGlobal(t, []zinc.EntityID{4, 5})
}

func TestEntityManagerDeleteEntities(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	e.DeleteEntities()
	test.DoesNotHaveEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerDeleteEntitiesGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	zinc.DeleteEntities()
	test.DoesNotHaveEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerEntities(t *testing.T) {
	// Setup
	e := zinc.NewEntityManager()
	// Assert
	test.CreateEntities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
	test.Entities(t, e, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerEntitiesGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.EntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
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

func TestEntityManagerResetGlobal(t *testing.T) {
	// Reset
	zinc.Reset()
	// Assert
	test.CreateEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	test.HasEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
	zinc.Reset()
	test.DoesNotHaveEntitiesGlobal(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerRegisterComponent(t *testing.T) {

	t.Run("register component type once", func(t *testing.T){

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()

		// Act
		do := func(){ 
			kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 20}) 
		}
		
		// Assert
		assert.NotPanics(t, do, "should not panic because component type has been registered")
	})

	t.Run("register component type twice", func(t *testing.T){

		// Arrange
		e := zinc.NewEntityManager()

		// Act
		do := func(){ kit.RegisterLocalPosition2ComponentWith(e) }
		do()

		// Assert
		assert.Panics(t, do, "must panic if you try to register a component with the same key again")
	})
}

func TestEntityManagerRegisterComponentGlobal(t *testing.T) {

	t.Run("register component type once", func(t *testing.T){

		// Setup
		zinc.ResetAll()
		kit.RegisterLocalPosition2Component()

		// Arrange
		id := zinc.CreateEntity()

		// Act
		do := func(){ 
			kit.SetLocalPosition2(id, kit.LocalPosition2Data{X: 10, Y: 20}) 
		}
		
		// Assert
		assert.NotPanics(t, do, "should not panic because component type has been registered")
	})

	t.Run("register component type twice", func(t *testing.T){

		// Setup
		zinc.ResetAll()

		// Act
		do := func(){ kit.RegisterLocalPosition2Component() }
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
	do := func(){ kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 20}) }

	// Assert
	assert.Panics(t, do, "should panic because component type has been deleted due to reset all")
}

func TestEntityManagerResetAllGlobal(t *testing.T) {

	// Setup
	zinc.ResetAll()
	kit.RegisterLocalPosition2Component()

	// Arrange
	id := zinc.CreateEntity()

	// Act
	zinc.ResetAll()
	do := func(){ kit.SetLocalPosition2(id, kit.LocalPosition2Data{X: 10, Y: 20}) }

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

func TestEntityManagerGroupGlobal(t *testing.T) {
	// Setup
	zinc.ResetAll()
	kit.RegisterLocalPosition2Component()
	// Assert
	test.GroupGlobal(t)
	test.GroupCountGlobal(t, 1)
	test.GroupGlobal(t)
	test.GroupCountGlobal(t,1)
}

func TestEntityManagerCollector(t *testing.T) {	
	// Setup
	e := zinc.NewEntityManager()
	// Arrange, Act
	c := e.CreateCollector(zinc.Added(kit.LocalPosition2Key))
	// Assert
	assert.NotNil(t, c, "entity manager must not return nil collector")
}

func TestEntityManagerCollectorGlobal(t *testing.T) {	
	// Setup
	zinc.ResetAll()
	// Arrange, Act
	c := zinc.CreateCollector(zinc.Added(kit.LocalPosition2Key))
	// Assert
	assert.NotNil(t, c, "entity manager must not return nil collector")
}