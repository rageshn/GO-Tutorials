package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      42,
		"moneypenny": 32,
	}

	m1 := m["james"]
	fmt.Println("Age of james = ", m1)

	if q, ok := m["Q"]; ok {
		fmt.Println("Age of Q = ", q)
	} else {
		fmt.Println("Q is not in map. but value of Q =", q)
	}
}
