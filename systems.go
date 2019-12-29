package atom

// System ...
type System interface{}

// Updater ...
type Updater interface {
	Update(dt float64)
}

// Cleanuper ...
type Cleanuper interface {
	Cleanup()
}

// Initializer ...
type Initializer interface {
	Initialize()
}

// Systems ...
type Systems interface {
	Initializer
	Updater
	Cleanuper
	Add(sys ...System)
}

type sys struct {
	initializer []Initializer
	updater     []Updater
	cleanuper   []Cleanuper
}

// NewSystems ...
func NewSystems() Systems {
	return &sys{}
}

// Add ...
func (s *sys) Add(sys ...System) {
	for _, sysv := range sys {
		switch v := sysv.(type) {
		case Initializer:
			s.initializer = append(s.initializer, v)
			break
		case Updater:
			s.updater = append(s.updater, v)
			break
		case Cleanuper:
			s.cleanuper = append(s.cleanuper, v)
			break
		}
	}
}

// Initialize ...
func (s *sys) Initialize() {
	for _, sys := range s.initializer {
		sys.Initialize()
	}
}

// Update ...
func (s *sys) Update(dt float64) {
	for _, sys := range s.updater {
		sys.Update(dt)
	}
}

// Cleanup ...
func (s *sys) Cleanup() {
	for _, sys := range s.cleanuper {
		sys.Cleanup()
	}
}
