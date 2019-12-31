package atom_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
)

func get(t *testing.T, p *atom.Pool, id atom.EntityID) {
	r := p.Get()
	if r != atom.EntityID(id) {
		t.Errorf("assert: want %v, got %v", id, r)
	}
}

func TestPoolNew(t *testing.T) {

	t.Run("new pool with nil factory returns nil pool", func(t *testing.T) {
		if p := atom.NewPool(nil); p != nil {
			t.Errorf("assert: new pool with nil factory doesn't return nil")
		}
	})

	t.Run("new pool with factory returns pool", func(t *testing.T) {
		if p := atom.NewPool(atom.NewEntityIDFactory()); p == nil {
			t.Errorf("assert: new pool with factory doesn't return pool")
		}
	})

}

func TestPool(t *testing.T) {

	p := atom.NewPool(atom.NewEntityIDFactory())

	// 1st Get
	// Should call factory
	// Should return 1
	get(t, p, 1)

	// 2nd Get
	// Should call factory
	// should return 2
	get(t, p, 2)
}

func TestPoolPut(t *testing.T) {

	p := atom.NewPool(func() interface{} {
		return atom.EntityID(1000000)
	})

	// Put
	p.Put(atom.EntityID(1))
	p.Put(atom.EntityID(2))
	p.Put(atom.EntityID(3))

	// 1st Get
	// Should not call factory
	// should return 3
	get(t, p, 3)

	// 2nd Get
	// Should not call factory
	// should return 2
	get(t, p, 2)

	// 3rd Get
	// Should not call factory
	// should return 1
	get(t, p, 1)

	// 4th Get
	// Should call factory
	// should return 1000000
	get(t, p, 1000000)
}
