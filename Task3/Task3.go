package main

import (
	"fmt"
	"sync"
	"time"
)

func Workers(work int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d has Started Work %d\n", work, job)
		time.Sleep(5 * time.Millisecond)
		fmt.Printf("Worker %d has Completed Work %d\n", work, job)
	}
}

func main() {
	var wg sync.WaitGroup
	no_workers := 3
	no_jobs := 5
	jobs := make(chan int)

	for i := 1; i <= no_workers; i++ {
		wg.Add(1)
		go Workers(i, jobs, &wg)
	}

	for j := 1; j <= no_jobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()

	fmt.Println("Jobs Completed Successfully....")
}
