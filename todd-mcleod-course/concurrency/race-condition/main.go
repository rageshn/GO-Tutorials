package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Data race occur because lots of go routines are accessing the same shared variable.
*/

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())

	counter := 0
	const goRoutinesCount = 100
	var wg sync.WaitGroup
	wg.Add(goRoutinesCount)

	for i := 0; i < goRoutinesCount; i++ {
		go func() {
			// Assign counter to a variable, increment it and assign back.
			v := counter
			runtime.Gosched() // Yields the processor so that other routines can run
			v += 1
			counter = v
			wg.Done()
		}()
		fmt.Println("GO-Routines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
