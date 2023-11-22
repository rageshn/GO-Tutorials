package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while reading file: ", err)
	} else {
		stat, err := file.Stat()
		if err != nil {
			fmt.Println("Error while getting file stats: ", err)
		} else {
			bs := make([]byte, stat.Size())
			_, err := bufio.NewReader(file).Read(bs)
			if err != nil {
				fmt.Println("Error while creating byute slice for file: ", err)
			} else {
				checkSum := sha256.Sum256(bs)
				fmt.Printf("checksum: %x \n", checkSum)
			}
		}

	}

}
