package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/stretchr/testify/assert"
)

func group(t *testing.T) {
	// Arrange, Act
	g := zinc.Group(zinc.AllOf(kit.LocalPosition2Key))

	// Assert
	assert.NotNil(t, g, "entity manager must not return nil group")
	assert.Equal(t, zinc.GroupCount(), 1, "group count must be equal to 1")
}

func groupCount(t *testing.T, groupCount int) {
	// Assert
	assert.Equalf(t, zinc.GroupCount(), groupCount, "group count does not match %d", groupCount)
}

func createEntities(t *testing.T, n []zinc.EntityID) {
	for _, nid := range n {
		// Act
		id := zinc.CreateEntity()
		
		// Assert
		assert.Equal(t, id, nid, "created entity id does match expected id")
	}
}

func deleteEntity(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		zinc.DeleteEntity(id)

		// Assert
		assert.NotContains(t, zinc.Entities(), id, "entity manager must not contain id")
	}
}

func entities(t *testing.T, n []zinc.EntityID) {
	// Assert
	assert.ElementsMatch(t, zinc.Entities(), n, "entity manager must contain ids")
}

func hasEntities(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := zinc.HasEntity(id)

		// Assert
		assert.True(t, has, "entity manager must have id")
	}
}

func doesNotHaveEntities(t *testing.T, n []zinc.EntityID) {
	for _, id := range n {
		// Act
		has := zinc.HasEntity(id)

		// Assert
		assert.False(t, has, "entity manager must not have id")
	}
}


func TestNewEntityManager(t *testing.T) {

	// Arrange, Act
	e := zinc.NewEntityManager() 

	// Assert
	assert.NotNil(t, e, "must not return nil")
}

func TestEntityManagerCreateEntity(t *testing.T) {
	zinc.Reset()
	createEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerHasEntity(t *testing.T) {
	zinc.Reset()
	createEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
	hasEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerDeleteEntity(t *testing.T) {
	zinc.Reset()
	createEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
	deleteEntity(t, []zinc.EntityID{1, 2, 3})
	doesNotHaveEntities(t, []zinc.EntityID{1, 2, 3})
	hasEntities(t, []zinc.EntityID{4, 5})
}

func TestEntityManagerDeleteEntities(t *testing.T) {
	zinc.Reset()
	createEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
	zinc.DeleteEntities()
	doesNotHaveEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerEntities(t *testing.T) {
	zinc.Reset()
	createEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
	entities(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerReset(t *testing.T) {
	zinc.Reset()
	createEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
	hasEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
	zinc.Reset()
	doesNotHaveEntities(t, []zinc.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerRegisterComponent(t *testing.T) {

	t.Run("register component type once", func(t *testing.T){

		// Arrange
		e := zinc.NewEntityManager()
		cmp := kit.NewLocalPosition2Component()
		ctx := e.RegisterComponent(kit.LocalPosition2Key, cmp)
		cmp.SetContext(ctx)
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

		do := func(){
			// Act
			cmp := kit.NewLocalPosition2Component()
			ctx := e.RegisterComponent(kit.LocalPosition2Key, cmp)
			cmp.SetContext(ctx)
		}
		
		do()

		// Assert
		assert.Panics(t, do, "must panic if you try to register a component with the same key again")
	})

}

func TestEntityManagerResetAll(t *testing.T) {

	// Arrange
	e := zinc.NewEntityManager()
	cmp := kit.NewLocalPosition2Component()
	ctx := e.RegisterComponent(kit.LocalPosition2Key, cmp)
	cmp.SetContext(ctx)
	id := e.CreateEntity()

	// Act
	e.ResetAll()
	do := func(){ 
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 20}) 
	}

	// Assert
	assert.Panics(t, do, "should panic because component type has been deleted due to reset all")
}


func TestEntityManagerGroup(t *testing.T) {
	zinc.Reset()
	group(t)
	groupCount(t, 1)
	group(t)
	groupCount(t, 1)
}

func TestEntityManagerCollector(t *testing.T) {
	zinc.Reset()

	// Arrange, Act
	c := zinc.CreateCollector(zinc.Added(kit.LocalPosition2Key))

	// Assert
	assert.NotNil(t, c, "entity manager must not return nil collector")
}