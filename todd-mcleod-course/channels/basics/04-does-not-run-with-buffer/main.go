package main

import "fmt"

func main() {

	// Buffer channel allows the specified number of values to be written to the channel, regardless of the read operation.
	// 1 indicates that channel can have 1 value in its buffer regardless of read.
	c := make(chan int, 1)

	// As the channel can have 1 value in its buffer, this will not be a blocking operation.
	c <- 42
	c <- 43 // But this will be a blocking operation. So the code will not run.

	fmt.Println(<-c)

}
