package main

import "fmt"

func main() {

	// Default value
	var num0 int

	// Explicit type
	// int size depends on the platform (32-bit or 64-bit)
	var num1 int = 1

	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// uint size depends on the platform (32-bit or 64-bit)
	var unsignedInt uint = 100500

	// uint8, uint16, uint32, uint64
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(bigInt, unsignedInt, unsignedBigInt)

	// Type is inferred automatically
	var num2 = 2

	fmt.Println(num0, num1, num2)

	// Short variable declaration (used inside functions only)
	num := 30

	// camelCase naming convention
	userIndex := 10

	// Multiple variable declarations
	var weight, height = 87, 180

	// Short declaration of multiple variables
	age, weight := 23, 67

	fmt.Println(num, weight, age, weight, userIndex, height)

	// float32, float64
	var pi float32 = 3.141
	var e = 2.718

	goldenRatio := 1.618

	fmt.Println(pi, e, goldenRatio)

	// bool
	var b bool // false by default
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b, isOk, success, cond)

	// complex64, complex128
	var c complex128 = 1.1 + 7.12i
	c2 := -1.1 + 7.12i

	fmt.Println(c, c2)
}
