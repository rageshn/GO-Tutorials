package main

import "fmt"

func main() {
	inputChannel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			inputChannel <- i
		}
		close(inputChannel)
	}()

	for v := range inputChannel {
		fmt.Println(v)
	}
}
