package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	x := doMath(10, 10, add)
	fmt.Println("Sum:", x)

	y := doMath(15, 10, subtract)
	fmt.Println("Diff:", y)

}
