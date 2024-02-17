package main

import (
	dog "command-line-arguments/home/ragesh/Documents/code/GO-Tutorials/todd-mcleod-course/ninja-level-12/dog/main.go"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	scooby := canine{
		name: "Scooby doo",
		age:  dog.Years(6),
	}
	fmt.Println(scooby)
}
