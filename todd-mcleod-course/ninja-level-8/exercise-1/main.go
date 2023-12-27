package main

import (
	"encoding/json"
	"fmt"
)

// Attributes should start with upper case for them to be exported to other packages
type user struct {
	FirstName string
	Age       int
}

func main() {
	u1 := user{
		FirstName: "ABC",
		Age:       50,
	}
	u2 := user{
		FirstName: "XYZ",
		Age:       100,
	}

	users := []user{u1, u2}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error while marshalling:", err)
	} else {
		fmt.Println(bs)
		fmt.Println(string(bs))
	}
}
