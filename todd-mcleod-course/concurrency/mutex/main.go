package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())

	counter := 0
	const goRoutinesCount = 100
	var wg sync.WaitGroup
	wg.Add(goRoutinesCount)

	var mtx sync.Mutex

	for i := 0; i < goRoutinesCount; i++ {
		go func() {
			mtx.Lock() // Lock the code so that other routines cannot run/update the variables.
			// Assign counter to a variable, increment it and assign back.
			v := counter
			runtime.Gosched() // Yields the processor so that other routines can run
			v += 1
			counter = v
			mtx.Unlock() // Unlocks the code block so that other routines can update the variable.
			wg.Done()
		}()
		fmt.Println("GO-Routines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
