package benchmark

import (
	"testing"

	"github.com/SirMetathyst/zinc"
)

func BenchmarkNewEntityManager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zinc.NewEntityManager()
	}
}

func BenchmarkCreateEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zinc.CreateEntity()
	}
}

func BenchmarkCreateDeleteEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := zinc.CreateEntity()
		zinc.DeleteEntity(id)
	}
}
