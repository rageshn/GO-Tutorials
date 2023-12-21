package main

import "fmt"

func printText(x string) {
	fmt.Println(x)
}

func showText(x string, f func(string)) {
	f(x)
}

func multiply(a int, b int) int {
	return a * b
}

func doOperation(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	showText("Hello", printText)
	v := doOperation(14, 7, multiply)
	fmt.Println(v)
}
