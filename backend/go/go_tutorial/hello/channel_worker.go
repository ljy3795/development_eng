package main

import (
	"fmt"
	"time"
)

// channel using worker pool


// These workers will receive work on the 'jobs' channel and send corresponding results on 'results'
func worker(id int, jobs<-chan int, results chan<-int){
	for j := range jobs {
		fmt.Println("Worker", id, "started Job", j)
		time.Sleep(1*time.Second)
		fmt.Println("Worker", id, "finished Job", j)
		results <- j * 2
		// fmt.Println("Worker ", id)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This setup 3 workers but blocked cuz of no jobs yet (results <- j*2 에서 blocked 됨. 하지만 아래 for loop는 지나감 (go써서))
	// --> that is, initial workder setup
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs and then close channel
	for j:=1; j<=5; j++{
		jobs <- j
	}
	close(jobs)

	for a :=1; a<=5; a++{
		<- results
	}
}