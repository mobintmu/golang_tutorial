package main

import (
	"fmt"
	"time"
)

func main() {

	go example("goroutine")

	example("direct")

	time.Sleep(time.Second)
	fmt.Println("done")

}

func example(s string) {
	for counter := 0; counter < 3; counter++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, " : ", counter)
	}
}
