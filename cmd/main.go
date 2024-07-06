package main

import (
	"fmt"

	"github.com/achen318/acplay.games/internal/sigma"
)

func main() {
	// For + and -, operands are [2, 100]
	ranges := sigma.Ranges{
		A_lower: 2, A_upper: 100,
		B_lower: 2, B_upper: 100,
	}

	eq := sigma.Addition(ranges)
	fmt.Printf("%d + %d = %d\n", eq.A, eq.B, eq.Result)

	eq = sigma.Subtraction(ranges)
	fmt.Printf("%d - %d = %d\n", eq.A, eq.B, eq.Result)

	// For * and /, operands are [2, 12]
	ranges.A_upper = 12

	eq = sigma.Multiplication(ranges)
	fmt.Printf("%d * %d = %d\n", eq.A, eq.B, eq.Result)

	eq = sigma.Division(ranges)
	fmt.Printf("%d / %d = %d\n", eq.A, eq.B, eq.Result)

}
