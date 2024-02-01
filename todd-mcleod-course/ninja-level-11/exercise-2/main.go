package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings string
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("Error while marshalling: %v", err)
		//return []byte{}, errors.New("Error while marshaling")
		return []byte{}, errors.New(fmt.Sprintf("Error while marshalling: %v", err))
	}
	return bs, err
}

func main() {
	p1 := person{
		First:   "ABC",
		Last:    "DEF",
		Sayings: "Hello! How are you?",
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
