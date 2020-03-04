package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func componentFunc(rkey *uint, rid *ZEntityID) func(key uint, id ZEntityID) {
	return func(key uint, id ZEntityID) {
		*rid = id
		*rkey = key
	}
}

func hasEntityFunc(rid *ZEntityID) func(id ZEntityID) bool {
	return func(id ZEntityID) bool {
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
	returnID := ZEntityID(0)

	expectedKey := uint(20)
	expectedID := ZEntityID(10)

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
	returnID := ZEntityID(0)

	expectedKey := uint(20)
	expectedID := ZEntityID(10)

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
	returnID := ZEntityID(0)

	expectedKey := uint(20)
	expectedID := ZEntityID(10)

	ctx := newContext(nil, componentFunc(&returnKey, &returnID), nil, nil)

	// Act
	ctx.ComponentDeleted(expectedKey, expectedID)

	// Assert
	assert.Equal(t, expectedKey, returnKey, "returned key does match expected key")
	assert.Equal(t, expectedID, returnID, "returned id does match expected id")
}

func TestContextHasEntity(t *testing.T) {

	// Arrange
	returnID := ZEntityID(0)
	returnState := false

	expectedID := ZEntityID(10)
	expectedState := true

	ctx := newContext(nil, nil, nil, hasEntityFunc(&returnID))

	// Act
	returnState = ctx.HasEntity(expectedID)

	// Assert
	assert.Equal(t, expectedState, returnState, "return state does match expected return state")
	assert.Equal(t, expectedID, returnID, "returned id does match expected id")
}
