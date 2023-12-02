package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)",
		"Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)",
		"Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)",
		"Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}

	for index := 0; index < len(xs); index++ {
		fmt.Println("Value at index", index, "is", xs[index])
	}
	fmt.Println("--------------------------------------")

	//Appending data to slice
	xs = append(xs, "This flavour is added to slice")

	for index := 0; index < len(xs); index++ {
		fmt.Println("Value at index", index, "is", xs[index])
	}
	fmt.Println("-------------------------------------")

	//Slicing a slice
	//[inclusive:exclusive]
	fmt.Println(xs[0:5])
	//[inclusive:]
	fmt.Println(xs[7:])
	//[:exclusive]
	fmt.Println(xs[:17])
	//[:] - all elements
	fmt.Println(xs[:])
	fmt.Println("------------------------------------")

	//Deleting from a slice e.g., remove element at index 4
	//Append the first 3 elements with elements from 5th index till end
	//Second parameters should be an element of type T, but as we are appending slice it will throw error.
	// ... opens the slice and places every element in second parameter
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)
	xi = append(xi[:4], xi[5:]...)
	fmt.Println(xi)
	fmt.Println("---------------------------------------")

	//Use make to create slice
	//make takes type, initial size, capacity of the underlying array
	xj := make([]int, 0, 10)
	fmt.Println("Original slice =", xj, "Length of slice =", len(xj), "Capacity of underlying array =", cap(xj))
	fmt.Println("----------------------------------------")
	xj = append(xj, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("slice after appending elements =", xj)
	xj = append(xj, 10, 11, 12, 13, 14, 15)
	fmt.Println("slice after appending additional elements =", xj)
	fmt.Println("Length of new slice =", len(xj), "Capacity of underlying array =", cap(xj))
	fmt.Println("-----------------------------------------")

	//Multi dimensional slice
	jb := []string{"james", "bond", "vodka martini", "chocolate"}
	mp := []string{"miss", "moneypenny", "champagne", "vannila"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
