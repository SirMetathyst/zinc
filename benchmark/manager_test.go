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
		zinc.Reset()
		zinc.CreateEntity()
	}
}
