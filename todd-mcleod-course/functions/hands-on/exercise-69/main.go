package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("This function is assigned to variable x")
	}
	y := func(a string) {
		fmt.Println("This function is assigned to variable y with parameter", a)
	}
	x()
	y("ABC")
}
