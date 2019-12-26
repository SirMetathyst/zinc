package atom_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomcommon"
)

// MatcherData ...
type MatcherData struct {
	allOf  []uint
	noneOf []uint
}

var matcherData = []MatcherData{
	{allOf: []uint{1, 2, 3}, noneOf: []uint{0, 10, 20}},
	{allOf: []uint{0, 10, 2}, noneOf: []uint{99, 82, 10}},
}

func TestAllOf(t *testing.T) {
	for it, ds := range matcherData {
		m := atom.AllOf(ds.allOf...)
		for i, v := range ds.allOf {
			if !m.HasAllOf(v) {
				t.Errorf("assert(%d-%d): want %v", it, i, v)
			}
		}
	}
}

func TestNoneOf(t *testing.T) {
	for it, ds := range matcherData {
		m := atom.NoneOf(ds.noneOf...)
		for i, v := range ds.noneOf {
			if !m.HasNoneOf(v) {
				t.Errorf("assert(%d-%d): want %v", it, i, v)
			}
		}
	}
}

func TestHash(t *testing.T) {

	t.Run("All of hash", func(t *testing.T) {
		m1 := atom.AllOf(atomcommon.Position2Key, atomcommon.Velocity2Key)
		m2 := atom.AllOf(atomcommon.Velocity2Key, atomcommon.Position2Key)
		if m1.Hash() != m2.Hash() {
			t.Errorf("assert: want %s = %s, got %s = %s", m1.Hash(), m1.Hash(), m1.Hash(), m2.Hash())
		}
	})

	t.Run("None of hash", func(t *testing.T) {
		m1 := atom.NoneOf(atomcommon.Position2Key, atomcommon.Velocity2Key)
		m2 := atom.NoneOf(atomcommon.Velocity2Key, atomcommon.Position2Key)
		if m1.Hash() != m2.Hash() {
			t.Errorf("assert: want %s = %s, got %s = %s", m1.Hash(), m1.Hash(), m1.Hash(), m2.Hash())
		}
	})

	t.Run("All/None of hash", func(t *testing.T) {
		m1 := atom.AllOf(atomcommon.Position2Key, atomcommon.Velocity2Key).NoneOf(atomcommon.Rotation2Key, atomcommon.Scale2Key)
		m2 := atom.AllOf(atomcommon.Velocity2Key, atomcommon.Position2Key).NoneOf(atomcommon.Scale2Key, atomcommon.Rotation2Key)
		if m1.Hash() != m2.Hash() {
			t.Errorf("assert: want %s = %s, got %s = %s", m1.Hash(), m1.Hash(), m1.Hash(), m2.Hash())
		}
	})
}

func TestMatch(t *testing.T) {

	t.Run("Match with non-existing", func(t *testing.T) {

		// Reset
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomcommon.SetPosition2(id, atomcommon.Position2Data{X: 10, Y: 10})

		// Act
		m := atom.AllOf(0)

		// Assert
		mv := m.Match(id)
		if mv != false {
			t.Errorf("assert: want %v, got %v", false, mv)
		}
	})

	t.Run("Match with all of", func(t *testing.T) {

		// Reset
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomcommon.SetPosition2(id, atomcommon.Position2Data{X: 10, Y: 10})

		// Act
		m := atom.AllOf(atomcommon.Position2Key)

		// Assert
		mv := m.Match(id)
		if mv != true {
			t.Errorf("assert: want %v, got %v", true, mv)
		}
	})

	t.Run("Match with all of/none of", func(t *testing.T) {

		// Setup
		atom.Reset()

		// Arrange
		id := atom.CreateEntity()
		atomcommon.SetPosition2(id, atomcommon.Position2Data{X: 10, Y: 10})
		atomcommon.SetVelocity2(id, atomcommon.Velocity2Data{X: 10, Y: 10})

		// Act
		m1 := atom.NoneOf(atomcommon.Velocity2Key)

		// Assert
		mv1 := m1.Match(id)
		if mv1 == true {
			t.Errorf("assert: want %v, got %v", false, mv1)
		}

		// Act
		m2 := atom.AllOf(atomcommon.Position2Key).NoneOf(atomcommon.Velocity2Key)

		// Assert
		mv2 := m2.Match(id)
		if mv2 == true {
			t.Errorf("assert: want %v, got %v", false, mv2)
		}
	})
}
