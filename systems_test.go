package zinc_test

import (
	"testing"

	"github.com/SirMetathyst/zinc"
	"github.com/stretchr/testify/assert"
)

// TestSystem ...
type TestSystem struct {
	Initialized bool
	Updated     bool
	DeltaTime   float64
	Cleanedup   bool
	Shutdownx   bool
}

// Initialize ...
func (sys *TestSystem) Initialize() {
	sys.Initialized = true
}

// Update ...
func (sys *TestSystem) Update(dt float64) {
	sys.Updated = true
	sys.DeltaTime = dt
}

// Cleanup ...
func (sys *TestSystem) Cleanup() {
	sys.Cleanedup = true
}

// Shutdown ...
func (sys *TestSystem) Shutdown() {
	sys.Shutdownx = true
}

func TestNewSystems(t *testing.T) {

	// Arrange, Act
	sys := zinc.NewSystems()

	// Assert
	assert.NotNil(t, sys, "NewSystems must not return nil")
}

func TestSystemAddInitialize(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)

	// Assert
	assert.Contains(t, sys.InitializeSystemsSlice(), s, "Systems does contain TestSystem")
	assert.Contains(t, sys.SystemsSlice(), s, "SystemsSlice() does contain TestSystem")
}

func TestSystemAddUpdate(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)

	// Assert
	assert.Contains(t, sys.UpdateSystemsSlice(), s, "UpdateSystemsSlice() does contain TestSystem")
	assert.Contains(t, sys.SystemsSlice(), s, "SystemsSlice() does contain TestSystem")
}

func TestSystemAddCleanup(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)

	// Assert
	assert.Contains(t, sys.CleanupSystemsSlice(), s, "CleanupSystemsSlice() does contain TestSystem")
	assert.Contains(t, sys.SystemsSlice(), s, "SystemsSlice() does contain TestSystem")
}

func TestSystemAddShutdown(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)

	// Assert
	assert.Contains(t, sys.ShutdownSystemsSlice(), s, "ShutdownSystemsSlice() does contain TestSystem")
	assert.Contains(t, sys.SystemsSlice(), s, "SystemsSlice() does contain TestSystem")
}

func TestSystemInitialize(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)
	sys.Initialize()

	// Assert
	assert.True(t, s.Initialized, "TestSystem did not initialize...")
}

func TestSystemUpdate(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}
	dt := float64(100)

	// Act
	sys.Add(s)
	sys.Update(dt)

	// Assert
	assert.True(t, s.Updated, "TestSystem did not update...")
	assert.Equal(t, dt, s.DeltaTime, "TestSystem did not return correct delta time")
}

func TestSystemCleanup(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)
	sys.Cleanup()

	// Assert
	assert.True(t, s.Cleanedup, "TestSystem did not cleanup...")
}

func TestSystemShutdown(t *testing.T) {

	// Arrange
	sys := zinc.NewSystems()
	s := &TestSystem{}

	// Act
	sys.Add(s)
	sys.Shutdown()

	// Assert
	assert.True(t, s.Shutdownx, "TestSystem did not shutdown...")
}
