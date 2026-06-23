package main

import "fmt"

func main() {
	// -------------------------------------------------------
	// Basic for loop
	// -------------------------------------------------------

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// -------------------------------------------------------
	// for as while
	// -------------------------------------------------------

	count := 0

	for count < 3 {
		fmt.Println("count:", count)
		count++
	}

	// -------------------------------------------------------
	// Infinite loop
	// -------------------------------------------------------

	x := 0

	for {
		if x == 3 {
			break
		}

		fmt.Println("x:", x)
		x++
	}

	// -------------------------------------------------------
	// continue
	// -------------------------------------------------------

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println("continue example:", i)
	}

	// -------------------------------------------------------
	// break
	// -------------------------------------------------------

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println("break example:", i)
	}

	// -------------------------------------------------------
	// Iterating over an array
	// -------------------------------------------------------

	arr := [3]int{10, 20, 30}

	for i := 0; i < len(arr); i++ {
		fmt.Println("array index:", i, "value:", arr[i])
	}

	for index, value := range arr {
		fmt.Println("array range:", index, value)
	}

	// -------------------------------------------------------
	// Iterating over a slice
	// -------------------------------------------------------

	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println("slice index:", i, "value:", nums[i])
	}

	for index, value := range nums {
		fmt.Println("slice range:", index, value)
	}

	// Index only
	for index := range nums {
		fmt.Println("slice index only:", index)
	}

	// Value only
	for _, value := range nums {
		fmt.Println("slice value only:", value)
	}

	// -------------------------------------------------------
	// Iterating over a map
	// -------------------------------------------------------

	ages := map[string]int{
		"John":  25,
		"Alice": 30,
		"Bob":   40,
	}

	for key, value := range ages {
		fmt.Println("map key:", key, "value:", value)
	}

	// Keys only
	for key := range ages {
		fmt.Println("map key only:", key)
	}

	// Values only
	for _, value := range ages {
		fmt.Println("map value only:", value)
	}

	// Important:
	// Map iteration order is not guaranteed.

	// -------------------------------------------------------
	// Iterating over a string
	// -------------------------------------------------------

	text := "Hello, 世界"

	// By bytes
	for i := 0; i < len(text); i++ {
		fmt.Println("byte index:", i, "byte value:", text[i])
	}

	// By runes
	for index, char := range text {
		fmt.Println("rune index:", index, "rune:", char, "string:", string(char))
	}

	// Important:
	// range over string returns:
	// 1. byte index
	// 2. rune value

	// -------------------------------------------------------
	// Nested loops
	// -------------------------------------------------------

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("i:", i, "j:", j)
		}
	}

	// -------------------------------------------------------
	// Loop labels
	// -------------------------------------------------------

OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break OuterLoop
			}

			fmt.Println("labeled loop:", i, j)
		}
	}

	// -------------------------------------------------------
	// Modifying slice values
	// -------------------------------------------------------

	values := []int{1, 2, 3}

	for i := range values {
		values[i] *= 2
	}

	fmt.Println("modified values:", values)

	// This does NOT modify the slice.
	for _, value := range values {
		value *= 2
	}

	fmt.Println("values after value copy:", values)

	// -------------------------------------------------------
	// Common mistake with range value copy
	// -------------------------------------------------------

	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{Name: "John", Age: 25},
		{Name: "Alice", Age: 30},
	}

	// Wrong if you want to modify original users.
	for _, user := range users {
		user.Age++
	}

	fmt.Println("users after wrong update:", users)

	// Correct way.
	for i := range users {
		users[i].Age++
	}

	fmt.Println("users after correct update:", users)
}
