package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Initialization function - runs after variable declaration and before main")
}

func main() {
	//SEQUENCE
	fmt.Println("This is the first sequence statement to execute")
	fmt.Println("This is the second sequence statement to execute")
	x := 40                               //This is the third sequence statement to execute
	y := 5                                //This is the fourth sequence statement to execute
	fmt.Printf("x = %v, y = %v \n", x, y) //This is the fifth sequence statement to execute

	// CONDITIONAL
	if x < 45 {
		fmt.Println("First conditional statement in first block")
	}

	if x < 45 {
		fmt.Println("First conditional statement in second block")
	} else {
		fmt.Println("Second conditional statement in second block")
	}

	if x < 45 {
		fmt.Println("First conditional statement in third block")
	} else if x == 45 {
		fmt.Println("Second conditional statement in third block")
	} else {
		fmt.Println("Third conditional statement in third block")
	}

	//LOGICAL
	if x < 45 && y < 45 {
		fmt.Println("First logical statement")
	}

	if x > 30 || x < 45 {
		fmt.Println("Second logical statement")
	}

	if x != 45 && y != 10 {
		fmt.Println("Third logical statement")
	}

	//STATEMENT STATEMENT IDIOM
	/*
		if x := f(); CONDITION {
			IF TRUE
		} else {
			IF FALSE
		}
	*/
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater than or equal to x which is %v \n", z, x)
	} else {
		fmt.Printf("z is %v and that is less than x which is %v \n", z, x)
	}

	//SWITCH
	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("Default case - First switch case")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 45:
		fmt.Println("x is 45")
	case 50:
		fmt.Println("x is 50")
	default:
		fmt.Println("Default case - Second switch case")
	}

	//no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough //Specifies that is must also execute next case
	case 45:
		fmt.Println("x is 45 (Printed because of fallthrough)")
	case 50:
		fmt.Println("x is 50")
	default:
		fmt.Println("Default case - Third switch case")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough //Specifies that is must also execute next case
	case 45:
		fmt.Println("x is 45 (Printed because of previous fallthrough)")
		fallthrough
	case 50:
		fmt.Println("x is 50 (Printed because of previous fallthrough)")
		fallthrough
	default:
		fmt.Println("Default case (Printed because of previous fallthrough)")
	}

	//FOR
	for i := 0; i < 5; i++ {
		fmt.Printf("First for loop - i = %v \n", i)
	}

	//Loops till y is less than 10
	for y < 10 {
		fmt.Printf("Second for loop - y = %v \n", y)
		y++
	}

	// Infinite for loop with break
	for {
		fmt.Printf("Third for loop - y = %v \n", y)
		if y > 20 {
			fmt.Println("Breaks when y > 20")
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("Printing only even numbers - %v \n", i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("------")
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %v, j = %v \n", i, j)
		}
	}

}
