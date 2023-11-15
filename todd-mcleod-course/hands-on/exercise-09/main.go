package main

import "fmt"

var x = 40

const y = 45

func main() {
	fmt.Printf("Value of x is %v and type is %T \n", x, x)
	fmt.Printf("Value of y is %v and type is %T \n", y, y)

	z := 50
	fmt.Printf("Value of z is %v and type is %T \n", z, z)
}
