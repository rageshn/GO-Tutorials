package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstName string
	age       int
}

type ByAge []person
type ByName []person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByName) Less(i, j int) bool {
	return b[i].firstName < b[j].firstName
}

func main() {
	p1 := person{"James", 40}
	p2 := person{"M", 70}
	p3 := person{"Q", 80}
	p4 := person{"Moneypenny", 35}
	p5 := person{"Dr.No", 75}

	people := []person{p1, p2, p3, p4, p5}
	fmt.Println(people)

	// sort.Sort function expects a type called "Interface" -> sort.Interface
	// type Interface interface {
	//		Len() int
	//		Less(i, j int) bool
	//		Swap(i, j int)
	// }

	// Interface implements Len(), Swap(), Less() functions.
	// As ByAge & ByName also implments those functions, it implicitly implements the "Interface"

	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
