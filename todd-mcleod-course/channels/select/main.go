package main

import "fmt"

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan int)

	go send(evenChannel, oddChannel, quitChannel)
	receive(evenChannel, oddChannel, quitChannel)
	fmt.Println("Exiting...")
}

// This takes send only channel
func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//quit <- 0
	close(quit)
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("From evenChannel:", v)
		case v := <-odd:
			fmt.Println("From oddChannel:", v)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("From quitChannel got", v, "ok =", ok)
				return
			} else {
				fmt.Println("From quitChannel:", v)
			}
		}
	}
}
