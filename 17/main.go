package main

import "fmt"

func main() {

	fmt.Println("counting")

	defer fmt.Println(9)
	defer fmt.Println(8)
	defer fmt.Println(7)
	defer fmt.Println(6)
	defer fmt.Println(5)

	fmt.Println("done")
}
