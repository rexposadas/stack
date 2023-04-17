package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Notes: [T int | float64] means you can pass in either an int or a float64.
func totalManual[T int | float64](a T, b T) T {
	return a + b
}

// Notes: Using lib for numbers, so we don't have to list types ourserlves like in totalManual.
func total[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {

	// Notes: passed in arguments needs to be of the same type
	fmt.Printf("totalManual: %d\n", totalManual(3, 2))

	fmt.Printf("totalManual: %0.2f\n", totalManual(5.8, 2.9))
}
