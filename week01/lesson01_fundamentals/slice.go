package main

import "fmt"

// Slice is a dynamic view over an underlying array.
// A slice has three main properties:
// 1. pointer to the underlying array
// 2. length
// 3. capacity

func main() {
	// Creating a slice
	nums := []int{1, 2, 3}
	fmt.Println("nums:", nums)

	// len = number of elements
	// cap = available space in the underlying array
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))

	// Nil slice
	var nilSlice []int
	fmt.Println("nilSlice:", nilSlice)
	fmt.Println("len:", len(nilSlice))
	fmt.Println("cap:", cap(nilSlice))
	fmt.Println("is nil:", nilSlice == nil)

	// Creating a slice with make
	// make([]type, length, capacity)
	buf := make([]int, 3, 5)
	fmt.Println("buf:", buf)
	fmt.Println("len:", len(buf))
	fmt.Println("cap:", cap(buf))

	// Updating elements
	buf[0] = 10
	buf[1] = 20
	buf[2] = 30
	fmt.Println("updated buf:", buf)

	// append adds elements to a slice
	buf = append(buf, 40)
	buf = append(buf, 50)
	fmt.Println("buf after append:", buf)
	fmt.Println("len:", len(buf))
	fmt.Println("cap:", cap(buf))

	// If capacity is not enough, append creates a new underlying array
	buf = append(buf, 60)
	fmt.Println("buf after capacity growth:", buf)
	fmt.Println("len:", len(buf))
	fmt.Println("cap:", cap(buf))

	// Copying a slice
	original := []int{1, 2, 3, 4, 5}

	copied := make([]int, len(original))
	n := copy(copied, original)

	fmt.Println("original:", original)
	fmt.Println("copied:", copied)
	fmt.Println("number of copied elements:", n)

	// Changing copied slice does not affect original slice
	copied[0] = 100

	fmt.Println("original after copied change:", original)
	fmt.Println("copied after change:", copied)

	// copy copies only the minimum length
	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)

	copiedCount := copy(dst, src)

	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
	fmt.Println("copiedCount:", copiedCount)
}
