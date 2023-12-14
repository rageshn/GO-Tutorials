package main

import (
	"fmt"
	"strconv"
)

/*
	Any type implementing String() method, internally implements Stringer interface

	type Stringer interface {
		String() string
	}
*/

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("Title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "Book Title",
	}
	var c count = 45
	fmt.Println(b)
	fmt.Println(c)
}
