package atom_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
	"github.com/stretchr/testify/assert"
	"github.com/SirMetathyst/atomkit"
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
		m := atom.AllOf(ds.allOf...)
		has := m.HasAllOf(ds.allOf...)

		// Assert
		assert.True(t, has, "has all of must return true")
	}
}

func TestMatcherAllOf(t *testing.T) {
	for _, ds := range matcherData {

		// Act
		m := atom.AllOf(ds.allOf...)

		// Assert
		assert.ElementsMatch(t, m.AllOfSlice(), ds.allOf, "returned all of slice does not match input data")
	}
}

func TestMatcherHasNoneOf(t *testing.T) {
	for _, ds := range matcherData {

		// Act
		m := atom.NoneOf(ds.noneOf...)
		has := m.HasNoneOf(ds.noneOf...)

		// Assert
		assert.True(t, has, "has none of must return true")
	}
}

func TestMatcherNoneOf(t *testing.T) {
	for _, ds := range matcherData {

		// Act
		m := atom.NoneOf(ds.noneOf...)

		// Assert
		assert.ElementsMatch(t, m.NoneOfSlice(), ds.noneOf, "returned none of slice does not match input data")
	}
}


func TestMatcherAllOfHash(t *testing.T) {

	// Arrange, Act
	m1 := atom.AllOf(atomkit.LocalPosition2Key, atomkit.Velocity2Key)
	m2 := atom.AllOf(atomkit.Velocity2Key, atomkit.LocalPosition2Key)

	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherNoneOfHash(t *testing.T) {

	// Arrange, Act
	m1 := atom.NoneOf(atomkit.LocalPosition2Key, atomkit.Velocity2Key)
	m2 := atom.NoneOf(atomkit.Velocity2Key, atomkit.LocalPosition2Key)

	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherHash(t *testing.T) {

	// Arrange, Act
	m1 := atom.AllOf(atomkit.LocalPosition2Key, atomkit.Velocity2Key).
		NoneOf(atomkit.LocalRotation2Key, atomkit.LocalScale2Key)

	m2 := atom.AllOf(atomkit.Velocity2Key, atomkit.LocalPosition2Key).
		NoneOf(atomkit.LocalScale2Key, atomkit.LocalRotation2Key)
	
	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}


func TestMatcherMatch(t *testing.T) {

	t.Run("non-existing key", func(t *testing.T) {

		// Reset
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomkit.SetLocalPosition2(id, atomkit.LocalPosition2Data{X: 10, Y: 10})
		atomkit.SetLocalRotation2(id, atomkit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := atom.AllOf(0)
		mv := m.Match(atom.Default(), id)

		// Assert
		assert.False(t, mv, "must return false if matcher contains non-existing key")
	})

	t.Run("all of", func(t *testing.T) {

		// Reset
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomkit.SetLocalPosition2(id, atomkit.LocalPosition2Data{X: 10, Y: 10})
		atomkit.SetLocalRotation2(id, atomkit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := atom.AllOf(atomkit.LocalPosition2Key, atomkit.LocalRotation2Key)
		mv := m.Match(atom.Default(), id)

		// Assert
		assert.True(t, mv, "must return true if matcher contains all of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Reset
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomkit.SetLocalPosition2(id, atomkit.LocalPosition2Data{X: 10, Y: 10})

		// Act
		m := atom.NoneOf(atomkit.LocalRotation2Key)
		mv := m.Match(atom.Default(), id)

		// Assert
		assert.True(t, mv, "must return true because matcher should not contain any of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Reset
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomkit.SetLocalPosition2(id, atomkit.LocalPosition2Data{X: 10, Y: 10})
		atomkit.SetLocalRotation2(id, atomkit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := atom.NoneOf(atomkit.LocalRotation2Key)
		mv := m.Match(atom.Default(), id)

		// Assert
		assert.False(t, mv, "must return false because matcher contains some of given keys")
	})

	t.Run("all of/none of", func(t *testing.T) {

		// Setup
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomkit.SetLocalPosition2(id, atomkit.LocalPosition2Data{X: 10, Y: 10})
		atomkit.SetLocalRotation2(id, atomkit.LocalRotation2Data{X: 10, Y: 10})

		// Act
		m := atom.AllOf(atomkit.LocalPosition2Key).NoneOf(atomkit.LocalRotation2Key)
		mv := m.Match(atom.Default(), id)

		// Assert
		assert.False(t, mv, "must return false to satisfy matcher")
	})
}
