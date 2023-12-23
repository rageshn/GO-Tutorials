package main

import "fmt"

/*
Method set is a set of method attached to a type. This is integral to how interfaces are implemented.
	1. The method set of a type T consists of all the methods with receiver type T.
		- These methods can be called using variable of type T
	2. The method set of a type *T consists of all the methods with receiver type *T ot T.
		- These methods can be called using variable of type *T
		- It can call methods of the corresponding non-pointer type as well.
*/

type dog struct {
	firstName string
}

func (d dog) walk() {
	fmt.Println(d.firstName, "is walking")
}

func (d *dog) run() {
	d.firstName = "XYZ"
	fmt.Println(d.firstName, "is running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{
		firstName: "ABC",
	}
	fmt.Println(d1)
	d1.walk()
	d1.run()
	//youngRun(d1) - This throws error

	d2 := &dog{
		firstName: "DEF",
	}
	fmt.Println(d2)
	d2.walk()
	d2.run()
	youngRun(d2)

}
