package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "red-hex-code",
		"green":  "green-hex-code",
		"blue":   "blue-hex-code",
		"white":  "white-hex-code",
		"yellow": "yellow-hex-code",
	}

	//var colorsDecl map[string]string
	//colorsMake := make(map[string]string)
	//colorsMake["white"] = "white-hex-code"
	//delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hexCode := range c {
		fmt.Println("Hex code for", color, "is", hexCode)
	}
}
