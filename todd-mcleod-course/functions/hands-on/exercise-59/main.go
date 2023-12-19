package main

import "fmt"

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar(i []int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func main() {

	xi := []int{1, 2, 3, 4, 5}
	x := foo(xi...)
	fmt.Println(x)

	y := bar(xi)
	fmt.Println(y)

}
