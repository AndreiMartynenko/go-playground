package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func student(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "student is thinking", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("student", workerNum, "came up with")
		out <- workerNum
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go student(ctx, i, result)
	}

	foundBy := <-result
	fmt.Println("the question was asked by a student", foundBy)
	finish()

	time.Sleep(time.Second)
}
