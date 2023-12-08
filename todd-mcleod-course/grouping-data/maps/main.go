package main

import "fmt"

func main() {
	// Creating map
	m := map[string]int{
		"a": 33,
		"b": 34,
		"c": 35,
	}
	fmt.Println(m)
	fmt.Println("Age of a =", m["a"])

	ms := make(map[string]int)
	ms["a"] = 33
	ms["b"] = 34
	ms["c"] = 35
	fmt.Println(ms)
	fmt.Println("Length of ma =", len(ms))

	for key, value := range ms {
		fmt.Printf("Key = %v, Value = %v \n", key, value)
	}
	for key := range ms {
		fmt.Printf("Key = %v\n", key)
	}

	// Delete key from map
	mi := map[string]int{
		"a": 33,
		"b": 34,
		"c": 35,
	}
	fmt.Println("Before deletion: ", mi)
	delete(mi, "c")
	fmt.Println("After deletion: ", mi)
	// Accessing key which doesn't exists returns 0
	fmt.Println("Accessing key which doesnt exists : ", mi["c"])

	//Comma OK idiom
	value, ok := m["d"]
	if ok {
		fmt.Println("Value = ", value)
	} else {
		fmt.Println("Key doesn't exists")
	}
	//Statement statement idiom
	if value, ok := m["a"]; ok {
		fmt.Println("Value = ", value)
	} else {
		fmt.Println("Key doesn't exists")
	}
}
