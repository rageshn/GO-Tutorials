package main

import "fmt"

/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier
use print verbs to show value and type
use print verbs to show numbers a decimal, binary, hexadecimal
use signed and unsigned int
*/

var a int

func main() {
	fmt.Println(a)

	b := 47
	fmt.Println(b)

	c, d := 97, "Hello"
	fmt.Println(c, d)

	var e float32 = 47.66
	fmt.Println(e)

	f, _ := 78, 56
	fmt.Println(f)

	g, h, i := "Hello world", 50, 67.98
	fmt.Printf("%v is of type %T \n", g, g)
	fmt.Printf("%v is of type %T \n", h, h)
	fmt.Printf("%v is of type %T \n", i, i)

	j, k, l := 747, 911, 90210
	fmt.Printf("Representation of %v is \t %d \t %b \t\t %#X \n", j, j, j, j)
	fmt.Printf("Representation of %v is \t %d \t %b \t\t %#X \n", k, k, k, k)
	fmt.Printf("Representation of %v is \t %d \t %b \t %#X \n", l, l, l, l)

	var m uint8 = 255
	var n int8 = 127
	fmt.Printf("%v is of type %T \n", m, m)
	fmt.Printf("%v is of type %T \n", n, n)

}
