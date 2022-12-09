// ======================================================================================================================
// using the Template design pattern from - https://refactoring.guru/design-patterns/template-method/go/example#example-0
// ======================================================================================================================
package main

import (
	"fmt"
	"reflect"
)

// Duck interface to set how flying and quacking should be handled
type IDuck interface {
	performFly() string
	performQuack() string
}

// Base duck "class"
type Duck struct {
	// duck IDuck
}

// Add the swim method to the Duck "class"
func (d Duck) Swim() string {
	return "swimming"
}

// create a new "class" that inherits from the Duck "class"
type Mallard struct {
	Duck
}

// Add performFly to the Mallard "class" to meet the Duck interface
func (m Mallard) performFly() string {
	return fly()
}

// use fly for ducks that can fly
func fly() string {
	return "flying"
}

// use cannot fly for ducks that cannot fly (ex wooden, rubber, etc)
func cannotFly() string {
	return "I cannot fly..."
}

// Add performQuack to the Mallard "class" to meet the Duck interface
func (m Mallard) performQuack() string {
	return quack()
}

// use quack for ducks that can quack
func quack() string {
	return "quack quack"
}

// use squeak for ducks that squeak (ex rubber)
func squeak() string {
	return "squeak squeak"
}

// use muteQuack for ducks that cannot quack (ex wooden)
func muteQuack() string {
	return "I cannot quack..."
}

// create a new "class" that inherits from the Duck "class"
type WoodenDuck struct {
	Duck
}

// Add performFly to the WoodenDuck "class" to meet the Duck interface
func (w WoodenDuck) performFly() string {
	return cannotFly()
}

// Add performQuack to the WoodenDuck "class" to meet the Duck interface
func (w WoodenDuck) performQuack() string {
	return muteQuack()
}

func main() {
	m := Mallard{}
	// use the method below if you don't need to loop over the ducks,
	// or want to use their methods individually
	// fmt.Println(m.performFly())
	// fmt.Println(m.performQuack())
	// fmt.Println(m.Swim())

	w := WoodenDuck{}
	// fmt.Println(w.performFly())

	// loop over the ducks - the ducks MUST implement IDuck (have all the methods from IDUCK)
	ducks := []IDuck{m, w}
	for _, duck := range ducks {
		fmt.Println(reflect.TypeOf(duck).Name(), duck.performFly())
	}
}
