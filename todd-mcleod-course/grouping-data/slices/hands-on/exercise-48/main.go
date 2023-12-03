package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "shaken, not stirred"}
	mp := []string{"miss", "moneypenny", "I am 008"}

	people := [][]string{jb, mp}

	for i, v := range people {
		fmt.Println("Index =", i, " Value =", v)
		for index, val := range v {
			fmt.Println("Index = ", index, "Record =", val)
		}
	}
}
