package main

import "fmt"

func main() {
	states := make([]string, 0, 36)
	states = append(states, "Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chhattisgarh",
		"Goa", "Gujarat", "Haryana", "Himachal Pradesh", "Jharkhand", "Karnataka", "Kerala",
		"Madhya Pradesh", "Maharashtra", "Manipur", "Meghalaya", "Mizoram", "Nagaland", "Odisha",
		"Punjab", "Rajasthan", "Sikkim", "Tamil Nadu", "Telangana", "Tripura", "Uttarakhand", "Uttar Pradesh",
		"West Bengal", "Andaman and Nicobar Islands", "Chandigarh", "Dadra and Nagar Haveli and Daman & Diu",
		"The Government of NCT of Delhi", "Jammu & Kashmir", "Ladakh", "Lakshadweep", "Puducherry")

	fmt.Printf("Length = %v, Capacity = %v \n", len(states), cap(states))

	for i := 0; i < 36; i++ {
		fmt.Println(states[i])
	}
}
