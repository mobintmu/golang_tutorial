package car

import "fmt"

type Car struct {
}

func (c Car) Start() {
	fmt.Println("Start")
}

func (c Car) drive() {
	fmt.Println("drive")
}
