# Zinc 

[![codecov](https://codecov.io/gh/SirMetathyst/zinc/branch/master/graph/badge.svg)](https://codecov.io/gh/SirMetathyst/zinc)
[![Build Status](https://travis-ci.com/SirMetathyst/zinc.svg?branch=master)](https://travis-ci.com/SirMetathyst/zinc)

`ZincECS` is an entity-component-system package inspired by Simon Schmid's [Entitas-CSharp](https://github.com/sschmid/Entitas-CSharp) and [Atom proof-of-concept](https://github.com/sschmid/Entitas-CSharp/issues/902) but for the go language. I had searched for an ECS package for golang but I couldn't find one which I liked from the short few that I did find. So I decided to write my own.

Zinc focuses mainly on modularity of components/systems while performance coming in as a close second. 

# Installation
This will install the Zinc CLI along with the `zinc` package.
```golang
go get github.com/SirMetathyst/zinc/...
```

# Quickstart

There isn't much we can do without any components so lets go ahead and generate some before we start.

```
zinc component add -p components -n position -d x:float32 -d y:float32 -o ./components
zinc component add -p components -n velocity -d x:float32 -d y:float32 -o ./components
```
So, What's going on here? We're calling the `Zinc` CLI and passing in some arguments. Firstly the `p` param tells the `Zinc` CLI that we want the generated file to have the package name of `components`. Then we're telling it we want a component with the name of `position` and defining some data types for that component. The format must be in `name:type` but there are no checks done for whether the type is valid. It just replaces a value in the template. Lastly we give it a folder we want our component files to live in. Now we have some components generated to play with.

```golang
package main

import (
    "github.com/SirMetathyst/zinc"

    // importing your components package will
    // automatically register component types
    // with the default entity manager
    // see generated files for how to do it manually
    // if required.
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
Zinc has a built-in way  to init/update/cleanup your systems

```golang
sys := zinc.NewSystems()
sys.Add(yourkit.NewPositionSystem())

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
I dont really have a contribution guideline. Just post an issue or pull request if you'd like to add or change something in `Zinc`. I generally welcome pull requests but don't be disappointed if it gets rejected. You can always fork it.

# Projects/Examples that I use `Zinc` in
- [Zinckit](https://github.com/SirMetathyst/zinckit) - Collection of Systems and Components for `Zinc` projects
<!--
- [Zincbird](https://github.com/SirMetathyst/zincbird) - Flappy birds clone written with `Zinc`
- [Zincpong](https://github.com/SirMetathyst/zincbird) - Pong clone written with `Zinc`-->

