package main

import "fmt"

func foo() {
	fmt.Println("FOO")
}

func bar() {
	fmt.Println("BAR")
}

func test() {
	fmt.Println("TEST")
}
func main() {
	defer foo()
	fmt.Println("Statement after calling foo()")
	bar()
	fmt.Println("Statement after calling bar()")
	defer test()
	fmt.Println("Statement after calling test()")
}
