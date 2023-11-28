package main

import "fmt"

func main() {
	x := []int{41, 42, 43, 44, 45}

	for index, value := range x {
		fmt.Println("index =", index, "Value =", value)
	}
}
