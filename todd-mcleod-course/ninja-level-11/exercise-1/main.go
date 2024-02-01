package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "ABC",
		Last:    "DEF",
		Sayings: []string{"Hello", "Hai"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Json marshalling error: ", err)
	}
	fmt.Println(string(bs))

}
