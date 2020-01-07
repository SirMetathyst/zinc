package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/stretchr/testify/assert"
	"github.com/SirMetathyst/zinc/kit"
)

var (
	matcherData = []struct {
		allOf  []uint
		noneOf []uint
	}{
		{allOf: []uint{1, 2, 3}, noneOf: []uint{0, 10, 20}},
		{allOf: []uint{0, 10, 2}, noneOf: []uint{99, 82, 10}},
	}
)

func TestMatcherHasAllOf(t *testing.T) {
	for _, ds := range matcherData {
		// Act
		m := zinc.AllOf(ds.allOf...)
		has := m.HasAllOf(ds.allOf...)
		// Assert
		assert.True(t, has, "has all of must return true")
	}
}

func TestMatcherAllOf(t *testing.T) {
	for _, ds := range matcherData {
		// Act
		m := zinc.AllOf(ds.allOf...)
		// Assert
		assert.ElementsMatch(t, m.AllOfSlice(), ds.allOf, "returned all of slice does not match input data")
	}
}

func TestMatcherHasNoneOf(t *testing.T) {
	for _, ds := range matcherData {
		// Act
		m := zinc.NoneOf(ds.noneOf...)
		has := m.HasNoneOf(ds.noneOf...)
		// Assert
		assert.True(t, has, "has none of must return true")
	}
}

func TestMatcherNoneOf(t *testing.T) {
	for _, ds := range matcherData {
		// Act
		m := zinc.NoneOf(ds.noneOf...)
		// Assert
		assert.ElementsMatch(t, m.NoneOfSlice(), ds.noneOf, "returned none of slice does not match input data")
	}
}


func TestMatcherAllOfHash(t *testing.T) {
	// Arrange, Act
	m1 := zinc.AllOf(kit.LocalPosition2Key, kit.Velocity2Key)
	m2 := zinc.AllOf(kit.Velocity2Key, kit.LocalPosition2Key)
	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherNoneOfHash(t *testing.T) {
	// Arrange, Act
	m1 := zinc.NoneOf(kit.LocalPosition2Key, kit.Velocity2Key)
	m2 := zinc.NoneOf(kit.Velocity2Key, kit.LocalPosition2Key)
	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherHash(t *testing.T) {
	// Arrange, Act
	m1 := zinc.AllOf(kit.LocalPosition2Key, kit.Velocity2Key).
		NoneOf(kit.LocalRotation2Key, kit.LocalScale2Key)

	m2 := zinc.AllOf(kit.Velocity2Key, kit.LocalPosition2Key).
		NoneOf(kit.LocalScale2Key, kit.LocalRotation2Key)
	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}


func TestMatcherMatch(t *testing.T) {
	
	t.Run("non-existing key", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(0)
		mv := m.Match(e, id)

		// Assert
		assert.False(t, mv, "must return false if matcher contains non-existing key")
	})

	t.Run("all of", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(kit.LocalPosition2Key, kit.LocalRotation2Key)
		mv := m.Match(e, id)

		// Assert
		assert.True(t, mv, "must return true if matcher contains all of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})

		// Act
		m := zinc.NoneOf(kit.LocalRotation2Key)
		mv := m.Match(e, id)

		// Assert
		assert.True(t, mv, "must return true because matcher should not contain any of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.NoneOf(kit.LocalRotation2Key)
		mv := m.Match(e, id)

		// Assert
		assert.False(t, mv, "must return false because matcher contains some of given keys")
	})

	t.Run("all of/none of", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.SetLocalPosition2X(e, id, kit.LocalPosition2Data{X: 10, Y: 10})
		kit.SetLocalRotation2X(e, id, kit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(kit.LocalPosition2Key).NoneOf(kit.LocalRotation2Key)
		mv := m.Match(e, id)

		// Assert
		assert.False(t, mv, "must return false to satisfy matcher")
	})
}
