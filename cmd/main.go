package main

import (
	"fmt"

	"github.com/achen318/acplay.games/internal/sigma"
)

func main() {
	a, b, sum := sigma.Addition(2)
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	a, b, diff := sigma.Subtraction(2)
	fmt.Printf("%d - %d = %d\n", a, b, diff)

	a, b, prod := sigma.Multiplication(2)
	fmt.Printf("%d * %d = %d\n", a, b, prod)

	a, b, quot := sigma.Division(2)
	fmt.Printf("%d / %d = %d\n", a, b, quot)
}
