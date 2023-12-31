package main

import "fmt"

func main() {
	sendOnlyChannel := make(chan<- int, 1) // Arrow indicates that its a send only channel and we cant receive data from that.
	fmt.Printf("%T\n", sendOnlyChannel)

	sendOnlyChannel <- 50

	//fmt.Println(<-sendOnlyChannel) //-- This does not work, as we cant receive data from this channel.

	receiveOnlyChannel := make(<-chan int)
	fmt.Printf("%T\n", receiveOnlyChannel)
	//receiveOnlyChannel <- 50 // -- This does not work as we cant send data to this channel.

	fmt.Println(<-receiveOnlyChannel) // This will also create a deadlock as channel writes are blocking.
}
