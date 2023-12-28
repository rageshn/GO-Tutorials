package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())

	var counter int64
	const goRoutinesCount = 100
	var wg sync.WaitGroup
	wg.Add(goRoutinesCount)

	for i := 0; i < goRoutinesCount; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // Gets the address to counter and increments the value by 1
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GO-Routines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GO-Routines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
