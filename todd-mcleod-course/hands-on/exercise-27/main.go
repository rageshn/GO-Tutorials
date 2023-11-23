package main

import "fmt"

func main() {
	i := 10

	for {
		if i < 0 {
			fmt.Println("Breaking loop")
			break
		}
		fmt.Println("i =", i)
		i -= 1
	}
}
