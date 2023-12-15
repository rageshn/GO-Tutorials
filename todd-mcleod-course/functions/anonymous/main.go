package main

/*
	func()(p parameters) (r return) { }()
*/

import "fmt"

func foo() {
	fmt.Println("FOO")
}

func main() {
	foo()

	func() {
		fmt.Println("Anonymous function")
	}()

	func(s string) {
		fmt.Println("My anonymous name is", s)
	}("ABC")

	// Assigning function to a variable
	x := func() {
		fmt.Println("Anonymous function called as x")
	}

	y := func(s string) {
		fmt.Println("Anonymous function called as y with parameter", s)
	}

	x()
	y("ABC")

}
