package main

import "fmt"

// Create an interface with the set of types that must be used and use them in function signature.
// Tilde in the datatypes denotes - Include all values of type int or float64, and also any values whose underlying value is of type int or float.
type myNumbers interface {
	~int | ~float64
}

// Use genric type 'T' by specifying the type set.
func addT[T myNumbers](a, b T) T {
	return a + b
}

// Create alias for underlying data types
type alias int

func main() {
	var num alias = 10
	fmt.Println(addT(num, 6))
	fmt.Println(addT(5.8, 7.8))
}
