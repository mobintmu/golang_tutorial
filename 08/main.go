package main

import (
	"fmt"
	"golang_tutorial/08/car"
	"math"
)

func main() {
	fmt.Println(math.Pi) // Export PI

	myCar := car.Car{}
	myCar.Start() //Export Function
	// myCar.drive() //Unexported Function
}
