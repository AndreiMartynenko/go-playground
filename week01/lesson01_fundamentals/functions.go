package main

import (
	"errors"
	"fmt"
)

// -------------------------------------------------------
// Simple function
// -------------------------------------------------------

func sayHello() {
	fmt.Println("Hello!")
}

// -------------------------------------------------------
// Function with parameters
// -------------------------------------------------------

func greet(name string) {
	fmt.Println("Hello,", name)
}

// Multiple parameters of the same type
func add(a, b int) int {
	return a + b
}

// -------------------------------------------------------
// Multiple return values
// -------------------------------------------------------

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

// -------------------------------------------------------
// Named return values
// -------------------------------------------------------

func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

// -------------------------------------------------------
// Ignoring return values
// -------------------------------------------------------

func fullName() (string, string) {
	return "John", "Doe"
}

// -------------------------------------------------------
// Variadic function
// -------------------------------------------------------

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

// -------------------------------------------------------
// Passing a slice to a variadic function
// -------------------------------------------------------

func average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	total := 0.0

	for _, n := range nums {
		total += n
	}

	return total / float64(len(nums))
}

// -------------------------------------------------------
// Functions are values
// -------------------------------------------------------

func multiply(a, b int) int {
	return a * b
}

// -------------------------------------------------------
// Higher-order function
// -------------------------------------------------------

func calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// -------------------------------------------------------
// Anonymous function
// -------------------------------------------------------

func main() {

	sayHello()

	greet("Andrew")

	fmt.Println(add(10, 20))

	result, err := divide(10, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	_, surname := fullName()

	fmt.Println(surname)

	area, perimeter := rectangle(5, 10)

	fmt.Println(area)
	fmt.Println(perimeter)

	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3, 4, 5))

	values := []float64{5, 10, 15}

	fmt.Println(average(values...))

	// Function assigned to a variable

	f := multiply

	fmt.Println(f(6, 7))

	// Anonymous function

	func() {
		fmt.Println("anonymous function")
	}()

	// Anonymous function with parameters

	func(name string) {
		fmt.Println("Hello", name)
	}("Andrew")

	// Function as an argument

	answer := calculate(10, 5, add)

	fmt.Println(answer)

	answer = calculate(10, 5, multiply)

	fmt.Println(answer)

	// Anonymous function as an argument

	answer = calculate(10, 5, func(a, b int) int {
		return a - b
	})

	fmt.Println(answer)
}
