package main

import "fmt"

//Closure - A piece of function or operation is enclosed inside another

// Incrementor returns a function and everytime it's called, the value will be incremented by 1
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	f := incrementor() // This returns a function
	fmt.Println(f())   // This calls the function and internally increments the value
	fmt.Println(f())

}
