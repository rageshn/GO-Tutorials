package main

import (
	"fmt"
	"log"
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

// Wrapper function - This can be called with any type implemeing stringer interface
func logInfo(s fmt.Stringer) {
	log.Println("LOGGED FROM CUSTOM FUNCTION: ", s.String())
}

func main() {
	b := book{
		title: "Book Title",
	}
	var c count = 45
	fmt.Println(b)
	fmt.Println(c)

	log.Println(b)
	log.Println(c)

	logInfo(b)
	logInfo(c)

}
