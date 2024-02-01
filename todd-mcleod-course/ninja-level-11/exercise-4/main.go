package main

import (
	"errors"
	"fmt"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math error: %v, %v, %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			lat:  "50.2299 N",
			long: "99.4565 E",
			err:  errors.New("Cannot compute square roor for negative number"),
		}
	}
	return math.Sqrt(f), nil
}

func main() {
	v, err := sqrt(-10.23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

}
