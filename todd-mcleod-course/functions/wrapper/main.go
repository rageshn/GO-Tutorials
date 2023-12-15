package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error while reading file: %s", err)
	}
	return xb, nil
}

func main() {
	xb, err := readFile("poem")
	if err != nil {
		fmt.Println("Error in readFile in main:", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}
