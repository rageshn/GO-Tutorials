package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("Value of x is %v \n", x)
	if x <= 100 {
		fmt.Println("Value is between 0 and 100")
	} else if x > 100 && x <= 200 {
		fmt.Println("Value is between 100 and 200")
	} else if x > 200 && x < 250 {
		fmt.Println("Value is between 200 and 250")
	}
}
