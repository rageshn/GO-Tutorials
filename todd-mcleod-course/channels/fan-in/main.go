package main

import (
	"fmt"
	"sync"
)

//Taking values from multiple channels and putting those values into one channel.

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	// Separate go routine to loop even channel and send to fanIn channel
	go func() {
		for i := range even {
			fanIn <- i
		}
		wg.Done()
	}()

	// Separate go routine to loop odd channel and send to fanIn channel
	go func() {
		for i := range odd {
			fanIn <- i
		}
		wg.Done()
	}()

	wg.Wait() // Wait till both wait groups completes
	close(fanIn)

}

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanIn)

	for i := range fanIn {
		fmt.Println(i)
	}
	fmt.Println("Exiting..")

}
