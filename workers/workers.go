package workers

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func RunWorker() {
	ch := make(chan int)
	qtsWorkers := 100

	// inicialize workers
	for i := 1; i <= qtsWorkers; i++ {
		go worker(i, ch)
	}

	// consume workers
	for i := 0; i < 1000000; i++ {
		ch <- i
	}
}
