package atom

import (
	"sort"
	"strconv"
)

type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Matcher ...
type Matcher struct {
	allOf         []uint
	noneOf        []uint
	hash          string
	entityManager *EntityManager
}

// NewMatcher ...
func NewMatcher(e *EntityManager) *Matcher {
	return &Matcher{entityManager: e}
}

func (m *Matcher) hasKeyInAllOf(key uint) bool {
	for _, v := range m.allOf {
		if v == key {
			return true
		}
	}
	return false
}

func (m *Matcher) hasKeyInNoneOf(key uint) bool {
	for _, v := range m.noneOf {
		if v == key {
			return true
		}
	}
	return false
}

// AllOf ...
func (m *Matcher) AllOf(keys ...uint) *Matcher {
	for _, inkey := range keys {
		if has := m.hasKeyInAllOf(inkey); !has {
			m.hash = ""
			m.allOf = append(m.allOf, inkey)
		}
	}
	return m
}

// NoneOf ...
func (m *Matcher) NoneOf(keys ...uint) *Matcher {
	for _, inkey := range keys {
		if has := m.hasKeyInNoneOf(inkey); !has {
			m.hash = ""
			m.noneOf = append(m.noneOf, inkey)
		}
	}
	return m
}

// Match ...
func (m *Matcher) Match(id EntityID) bool {
	if len(m.allOf) == 0 {
		return false
	}
	for _, k := range m.allOf {
		c := m.entityManager.Component(k)
		v := c.HasEntity(id)
		if v != true {
			return false
		}
	}
	for _, k := range m.noneOf {
		c := m.entityManager.Component(k)
		v := c.HasEntity(id)
		if v == true {
			return false
		}
	}
	return true
}

// Hash ...
func (m *Matcher) Hash() string {
	if m.hash != "" {
		return m.hash
	}
	sort.Sort(UintSlice(m.allOf))
	sort.Sort(UintSlice(m.noneOf))
	allOfStr := ""
	for _, v := range m.allOf {
		str := strconv.Itoa(int(v))
		allOfStr += str
	}
	noneOfStr := ""
	for _, v := range m.noneOf {
		str := strconv.Itoa(int(v))
		noneOfStr += str
	}
	m.hash = allOfStr + "-" + noneOfStr
	return m.hash
}

// AllOf ...
func AllOf(keys ...uint) *Matcher {
	return NewMatcher(Default()).AllOf(keys...)
}

// NoneOf ...
func NoneOf(keys ...uint) *Matcher {
	return NewMatcher(Default()).NoneOf(keys...)
}
