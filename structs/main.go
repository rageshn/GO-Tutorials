package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	ragesh := person{
		firstName: "Ragesh",
		lastName:  "N",
		contact: contactInfo{
			email:   "r.n@gmail.com",
			pinCode: 620012,
		},
	}

	// & gets the memory address (pointer) to the object
	rageshPointer := &ragesh
	rageshPointer.updateFirstName("Raki")
	ragesh.print()

	// If the function's receiver is a pointer to a type,
	// the function can be called on the pointer, as well as the object of that type
	ragesh.updateFirstName("Raki_New")
	ragesh.print()
}

// call the function on the pointer to person type and update the object in the pointer
// If * is placed before a type, then it represents the pointer to that type
// If * is placed before a actual pointer, then it refers to the value in that address
func (pointerAddress *person) updateFirstName(newFirstName string) {
	// * gives the value in the address
	(*pointerAddress).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
