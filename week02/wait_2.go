package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func startWorker(in int, waiter *sync.WaitGroup) {
	defer waiter.Done() // wait_2.go decreases the counter by 1
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		time.Sleep(time.Millisecond) // try removing this sleep.
	}
}

func main() {
	wg := &sync.WaitGroup{} // wait_2.go initialize the group
	for i := 0; i < goroutinesNum; i++ {
		// wg.Add it needs to be called in the goroutine that spawns the workers.
		// otherwise, another goroutine might not start in time, and `wait` might execute
		wg.Add(1) // wait_2.go adding
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait() // wait_2.go we wait until waiter.Done() sets the counter to 0.

	fmt.Println(11111)

}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}
