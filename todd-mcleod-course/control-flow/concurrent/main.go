package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Create two cahnnels
	channel1 := make(chan int)
	channel2 := make(chan int)

	// Create durations for go routines and sleep for that many milliseconds
	duration1 := time.Duration(rand.Int63n(250))
	duration2 := time.Duration(rand.Int63n(250))

	//Go routines
	go func() {
		time.Sleep(duration1 * time.Millisecond)
		fmt.Printf("Sleep for %v milliseconds \n", duration1)
		channel1 <- 41
	}()

	go func() {
		time.Sleep(duration2 * time.Millisecond)
		fmt.Printf("Sleep for %v milliseconds \n", duration2)
		channel2 <- 42
	}()

	//Select chooses which set of possible send or receive operations will proceed
	select {
	case v1 := <-channel1:
		fmt.Printf("Value from channel1: %v \n", v1)
	case v2 := <-channel2:
		fmt.Printf("Value from channel2: %v \n", v2)
	}
}
