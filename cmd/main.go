package main

import (
	"fmt"

	"github.com/achen318/acplay.games/internal/sigma"
)

func main() {
	a, b, sum := sigma.Addition(3)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}
