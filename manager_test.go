package atom

import "testing"

func addEntities(n []int, e *EntityManager, t *testing.T) {
	for _, x := range n {
		id := e.CreateEntity()
		if id != EntityID(x) {
			t.Errorf("addEntities: want %d, got %v", x, id)
		}
	}
}

func deleteEntities(n []int, e *EntityManager, t *testing.T) {
	for _, x := range n {
		e.DeleteEntity(EntityID(x))
	}
}

func hasEntities(n []int, e *EntityManager, t *testing.T) {
	for _, x := range n {
		v := e.HasEntity(EntityID(x))
		if !v {
			t.Errorf("hasEntities: EntityID(%d) does not exist", x)
		}
	}
}

func doesNotHaveEntities(n []int, e *EntityManager, t *testing.T) {
	for _, x := range n {
		v := e.HasEntity(EntityID(x))
		if v {
			t.Errorf("doesNotHaveEntities: EntityID(%d) should not exist", x)
		}
	}
}

func TestEntityManager(t *testing.T) {
	e := NewEntityManager()
	if e == nil {
		t.Errorf("NewEntityManager(): should not return nil")
	}
	addEntities([]int{1, 2, 3, 4, 5}, e, t)
	hasEntities([]int{1, 2, 3, 4, 5}, e, t)
	doesNotHaveEntities([]int{0, 6, 7, 8}, e, t)
	deleteEntities([]int{1, 2, 3}, e, t)
	hasEntities([]int{4, 5}, e, t)
	doesNotHaveEntities([]int{1, 2, 3}, e, t)
	addEntities([]int{3, 2, 1, 6}, e, t)
}

func BenchmarkNewEntityManager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEntityManager()
	}
}

func BenchmarkCreateEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := NewEntityManager()
		e.CreateEntity()
	}
}
