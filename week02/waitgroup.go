package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

const (
	iterationsNum = 6
	goroutinesNum = 6
)

func formatWork(workerID, iteration int) string {
	return fmt.Sprintf(
		"worker %d: %s%d\n",
		workerID,
		strings.Repeat(" ", iteration),
		iteration,
	)
}

func doWork(workerID int, wg *sync.WaitGroup) {
	// Decrease the WaitGroup counter when the goroutine finishes.
	defer wg.Done()

	for iteration := 0; iteration < iterationsNum; iteration++ {
		fmt.Print(formatWork(workerID, iteration))

		// Yield execution and allow the Go scheduler
		// to run another goroutine.
		//
		// This is useful for learning, but rarely used in production code.
		runtime.Gosched()
	}
}

func main() {
	var wg sync.WaitGroup

	for workerID := 0; workerID < goroutinesNum; workerID++ {
		// Increase the WaitGroup counter before starting a goroutine.
		wg.Add(1)

		go doWork(workerID, &wg)
	}

	// Wait until all goroutines call wg.Done().
	wg.Wait()
}
