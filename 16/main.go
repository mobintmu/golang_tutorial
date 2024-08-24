package main

import "fmt"

func main() {

	//Check For Specific Elements in a Map

	my := make(map[string]int)

	my["one"] = 1

	val, ok := my["one"]

	if ok {
		fmt.Println("Value: ", val, "Present? ", ok)
	}

	for counter, value := range my {
		fmt.Println("Key: ", counter, "Value: ", value)
	}

}
