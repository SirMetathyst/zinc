package zinc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTriggerMatcher(t *testing.T) {
	
	// Arrange
	m := AllOf(0, 1)
	et := Added(0, 1)

	// Act
	v := et.Matcher()

	// Assert
	assert.NotNil(t, v, "must not be nil")
	assert.Equal(t, v.Hash(), m.Hash(), "must return identical matcher")
}

func TestEventTriggerGroupEvent(t *testing.T) {
	
	// Arrange
	g := GroupEventAdded
	et := Added(0, 1)

	// Act
	v := et.GroupEvent()

	// Assert
	assert.Equal(t, v, g, "must return given group event")
}

func TestEventTriggerAdded(t *testing.T) {
	
	// Arrange
	m := AllOf(10, 20, 30)
	g := GroupEventAdded
	et := Added(10, 20, 30)

	// Act
	etm := et.Matcher()
	etg := et.GroupEvent()

	// Assert
	assert.Equal(t, etm.Hash(), m.Hash(), "must return matcher hash with given keys")
	assert.Equal(t, etg,  g, "must return group event added")
}

func TestEventTriggerUpdated(t *testing.T) {
	
	// Arrange
	m := AllOf(10, 20, 30)
	g := GroupEventUpdated
	et := Updated(10, 20, 30)

	// Act
	etm := et.Matcher()
	etg := et.GroupEvent()

	// Assert
	assert.Equal(t, etm.Hash(), m.Hash(), "must return matcher hash with given keys")
	assert.Equal(t, etg,  g, "must return group event updated")
}

func TestEventTriggerDeleted(t *testing.T) {
	
	// Arrange
	m := AllOf(10, 20, 30)
	g := GroupEventDeleted
	et := Deleted(10, 20, 30)

	// Act
	etm := et.Matcher()
	etg := et.GroupEvent()

	// Assert
	assert.Equal(t, etm.Hash(), m.Hash(), "must return matcher hash with given keys")
	assert.Equal(t, etg,  g, "must return group event deleted")
}