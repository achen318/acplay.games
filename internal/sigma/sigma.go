package sigma

import (
	"math/rand/v2"
)

type Ranges struct {
	A_lower, A_upper, B_lower, B_upper int
}

type Equation struct {
	A, B, Result int
}

func randRange(lower, upper int) int {
	// Returns a random integer in [lower, upper]
	return lower + rand.IntN(upper-lower+1)
}

func Addition(ran Ranges) (eq Equation) {
	eq.A = randRange(ran.A_lower, ran.A_upper)
	eq.B = randRange(ran.B_lower, ran.B_upper)
	eq.Result = eq.A + eq.B

	return
}

func Subtraction(ran Ranges) (eq Equation) {
	// Subtraction is the inverse of addition
	eq = Addition(ran)
	eq.A, eq.Result = eq.Result, eq.A

	return
}

func Multiplication(ran Ranges) (eq Equation) {
	eq.A = randRange(ran.A_lower, ran.A_upper)
	eq.B = randRange(ran.B_lower, ran.B_upper)
	eq.Result = eq.A * eq.B

	return
}

func Division(ran Ranges) (eq Equation) {
	// Division is the inverse of multiplication
	eq = Multiplication(ran)
	eq.A, eq.Result = eq.Result, eq.A

	return
}
