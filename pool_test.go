package atom

import "testing"

func TestPoolGet(t *testing.T) {

	p := NewPool(NewEntityIDFactory())

	// 1st Get
	// Should call factory
	// Should return 1
	r1 := p.Get()
	if r1 != EntityID(1) {
		t.Errorf("p.Get: want %d, got %v", 1, r1)
	}

	// 2nd Get
	// Should call factory
	// should return 2
	r2 := p.Get()
	if r2 != EntityID(2) {
		t.Errorf("p.Get: want %d, got %v", 2, r2)
	}
}

func TestPoolPut(t *testing.T) {

	p := NewPool(func() interface{} {
		return -1
	})

	// Put
	p.Put(1)
	p.Put(2)
	p.Put(3)

	// 1st Get
	// Should not call factory
	// should return 3
	r1 := p.Get()
	if r1 != 3 {
		t.Errorf("p.Get: want %d, got %v", 3, r1)
	}

	// 2nd Get
	// Should not call factory
	// should return 2
	r2 := p.Get()
	if r2 != 2 {
		t.Errorf("p.Get: want %d, got %v", 2, r2)
	}

	// 3rd Get
	// Should not call factory
	// should return 1
	r3 := p.Get()
	if r3 != 1 {
		t.Errorf("p.Get: want %d, got %v", 1, r3)
	}

	// 4th Get
	// Should not call factory
	// should return -1
	r4 := p.Get()
	if r4 != -1 {
		t.Errorf("p.Get: want %d, got %v", -1, r4)
	}
}
