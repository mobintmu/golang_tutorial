package main

import (
	"fmt"
	"golang_tutorial/27/safe"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	s := safe.New()

	for counter := 0; counter < 1000; counter++ {

		wg.Add(2)
		go func() {
			defer wg.Done()
			s.Increment("a")
		}()

		go func() {
			defer wg.Done()
			s.IncrementByValue("b", 5)
		}()
	}

	wg.Wait()
	fmt.Println("a", s.Get("a"))
	fmt.Println("b", s.Get("b"))

}
