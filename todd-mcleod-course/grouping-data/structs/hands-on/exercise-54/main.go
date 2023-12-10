package main

import "fmt"

func main() {
	type person struct {
		firstName    string
		lastName     string
		favIceCreams []string
	}

	p1 := person{
		firstName:    "James",
		lastName:     "Bond",
		favIceCreams: []string{"Vannila", "Chocolate"},
	}
	p2 := person{
		firstName:    "Dr",
		lastName:     "No",
		favIceCreams: []string{"Black Current", "Pista"},
	}

	m := make(map[string]person)
	m[p1.lastName] = p1
	m[p2.lastName] = p2

	m1 := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)
	fmt.Println(m1)

	for k, v := range m {
		fmt.Printf("Last name: %v, Values: %v \n", k, v)
		for _, val := range v.favIceCreams {
			fmt.Println(val)
		}
	}
}
