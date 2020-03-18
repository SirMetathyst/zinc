# Zinc 

[![codecov](https://codecov.io/gh/SirMetathyst/zinc/branch/master/graph/badge.svg)](https://codecov.io/gh/SirMetathyst/zinc)
[![Build Status](https://travis-ci.com/SirMetathyst/zinc.svg?branch=master)](https://travis-ci.com/SirMetathyst/zinc)

`ZincECS` is an entity-component-system package inspired by Simon Schmid's [Entitas-CSharp](https://github.com/sschmid/Entitas-CSharp) and [Atom proof-of-concept](https://github.com/sschmid/Entitas-CSharp/issues/902) but for the go language. `ZincECS` uses code-generation to achieve a nice API similar to EntitasECS using the built-in `ZincCLI`. This package puts focus on modularity and ease of use with performance coming in as a close second. 

# Installation
This will install the `ZincCLI` along with the `zinc` package.
```golang
go get github.com/SirMetathyst/zinc/...
```

# Quickstart

There isn't much we can do without any components so lets go ahead and generate some before we start.

```
zinc <component> -package|-import|-name|-var|-extras|-unique
---
zinc component -package components -name position -var x:float32 -var y:float32
zinc component -package components -name velocity -var x:float32 -var y:float32
```
We first execute the `ZincCLI` and pass in some arguments. The `-package` argument means we want our generated component to have the package name of `components` and `-name` specifies the name of the component with `-var` being the data. The format must be in `name:type` or `type`.

## Components

```golang
package main

import (
    "github.com/SirMetathyst/zinc"

    // importing your component package will
    // automatically register component types
    // with the default entity manager
    // see generated files for how to do it manually
    // if required. The convention recommended is to have the "kit"
    // suffix for packages which extend zinc funtionality.
    "path/to/yourkit"
)

func main() {

    // create an entity
    // uses the built-in default entity manager
    id := zinc.CreateEntity()

    // Adding a component on an entity.
    // The `Add` method could return an error which will result in 
    // ErrEntityComponentAlreadyExists or ErrEntityNotFound
    err := yourkit.AddPosition(id, yourkit.ZPositionData{10, 10})

    // Deleting a component on an entity.
    // The `Delete` method could return an error which will result in 
    // ErrEntityComponentNotFound or ErrEntityNotFound
    err := yourkit.DeletePosition(id)

    // Updating a component on an entity.
    // The `Update` method could return an error which will result in 
    // ErrEntityComponentNotFound or ErrEntityNotFound
    err := yourkit.UpdatePosition(id, yourkit.ZPositionData{10, 10})

    // Updating a component on an entity silently.
    // The `UpdateSilently` method will update the component
    // on an entity but will not notify groups. This method 
    // could return an error which will result in 
    // ErrEntityComponentNotFound or 
    // ErrEntityNotFound
    err := yourkit.UpdatePositionSilently(id, yourkit.ZPositionData{10, 10})

    // Setting a component on an entity.
    // The `Set` method will call the `Add` method if the component
    // has not been added to an entity and if it already exists it will
    // call `Update` instead and return an error if any
    err := yourkit.SetPosition(id, yourkit.ZPositionData{10, 10})

    // getting a component from an entity.
    // It can return an error which will result in 
    // ErrEntityComponentNotFound or ErrEntityNotFound
    position, err := yourkit.Position(id)

    // You can generate an optional API which prefix as "Must"
    // which allows you to write code without checking the errors
    // yourself however these methods expect the error to be nil and
    // will panic if they are not
    pos := yourkit.MustPosition(id) 

    // You can also pass in a different entity manager
    // through the methods which end in X
    pos := yourkit.MustPositionX(myOtherEntityManager, id)
}
```
```golang
// Full list of possible functions for generated component `position`

AddPositionX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) error
MustAddPositionX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) // optional extra
AddPosition(id zinc.ZEntityID, data ZPositionData) error
MustAddPosition(id zinc.ZEntityID, data ZPositionData) // optional extra

UpdatePositionSilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) error
MustUpdatePositionSilentlyX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) // optional extra
UpdatePositionSilently(id zinc.ZEntityID, data ZPositionData) error
MustUpdatePositionSilently(id zinc.ZEntityID, data ZPositionData) // optional extra
UpdatePositionX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) error
MustUpdatePositionX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) // optional extra
UpdatePosition(id zinc.ZEntityID, data ZPositionData) error
MustUpdatePosition(id zinc.ZEntityID, data ZPositionData) // optional extra

SetPositionX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) error
MustSetPositionX(e *zinc.ZEntityManager, id zinc.ZEntityID, data ZPositionData) // optional extra
SetPosition(id zinc.ZEntityID, data ZPositionData) error
MustSetPosition(id zinc.ZEntityID, data ZPositionData) // optional extra

HasPositionX(e *zinc.ZEntityManager, id zinc.ZEntityID) bool
HasPosition(id zinc.ZEntityID) bool 

Position2X(e *zinc.ZEntityManager, id zinc.ZEntityID) (ZPositionData, error)
MustPositionX(e *zinc.ZEntityManager, id zinc.ZEntityID) ZPositionData // optional extra
Position(id zinc.ZEntityID) (ZPositionData, error)
MustPosition(id zinc.ZEntityID) ZPositionData // optional extra

DeletePositionX(e *zinc.ZEntityManager, id zinc.ZEntityID) error
MustDeletePositionX(e *zinc.ZEntityManager, id zinc.ZEntityID) // optional extra
DeletePosition(id zinc.ZEntityID) error
MustDeletePosition(id zinc.ZEntityID) // optional extra
```

You can also generate unique components not bound to a specific entity.

```
zinc <component> -package|-import|-name|-var|-extras|-unique
---
zinc component -package components -name active -unique -var 
```



## Entity Manager
```golang
/////////////////////
// Create components 
/////////////////////

// ...

// get all created entity ids in the default
// entity manager instance
entities := zinc.Entities()

// TODO: add examples for the rest 
```

## Groups

```golang
// create a matcher 
matcher := zinc.AllOf(yourkit.ZPosition, yourkit.ZVelocity)

// create a group
group1 := zinc.Group(matcher)

// TODO: complete section
```

## Collector

```golang

```

## Systems
Zinc has a simple, built-in way  to init/update/cleanup your systems

```golang
sys := zinc.NewSystems()
sys.Add(yourkit.NewPositionSystem())
sys.Add(yourkit.NewAllSystems()...)

// init systems, must have `Initialize()` method
sys.Initialize()

// update systems, must have `Update(dt float64)` method
sys.Update(deltaTime)

// cleanup systems, must have `Cleanup()` method
sys.Cleanup()

// shutdown systems, must have `Shutdown()` method
sys.Shutdown()

```
then you can implement a system for moving your position components around. add the system to `Systems` and loop through those and that's your game loop. You can even have systems for drawing things in a different `Systems` instance and execute them at different times. You can even return `Systems` of `Systems` and update them as a group of systems as long as that type has the supported method.
```golang

import (
    "github.com/SirMetathyst/zinc"
    "github.com/xxx/yourkit"
)

// PositionSystem ...
type PositionSystem struct {
	group   zinc.ZGroup
	em      *zinc.ZEntityManager
}

// NewPositionSystem ...
func NewPositionSystem() *PositionSystem {
	return &PositionSystem{
		em:     zinc.Default(),
		group:  zinc.Default().Group(zinc.AllOf(yourkit.ZPosition, yourkit.ZVelocity)),
	}
}

// NewPositionSystemWith ...
func NewPositionSystemWith(em *zinc.ZEntityManager) *PositionSystem {
	return &PositionSystem{
		em:     em,
		group:  em.Group(zinc.AllOf(yourkit.ZPosition, yourkit.ZVelocity)),
	}
}

// Update ...
func (s PositionSystem) Update(dt float64) {
	for _, id := range s.group.Entities() {
		velocity := yourkit.MustVelocityX(s.em, id)
		position := yourkit.MustPositionX(s.em, id)
		position.X += velocity.X * dt
		position.Y += velocity.Y * dt
		yourkit.MustSetPositionX(s.em, id, position)
	}
}
```

# Contributing
I dont really have a contribution guideline. Just post an issue or pull request if you'd like to add or change something in `ZincECS`. I generally welcome pull requests but don't be disappointed if it gets rejected. You can always fork it.

# Projects/Examples I use ZincECS in
- [Zinckit](https://github.com/SirMetathyst/zinckit) - Collection of Systems and Components for `ZincECS` projects
<!--
- [Zincbird](https://github.com/SirMetathyst/zincbird) - Flappy birds clone written with `ZincECS`
- [Zincpong](https://github.com/SirMetathyst/zincbird) - Pong clone written with `ZincECS`-->

