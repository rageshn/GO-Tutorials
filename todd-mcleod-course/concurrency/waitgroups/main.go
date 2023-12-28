package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("FOO:", i)
	}
	wg.Done() //Once this is marked as done, it will reduce the go routines count by 1.
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("BAR:", i)
	}
}

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())

	wg.Add(1) // Indicates we have to include one go routine and have to wait till it finishes.

	go foo() //Creates a new go routine
	bar()
	fmt.Println("GO-Routines:", runtime.NumGoroutine())

	wg.Wait() // Waits till all go routines added to the wait group finishes.
}
