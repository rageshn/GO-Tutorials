package main

import "fmt"

type person struct {
	firstName string
	age       int
}

func (p *person) speak() {
	fmt.Println("Hello! My name is", p.firstName, "and I am", p.age, "years old.")
}

type human interface {
	speak()
}

func speakSomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		firstName: "ABC",
		age:       50,
	}

	speakSomething(&p1)
	//speakSomething(p1) // This does not work as person does not implement human

}
