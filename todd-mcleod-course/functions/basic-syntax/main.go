package main

import "fmt"

type person struct {
	firstName string
}

func main() {
	foo()
	bar("Hai")
	v := test("Aloha")
	fmt.Println(v)
	name, age := getDetails("A", 5)
	fmt.Println("Name =", name, ", New age =", age)

	totalSum := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("Total sum: ", totalSum)

	//Unfurling a slice
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//Unfluring the slice - "..." near the slice opens it to individual elements.
	total := sum(values...)
	fmt.Println("Total Sum: ", total)

	//Defer - Defers the execution till the parent/surrounding function exits.
	//If a parent function has more than one defer functions, then they all will be added to a stack and called in LIFO fashion.
	//i.e., executed from the end to start
	defer fooDeferred()
	bar("After defer statement")

	//Methods - function with receivers
	p1 := person{
		firstName: "A",
	}
	p2 := person{
		firstName: "B",
	}
	p1.speak()
	p2.speak()
}

/*
 func (r receiver) functionName(p parameters) (r returnType) { ... }
*/

// No Parameters, No returns
func foo() {
	fmt.Println("Hello from foo")
}

// 1 Parameter, No returns
func bar(s string) {
	fmt.Println("Hello from bar. You have passed", s, "as parameter")
}

// 1 Parameter, 1 return
func test(s string) string {
	return fmt.Sprint("Hello from test. You have passed ", s, " as parameter")
}

// 2 Parameters, 2 returns
func getDetails(name string, age int) (string, int) {
	newAge := age * 7
	return fmt.Sprint("Hello from getDetails. Name = ", name), newAge

}

// Variadic parameters - Takes unlimited number of parameters. Represented by "..." . This must be the final parameter in function signature.
func sum(i ...int) int {
	fmt.Println("Parameter: ", i)
	fmt.Printf("Parameter type: %T\n", i)
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func fooDeferred() {
	fmt.Println("Hi from deferred function. This will be executed last.")
}

func (p person) speak() {
	fmt.Println("Hi! I am ", p.firstName)
}
