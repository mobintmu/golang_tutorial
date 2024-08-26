package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()
		fmt.Println("function one")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("finish function one")

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("function two")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("finish function two")
	}()

	wg.Wait()
	fmt.Println("done main")
}
