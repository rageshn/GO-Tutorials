package main

import "fmt"

func main() {
	p := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "A",
		friends: map[string]int{
			"B": 5,
			"C": 10,
		},
		favDrinks: []string{
			"Vodka", "Wine",
		},
	}

	fmt.Println(p)
	fmt.Printf("Name: %v, Friends: %v, Favourite Drinks: %v \n", p.first, p.friends, p.favDrinks)
}
