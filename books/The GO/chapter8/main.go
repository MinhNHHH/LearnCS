package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a task to be done
type Job struct {
	id    int
	value int
}

// Worker function that processes jobs
func worker(id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.id)
		// Simulate work by sleeping
		time.Sleep(time.Second)
		// Send the result to the results channel
		results <- job.value * 2
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	var wg sync.WaitGroup

	jobs := make(chan Job, numJobs)
	results := make(chan int, numJobs)

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{id: j, value: j}
	}
	close(jobs) // Close the job channel to signal no more jobs

	// Wait for all workers to complete
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println("All jobs processed")
}
