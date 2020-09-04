package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type product struct {
	category	string
	sn			string
	err			error
}

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Println(err)
	}

	_, err = devide(1, 0)
	if err != nil {
		log.Println(err)
	}

	_, err = getProduct(true) 
	if err != nil {
		log.Println(err)
	}
}

// By adding this Error() method to the type product the type product implicitly becomes of type error also.
// Therefore, it can be passed and run as an error.
func (p product) Error() string {
	return fmt.Sprintf("There is an error with this product: %v %v %v", p.category, p.sn, p.err)
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

func getProduct(isDefective bool) (product, error) {
	if isDefective {
		return product{}, product{"electronic good", "235ab7xd", errors.New("This product is defective")}
	}
	return product{"electronic good", "235ab7xd", nil}, nil
}