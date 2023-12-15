package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type person struct {
	firstName string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.firstName))
	return err
}

func main() {
	p := person{
		firstName: "ABC",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())

}
