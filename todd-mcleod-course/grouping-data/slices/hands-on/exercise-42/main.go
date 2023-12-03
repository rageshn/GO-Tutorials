package main

import "fmt"

func main() {
	ai := [5]int{}
	for i := 0; i < 5; i++ {
		ai[i] = i
	}
	fmt.Println(ai)
	for i, v := range ai {
		fmt.Printf("Index = %v , Value = %v , Type = %T \n", i, v, v)
	}
}
