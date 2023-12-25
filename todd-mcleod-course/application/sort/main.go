package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{45, 98, 9, 67, 23, 37, 59, 39, 25, 45, 10, 73}
	xs := []string{"James", "Q", "M", "No", "Moneypenny"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi) // Sorts the underlying array, instead of returning the sorted slice/array
	fmt.Println("After sorting:", xi)

	sort.Strings(xs) // Sorts the underlying array, instead of returning the sorted slice/array
	fmt.Println("After sorting:", xs)
}
