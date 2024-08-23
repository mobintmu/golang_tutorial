package main

import "fmt"

func main() {
	var b [5]int
	fmt.Println("emp:", b)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr2)

	fmt.Println(arr2[0])
	fmt.Println(arr2[len(arr2)-1])

	fmt.Println("______________________________________________")

	//In go, Arrays and Slices are slightly different and cannot be used interchangeably;
	//however, you can make a slice from an array easily using the [:] operator.

	a := [4]int{1, 2, 3, 4} // "a" has type [4]int (array of 4 ints)
	x := a[:]               // "x" has type []int (slice of ints) and length 4
	y := a[:2]              // "y" has type []int, length 2, values {1, 2}
	z := a[2:]              // "z" has type []int, length 2, values {3, 4}
	m := a[1:3]             // "m" has type []int, length 2, values {2, 3}

	fmt.Println(a, x, y, z, m)

	fmt.Println("________________________SLICES_________________")

	//In Go, arrays and slices are data structures that consist of an ordered sequence of elements.

	//Although arrays and slices in Go are both ordered sequences of elements, there are significant differences between the two.
	//An array in Go is a data structure that consists of an ordered sequence of elements that has its capacity defined at creation time.
	// Once an array has allocated its size, the size can no longer be changed.
	// A slice, on the other hand, is a variable length version of an array, providing more flexibility for developers using these data structures.
	//Slices constitute what you would think of as arrays in other languages.

	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3)

	//ERROR
	//arr3 = append(arr3, 4)

	mySlice := []int{}

	fmt.Println(mySlice)

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice = append(mySlice, 1)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice = make([]int, 3)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice = make([]int, 3, 6)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fmt.Println("________________________END")

	//A slice is a data type in Go that is a mutable, or changeable, ordered sequence of elements.
	//Since the size of a slice is variable, there is a lot more flexibility when using them.

	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp", "anemone"}

	fmt.Println(seaCreatures)

	fmt.Println(len(seaCreatures))

	seaCreatures = append(seaCreatures, "hi")

	fmt.Println(seaCreatures)

	// create with size

	slice1 := make([]int, 5)
	fmt.Println(slice1)

	//nil error
	//fmt.Println(slice1[5])

	//If you want to pre-allocate the memory for a certain capacity, you can pass in a third argument to make():

	oceans := make([]int, 3, 10)
	fmt.Println(oceans)

	oceans = append(oceans, 1)

	fmt.Println(oceans)

	// Converting from an Array to a Slice

	var A = [5]int32{1, 2, 3, 4, 5}

	B := A[:]

	fmt.Println(B)

	B = append(B, 6)

	fmt.Println(B)

	M := A[1:3]
	fmt.Println(M)

	//To combine two slices together, you can use append(),
	//but you must expand the second argument to append using the ... expansion syntax:

	fmt.Println("_________________APPEND____________________")

	C := []int{1, 2, 3}
	D := []int{4, 5, 6}

	C = append(C, D...)

	fmt.Println(C)

	E := []int{1, 2, 3}
	F := []int{4, 5, 6}

	G := []int{}

	G = append(E, F[0])

	fmt.Println(G)

	mySlice1 := []int{1, 2, 3}
	mySlice2 := []int{4, 5, 6}
	mySlice3 := append(mySlice1, mySlice2...)

	fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	fmt.Printf("mySlice3=%v\n", mySlice3)
	fmt.Printf("length=%d\n", len(mySlice3))
	fmt.Printf("capacity=%d\n", cap(mySlice3))

}
