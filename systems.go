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

// Shutdown ...
type Shutdown interface {
	Shutdown()
}

// ZSystems ...
type ZSystems struct {
	systems    []interface{}
	initialize []Initialize
	update     []Update
	cleanup    []Cleanup
	shutdown   []Shutdown
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
		add := false
		if v, ok := sysv.(Initialize); ok {
			s.initialize = append(s.initialize, v)
			add = true
		}
		if v, ok := sysv.(Update); ok {
			s.update = append(s.update, v)
			add = true
		}
		if v, ok := sysv.(Cleanup); ok {
			s.cleanup = append(s.cleanup, v)
			add = true
		}
		if v, ok := sysv.(Shutdown); ok {
			s.shutdown = append(s.shutdown, v)
			add = true
		}
		if add {
			s.systems = append(s.systems, sysv)
		}
	}
}

// SystemsSlice returns a slice of all system
// types that have been added.
func (s *ZSystems) SystemsSlice() []interface{} {
	return s.systems
}

// InitializeSystemsSlice returns a slice of initialize
// systems.
func (s *ZSystems) InitializeSystemsSlice() []Initialize {
	return s.initialize
}

// UpdateSystemsSlice returns a slice of update
// systems.
func (s *ZSystems) UpdateSystemsSlice() []Update {
	return s.update
}

// CleanupSystemsSlice returns a slice of Cleanup
// systems.
func (s *ZSystems) CleanupSystemsSlice() []Cleanup {
	return s.cleanup
}

// ShutdownSystemsSlice returns a slice of shutdown
// systems.
func (s *ZSystems) ShutdownSystemsSlice() []Shutdown {
	return s.shutdown
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

// Shutdown calls the shutdown method
// on all systems in the systems list.
func (s *ZSystems) Shutdown() {
	for _, sys := range s.shutdown {
		sys.Shutdown()
	}
}
