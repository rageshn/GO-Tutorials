package main

import "fmt"

type dog struct {
	firstName string
}

func (d dog) walk() {
	fmt.Println(d.firstName, "is walking")
}

func (d dog) run() {
	d.firstName = "XYZ"
	fmt.Println(d.firstName, "is running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{
		firstName: "ABC",
	}
	fmt.Println(d1)
	d1.walk()
	d1.run()
	youngRun(d1)

	d2 := &dog{
		firstName: "DEF",
	}
	fmt.Println(d2)
	d1.walk()
	d2.run()
	youngRun(d2)

}
