package main

import "fmt"

func main() {

	//GO represents strings as byte slices using under the hood. This means you can access different indexes of a string like you would for a slice variable.
	//A byte slice is a slice whose underlying type is a slice. Byte slices are more like lists of bytes that represent UTF-8 encodings of Unicode code points.

	myString := "Hello String"
	// printBytes(myString)
	// runeFunc(myString)
	immutable(myString)
}

func printBytes(s string) {

	fmt.Printf("Bytes: ")

	for i := 0; i < len(s); i++ {

		fmt.Print(s[i], " ")
		fmt.Printf("%x ", s[i])
		fmt.Println("___")

	}

}

func runeFunc(s string) {

	//1. Rune Data Type:
	//In Go, a rune is represented using the `rune` data type, which is an alias for the `int32` type.
	// This data type allows you to store a Unicode code point, which is a 32-bit integer, making it suitable for handling characters from all languages.

	for i, r := range s {
		fmt.Printf("%c starts at byte %d\n", r, i)
	}
}

func immutable(s string) {

	//Go string immutability principle

	original := "hello"
	modified := original

	modified = "world"

	fmt.Println(original) // Outputs: hello
	fmt.Println(modified) // Outputs: world
}
