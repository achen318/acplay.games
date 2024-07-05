package sigma

import (
	"math"
	"math/rand/v2"
)

func Addition(digits int) (a, b, sum int) {
	upper := int(math.Pow10(digits))

	a = rand.IntN(upper)
	b = rand.IntN(upper)
	sum = a + b

	return
}

func Subtraction(digits int) (a int, b int, diff int) {
	return 1, 2, 3
}

func Multiplication(digits int) (a int, b int, prod int) {
	return 1, 2, 3
}

func Division(digits int) (a int, b int, quot int) {
	return 1, 2, 3
}
