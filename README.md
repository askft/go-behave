**<h1 align=center>Go Behave</h1>**

**<p align=center>An extensible Behavior Tree library in Go.</p>**

## Introduction

A behavior tree is a formalism for describing the behavior of an autonomous entity such as a robot or a non-player character in a video game. Behavior trees, by their nature, allow for highly modular design thanks to the composability of nodes, and the formalism makes task switching and state management trivial.

A behavior tree is a directed rooted tree with at most three categories of nodes: _composite nodes_, _decorator nodes_ and _leaf nodes_. Each category can contain many different _types_ of nodes. A _tick_ is sent from the root with a certain frequency, making a pre-order traversal of the tree. Each node type provides a different algorithm for processing the tick, explained below. Once a tick has been processed, the node returns a status to its parent - either _Failure_, _Running_ or _Success_.

Leaf nodes require a data context, which will be propagated throughout the tree via the `Tick` procedure. The context usually contains a reference to the owner of the tree, i.e. the entity for which the tree describes a behavior, and a _store_ or _blackboard_ which is a data storage system that can be shared between different entities - in a game this store would contain e.g. information about the game world, positions of all entities, and so on.

### Composite nodes

A composite node has a type and one or more children. Below are some common types of composite nodes:

#### Sequence

Tick children in order. Succeeds only if each child succeeds. Fails as soon as any child fails. Returns Running if the currently executing child returns Running.

#### Selector

Tick children in order. Succeeds as soon as a child succeeds. Fails only if each child fails. Returns Running if the currently executing child returns Running.

#### Random Sequence and Random Selector

Same as sequence and selector, except the order of the children is randomized before each time the node runs.

### Decorator nodes

A decorator node has a type and one child. Below are some common types of decorator nodes.

#### Inverter

The inverter node inverts (negates) the the result of its child, or returns Running if the child returns Running.

#### Repeater

A repeater node will re-tick its child when it returns Success or Failure. The number of times the child is re-ticked can be limited or unlimited. Often used to wrap the root of the tree in order to make the tree run indefinitely.

#### Delayer

A delayer node will always return Running during a certain amount of time, after which it tick its child and return its status.

### Leaf nodes

A leaf node, also called execution node, action node or task node, is normally specifically tailored to the application at hand. In a robotics context the task might be "Pick Up Object" or "Move Arm"; in a video game context it might be "Find Nearest Target" or "Attack Target".

A leaf node can also be a _condition node_ which returns Success or Failure depending on whether some condition is fulfilled in the environment.

## Usage

Please see the `examples` package. An extensive explanation of the library can be found below.

### Defining custom nodes

The library offers a set of common node types, although it's easy to make your own. The available common node types can be found in the `composite`, `decorator` and `action` packages. These are _not_ required for you to use the rest of the library, but greatly simplify usage.

Defining a custom node means having a `struct` embedding a pointer to a `core.T` where `T` is either `Composite`, `Decorator` or `Action`, and having the following methods defined on the type:

```go
Start(*Context)
Tick(*Context) Status
Stop(*Context)
String() string
```

The struct may also contain other fields that will be initialized in the node's _constructor_, which you also need to create. The constructor function needs have a type equal to one of `CompositeFn`, `DecoratorFn` or `ActionFn` (see [core/types.go](https://github.com/AlexanderSkafte/BehaviorTree/blob/master/core/types.go)). An example can be seen in [decorator/repeater.go](https://github.com/AlexanderSkafte/BehaviorTree/blob/master/decorator/repeater.go) (or any other type in the `composite`, `decorator` or `action` packages).

### Defining behavior trees

You may define a behavior tree in one of two ways: via a definition string that compiles to a behavior tree node, or by declaring the node directly in Go code. Please see [examples/behaviortree/](https://github.com/AlexanderSkafte/BehaviorTree/tree/master/examples/behaviortree), where both methods are used. Defining the behavior tree via a definition string requires you to register the functions in a `Registry` in order for the parser to recognize them. [registry/Registry.go](https://github.com/AlexanderSkafte/BehaviorTree/blob/master/registry/registry.go) offers a `NewDefault` function that will return a registry where the available common nodes have been registered. Please refer to that function when you want to find out how to register your own nodes.

The `BehaviorTree` type contains a reference to a root node, a reference to its owner, a reference to a `Registry` [\*] and a reference to a `Store` interface. The library offers a `Blackboard` ([store/Blackboard.go](https://github.com/AlexanderSkafte/BehaviorTree/blob/master/store/blackboard.go)) which implements the `Store` interface. A `BehaviorTree` is created via the `NewBehaviorTree` function which takes a `Config` type containing the types mentioned above. See [behaviortree.go](https://github.com/AlexanderSkafte/BehaviorTree/blob/master/behaviortree.go).

<p style="font-size:0.8em">[*] TODO: Will be removed as it's not really necessary if the used precompiles the definition string into a node</p>

## Important notice

The library is still in the development phase, so the API is not stable, and I can not guarantee that the implementation is bug-free.

## Installation

`$ go get github.com/alexanderskafte/go-behave`

## License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.
