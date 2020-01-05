package zinc

// S ...
type S interface{}

// U ...
type U interface {
	Update(dt float64)
}

// CU ...
type CU interface {
	Cleanup()
}

// I ...
type I interface {
	Initialize()
}

// SS ...
type SS interface {
	// Initialize ...
	I
	// Update ...
	U
	// Cleanup ...
	CU
	Add(sys ...S)
}

type sys struct {
	initialize []I
	update    []U
	cleanup   []CU
}

// NewSystems ...
// TODO: Write TEST
func NewSystems() SS {
	return &sys{}
}

// Add ...
// TODO: Write TEST
func (s *sys) Add(sys ...S) {
	for _, sysv := range sys {
		switch v := sysv.(type) {
		case I:
			s.initialize = append(s.initialize, v)
			break
		case U:
			s.update = append(s.update, v)
			break
		case CU:
			s.cleanup = append(s.cleanup, v)
			break
		}
	}
}

// Initialize ...
// TODO: Write TEST
func (s *sys) Initialize() {
	for _, sys := range s.initialize {
		sys.Initialize()
	}
}

// Update ...
// TODO: Write TEST
func (s *sys) Update(dt float64) {
	for _, sys := range s.update {
		sys.Update(dt)
	}
}

// Cleanup ...
// TODO: Write TEST
func (s *sys) Cleanup() {
	for _, sys := range s.cleanup {
		sys.Cleanup()
	}
}
