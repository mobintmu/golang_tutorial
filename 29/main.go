package main

import (
	"fmt"
	"golang_tutorial/29/worker"
)

func main() {

	const numberJobs = 10

	jobs := make(chan int, numberJobs)
	results := make(chan int, numberJobs)

	defer close(jobs)
	defer close(results)

	for w := 1; w <= 3; w++ {
		work := worker.New(w, w+1, jobs, results)
		go work.Job()
	}

	for j := 1; j <= numberJobs; j++ {
		jobs <- j
	}

	for a := 1; a <= numberJobs; a++ {
		msg := <-results
		fmt.Println("main msg :", msg)
	}

}
