package zinc

type p struct {
	s []interface{}
	f func() interface{}
}

func newPool(f func() interface{}) *p {
	if f == nil {
		return nil
	}
	return &p{f: f}
}

// Get ...
func (p *p) Get() interface{} {
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
func (p *p) Put(v interface{}) {
	p.s = append(p.s, v)
}
