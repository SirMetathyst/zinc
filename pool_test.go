package atom

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func get(t *testing.T, p *p, id EntityID) {

	// Act
	r := p.Get()

	// Assert
	assert.Equal(t, id, r)
}

func TestNewPool(t *testing.T) {

	// Arrange, Act
	p := newPool(nil);

	// Assert
	assert.Nil(t, p, "pool with nil factory must return nil")
}

func TestNewPoolWithFactory(t *testing.T) {

	// Arrange, Act
	p := newPool(newEntityIDFactory());

	// Assert
	assert.NotNil(t, p, "pool with factory must not return nil")
}

func TestPoolGet(t *testing.T) {

	// Arrange
	p := newPool(newEntityIDFactory())

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

	// Arrange
	p := newPool(func() interface{} {
		return EntityID(1000000)
	})

	p.Put(EntityID(1))
	p.Put(EntityID(2))
	p.Put(EntityID(3))

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
