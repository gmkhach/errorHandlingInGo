package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}

	// n is the number of byes written
	// The reason it prints out 6 and not 5 is because of the new line character that Println() is printing	
	fmt.Println(n)

	// Now let's use the Scan method to take some input
	var str1, str2 string

	fmt.Print("String 1: ")
	_, err = fmt.Scan(&str1)
	if err != nil {
		panic(err)
	}

	fmt.Print("String 2: ")
	_, err = fmt.Scan(&str2)
	if err != nil {
		panic(err)
	}

	fmt.Println(str1, str2)
}