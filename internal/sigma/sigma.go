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
	upper := int(math.Pow10(digits))

	a = rand.IntN(upper)
	b = rand.IntN(upper)

	// Ensure positive difference
	if a < b {
		a, b = b, a
	}

	diff = a - b

	return
}

func Multiplication(digits int) (a int, b int, prod int) {
	upper := int(math.Pow10(digits))

	a = rand.IntN(upper)
	b = rand.IntN(upper)
	prod = a * b

	return
}

func Division(digits int) (a int, b int, quot int) {
	upper := int(math.Pow10(digits))

	a = rand.IntN(upper)

	// Ensure integer quotient
	sqrt := int(math.Sqrt(float64(a)))
	factors := make([]int, 0, 2*sqrt) // 2sqrt(a) upper bound

	for i := 1; i <= sqrt; i++ {
		if a%i == 0 {
			factors = append(factors, i)
		}
	}

	b = factors[rand.IntN(len(factors))]

	quot = a / b

	return
}
