package main

import (
	"fmt"
	"io"
	"os"
)

/*
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
*/

type person struct {
	firstName string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.firstName))
	return err
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
	}
	defer f.Close()

	s := []byte("Hello, How are you?")
	bytesWritten, writeError := f.Write(s)
	if writeError != nil {
		fmt.Println("Error while writting to file: ", writeError)
	}
	fmt.Println("Bytes written: ", bytesWritten)
}
