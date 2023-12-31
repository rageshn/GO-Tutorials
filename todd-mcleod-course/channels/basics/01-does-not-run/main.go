package main

import "fmt"

func main() {
	c := make(chan int)

	// Add a value to channel. But this does not work, as the send & receive to a channel must happen at the same time.
	// Here it's trying to add to channel, but the receiver is not ready.
	c <- 42

	fmt.Println(<-c) // Take the value out of channel
}
