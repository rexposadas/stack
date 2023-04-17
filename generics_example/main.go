package main

import "fmt"

// Notes:
// [T int | float64] means you can pass in either an int or a float64.
func total[T int | float64](a T, b T) T {
	return a + b
}

func main() {

	// Notes: passed in arguments needs to be of the same type
	fmt.Printf("total: %d\n", total(3, 2))

	fmt.Printf("total: %0.2f\n", total(5.8, 2.9))
}
