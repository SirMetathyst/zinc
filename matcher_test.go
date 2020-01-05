package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/stretchr/testify/assert"
	"github.com/SirMetathyst/zinckit"
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
	m1 := zinc.AllOf(zinckit.LocalPosition2Key, zinckit.Velocity2Key)
	m2 := zinc.AllOf(zinckit.Velocity2Key, zinckit.LocalPosition2Key)

	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherNoneOfHash(t *testing.T) {

	// Arrange, Act
	m1 := zinc.NoneOf(zinckit.LocalPosition2Key, zinckit.Velocity2Key)
	m2 := zinc.NoneOf(zinckit.Velocity2Key, zinckit.LocalPosition2Key)

	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherHash(t *testing.T) {

	// Arrange, Act
	m1 := zinc.AllOf(zinckit.LocalPosition2Key, zinckit.Velocity2Key).
		NoneOf(zinckit.LocalRotation2Key, zinckit.LocalScale2Key)

	m2 := zinc.AllOf(zinckit.Velocity2Key, zinckit.LocalPosition2Key).
		NoneOf(zinckit.LocalScale2Key, zinckit.LocalRotation2Key)
	
	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}


func TestMatcherMatch(t *testing.T) {

	t.Run("non-existing key", func(t *testing.T) {

		// Reset
		zinc.Reset()

		// Arrange
		id := zinc.CreateEntity()
		zinckit.SetLocalPosition2(id, zinckit.LocalPosition2Data{X: 10, Y: 10})
		zinckit.SetLocalRotation2(id, zinckit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(0)
		mv := m.Match(zinc.Default(), id)

		// Assert
		assert.False(t, mv, "must return false if matcher contains non-existing key")
	})

	t.Run("all of", func(t *testing.T) {

		// Reset
		zinc.Reset()

		// Arrange
		id := zinc.CreateEntity()
		zinckit.SetLocalPosition2(id, zinckit.LocalPosition2Data{X: 10, Y: 10})
		zinckit.SetLocalRotation2(id, zinckit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(zinckit.LocalPosition2Key, zinckit.LocalRotation2Key)
		mv := m.Match(zinc.Default(), id)

		// Assert
		assert.True(t, mv, "must return true if matcher contains all of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Reset
		zinc.Reset()

		// Arrange
		id := zinc.CreateEntity()
		zinckit.SetLocalPosition2(id, zinckit.LocalPosition2Data{X: 10, Y: 10})

		// Act
		m := zinc.NoneOf(zinckit.LocalRotation2Key)
		mv := m.Match(zinc.Default(), id)

		// Assert
		assert.True(t, mv, "must return true because matcher should not contain any of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Reset
		zinc.Reset()

		// Arrange
		id := zinc.CreateEntity()
		zinckit.SetLocalPosition2(id, zinckit.LocalPosition2Data{X: 10, Y: 10})
		zinckit.SetLocalRotation2(id, zinckit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.NoneOf(zinckit.LocalRotation2Key)
		mv := m.Match(zinc.Default(), id)

		// Assert
		assert.False(t, mv, "must return false because matcher contains some of given keys")
	})

	t.Run("all of/none of", func(t *testing.T) {

		// Setup
		zinc.Reset()

		// Arrange
		id := zinc.CreateEntity()
		zinckit.SetLocalPosition2(id, zinckit.LocalPosition2Data{X: 10, Y: 10})
		zinckit.SetLocalRotation2(id, zinckit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(zinckit.LocalPosition2Key).NoneOf(zinckit.LocalRotation2Key)
		mv := m.Match(zinc.Default(), id)

		// Assert
		assert.False(t, mv, "must return false to satisfy matcher")
	})
}
