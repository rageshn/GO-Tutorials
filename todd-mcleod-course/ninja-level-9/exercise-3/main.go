package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const goRoutinesCount = 100
	var wg sync.WaitGroup
	counter := 0
	wg.Add(goRoutinesCount)

	for i := 0; i < goRoutinesCount; i++ {
		func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
