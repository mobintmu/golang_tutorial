package main

import (
	"fmt"
	"time"
)

func main() {

	myChannel := make(chan string)

	go example("goroutine", myChannel)

	msg := <-myChannel
	fmt.Println(msg)
}

func example(s string, myChannel chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("function")
	myChannel <- s
}
