package main

import "fmt"

func main() {
	c := make(chan int)

	// Writing to a channel blocks the operation but as its in a separate go routine,
	// the write to channel and read from main routine happens at the same time. So this works without interruption.
	go func() {
		c <- 42
	}()

	fmt.Println(<-c) // Take the value out of channel
}
