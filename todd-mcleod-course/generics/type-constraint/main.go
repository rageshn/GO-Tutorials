package main

import "fmt"

// Two separate functions are needed to add int and float. Instead of two functions, we can use generic function.

func addI(a int, b int) int {
	return a + b
}

func addF(a float64, b float64) float64 {
	return a + b
}

// Use genric type 'T' by specifying which data types must be accepted. The inputs and return value must also be of type T
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(5, 6))
	fmt.Println(addF(5.8, 7.8))

	fmt.Println(addT(5, 6))
	fmt.Println(addT(5.8, 7.8))
}
