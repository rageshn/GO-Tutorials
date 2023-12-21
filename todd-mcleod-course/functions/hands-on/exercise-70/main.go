package main

import "fmt"

func foo() func() {
	return func() {
		fmt.Println("Hai from inner function")
	}
}

func bar() func(string) {
	return func(y string) {
		fmt.Println("Hai from inner function with", y, "as parameter")
	}
}

func main() {
	f := foo()
	f()

	b := bar()
	b("Test")
}
