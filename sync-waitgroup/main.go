package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	// some task is happening
	fmt.Printf("worker %d finished\n", i)
}

func main() {
	fmt.Println("Leaning wait group...")

	var wg sync.WaitGroup

	//start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) //increment the waitgroup counter]
		go worker(i, &wg)
	}
	//wait for all workers to finish
	wg.Wait()
	//main goroutine completes the task and exits gracefully  [no need to call wg.Done() here]
	fmt.Println("workers task completed successfully")

}
