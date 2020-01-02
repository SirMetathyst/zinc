package benchmark

import (
	"testing"

	"github.com/SirMetathyst/atom"
)

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
