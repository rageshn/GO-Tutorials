package main

import "fmt"

/*
func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}*/

func main() {
	greeting := "Hi There!"

	go (func() {
		fmt.Println(greeting)
	})()
}
