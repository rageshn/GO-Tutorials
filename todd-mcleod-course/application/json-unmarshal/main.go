package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {

	jsonString := `[{"FirstName":"James","LastName":"Bond","Age":30},{"FirstName":"Miss","LastName":"Moneypenny","Age":25}]`
	bs := []byte(jsonString)
	fmt.Println(bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error while unmarshalling:", err)
	} else {
		fmt.Println(people)
	}
}
