package main

import "fmt"

// everything is passed by value in Go
func myCode(A []string) {

	A = append(A, "Hello")

}

// everything is passed by value in Go
func AppendArray(B [10]string) {

	B[0] = "hi"
}

func AppendPointer(A *[]string) {
	*A = append(*A, "Hello")
}

func ByReference() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func main() {

	A := []string{"hi", "bye"}

	myCode(A)

	fmt.Println(A)

	B := [10]string{"hi", "bye"}
	AppendArray(B)

	fmt.Println(B)

	AppendPointer(&A)
	fmt.Println(A)

}
