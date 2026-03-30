package main

import (
	"fmt"
	"slices"
)

func main() {
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", original)

	s1 := original[2:5] // low to high [excluding high]
	fmt.Printf("s1: %v, %d\n", s1, len(s1))

	s2 := original[:5]
	fmt.Printf("s2: %v\n", s2) // beging excluding end

	s3 := original[2:]
	fmt.Printf("s3: %v\n", s3) //print everything to the end

	slices.Contains(original, 8)
	fmt.Printf("slices.Containss(%v, 10)\n", s3)

	original = append(original, 15)
	fmt.Printf("Original: %v\n", original)

}
