package main

import "fmt"

type customErr struct {
	msg string
}

/*
If any type implements Error() with string return, it implicitely implements error interface
*/
func (c customErr) Error() string {
	return c.msg
}

func foo(e error) {
	fmt.Println(e)
}
func main() {

	c1 := customErr{
		msg: "Custom error message",
	}
	foo(c1)

}
