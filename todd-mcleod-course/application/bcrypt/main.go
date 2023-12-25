package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error while bcrypting:", err)
	}
	fmt.Println(bs)

	loginPassword := "password123"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("Password incorrect")
		return
	} else {
		fmt.Println("Login successful")
	}

	fmt.Println(bs)
}
