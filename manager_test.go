package atom_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomkit"
	"github.com/stretchr/testify/assert"
)

func group(t *testing.T) {
	// Arrange, Act
	g := atom.Group(atom.AllOf(atomkit.LocalPosition2Key))

	// Assert
	assert.NotNil(t, g, "entity manager must not return nil group")
	assert.Equal(t, atom.GroupCount(), 1, "group count must be equal to 1")
}

func groupCount(t *testing.T, groupCount int) {
	// Assert
	assert.Equalf(t, atom.GroupCount(), groupCount, "group count does not match %d", groupCount)
}

func createEntities(t *testing.T, n []atom.EntityID) {
	for _, nid := range n {
		// Act
		id := atom.CreateEntity()
		
		// Assert
		assert.Equal(t, id, nid, "created entity id does match expected id")
	}
}

func deleteEntity(t *testing.T, n []atom.EntityID) {
	for _, id := range n {
		// Act
		atom.DeleteEntity(id)

		// Assert
		assert.NotContains(t, atom.Entities(), id, "entity manager must not contain id")
	}
}

func entities(t *testing.T, n []atom.EntityID) {
	// Assert
	assert.ElementsMatch(t, atom.Entities(), n, "entity manager must contain ids")
}

func hasEntities(t *testing.T, n []atom.EntityID) {
	for _, id := range n {
		// Act
		has := atom.HasEntity(id)

		// Assert
		assert.True(t, has, "entity manager must have id")
	}
}

func doesNotHaveEntities(t *testing.T, n []atom.EntityID) {
	for _, id := range n {
		// Act
		has := atom.HasEntity(id)

		// Assert
		assert.False(t, has, "entity manager must not have id")
	}
}


func TestNewEntityManager(t *testing.T) {

	// Arrange, Act
	e := atom.NewEntityManager() 

	// Assert
	assert.NotNil(t, e, "must not return nil")
}

func TestEntityManagerCreateEntity(t *testing.T) {
	atom.Reset()
	createEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerHasEntity(t *testing.T) {
	atom.Reset()
	createEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
	hasEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerDeleteEntity(t *testing.T) {
	atom.Reset()
	createEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
	deleteEntity(t, []atom.EntityID{1, 2, 3})
	doesNotHaveEntities(t, []atom.EntityID{1, 2, 3})
	hasEntities(t, []atom.EntityID{4, 5})
}

func TestEntityManagerDeleteEntities(t *testing.T) {
	atom.Reset()
	createEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
	atom.DeleteEntities()
	doesNotHaveEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerEntities(t *testing.T) {
	atom.Reset()
	createEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
	entities(t, []atom.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerReset(t *testing.T) {
	atom.Reset()
	createEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
	hasEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
	atom.Reset()
	doesNotHaveEntities(t, []atom.EntityID{1, 2, 3, 4, 5})
}

func TestEntityManagerRegisterComponent(t *testing.T) {

	t.Run("register component type once", func(t *testing.T){

		// Arrange
		e := atom.NewEntityManager()
		cmp := atomkit.NewLocalPosition2Component()
		ctx := e.RegisterComponent(atomkit.LocalPosition2Key, cmp)
		cmp.SetContext(ctx)
		id := e.CreateEntity()

		// Act
		do := func(){ 
			atomkit.SetLocalPosition2X(e, id, atomkit.LocalPosition2Data{X: 10, Y: 20}) 
		}
		
		// Assert
		assert.NotPanics(t, do, "should not panic because component type has been registered")
	})

	t.Run("register component type twice", func(t *testing.T){

		// Arrange
		e := atom.NewEntityManager()

		do := func(){
			// Act
			cmp := atomkit.NewLocalPosition2Component()
			ctx := e.RegisterComponent(atomkit.LocalPosition2Key, cmp)
			cmp.SetContext(ctx)
		}
		
		do()

		// Assert
		assert.Panics(t, do, "must panic if you try to register a component with the same key again")
	})

}

func TestEntityManagerResetAll(t *testing.T) {

	// Arrange
	e := atom.NewEntityManager()
	cmp := atomkit.NewLocalPosition2Component()
	ctx := e.RegisterComponent(atomkit.LocalPosition2Key, cmp)
	cmp.SetContext(ctx)
	id := e.CreateEntity()

	// Act
	e.ResetAll()
	do := func(){ 
		atomkit.SetLocalPosition2X(e, id, atomkit.LocalPosition2Data{X: 10, Y: 20}) 
	}

	// Assert
	assert.Panics(t, do, "should panic because component type has been deleted due to reset all")
}


func TestEntityManagerGroup(t *testing.T) {
	atom.Reset()
	group(t)
	groupCount(t, 1)
	group(t)
	groupCount(t, 1)
}

func TestEntityManagerCollector(t *testing.T) {
	atom.Reset()

	// Arrange, Act
	c := atom.CreateCollector(atom.Added(atomkit.LocalPosition2Key))

	// Assert
	assert.NotNil(t, c, "entity manager must not return nil collector")
}