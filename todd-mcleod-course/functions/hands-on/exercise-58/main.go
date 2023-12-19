package main

import "fmt"

func foo() int {
	return 5
}

func bar() (int, string) {
	return 5, "bar"
}

func main() {
	f := foo()
	fmt.Println(f)

	b, c := bar()
	fmt.Println(b, c)

}
