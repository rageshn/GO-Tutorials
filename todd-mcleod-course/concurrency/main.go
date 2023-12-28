package main

import (
	"fmt"
	"runtime"
)

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("FOO:", i)
	}
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

	go foo() //Creates a new go routine
	bar()
	fmt.Println("GO-Routines:", runtime.NumGoroutine())

}
