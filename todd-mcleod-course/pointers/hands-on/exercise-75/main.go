package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "ABC"
	q := "DEF"
	r := "GHI"
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n

}

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
