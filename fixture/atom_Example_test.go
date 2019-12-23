package fixture

import (
	"testing"

	"github.com/SirMetathyst/atom"
)

func TestSetComponent(t *testing.T) {

	// Setup
	atom.Reset()

	// Arrange
	id := atom.CreateEntity()
	inData1 := ExampleData{Value: 10}

	// Act
	SetExample(id, inData1)

	// Assert
	outData := Example(id)
	if outData != inData1 {
		t.Errorf("assert: want %v, got %v", inData1, outData)
	}

	// Arrange
	inData2 := ExampleData{Value: 10}

	// Act
	SetExample(id, inData2)

	// Assert
	outData = Example(id)
	if outData != inData2 {
		t.Errorf("assert: want %v, got %v", inData2, outData)
	}
}

func TestDeleteComponent(t *testing.T) {

	// Setup
	atom.Reset()

	// Arrange
	id := atom.CreateEntity()
	inData1 := ExampleData{Value: 10}

	// Act
	SetExample(id, inData1)

	// Assert
	v := HasExample(id)
	if v != true {
		t.Errorf("assert: want %v, got %v", true, v)
	}

	// Act
	DeleteExample(id)

	// Assert
	v = HasExample(id)
	if v != false {
		t.Errorf("assert: want %v, got %v", false, v)
	}
}

func BenchmarkSetComponent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := atom.NewEntityManager()
		id := e.CreateEntity()
		SetExample(id, ExampleData{10})
	}
}

func BenchmarkDeleteComponent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := atom.NewEntityManager()
		id := e.CreateEntity()
		SetExample(id, ExampleData{10})
		DeleteExample(id)
	}
}

func BenchmarkGetComponent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := atom.NewEntityManager()
		id := e.CreateEntity()
		SetExample(id, ExampleData{10})
		Example(id)
	}
}
