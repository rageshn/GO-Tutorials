package main

import "fmt"

func main() {
	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine engine
		make   int
		model  string
		doors  int
		color  string
	}

	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  2023,
		model: "New",
		doors: 4,
		color: "Black",
	}
	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  2022,
		model: "Old",
		doors: 4,
		color: "Grey",
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.engine)
	fmt.Println(v2.color)
}
