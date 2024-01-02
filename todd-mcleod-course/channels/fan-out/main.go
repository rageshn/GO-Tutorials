package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Exiting...")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(value int) {
			c2 <- timeConsumingWork(value)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)

}

func timeConsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
