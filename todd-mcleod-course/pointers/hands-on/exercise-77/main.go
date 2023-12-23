package main

import "fmt"

type person struct {
	firstName string
}

func changeName(p person) string {
	p.firstName = "Hello"
	return "Hello"
}

func change(p *person) {
	p.firstName = "Hai"
}

func main() {
	p1 := person{
		firstName: "ABC",
	}
	p2 := person{
		firstName: "XYZ",
	}

	changeName(p1)
	fmt.Println(p1)

	change(&p2)
	fmt.Println(p2)

}
