// Package math provides elementary math functions.
package math

import (
	"golang.org/x/exp/constraints"
)

// Number is a constraint that permits any integer or float type.
//
// Relies on [golang.org/x/exp/constraints.Integer] and [golang.org/x/exp/constraints.Float].
type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns a sum of its two parameters.
// Supports parameters of any integer or float type.
//
// More information on addition can be found at [Math is Fun].
//
// [Math is Fun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
