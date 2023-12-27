package main

import (
	"fmt"
	"sort"
)

type user struct {
	FirstName string
	LastName  string
	Age       int
	Sayings   []string
}

type ByAge []user
type ByLast []user

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByLast) Len() int {
	return len(b)
}
func (b ByLast) Less(i, j int) bool {
	return b[i].LastName < b[j].LastName
}
func (b ByLast) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	u1 := user{
		FirstName: "James",
		LastName:  "Bond",
		Age:       30,
		Sayings:   []string{"Shaken, not stirred", "I am Bond, James Bond."},
	}
	u2 := user{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       25,
		Sayings:   []string{"Hello James", "M will see you now"},
	}
	users := []user{u1, u2}
	sort.Sort(ByAge(users))
	fmt.Println(users)

	sort.Sort(ByLast(users))
	fmt.Println(users)

	for _, u := range users {
		fmt.Println(u.FirstName, u.LastName)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println(v)
		}
	}
}
