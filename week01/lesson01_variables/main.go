package main

import "fmt"

func main() {

	// value by default
	var num0 int

	// value after initialization
	var num1 int = 1

	// skip the type of variable
	var num2 = 2

	fmt.Println(num0, num1, num2)

	// short version
	// if I want to return something from the func
	// Impossible to update
	num := 30

	//camelCase
	userIndex := 10

	// initialization of several variables

	var weight, height = 87, 180

	// short version of init several variables

	age, weight := 23, 67

	fmt.Println(num, weight, age, weight, userIndex, height)

}
