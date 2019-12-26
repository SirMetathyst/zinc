package atom

import (
	"sort"
	"strconv"
)

// UintSlice ...
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

// HasAllOf ...
func (m *Matcher) HasAllOf(keys ...uint) bool {
	for _, v := range m.allOf {
		for _, k := range keys {
			if v == k {
				return true
			}
		}
	}
	return false
}

// HasNoneOf ...
func (m *Matcher) HasNoneOf(keys ...uint) bool {
	for _, v := range m.noneOf {
		for _, k := range keys {
			if v == k {
				return true
			}
		}
	}
	return false
}

// AllOf ...
func (m *Matcher) AllOf(keys ...uint) *Matcher {
	for _, key := range keys {
		if has := m.HasAllOf(key); !has {
			m.hash = ""
			m.allOf = append(m.allOf, key)
		}
	}
	return m
}

// NoneOf ...
func (m *Matcher) NoneOf(keys ...uint) *Matcher {
	for _, key := range keys {
		if has := m.HasNoneOf(key); !has {
			m.hash = ""
			m.noneOf = append(m.noneOf, key)
		}
	}
	return m
}

func (m *Matcher) match(key uint, id EntityID) bool {
	if c, ok := m.entityManager.Component(key); ok {
		return c.HasEntity(id)
	}
	return false
}

// Match ...
func (m *Matcher) Match(id EntityID) bool {
	if len(m.allOf) == 0 {
		return false
	}
	for _, k := range m.allOf {
		if !m.match(k, id) {
			return false
		}
	}
	for _, k := range m.noneOf {
		if m.match(k, id) {
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

// AllOfX ...
func AllOfX(e *EntityManager, keys ...uint) *Matcher {
	return NewMatcher(e).AllOf(keys...)
}

// AllOf ...
func AllOf(keys ...uint) *Matcher {
	return AllOfX(Default(), keys...)
}

// NoneOfX ...
func NoneOfX(e *EntityManager, keys ...uint) *Matcher {
	return NewMatcher(e).NoneOf(keys...)
}

// NoneOf ...
func NoneOf(keys ...uint) *Matcher {
	return NoneOfX(Default(), keys...)
}
