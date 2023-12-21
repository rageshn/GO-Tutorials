package main

import "fmt"

func exponent(a float64) func() float64 {
	x := a
	return func() float64 {
		x *= x
		return x
	}
}

func main() {
	exp := exponent(2)

	fmt.Println(exp())
	fmt.Println(exp())
	fmt.Println(exp())
	fmt.Println(exp())
	fmt.Println(exp())

}
