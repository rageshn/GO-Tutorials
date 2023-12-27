package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	FirstName string
	LastName  string
	Age       int
	Sayings   []string
}

func main() {

	u1 := user{
		FirstName: "James",
		LastName:  "Bond",
		Age:       30,
		Sayings:   []string{"Shaken, not stirred", "I am Bond, James Bond."},
	}
	u2 := user{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       25,
		Sayings:   []string{"Hello James", "M will see you now"},
	}
	users := []user{u1, u2}

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(users)
	if err != nil {
		fmt.Println("Error while encoding:", err)
	}
}
