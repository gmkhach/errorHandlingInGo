package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}

	// n is the number of byes written
	// The reason it prints out 6 and not 5 is because of the new line character that Println() is printing	
	fmt.Println(n)

	// Using the Scan method to take some input
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

	// Now let's print out what's been collected
	fmt.Println(str1, str2)

	// This implements the Writer interface
	file, err := os.Create("greetings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// When we open a file we want to close it right away
	defer file.Close()

	// This implements the Reader interface
	reader := strings.NewReader("Hello from greetings.txt")

	// The Copy() method of package io takes a writer and a reader and make a file
	io.Copy(file, reader)

	// Now let's take a look at another build-in reader
	file, err = os.Open("greetings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	xb, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(xb))
}