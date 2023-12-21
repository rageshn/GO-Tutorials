package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hai from anonymous function")
	}()

	func(x string) {
		fmt.Println("Parameter to anonymous function:", x)
	}("ABC")
}
