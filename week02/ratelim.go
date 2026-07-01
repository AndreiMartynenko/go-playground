package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2 // Maximum number of workers allowed simultaneously.
)

func startWorker(workerID int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	// Acquire a slot.
	// If the channel is full, this goroutine blocks until a slot becomes available.
	quotaCh <- struct{}{}

	// Always release resources when the goroutine finishes.
	defer wg.Done()
	defer func() {
		<-quotaCh // Release the slot.
	}()

	for iteration := 0; iteration < iterationsNum; iteration++ {
		fmt.Print(formatWork(workerID, iteration))

		// Demonstration only.
		// Every second iteration, voluntarily release the slot
		// and immediately acquire it again.
		if iteration%2 == 0 {
			<-quotaCh             // Release slot.
			quotaCh <- struct{}{} // Acquire slot again.
		}

		// Yield execution to another goroutine.
		// Mostly useful for demonstrations.
		runtime.Gosched()
	}
}

func main() {
	var wg sync.WaitGroup

	// Buffered channel used as a semaphore.
	quotaCh := make(chan struct{}, quotaLimit)

	for workerID := 0; workerID < goroutinesNum; workerID++ {
		wg.Add(1)
		go startWorker(workerID, &wg, quotaCh)
	}

	// Wait until all workers finish.
	wg.Wait()
}

func formatWork(workerID, iteration int) string {
	return fmt.Sprintf(
		"%s█%sth %d iter %d %s\n",
		strings.Repeat("  ", workerID),
		strings.Repeat("  ", goroutinesNum-workerID),
		workerID,
		iteration,
		strings.Repeat("■", iteration),
	)
}
