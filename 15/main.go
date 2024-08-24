package main

import "fmt"

func main() {

	var my = make(map[string]int)

	my["one"] = 1

	fmt.Println(my)

	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
