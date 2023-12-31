package main

import "fmt"

func main() {
	c := gen()

	receive(c)
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

}
