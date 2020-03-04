package zinc

// Update ...
type Update interface {
	Update(dt float64)
}

// Cleanup ...
type Cleanup interface {
	Cleanup()
}

// Initialize ...
type Initialize interface {
	Initialize()
}

// ZSystems ...
type ZSystems struct {
	initialize []Initialize
	update     []Update
	cleanup    []Cleanup
}

// NewSystems returns a new systems container
// and returns it.
func NewSystems() *ZSystems {
	return &ZSystems{}
}

// Add takes in one or more types
// and adds them to lists based on what
// system methods have been implemented.
func (s *ZSystems) Add(sys ...interface{}) {
	for _, sysv := range sys {
		switch v := sysv.(type) {
		case Initialize:
			s.initialize = append(s.initialize, v)
			break
		case Update:
			s.update = append(s.update, v)
			break
		case Cleanup:
			s.cleanup = append(s.cleanup, v)
			break
		}
	}
}

// Initialize calls the initialize method
// on all systems in the systems list.
func (s *ZSystems) Initialize() {
	for _, sys := range s.initialize {
		sys.Initialize()
	}
}

// Update calls the update method
// on all systems in the systems list.
func (s *ZSystems) Update(dt float64) {
	for _, sys := range s.update {
		sys.Update(dt)
	}
}

// Cleanup calls the clean up method
// on all systems in the systems list.
func (s *ZSystems) Cleanup() {
	for _, sys := range s.cleanup {
		sys.Cleanup()
	}
}
