package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	entityListData = []EntityID {1, 2000, 20}
)

func TestNewEntityList(t *testing.T) {

	// Arrange, Act
	el := newEntityList()

	// Assert
	assert.NotNil(t, el, "must not return nil")
}

func TestEntityListAddEntity(t *testing.T) {

	// Arrange
	el := newEntityList()

	for _, v := range entityListData {

		// Act
		v1 := el.AddEntity(v)
		v2 := el.AddEntity(v)

		// Assert
		assert.Equal(t, true, v1, "adding a new entity that is not in the list should return true")
		assert.Equal(t, false, v2, "adding an entity that has already been added should return false")
	}
}

func TestEntityListDeleteEntity(t *testing.T) {

	// Arrange
	el := newEntityList()

	for _, va := range entityListData {
		el.AddEntity(va)
	}

	for _, v := range entityListData {

		// Act
		v1 := el.DeleteEntity(v)
		v2 := el.DeleteEntity(v)

		// Assert
		assert.Equal(t, true, v1, "deleting an entity that has been added previously will return true")
		assert.Equal(t, false, v2, "deleting an entity that does not exist should return false")
	}
}

func TestEntityListHasEntity(t *testing.T) {

	// Arrange
	el := newEntityList()

	for _, va := range entityListData {
		el.AddEntity(va)
	}

	for _, v := range entityListData {

		// Act
		has := el.HasEntity(v)

		// Assert
		assert.Equalf(t, true, has, "must return true for id: %d", v)
	}
}


func TestEntityListEntities(t *testing.T) {

	// Arrange
	el := newEntityList()

	for _, v := range entityListData {
		el.AddEntity(v)
	}

	// Act
	entities := el.Entities()

	// Assert
	assert.ElementsMatch(t, entities, entityListData, "returned entities slice does not match input data")
}