package main

import "fmt"

func main() {
	favourites := make(map[string][]string)
	favourites["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	favourites["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	favourites["no_dr"] = []string{"cats", "icecream", "sunsets"}

	for k, v := range favourites {
		for i, val := range v {
			fmt.Printf("Name: %v, Index: %d, Value: %v \n", k, i, val)
		}
	}
}
