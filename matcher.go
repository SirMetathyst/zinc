package atom

// Matcher ...
type Matcher struct {
	allOf         []uint
	noneOf        []uint
	hash          uint
	entityManager *EntityManager
}

// NewMatcher ...
func NewMatcher(e *EntityManager) *Matcher {
	return &Matcher{entityManager: e}
}

// HasAllOf ...
func (m *Matcher) HasAllOf(keys ...uint) bool {
	c := 0
	for _, v := range m.allOf {
		for _, k := range keys {
			if v == k {
				c++
			}
		}
	}
	return c >= len(keys)
}

// HasNoneOf ...
func (m *Matcher) HasNoneOf(keys ...uint) bool {
	c := 0
	for _, v := range m.noneOf {
		for _, k := range keys {
			if v == k {
				c++
			}
		}
	}
	return c >= len(keys)
}

// AllOf ...
func (m *Matcher) AllOf(keys ...uint) *Matcher {
	for _, key := range keys {
		if has := m.HasAllOf(key); !has {
			m.allOf = append(m.allOf, key)
		}
	}
	m.updateHash()
	return m
}

// NoneOf ...
func (m *Matcher) NoneOf(keys ...uint) *Matcher {
	for _, key := range keys {
		if has := m.HasNoneOf(key); !has {
			m.noneOf = append(m.noneOf, key)
		}
	}
	m.updateHash()
	return m
}

func (m *Matcher) updateHash() {
	m.hash = hash(647,
		hash(653, m.allOf...),
		hash(661, m.noneOf...))
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
func (m *Matcher) Hash() uint {
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

func hash(factor uint, x ...uint) uint {
	var hash uint
	for _, v := range x {
		hash ^= uint(v)
	}
	hash ^= uint(len(x)) * factor
	return hash
}
