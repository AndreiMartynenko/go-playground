package main

import "fmt"

// defer is often used for cleanup:
// closing files, unlocking mutexes, closing database connections, etc.
//
// defer can also be used with recover to handle panics.
//
// panic stops the normal execution of the current goroutine.
// After panic, deferred functions are executed in LIFO order.
// LIFO = last in, first out.

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recovered in FIRST defer:", err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recovered in SECOND defer:", err)
			panic("second panic")
		}
	}()

	fmt.Println("Some useful work")

	panic("something bad happened")

	// This line is unreachable.
	// return
}

func main() {
	deferTest()
}
