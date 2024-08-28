package main

import (
	"fmt"
	"time"
)

func main() {

	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			channelOne <- "hi one"
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			channelTwo <- "bye two"
		}

	}()

	for {
		select {
		case msg := <-channelOne:
			fmt.Println("recive :", msg)
		case msg := <-channelTwo:
			fmt.Println("recive :", msg)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}
	}

}
