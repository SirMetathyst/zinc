package zinc

// ZMatcher ...
type ZMatcher struct {
	allOf  []uint
	noneOf []uint
	hash   uint
}

// NewMatcher ...
func NewMatcher() *ZMatcher {
	return &ZMatcher{}
}

// HasAllOf ...
func (m *ZMatcher) HasAllOf(keys ...uint) bool {
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
func (m *ZMatcher) HasNoneOf(keys ...uint) bool {
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

// AllOfSlice ...
func (m *ZMatcher) AllOfSlice() []uint {
	return m.allOf
}

// AllOf ...
func (m *ZMatcher) AllOf(keys ...uint) *ZMatcher {
	for _, key := range keys {
		if has := m.HasAllOf(key); !has {
			m.allOf = append(m.allOf, key)
		}
	}
	m.updateHash()
	return m
}

// NoneOfSlice ...
func (m *ZMatcher) NoneOfSlice() []uint {
	return m.noneOf
}

// NoneOf ...
func (m *ZMatcher) NoneOf(keys ...uint) *ZMatcher {
	for _, key := range keys {
		if has := m.HasNoneOf(key); !has {
			m.noneOf = append(m.noneOf, key)
		}
	}
	m.updateHash()
	return m
}

func (m *ZMatcher) updateHash() {
	m.hash = hash(647,
		hash(653, m.allOf...),
		hash(661, m.noneOf...))
}

func (m *ZMatcher) match(e *ZEntityManager, key uint, id ZEntityID) bool {
	c := e.Component(key)
	return c.HasEntity(id)
}

// Match ...
func (m *ZMatcher) Match(e *ZEntityManager, id ZEntityID) bool {
	for _, k := range m.allOf {
		if !m.match(e, k, id) {
			return false
		}
	}
	for _, k := range m.noneOf {
		if m.match(e, k, id) {
			return false
		}
	}
	return true
}

// Hash ...
func (m *ZMatcher) Hash() uint {
	return m.hash
}

// AllOf ...
func AllOf(keys ...uint) *ZMatcher {
	return NewMatcher().AllOf(keys...)
}

// NoneOf ...
func NoneOf(keys ...uint) *ZMatcher {
	return NewMatcher().NoneOf(keys...)
}

func hash(factor uint, x ...uint) uint {
	var hash uint
	for _, v := range x {
		hash ^= uint(v)
	}
	hash ^= uint(len(x)) * factor
	return hash
}
