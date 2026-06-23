package main

import "fmt"

// Constants

const pi = 3.141

const (
	hello = "Hello"
	e     = 2.718
)

const (
	zero  = iota
	_     // skip this value
	three // = 2
)

const (
	_       = iota             // skip the first value
	KB uint = 1 << (10 * iota) // 1024
	MB                         // 1048576
)

const (
	// Untyped constant
	year = 2026

	// Typed constant
	yearTyped int = 2026
)

func main() {
	var month int32 = 12

	fmt.Println(month + year)
	fmt.Println(yearTyped)
}
