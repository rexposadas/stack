package main

import (
	"testing"
)

// An interface is a type defined by a set of methods.
type Animal interface {
	// structs must implement these to be considered a type of Animal.
	Eat()
	Size()
}

type Cat struct{}

func (c Cat) Eat()  {}
func (c Cat) Size() {}

type Table struct{}   // does not implement eat() so is not an Animal type
func (t Table) Size() {}

func Breed(a Animal) {
	a.Eat()
}

func TestSimple(t *testing.T) {

	c := Cat{} // cat is a type of animal
	Breed(c)   // Breed takes in an animal type, which cat is.

	// Will produce compile issues since a table is not an animal
	// t := Table{}
	// Breed(t)
}
