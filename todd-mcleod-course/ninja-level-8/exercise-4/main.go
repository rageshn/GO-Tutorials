package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{67, 9, 45, 76, 89, 43, 37, 64, 23, 223, 1, 0}
	xs := []string{"random", "rainbow", "absolute", "zebra", "jeopardize"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
