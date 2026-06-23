package main

import "fmt"

func main() {
	// The size of an array is part of its type.
	// For example, [3]int and [4]int are different types.

	var a1 [3]int // [0, 0, 0]

	fmt.Println("a1:", a1)
	fmt.Printf("a1 with %%v: %v\n", a1)
	fmt.Printf("a1 with %%#v: %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool // [false, false, false, false]

	fmt.Println("a2:", a2)

	// Let Go determine the array size during initialization.
	a3 := [...]int{1, 2, 3}

	fmt.Println("a3:", a3)

	// Arrays are value types.
	// When you copy an array, all elements are copied.
	a4 := a3
	a4[0] = 100

	fmt.Println("a3 after copy:", a3)   // [1, 2, 3]
	fmt.Println("a4 after change:", a4) // [100, 2, 3]

	// Accessing elements by index.
	fmt.Println("first element:", a3[0])
	fmt.Println("last element:", a3[len(a3)-1])
}
