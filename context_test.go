package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func componentFunc(rkey *uint, rid *EntityID) func(key uint, id EntityID) {
	return func(key uint, id EntityID) {
		*rid = id
		*rkey = key
	}
}

func hasEntityFunc(rid *EntityID) func(id EntityID) bool {
	return func(id EntityID) bool {
		*rid = id
		return true
	}
}


func TestNewContext(t *testing.T) {
	
	// Arrange, Assert
	ctx := newContext(nil, nil, nil, nil)

	// Assert
	assert.NotNil(t, ctx, "must not return nil")
}

func TestContextComponentAdded(t *testing.T) {
	
	// Arrange
	returnKey := uint(0)
	returnID := EntityID(0)

	expectedKey := uint(20)
	expectedID := EntityID(10)

	ctx := newContext(componentFunc(&returnKey, &returnID), nil, nil, nil)

	// Act
	ctx.ComponentAdded(expectedKey, expectedID)

	// Assert
	assert.Equal(t, expectedKey, returnKey, "returned key does match expected key")
	assert.Equal(t, expectedID, returnID, "returned id does match expected id")
}

func TestContextComponentUpdated(t *testing.T) {
	
	// Arrange
	returnKey := uint(0)
	returnID := EntityID(0)

	expectedKey := uint(20)
	expectedID := EntityID(10)

	ctx := newContext(nil, nil, componentFunc(&returnKey, &returnID), nil)

	// Act
	ctx.ComponentUpdated(expectedKey, expectedID)

	// Assert
	assert.Equal(t, expectedKey, returnKey, "returned key does match expected key")
	assert.Equal(t, expectedID, returnID, "returned id does match expected id")
}


func TestContextComponentDeleted(t *testing.T) {
	
	// Arrange
	returnKey := uint(0)
	returnID := EntityID(0)

	expectedKey := uint(20)
	expectedID := EntityID(10)

	ctx := newContext(nil, componentFunc(&returnKey, &returnID), nil, nil)

	// Act
	ctx.ComponentDeleted(expectedKey, expectedID)

	// Assert
	assert.Equal(t, expectedKey, returnKey, "returned key does match expected key")
	assert.Equal(t, expectedID, returnID, "returned id does match expected id")
}

func TestContextHasEntity(t *testing.T) {
	
	// Arrange
	returnID := EntityID(0)
	returnState := false

	expectedID := EntityID(10)
	expectedState := true

	ctx := newContext(nil, nil, nil, hasEntityFunc(&returnID))

	// Act
	returnState = ctx.HasEntity(expectedID)

	// Assert
	assert.Equal(t, expectedState, returnState, "return state does match expected return state")
	assert.Equal(t, expectedID, returnID, "returned id does match expected id")
}
