package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      42,
		"moneypenny": 32,
	}

	for key, value := range m {
		fmt.Println("name = ", key, "age = ", value)
	}
}
