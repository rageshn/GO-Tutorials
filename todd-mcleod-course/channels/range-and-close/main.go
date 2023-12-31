package main

import "fmt"

// Takes send only channel as parameter
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// If the channel is not closed, it will create a deadlock as the range operation will wait for reading data.
	close(c)
}

func main() {
	c := make(chan int)

	// Channels are reference types, so the underlying address will be sent as parameter.

	// Creates a new go routine to send the data to channel.
	go send(c)

	// Loops through the channel and reads the value.
	// As soon as the value is written to the channel in the other go routine, it will be read from the main routine.
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Exiting...")
}
