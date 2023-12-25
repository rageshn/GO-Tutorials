package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       50,
	}

	p2 := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       30,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error while marshalling people:", err)
	}
	fmt.Println(b)
	fmt.Println(string(b)) // This gives empty string as the field names are in lower case. Variables in lower case cannot be exported to other packages.

	sa1 := secretAgent{
		FirstName: "James",
		LastName:  "Bond",
		Age:       30,
	}
	sa2 := secretAgent{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       25,
	}

	sa := []secretAgent{sa1, sa2}
	bs, err := json.Marshal(sa)
	if err != nil {
		fmt.Println("Error while marshalling secret agent:", err)
	}
	fmt.Println(string(bs)) // This prints all the attributes as they are in Upper case.
}
