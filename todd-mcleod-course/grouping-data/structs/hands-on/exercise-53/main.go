package main

import (
	"fmt"
)

func main() {
	type person struct {
		firstName    string
		lastName     string
		favIceCreams []string
	}

	p1 := person{
		firstName:    "James",
		lastName:     "Bond",
		favIceCreams: []string{"Vannila", "Chocolate"},
	}
	p2 := person{
		firstName:    "Dr",
		lastName:     "No",
		favIceCreams: []string{"Black Current", "Pista"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.favIceCreams {
		fmt.Println(v)
	}
	for _, v := range p2.favIceCreams {
		fmt.Println(v)
	}
}
