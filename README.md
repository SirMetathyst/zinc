# Zinc 

[![codecov](https://codecov.io/gh/SirMetathyst/zinc/branch/master/graph/badge.svg)](https://codecov.io/gh/SirMetathyst/zinc)
[![Build Status](https://travis-ci.com/SirMetathyst/zinc.svg?branch=master)](https://travis-ci.com/SirMetathyst/zinc)

`ZincECS` is an entity-component-system package inspired by Simon Schmid's [Entitas-CSharp](https://github.com/sschmid/Entitas-CSharp) and [Atom proof-of-concept](https://github.com/sschmid/Entitas-CSharp/issues/902) but for the go language. ZincECS uses code-generation to achieve a nice API similar to EntitasECS using the built-in ZincCLI. This package puts focus on modularity and ease of use with performance coming in as a close second. 

# Installation
This will install the ZincCLI along with the `zinc` package.
```golang
go get github.com/SirMetathyst/zinc/...
```

# Quickstart

There isn't much we can do without any components so lets go ahead and generate some before we start.

```
zinc component add -p components -n position -d x:float32 -d y:float32 -o ./components
zinc component add -p components -n velocity -d x:float32 -d y:float32 -o ./components
```
We first call the `ZincCLI` and pass in some arguments. The `p` argument tells the `ZincCLI` that we want our generated component to have the package name of `components`. Then we tell it we want a component with the name of `position` and specify the data. The format must be in `name:type` but no checks are done to ensure valid go code. Lastly, we give it a folder where we want our component files to be generated in. Now we have some components generated to play with.

```golang
package main

import (
    "github.com/SirMetathyst/zinc"

    // importing your component package will
    // automatically register component types
    // with the default entity manager
    // see generated files for how to do it manually
    // if required. The convention I use is to call it a "kit"
    // where the root contains all your components and /systems
    // contains all logic for those components
    "xxx/xxx/to/yourkit"
)

func main() {

    // create an entity
    // uses the built-in default entity manager
    id := zinc.CreateEntity()

    // we can already use our component types 
    // setting a component will add or update it
    yourkit.SetPosition(id, yourkit.PositionData{10, 10})


    // get position with id
    pos := yourkit.Position(id)


    // there is also an API for passing in a different entity manager
    // these end in X
    pos := yourkit.PositionX(entityManager, id)

    // getting all entity ids in the entity manager
    entities := zinc.Entities()


    // getting groups of entities with specific components
    // will return a entity group that has position and velocity component
    group1 := zinc.Group(zinc.AllOf(yourkit.PositionKey, yourkit.VelocityKey))

    // ids of entities in group
    group1.Entities()

    // does the group have an entity?
    group1.HasEntity(id)

    // does the component exist for entity
    ok := yourkit.HasPosition(id)

    // delete the position component 
    yourkit.DeletePosition(id)

    // will return a entity group that has position but not velocity component
    group2 := zinc.Group(zinc.AllOf(yourkit.PositionKey).NoneOf(yourkit.VelocityKey))  
}
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

```
then you can implement a system for moving your position components around. add the system to `Systems` and loop through those and that's your game loop. You can even have systems for drawing things in a different `Systems` instance and execute them at different times. You can even return `Systems` of `Systems` and update them as a group of systems as long as that type has the supported method.
```golang

import (
    "github.com/SirMetathyst/zinc"
    "github.com/xxx/yourkit"
)

// PositionSystem ...
type PositionSystem struct {
	g         zinc.G
	em *zinc.EntityManager
}

// NewPositionSystem ...
func NewPositionSystem() *PositionSystem {
	return &PositionSystem{
		em: zinc.Default(),
		g:  zinc.Default().Group(zinc.AllOf(yourkit.PositionKey, yourkit.VelocityKey)),
	}
}

// NewPositionSystemWith ...
func NewPositionSystemWith(em *zinc.EntityManager) *PositionSystem {
	return &PositionSystem{
		em: em,
		g:  em.Group(zinc.AllOf(yourkit.PositionKey, yourkit.VelocityKey)),
	}
}

// Update ...
func (s PositionSystem) Update(dt float64) {
	for _, id := range s.g.Entities() {
		velocity := yourkit.VelocityX(s.em, id)
		position := yourkit.PositionX(s.em, id)
		position.X += velocity.X * dt
		position.Y += velocity.Y * dt
		yourkit.SetPositionX(s.em, id, position)
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

