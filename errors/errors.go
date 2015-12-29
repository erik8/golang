package main

import (
	"errors"
	"fmt"
	"math"
)

func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return math.Sqrt(f), nil
	}
}

func main() {
	// My square root and error handling
	f, err := MySqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

	f, err = MySqrt(27)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

}
