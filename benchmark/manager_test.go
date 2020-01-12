package benchmark

import (
	"testing"
	"time"

	"github.com/SirMetathyst/zinc"
)

func BenchmarkNewEntityManager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zinc.NewEntityManager()
	}
}

func BenchmarkCreateEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		zinc.Reset()
		b.StartTimer()
		zinc.CreateEntity()
	}
}

func BenchmarkCreateDeleteEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := zinc.CreateEntity()
		zinc.DeleteEntity(id)
	}
}
