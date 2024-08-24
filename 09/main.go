package main

import (
	"fmt"
	"golang_tutorial/09/car"
)

func main() {

	myCar := car.NewCar()
	myCar.AddData("HI")
	myCar.AddData("Bye")
	fmt.Println("____________________")
	fmt.Printf("length : %d", myCar.Length())
	fmt.Printf("cap: %d", myCar.Cap())

	data := myCar.Data()

	data = append(data, "Hello")

	fmt.Println("_______________________________")
	fmt.Printf("length : %d", myCar.Length())
	fmt.Printf("cap: %d", myCar.Cap())

	fmt.Println("___________________________")

	dataArray := myCar.DataArrayPointer()

	dataArray[0] = "jump"

	myCar.PrintData()

}
