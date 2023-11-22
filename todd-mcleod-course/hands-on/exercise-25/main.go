package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
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
	}
}
