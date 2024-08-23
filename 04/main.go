package main

import "fmt"

func main() {

	/*
		In the provided Go code, an array named [`arr1`] is initialized with six integer elements: 9, 10, 11, 12, 13, and 14.
		Arrays in Go have a fixed size, which in this case is specified as 6.
		The array is defined using the syntax `[6]int`, indicating it holds six integers.
		Next, a slice named [`mySlice1`]  is created from the array [`arr1`].
		Slices in Go are more flexible than arrays as they can dynamically change in size.
		The slice [`mySlice1`] is created using the expression [`arr1[1:5]`].
		Overall, this code demonstrates how to create a slice from an array, and how to inspect its contents, length, and capacity using formatted print statements in Go.
	*/
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	mySlice1 := arr1[1:5]                 // Slice array
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n", cap(mySlice1))

	mySlice2 := arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("mySlice2 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))

	mySlice2 = append(mySlice2, 20) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))

	mySlice2 = append(mySlice2, 21) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))

	mySlice2 = append(mySlice2, 22) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))

	mySlice2 = append(mySlice2, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))

	mySlice2 = append(mySlice2, 24, 25, 26, 27, 28) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))



	
}
