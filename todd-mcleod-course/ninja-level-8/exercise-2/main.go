package main

import (
	"encoding/json"
	"fmt"
)

type MI6 struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	jsonString := `[{"First": "James", "Last": "Bond", "Age": 30, "Sayings": ["Shaken, not stirred", "I am Bond, James Bond."]}, {"First": "Miss", "Last": "Moneypenny", "Age": 25, "Sayings": ["Hello James", "M will see you now"]}]`
	fmt.Println(jsonString)

	bs := []byte(jsonString)

	var people []MI6
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error while unmarshalling:", err)
	} else {
		fmt.Println(people)
	}

}
