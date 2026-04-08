package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	//explain regex in go and the regexp package
	//In Go, regular expressions are supported through the "regexp" package, which provides functions for compiling and executing regular expressions. The package allows you to define patterns for matching strings, extracting substrings, and performing replacements. Regular expressions in Go use a syntax similar to other programming languages, making it easy to learn and use for string manipulation tasks.

	//example of using regex in Go

	text1 := "The quick brown fox jumps over the lazy dog"
	text2 := "Go is a great programming language"

	text3 := "Products Codes : P123, XYZ789, DEF456"

	// Check if text1 contains the word "fox"
	regGo, err := regexp.Compile("fox")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		os.Exit(1)
	}

	// MatchString returns true if the string matches the compiled pattern
	fmt.Printf("Does text1 contain 'fox'? %t\n", regGo.MatchString(text1)) // Output: true

	// Check if text2 contains the word "Python"
	regPython, err := regexp.Compile("Python")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		os.Exit(1)
	}

	fmt.Printf("Does text2 contain 'Python'? %t\n", regPython.MatchString(text2)) // Output: false


	// Extract product codes from text3
	// MustCompile is like Compile but panics if the expression cannot be parsed.
	// It simplifies initialization of global variables or static patterns.
	regProductCode  := regexp.MustCompile(`P\d+`);

	// FindString returns the first occurrence of the match
	firstProduct := regProductCode.FindString(text3)
	fmt.Printf("First product code found: %s\n", firstProduct)

	// FindAllString returns a slice of all matches. -1 means find all occurrences.
	allProductCodes := regProductCode.FindAllString(text3, -1)
	fmt.Printf("All product codes found: %+v\n", allProductCodes)


}
