package main

import "fmt"

type person struct {
	firstName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("My name is ", p.firstName)
}

func (s secretAgent) speak() {
	fmt.Println("I am secret agent ", s.person.firstName)
	fmt.Println("LTK: ", s.licenseToKill)
}

func main() {
	p1 := person{
		firstName: "A",
	}
	p2 := person{
		firstName: "B",
	}
	p1.speak()
	p2.speak()

	s1 := secretAgent{
		person: person{
			firstName: "SA1",
		},
		licenseToKill: true,
	}
	s1.speak()
	fmt.Println("------------------------------")

	// Using interface methods
	// person and secretAgent types are implemeting speak(), and interface human also has speak()
	// person and secretAgent are also of type human.
	saySomething(p2)
	saySomething(s1)

}
