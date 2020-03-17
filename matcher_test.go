package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinc/kit"
	"github.com/stretchr/testify/assert"
)

func TestMatcherHasAllOf(t *testing.T) {

	// Arrange
	data := []struct {
		allOf  []uint
		noneOf []uint
	}{
		{allOf: []uint{1, 2, 3}, noneOf: []uint{0, 10, 20}},
		{allOf: []uint{0, 10, 2}, noneOf: []uint{99, 82, 10}},
	}

	for _, d := range data {

		// Act
		m := zinc.AllOf(d.allOf...)
		has := m.HasAllOf(d.allOf...)

		// Assert
		assert.True(t, has, "has all of must return true")
	}
}

func TestMatcherAllOf(t *testing.T) {

	// Arrange
	data := []struct {
		allOf  []uint
		noneOf []uint
	}{
		{allOf: []uint{1, 2, 3}, noneOf: []uint{0, 10, 20}},
		{allOf: []uint{0, 10, 2}, noneOf: []uint{99, 82, 10}},
	}

	for _, d := range data {

		// Act
		m := zinc.AllOf(d.allOf...)

		// Assert
		assert.ElementsMatch(t, m.AllOfSlice(), d.allOf, "returned all of slice does not match input data")
	}
}

func TestMatcherHasNoneOf(t *testing.T) {

	// Arrange
	data := []struct {
		allOf  []uint
		noneOf []uint
	}{
		{allOf: []uint{1, 2, 3}, noneOf: []uint{0, 10, 20}},
		{allOf: []uint{0, 10, 2}, noneOf: []uint{99, 82, 10}},
	}

	for _, d := range data {

		// Act
		m := zinc.NoneOf(d.noneOf...)
		has := m.HasNoneOf(d.noneOf...)

		// Assert
		assert.True(t, has, "has none of must return true")
	}
}

func TestMatcherNoneOf(t *testing.T) {

	// Arrange
	data := []struct {
		allOf  []uint
		noneOf []uint
	}{
		{allOf: []uint{1, 2, 3}, noneOf: []uint{0, 10, 20}},
		{allOf: []uint{0, 10, 2}, noneOf: []uint{99, 82, 10}},
	}

	for _, d := range data {

		// Act
		m := zinc.NoneOf(d.noneOf...)

		// Assert
		assert.ElementsMatch(t, m.NoneOfSlice(), d.noneOf, "returned none of slice does not match input data")
	}
}

func TestMatcherAllOfHash(t *testing.T) {

	// Arrange, Act
	m1 := zinc.AllOf(kit.ZLocalPosition2, kit.ZVelocity2)
	m2 := zinc.AllOf(kit.ZVelocity2, kit.ZLocalPosition2)

	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherNoneOfHash(t *testing.T) {

	// Arrange, Act
	m1 := zinc.NoneOf(kit.ZLocalPosition2, kit.ZVelocity2)
	m2 := zinc.NoneOf(kit.ZVelocity2, kit.ZLocalPosition2)

	// Assert
	assert.Equal(t, m1.Hash(), m2.Hash(), "must share identical hash")
}

func TestMatcherHash(t *testing.T) {

	// Arrange, Act
	m1 := zinc.AllOf(kit.ZLocalPosition2, kit.ZVelocity2).NoneOf(kit.ZLocalRotation2, kit.ZLocalScale2)
	m2 := zinc.AllOf(kit.ZVelocity2, kit.ZLocalPosition2).NoneOf(kit.ZLocalScale2, kit.ZLocalRotation2)

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
		kit.AddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})
		kit.AddLocalRotation2X(e, id, kit.ZLocalRotation2Data{X: 10, Y: 10})

		// Act
		f := func() {
			m := zinc.AllOf(0)
			m.Match(e, id)
		}

		// Assert
		assert.Panics(t, f, "must panic if matcher attempts to match non-existing key")
	})

	t.Run("all of", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.AddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})
		kit.AddLocalRotation2X(e, id, kit.ZLocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(kit.ZLocalPosition2, kit.ZLocalRotation2)
		mv := m.Match(e, id)

		// Assert
		assert.True(t, mv, "must return true if matcher contains all of given keys")
	})

	t.Run("none of", func(t *testing.T) {

		// Setup
		e := zinc.NewEntityManager()
		kit.RegisterLocalPosition2ComponentWith(e)
		kit.RegisterLocalRotation2ComponentWith(e)

		// Arrange
		id := e.CreateEntity()
		kit.AddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})

		// Act
		m := zinc.NoneOf(kit.ZLocalRotation2)
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
		kit.AddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})
		kit.AddLocalRotation2X(e, id, kit.ZLocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.NoneOf(kit.ZLocalRotation2)
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
		kit.AddLocalPosition2X(e, id, kit.ZLocalPosition2Data{X: 10, Y: 10})
		kit.AddLocalRotation2X(e, id, kit.ZLocalRotation2Data{X: 10, Y: 10})

		// Act
		m := zinc.AllOf(kit.ZLocalPosition2).NoneOf(kit.ZLocalRotation2)
		mv := m.Match(e, id)

		// Assert
		assert.False(t, mv, "must return false to satisfy matcher")
	})
}
