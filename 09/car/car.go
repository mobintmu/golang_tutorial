package car

import "fmt"

type Car struct {
	data      []string
	dataArray [10]string
}

func NewCar() *Car {
	return &Car{
		data:      make([]string, 10),
		dataArray: [10]string{},
	}
}

func (c *Car) AddData(d string) {
	c.data = append(c.data, d)
	c.dataArray[0] = d
}

func (c *Car) Data() []string {
	return c.data
}

func (c *Car) DataArray() [10]string {
	return c.dataArray
}

func (c *Car) IsExists(d string) bool {

	for _, v := range c.data {
		if v == d {
			return true
		}
	}

	return false
}

func (c *Car) Remove(s string) {
	for i, v := range c.data {
		if v == s {
			//remove data[i]
			c.data = append(c.data[:i], c.data[i+1:]...)
		}
	}
}

func (c *Car) Length() int {
	return len(c.data)
}

func (c *Car) Cap() int {
	return cap(c.data)
}

func (c *Car) DataArrayPointer() *[10]string {
	return &c.dataArray
}

func (c *Car) PrintData() {
	fmt.Println(c.dataArray)
}
