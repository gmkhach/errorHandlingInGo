package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Println(err)
	}

	_, err = devide(1, 0)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		// The float64 type in go has a IEEE 754 64-bit floating-point numbers. 
		// math.NaN() returns an IEEE 754 "not-a-number" value. 
		// When we need to return a correct type but cannot return any actual number, because of the type 
		// of the error as the one at hand, we can use math.NaN() to generate the correct return type.
		return math.NaN(), errors.New("Error: can't take a square root of a negative number.")
	}
	return 1984, nil
}

// The same can be accomplished with fmt.Errorf(), which literally returns the same errors.New()
func devide(n1 float64, n2 float64) (float64, error) {
	if n2 == 0 {
		return math.NaN(), fmt.Errorf("Error: can't devide over zero.")
	}
	return n1 / n2, nil
}