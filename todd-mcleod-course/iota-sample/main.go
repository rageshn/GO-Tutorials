package main

import "fmt"

/*
 iota is an identifier which automatically creates constant values starting from 0
*/

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

/*
 Declaring it once at the beginning, automatically increments the values by 1
*/

const (
	a = iota
	b
	c
	d
	e
	f
	g
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(a, b, c, d, e, f)
	// Bit shifting
	fmt.Printf("%d \t %b \n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b \n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b \n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b \n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b \n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b \n", 1<<f, 1<<f)
	fmt.Printf("%d \t %b \n", 1<<g, 1<<g)
}
