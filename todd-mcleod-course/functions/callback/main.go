package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// This function take two values and one function as parameter e.g., add, sub etc.
// This applies the passed function on the two values
func doOperation(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	x := doOperation(1, 2, add) //Pass add function as a parameter and that function will be applied on two values
	fmt.Println("Sum:", x)

	y := doOperation(67, 89, sub)
	fmt.Println("Sub:", y)

}
