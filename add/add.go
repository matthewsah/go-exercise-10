// Package add is a package that contains the mathematical
// function Add which adds two integers together
// in the future there could be more adding functions which
// could add numbers of different types
package add

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

// Add takes in two numbers and returns their sum
func Add[N Number](num1 N, num2 N) N {
	return num1 + num2
}
