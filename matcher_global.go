package zinc

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
