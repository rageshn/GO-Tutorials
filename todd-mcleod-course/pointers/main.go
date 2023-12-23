package main

import "fmt"

// This methods takes the pointer to the int (i.e., memory address) and updates the value at that location
func intDelta(n *int) {
	fmt.Println("Address of parameter in intDelta:", n)
	*n = 50
}

// Slice is a reference type, so the slice in the function parameter also points to the same address.
func sliceDelta(i []int) {
	fmt.Printf("Address of parameter in sliceDelta: %p \n", i)
	i[0] = 100
}

func mapDelta(m map[int]string, s string) {
	fmt.Printf("Address of parameter in mapDelta: %p \n", m)
	m[0] = s
}

// Value semantics
func addOne(v int) int {
	v += 1
	return v
}

// Pointer semantics
func addOnePointer(i *int) {
	*i += 1
}

func main() {
	x := 5
	fmt.Println(x)   //Prints the value
	fmt.Println(&x)  //Prints the memory address which holds the value. '&' indicates the memory address where the value is stored.
	fmt.Println(*&x) //'*' Retrieves the value from the memory address. combination of '*&' is same as printing the value directly.

	fmt.Printf("%T\n", &x) //Gets the pointer to the type of value stored in the memory address. i.e., pointer to int (*int)

	y := &x // Assigning the memory address to a variable.
	fmt.Println(y)

	*y = 10 //Altering the value at the memory address directly. This updates the x value too as they are both same memory address.
	fmt.Println(x)
	fmt.Println(y)

	intDelta(&x)
	fmt.Println("After updating the int value at the memory address:", x)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	fmt.Printf("Address of original slice: %p \n", xi)
	sliceDelta(xi) //As slice is a reference type, the object can be directly passed to the function.
	fmt.Println("After updating the slice value at the memory address:", xi)

	mi := make(map[int]string)
	mi[0] = "My"
	mi[1] = "Name is"
	mi[2] = "Bond"
	mi[3] = "James Bond"
	fmt.Println(mi)
	fmt.Printf("Address of original map: %p \n", mi)
	mapDelta(mi, "Hello")
	fmt.Println("After updating the map value at the memory address:", mi)

	a := 5
	fmt.Println(a)
	addOne(a)
	fmt.Println(a)
	addOnePointer(&a)
	fmt.Println(a)

}
