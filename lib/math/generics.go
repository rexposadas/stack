package math

import "golang.org/x/exp/constraints"

// Sum .
// Notes: Using lib for numbers, so we don't have to list types ourselves like in totalManual.
func Sum[T constraints.Ordered](a ...T) T {
	var total T
	for _, v := range a {
		total += v
	}

	return total
}
