package worker

import (
	"fmt"
	"time"
)

type Worker struct {
	id        int
	jobs      <-chan int
	result    chan<- int
	timeSpend int
}

func New(id int, timeSpend int, jobs <-chan int, result chan<- int) *Worker {
	return &Worker{
		id:        id,
		timeSpend: timeSpend,
		jobs:      jobs,
		result:    result,
	}
}

// chan   // read-write
// <-chan // read only
// chan<- // write only

func (w *Worker) Job() {

	for counter := range w.jobs {
		fmt.Println("worker ", w.id, "started , work :", counter)
		time.Sleep(time.Duration(w.timeSpend) * time.Second)
		fmt.Println("worker ", w.id, "finished ,work :", counter)
		w.result <- counter * 2
	}
}
