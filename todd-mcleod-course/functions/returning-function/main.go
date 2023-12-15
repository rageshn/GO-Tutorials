package main

import "fmt"

func foo() int {
	return 50
}

// Function which return another function which internally returns an integer
func bar() func() int {
	return func() int {
		return 100
	}
}

func main() {
	x := foo()
	fmt.Println(x)

	y := bar() // This returns a function
	ret := y() //Calling the returned function returns the integer value
	fmt.Println(ret)
}
