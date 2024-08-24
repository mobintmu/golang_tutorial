package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 10
}

func main() {

	mySlice := []int{20, 30, 40}
	modifySlice(mySlice)
	fmt.Println(mySlice) // Output: [10 30 40]

}
