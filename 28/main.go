package main

import "fmt"

func main() {

	messageFirst := make(chan string, 2)
	messageSecond := make(chan string, 2)

	ping(messageFirst, "ping")
	pong(messageFirst, messageSecond)

	msg := <-messageSecond
	fmt.Println(msg)

}

func ping(channel chan<- string, msg string) {
	fmt.Println("func ping")
	channel <- msg
}

func pong(channel <-chan string, channel2 chan<- string) {
	fmt.Println("func pong")
	msg := <-channel
	channel2 <- msg
}
