package zinc

// ZMatcher ...
type ZMatcher struct {
	allOf  []uint
	noneOf []uint
	hash   uint
}

// NewMatcher creates a new matcher and returns it.
func NewMatcher() *ZMatcher {
	return &ZMatcher{}
}

// HasAllOf returns true if all of the given keys
// were added to the matcher through AllOf method.
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

// HasNoneOf returns true if all of the given keys
// were added to the matcher through NoneOf method.
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

// AllOfSlice returns a slice of component
// keys an entity should contain.
func (m *ZMatcher) AllOfSlice() []uint {
	return m.allOf
}

// AllOf includes the given component keys
// for matching an entity.
func (m *ZMatcher) AllOf(keys ...uint) *ZMatcher {
	for _, key := range keys {
		if has := m.HasAllOf(key); !has {
			m.allOf = append(m.allOf, key)
		}
	}
	return m
}

// NoneOfSlice returns a slice of component
// keys an entity should not contain.
func (m *ZMatcher) NoneOfSlice() []uint {
	return m.noneOf
}

// NoneOf excludes the given component keys
// for matching an entity.
func (m *ZMatcher) NoneOf(keys ...uint) *ZMatcher {
	for _, key := range keys {
		if has := m.HasNoneOf(key); !has {
			m.noneOf = append(m.noneOf, key)
		}
	}
	return m
}

// Match returns true if the given entity contains
// the components keys added through AllOf and does not
// contain the components added through NoneOf.
func (m *ZMatcher) Match(e *ZEntityManager, id ZEntityID) bool {
	for _, k := range m.allOf {
		if !e.Component(k).HasEntity(id) {
			return false
		}
	}
	for _, k := range m.noneOf {
		if e.Component(k).HasEntity(id) {
			return false
		}
	}
	return true
}

// Hash returns a hash of the current matcher.
func (m *ZMatcher) Hash() uint {
	return hash(647, hash(653, m.allOf...), hash(661, m.noneOf...))
}

// AllOf includes the given component keys
// for matching an entity.
func AllOf(keys ...uint) *ZMatcher {
	return NewMatcher().AllOf(keys...)
}

// NoneOf excludes the given component keys
// for matching an entity.
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
