package main

import "fmt"

// A subslice is a slice created from another slice.
// It references the same underlying array.

func main() {
	buf := []int{1, 2, 3, 4, 5}

	fmt.Println("buf:", buf)
	fmt.Println("len:", len(buf))
	fmt.Println("cap:", cap(buf))

	// Basic slicing
	sl1 := buf[1:4] // elements from index 1 to index 3
	sl2 := buf[:2]  // elements from the beginning to index 1
	sl3 := buf[2:]  // elements from index 2 to the end
	sl4 := buf[:]   // all elements

	fmt.Println("sl1:", sl1) // [2 3 4]
	fmt.Println("sl2:", sl2) // [1 2]
	fmt.Println("sl3:", sl3) // [3 4 5]
	fmt.Println("sl4:", sl4) // [1 2 3 4 5]

	// Subslice shares the same underlying array with the original slice
	sl1[0] = 100

	fmt.Println("buf after sl1[0] = 100:", buf)
	fmt.Println("sl1:", sl1)

	// This means:
	// sl1[0] is the same underlying element as buf[1]

	// append to subslice can be dangerous
	buf = []int{1, 2, 3, 4, 5}

	part := buf[1:3] // [2 3]
	fmt.Println("part before append:", part)
	fmt.Println("buf before append:", buf)

	part = append(part, 999)

	fmt.Println("part after append:", part)
	fmt.Println("buf after append:", buf)

	// Why did buf change?
	// Because part still had enough capacity.
	// append reused the same underlying array.

	// To prevent this, use full slice expression:
	// s[low:high:max]
	buf = []int{1, 2, 3, 4, 5}

	safePart := buf[1:3:3] // len = 2, cap = 2
	fmt.Println("safePart before append:", safePart)
	fmt.Println("buf before safe append:", buf)

	safePart = append(safePart, 999)

	fmt.Println("safePart after append:", safePart)
	fmt.Println("buf after safe append:", buf)

	// Now buf did not change because append had to create a new underlying array.

	// If you need a real independent copy, use make + copy.
	buf = []int{1, 2, 3, 4, 5}

	sub := buf[1:4] // [2 3 4]

	independent := make([]int, len(sub))
	copy(independent, sub)

	independent[0] = 777

	fmt.Println("buf after independent change:", buf)
	fmt.Println("sub:", sub)
	fmt.Println("independent:", independent)
}
