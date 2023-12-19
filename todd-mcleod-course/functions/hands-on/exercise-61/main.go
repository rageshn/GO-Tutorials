package main

import "fmt"

type person struct {
	firstName string
	age       int
}

func (p person) speak() {
	fmt.Println("Hi, My name is", p.firstName, "and I am", p.age, "years old")

}

func main() {
	p := person{
		firstName: "A",
		age:       50,
	}
	p.speak()

}
