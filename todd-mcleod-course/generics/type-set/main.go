package main

import "fmt"

// Create an interface with the set of types that must be used and use them in function signature
type myNumbers interface {
	int | float64
}

// Use genric type 'T' by specifying the type set.
func addT[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(5, 6))
	fmt.Println(addT(5.8, 7.8))
}
