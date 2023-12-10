package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

type secretAgent struct {
	person
	licenceToKill bool
}

func main() {
	p1 := person{
		first_name: "James",
		last_name:  "Bond",
		age:        50,
	}

	p2 := person{
		first_name: "Miss",
		last_name:  "Moneypenny",
		age:        30,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	// Embedded structs
	sa1 := secretAgent{
		person: person{
			first_name: "James",
			last_name:  "Bond",
			age:        50,
		},
		licenceToKill: true,
	}

	sa2 := secretAgent{
		person: person{
			first_name: "Miss",
			last_name:  "Moneypenny",
			age:        30,
		},
		licenceToKill: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	//Anonymous struct - struct is defined without an identifier
	ap1 := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "James",
		last_name:  "Bond",
		age:        50,
	}
	fmt.Println(ap1)

}
