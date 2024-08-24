package main

import (
	"fmt"
	"strings"
)

func main() {

	// slice()
	search()

}

func slice() {
	str := "Hello, World!"
	// Slicing from index 7 up to, but not including, index 12
	substring := str[7:12]
	fmt.Println(substring) // Outputs: "World"
	fmt.Println(len(substring))
}

func search() {
	// The string in which we want to search
	haystack := "Hello, Golang World!"

	// The substring we want to search for
	needle := "Golang"

	// Use the strings.Contains function to check if the haystack contains the needle
	if strings.Contains(haystack, needle) {
		fmt.Println("Found:", needle)
	} else {
		fmt.Println(needle, "not found!")
	}
}
