package main

import "fmt"

func change(my map[string]int) {
	my["three"] = 3
}

func main() {

	my := map[string]int{}

	my["one"] = 1
	my["two"] = 2

	change(my)

	fmt.Println(my)
}
