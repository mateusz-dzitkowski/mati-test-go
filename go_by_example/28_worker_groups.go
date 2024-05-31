package main

import (
	"fmt"
	"time"
)

func worker28(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker:", id, "started job:", j)
		time.Sleep(time.Second)
		fmt.Println("Worker:", id, "finished job:", j)
		results <- j * j
	}
}

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := range 3 {
		go worker28(w, jobs, results)
	}
	for j := range numJobs {
		jobs <- j
	}
	close(jobs)
	for range numJobs {
		<-results
	}
}
