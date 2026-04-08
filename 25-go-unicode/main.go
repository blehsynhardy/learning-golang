package main

import (
	"fmt";
	"unicode"
)



func main() {

	//explain unicode in go and runes
	//In Go, Unicode is supported through the use of runes, which are aliases for int32. A rune represents a single Unicode code point, allowing Go to handle a wide range of characters from various languages and symbol sets. The standard library provides functions for working with runes, such as converting between runes and strings, and iterating over Unicode characters in a string.

	//example of using runes in Go
	data := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	str := "Hello, 世界"
	for _, r := range str {
		fmt.Printf("Rune: %c, Unicode: U+%04X\n", r, r)
	}


	for _, r := range data {
		fmt.Printf("Rune: %c, unicode.isLower: %t\n", r, unicode.IsLower(r))
	}

}