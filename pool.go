package atom

// FactoryFunc ...
type FactoryFunc func() interface{}

// Pool ...
type Pool struct {
	s []interface{}
	f FactoryFunc
}

// NewPool ...
func NewPool(f FactoryFunc) *Pool {
	if f == nil {
		return nil
	}
	return &Pool{f: f}
}

// Get ...
func (p *Pool) Get() interface{} {
	l := len(p.s)
	if l == 0 {
		return p.f()
	}
	i := l - 1
	v := p.s[i]
	p.s = p.s[:i]
	return v
}

// Put ...
func (p *Pool) Put(v interface{}) {
	p.s = append(p.s, v)
}
