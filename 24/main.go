package main

import (
	"fmt"
	"sync"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func getRequestData(city string, message chan<- string, wg *sync.WaitGroup) {

	defer timer("getRequestData " + city)()
	defer wg.Done()

	time.Sleep(200 * time.Microsecond)
	data := make(map[string]int)
	data[city] = 100

	message <- fmt.Sprintf("city: %s, data: %v", city, data)

}

func main() {

	defer timer("main")()

	message := make(chan string)

	var wg sync.WaitGroup

	cites := []string{"tehran", "isfahan", "rasht"}

	for _, city := range cites {
		wg.Add(1)
		go getRequestData(city, message, &wg)
		// fmt.Println(result)
	}

	go func() {
		wg.Wait()
		close(message)
	}()

	for msg := range message {
		fmt.Println(msg)
	}

}
