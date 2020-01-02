# Atom 
`Atom` is an ECS (Entity-Component-System) package inspired by Simon Schmid's [Entitas-CSharp](https://github.com/sschmid/Entitas-CSharp) and [Atom proof-of-concept](https://github.com/sschmid/Entitas-CSharp/issues/902) but for the go language. I had searched for an ECS package for golang but I couldn't find one which I liked from the short few that I did find. So I decided to write my own.

atom focuses on modularity of components and systems while being as small as possible. Coming in as a close second is performance. 

# Contributing
I dont really have a contribution guideline. Just post an issue or pull request if you'd like to add or change something in `Atom`. I generally welcome pull requests but don't be disappointed if it gets rejected. You can always fork it.

# Installation
This will install the atom CLI along with the `Atom` package.
```golang
go get github.com/SirMetathyst/atom/cmd/...
```

# Quickstart

There isn't much we can do without any components so lets go ahead and generate some before we start.

```
atom component add -p components -n position -d x:float32 -d y:float32 -o ./components
atom component add -p components -n velocity -d x:float32 -d y:float32 -o ./components
```
So, What's going on here? We're calling the `Atom` CLI and passing in some arguments. Firstly the `p` param tells the `Atom` CLI that we want the generated file to have the package name of `components`. Then we're telling it we want a component with the name of `position` and defining some data types for that component. The format must be in `name:type` but there are no checks done for whether the type is valid. It just replaces a value in the template. Lastly we give it a folder we want our component files to live in. Now we have some components generated to play with.

```golang
package main

import (
    "github.com/SirMetathyst/atom"

    // importing your components package will
    // automatically register component types
    // with the default entity manager
    // see generated files for how to do it manually
    // if required.
    "xxx/xxx/to/your/yourkit"
)

func main() {

    // create an entity
    // uses the built-in default entity manager
    id := atom.CreateEntity()

    // we can already use our component types 
    // setting a component will add or update it
    yourkit.SetPosition(id, yourkit.PositionData{10, 10})


    // get position with id
    pos := yourkit.Position(id)


    // there is also an API for passing in a different entity manager
    // these end in X
    pos := yourkit.PositionX(entityManager, id)

    // getting all entity ids in the entity manager
    entities := atom.Entities()


    // getting groups of entities with specific components
    // will return a entity group that has position and velocity component
    group1 := atom.Group(atom.AllOf(yourkit.PositionKey, yourkit.VelocityKey))

    // ids of entities in group
    group1.Entities()

    // does the group have an entity?
    group1.HasEntity(id)

    // does the component exist for entity
    ok := atom.HasPosition(id)

    // delete the position component 
    atom.DeletePosition(id)

    // will return a entity group that has position but not velocity component
    group2 := atom.Group(atom.AllOf(yourkit.PositionKey).NoneOf(yourkit.VelocityKey))  
}
```

## Systems
Atom has a built in way  to init/update/cleanup your systems

```golang
sys := atom.NewSystems()
sys.Add(yourkit.NewPositionSystem())

// init systems, must have `Initialize()` method
sys.Initialize()

// update systems, must have `Update(dt float64)` method
sys.Update(deltaTime)

// cleanup systems, must have `Cleanup()` method
sys.Cleanup()

```
then you can implement a system for moving your position components around. add the system to a slice and loop through those and that's your game loop. You can even have systems for drawing things in a different slice and execute them at different times.
```golang

import (
    "github.com/SirMetathyst/atom"
    "github.com/xxx/yourkit"
)

// PositionSystem ...
type PositionSystem struct {
	g         atom.G
	em *atom.EntityManager
}

// NewPositionSystem ...
func NewPositionSystem() *PositionSystem {
	return &PositionSystem{
		em: atom.Default(),
		g:  atom.Default().Group(atom.AllOf(yourkit.PositionKey, yourkit.VelocityKey)),
	}
}

// NewPositionSystemWith ...
func NewPositionSystemWith(em *atom.EntityManager) *PositionSystem {
	return &PositionSystem{
		em: em,
		g:  em.Group(atom.AllOf(yourkit.PositionKey, yourkit.VelocityKey)),
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

# Projects/Examples that I use `Atom` in
- [Atomkit](https://github.com/SirMetathyst/atomkit) - Collection of Systems and Components for `Atom` projects
- [Atombird](https://github.com/SirMetathyst/atombird) - Flappy birds clone written with `Atom`

