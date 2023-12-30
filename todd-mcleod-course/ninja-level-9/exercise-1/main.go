package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("FOO:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("BAR:", i)
	}
	wg.Done()
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("TEST:", i)
	}
}

func main() {
	wg.Add(2)
	fmt.Println("Foo called in separate go routine")
	go foo()
	fmt.Println("Bar called in separate go routine")
	go bar()
	fmt.Println("Test called in separate go routine")
	test()
	wg.Wait()
}
