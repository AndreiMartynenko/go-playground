package main

import "fmt"

func main() {
	a := 2
	b := &a // b points to the address of a

	fmt.Println(a)  // 2
	fmt.Println(&a) // memory address, e.g. 0xc0000120a8
	fmt.Println(b)  // same address as &a
	fmt.Println(*b) // 2 (dereferencing the pointer)

	*b = 3  // a = 3
	c := &a // c points to a

	// Create a pointer to a new int value.
	// The default value is 0.
	d := new(int)

	*d = 12
	*c = *d // a = 12

	*d = 13 // a does not change

	c = d   // now c points to the same value as d
	*c = 14 // c = 14, d = 14, a = 12

	fmt.Println(a, *b, *c, *d)
}
