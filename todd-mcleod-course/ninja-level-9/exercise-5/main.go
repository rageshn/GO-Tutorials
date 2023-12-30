package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	const goRoutinesCount = 100
	var wg sync.WaitGroup
	var counter int64
	wg.Add(goRoutinesCount)

	for i := 0; i < goRoutinesCount; i++ {
		func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
