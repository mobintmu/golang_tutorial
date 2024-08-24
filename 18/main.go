package main

import "fmt"

func main() {

	var p *int

	p = new(int)

	*p = 10

	fmt.Println(p)
	fmt.Println(*p)

	a := 32

	b := &a

	*b = 20

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(&a)

}
