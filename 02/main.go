package main

import "fmt"

var b int

func main() {

	//Comment

	/*
		multi lines comments
	*/

	//Go is statically typed,
	var a int
	a = 1
	fmt.Println(a)

	b = 10
	fmt.Println(b)

	// Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).
	// Note: It is not possible to declare a variable using :=, without assigning a value to it.

	c := 10
	fmt.Println(c)

	//error
	//c = 10.1

	//string

	var str string = "mobin" //type is string
	fmt.Println(str)

	const str2 string = "John"

	//ERROR
	//	str2 = "Doe"

	fmt.Println("____________________")
	var data string
	fmt.Println(data)

	var data2 int
	fmt.Println(data2)

	var data3 bool
	fmt.Println(data3)

	const data4 int = 33
	//Error
	//data3 = 45

	{
		data5 := "hi"
		fmt.Println(data5)
	}

	//Error undefined
	// data5 = "bye"

	//Multi variable

	var i, j string = "Hello", "World"

	fmt.Println(i, j)

	// Data types

	/*
		Type 	Size 	Range
		uint 	Depends on platform:
				32 bits in 32 bit systems and
				64 bit in 64 bit systems 	0 to 4294967295 in 32 bit systems and
				0 to 18446744073709551615 in 64 bit systems
		uint8 	8 bits/1 byte 	0 to 255
		uint16 	16 bits/2 byte 	0 to 65535
		uint32 	32 bits/4 byte 	0 to 4294967295
		uint64 	64 bits/8 byte 	0 to 18446744073709551615
	*/

	var value1 uint = 10
	fmt.Println(value1)

}
