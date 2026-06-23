package main

import "fmt"

// Functions are first-class values in Go.
// It means that we can:
// 1. assign a function to a variable;
// 2. pass a function as an argument;
// 3. return a function from another function;
// 4. store a function in a struct field.

func doNothing() {
	fmt.Println("do nothing")
}

func main() {
	// Anonymous function
	func(in string) {
		fmt.Println("anonymous function output:", in)
	}("nobody")

	// Assigning an anonymous function to a variable
	printer := func(in string) {
		fmt.Println("printer output:", in)
	}

	printer("as variable")

	// Defining a function type
	type stringFunc func(string) error

	// Function takes a callback
	worker := func(callback stringFunc) {
		err := callback("as callback")
		if err != nil {
			fmt.Println("error:", err)
		}
	}

	worker(func(in string) error {
		fmt.Println("callback output:", in)
		return nil
	})

	// Function returns a closure
	prefixer := func(prefix string) func(string) {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	successLogger := prefixer("success")
	successLogger("expected behaviour")

	doNothing()
}
