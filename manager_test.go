package atom_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
)

func addEntities(n []int, t *testing.T) {
	for _, x := range n {
		id := atom.CreateEntity()
		if id != atom.EntityID(x) {
			t.Errorf("addEntities: want %d, got %v", x, id)
		}
	}
}

func deleteEntities(n []int, t *testing.T) {
	for _, x := range n {
		atom.DeleteEntity(atom.EntityID(x))
	}
}

func hasEntities(n []int, t *testing.T) {
	for _, x := range n {
		v := atom.HasEntity(atom.EntityID(x))
		if !v {
			t.Errorf("hasEntities: EntityID(%d) does not exist", x)
		}
	}
}

func doesNotHaveEntities(n []int, t *testing.T) {
	for _, x := range n {
		v := atom.HasEntity(atom.EntityID(x))
		if v {
			t.Errorf("doesNotHaveEntities: EntityID(%d) should not exist", x)
		}
	}
}

func TestEntityManager(t *testing.T) {
	if atom.NewEntityManager() == nil {
		t.Errorf("NewEntityManager(): should not return nil")
	}
	atom.Reset()
	addEntities([]int{1, 2, 3, 4, 5}, t)
	hasEntities([]int{1, 2, 3, 4, 5}, t)
	doesNotHaveEntities([]int{0, 6, 7, 8}, t)
	deleteEntities([]int{1, 2, 3}, t)
	hasEntities([]int{4, 5}, t)
	doesNotHaveEntities([]int{1, 2, 3}, t)
	addEntities([]int{3, 2, 1, 6}, t)
}

func BenchmarkNewEntityManager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atom.NewEntityManager()
	}
}

func BenchmarkCreateEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atom.Reset()
		atom.CreateEntity()
	}
}
