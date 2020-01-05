package atom_test

import (
	"testing"

	"github.com/SirMetathyst/atom"
)

func TestContext(t *testing.T) {

	t.Run("component added", func(t *testing.T) {
		key := uint(10)
		rkey := uint(0)
		id := atom.EntityID(2)
		rid := atom.EntityID(0)
		called := false
		ctx := atom.NewContext(func(key uint, id atom.EntityID) {
			rid = id
			rkey = key
			called = true
		}, nil, nil, nil)
		ctx.ComponentAdded(key, id)
		if !called {
			t.Errorf("assert: context component added should have been called")
		}
		if rid != id {
			t.Errorf("assert: want %v, got %v", id, rid)
		}
		if rkey != key {
			t.Errorf("assert: want %v, got %v", key, rkey)
		}
	})

	t.Run("component deleted", func(t *testing.T) {
		key := uint(10)
		rkey := uint(0)
		id := atom.EntityID(2)
		rid := atom.EntityID(0)
		called := false
		ctx := atom.NewContext(nil, func(key uint, id atom.EntityID) {
			rid = id
			rkey = key
			called = true
		}, nil, nil)
		ctx.ComponentDeleted(key, id)
		if !called {
			t.Errorf("assert: context component deleted should have been called")
		}
		if rid != id {
			t.Errorf("assert: want %v, got %v", id, rid)
		}
		if rkey != key {
			t.Errorf("assert: want %v, got %v", key, rkey)
		}
	})

	t.Run("component updated", func(t *testing.T) {
		key := uint(10)
		rkey := uint(0)
		id := atom.EntityID(2)
		rid := atom.EntityID(0)
		called := false
		ctx := atom.NewContext(nil, nil, func(key uint, id atom.EntityID) {
			rid = id
			rkey = key
			called = true
		}, nil)
		ctx.ComponentUpdated(key, id)
		if !called {
			t.Errorf("assert: context component updated should have been called")
		}
		if rid != id {
			t.Errorf("assert: want %v, got %v", id, rid)
		}
		if rkey != key {
			t.Errorf("assert: want %v, got %v", key, rkey)
		}
	})

	t.Run("has entity", func(t *testing.T) {
		id := atom.EntityID(2)
		rid := atom.EntityID(0)
		ctx := atom.NewContext(nil, nil, nil, func(id atom.EntityID) bool {
			rid = id
			return true
		})
		if !ctx.HasEntity(id) {
			t.Errorf("assert: context has entity should have been called")
		}
		if rid != id {
			t.Errorf("assert: want %v, got %v", id, rid)
		}
	})
}
