package atom

import (
	"testing"
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
	for it, d := range matcherData {
		m := AllOf(d.allOf...)
		for i, _ := range m.allOf {
			if m.allOf[i] != d.allOf[i] {
				t.Errorf("assert(%d-%d): want %v, got %v", it, i, d.allOf[i], m.allOf[i])
			}
		}
	}
}

func TestNoneOf(t *testing.T) {
	for it, d := range matcherData {
		m := NoneOf(d.noneOf...)
		for i, _ := range m.noneOf {
			if m.noneOf[i] != d.noneOf[i] {
				t.Errorf("assert(%d-%d): want %v, got %v", it, i, d.noneOf[i], m.noneOf[i])
			}
		}
	}
}

// Mock component for Matcher.Match

const cKey uint = 31162

type cData struct{}

type c struct {
	context Context
	d       map[EntityID]cData
}

func newC() *c {
	return &c{
		d: make(map[EntityID]cData),
	}
}

// EntityDeleted ...
func (x *c) EntityDeleted(id EntityID) {
	delete(x.d, id)
}

// HasEntity ...
func (x *c) HasEntity(id EntityID) bool {
	_, ok := x.d[id]
	return ok
}

// Set ...
func (x *c) Set(id EntityID, example cData) {
	if x.context.HasEntity(id) {
		if x.HasEntity(id) {
			x.d[id] = example
			x.context.ComponentUpdated(cKey, id)
		} else {
			x.d[id] = example
			x.context.ComponentAdded(cKey, id)
		}
	}
}

func TestMatch(t *testing.T) {

	// Setup
	cmp := newC()
	context := Default().RegisterComponent(cKey, cmp)
	cmp.context = context

	// Arrange
	id := CreateEntity()
	v := Default().Component(cKey)
	c := v.(*c)
	c.Set(id, cData{})

	// Act
	m := AllOf(cKey)

	// Assert
	mv := m.Match(id)
	if mv != true {
		t.Errorf("assert: want %v, got %v", true, mv)
	}
}
