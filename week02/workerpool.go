package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched() // try to comment it out
	}
	printFinishWork(workerNum)
}

func main() {
	runtime.GOMAXPROCS(0)               // try from 0 (all possibilities) and 1
	worketInput := make(chan string, 2) // try to increase the size of the channel
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, worketInput)
	}

	months := []string{"January", "February", "March",
		"April", "May", "June",
		"July", "August", "September",
		"October", "November", "December",
	}

	for _, monthName := range months {
		worketInput <- monthName
	}
	close(worketInput) // try to comment it out

	time.Sleep(time.Millisecond)
}

func formatWork(in int, input string) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"recieved", input)
}

func printFinishWork(in int) {
	fmt.Println(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"===", in,
		"finished")
}
