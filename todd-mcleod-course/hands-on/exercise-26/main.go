package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 20
	for i > 0 {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("Iteration: ", i, "Value of x is 0")
		case 1:
			fmt.Println("Iteration: ", i, "Value of x is 1")
		case 2:
			fmt.Println("Iteration: ", i, "Value of x is 2")
		case 3:
			fmt.Println("Iteration: ", i, "Value of x is 3")
		case 4:
			fmt.Println("Iteration: ", i, "Value of x is 4")
		}
		i -= 1
	}
}
