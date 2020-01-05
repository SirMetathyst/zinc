package zinc

// M ...
type M interface {
	
	HasAllOf(keys ...uint) bool
	HasNoneOf(keys ...uint) bool

	AllOfSlice() []uint
	AllOf(keys ...uint) M
	NoneOfSlice() []uint
	NoneOf(keys ...uint) M

	Match(e *EntityManager, id EntityID) bool 
	Hash() uint 
}


// Matcher ...
type m struct {
	allOf         []uint
	noneOf        []uint
	hash          uint
}

// NewMatcher ...
func NewMatcher() M {
	return &m{}
}

// HasAllOf ...
func (m *m) HasAllOf(keys ...uint) bool {
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
func (m *m) HasNoneOf(keys ...uint) bool {
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
func (m *m) AllOfSlice() []uint {
	return m.allOf
}

// AllOf ...
func (m *m) AllOf(keys ...uint) M {
	for _, key := range keys {
		if has := m.HasAllOf(key); !has {
			m.allOf = append(m.allOf, key)
		}
	}
	m.updateHash()
	return m
}

// NoneOfSlice ...
func (m *m) NoneOfSlice() []uint {
	return m.noneOf
}

// NoneOf ...
func (m *m) NoneOf(keys ...uint) M {
	for _, key := range keys {
		if has := m.HasNoneOf(key); !has {
			m.noneOf = append(m.noneOf, key)
		}
	}
	m.updateHash()
	return m
}

func (m *m) updateHash() {
	m.hash = hash(647,
		hash(653, m.allOf...),
		hash(661, m.noneOf...))
}

func (m *m) match(e *EntityManager, key uint, id EntityID) bool {
	if c, ok := e.Component(key); ok {
		return c.HasEntity(id)
	}
	return false
}

// Match ...
func (m *m) Match(e *EntityManager, id EntityID) bool {
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
func (m *m) Hash() uint {
	return m.hash
}

// AllOf ...
func AllOf(keys ...uint) M {
	return NewMatcher().AllOf(keys...)
}

// NoneOf ...
func NoneOf(keys ...uint) M {
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
