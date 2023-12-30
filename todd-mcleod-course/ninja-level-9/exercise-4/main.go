package main

import (
	"fmt"
	"sync"
)

func main() {

	const goRoutinesCount = 100
	var wg sync.WaitGroup
	counter := 0
	wg.Add(goRoutinesCount)
	var mtx sync.Mutex

	for i := 0; i < goRoutinesCount; i++ {
		func() {
			mtx.Lock()
			v := counter
			v++
			counter = v
			mtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
