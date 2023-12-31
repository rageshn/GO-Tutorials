package main

import "fmt"

// Takes send only channel as parameter
func send(c chan<- int) {
	c <- 50
}

// Takes receive only channel as parameter
func receive(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)

	// Channels are reference types, so the underlying address will be sent as parameter.

	// Creates a new go routine to send the data to channel.
	go send(c)

	// As this is part of the main routine, this will wait and reads the data from the channel.
	receive(c)

	fmt.Println("Exiting...")
}
