package main

import "fmt"

func main() {
	// -------------------------------------------------------
	// if
	// -------------------------------------------------------

	age := 20

	if age >= 18 {
		fmt.Println("adult")
	}

	// -------------------------------------------------------
	// if else
	// -------------------------------------------------------

	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}

	// -------------------------------------------------------
	// if else if else
	// -------------------------------------------------------

	score := 85

	if score >= 90 {
		fmt.Println("grade A")
	} else if score >= 80 {
		fmt.Println("grade B")
	} else if score >= 70 {
		fmt.Println("grade C")
	} else {
		fmt.Println("grade D")
	}

	// -------------------------------------------------------
	// if with short statement
	// -------------------------------------------------------

	if value := 10; value > 5 {
		fmt.Println("value is greater than 5")
	}

	// value is available only inside this if block.
	// fmt.Println(value) // not allowed

	// -------------------------------------------------------
	// if with map lookup
	// -------------------------------------------------------

	users := map[string]int{
		"John":  25,
		"Alice": 30,
	}

	if age, ok := users["John"]; ok {
		fmt.Println("John exists, age:", age)
	} else {
		fmt.Println("John does not exist")
	}

	// -------------------------------------------------------
	// Boolean operators
	// -------------------------------------------------------

	isLoggedIn := true
	isAdmin := false

	if isLoggedIn && isAdmin {
		fmt.Println("admin panel")
	}

	if isLoggedIn || isAdmin {
		fmt.Println("user has access")
	}

	if !isAdmin {
		fmt.Println("user is not admin")
	}

	// -------------------------------------------------------
	// switch
	// -------------------------------------------------------

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "Friday":
		fmt.Println("almost weekend")
	case "Saturday", "Sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("regular day")
	}

	// -------------------------------------------------------
	// switch without expression
	// -------------------------------------------------------

	temperature := 25

	switch {
	case temperature < 0:
		fmt.Println("freezing")
	case temperature < 20:
		fmt.Println("cold")
	case temperature < 30:
		fmt.Println("warm")
	default:
		fmt.Println("hot")
	}

	// -------------------------------------------------------
	// switch with short statement
	// -------------------------------------------------------

	switch status := 200; status {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Internal Server Error")
	default:
		fmt.Println("Unknown status")
	}

	// -------------------------------------------------------
	// No automatic fallthrough
	// -------------------------------------------------------

	number := 2

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// In Go, switch automatically stops after the matching case.
	// There is no need to write break.

	// -------------------------------------------------------
	// fallthrough
	// -------------------------------------------------------

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other number")
	}

	// fallthrough forces execution of the next case.
	// It is rarely used in real Go code.

	// -------------------------------------------------------
	// switch with multiple conditions
	// -------------------------------------------------------

	role := "admin"

	switch role {
	case "admin", "owner":
		fmt.Println("full access")
	case "manager":
		fmt.Println("limited access")
	case "user":
		fmt.Println("basic access")
	default:
		fmt.Println("no access")
	}

	// -------------------------------------------------------
	// type switch
	// -------------------------------------------------------

	var value any = "hello"

	switch v := value.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	case bool:
		fmt.Println("bool:", v)
	default:
		fmt.Println("unknown type")
	}

	// -------------------------------------------------------
	// if vs switch
	// -------------------------------------------------------

	httpStatus := 201

	if httpStatus >= 200 && httpStatus < 300 {
		fmt.Println("success")
	} else if httpStatus >= 400 && httpStatus < 500 {
		fmt.Println("client error")
	} else if httpStatus >= 500 {
		fmt.Println("server error")
	}

	switch {
	case httpStatus >= 200 && httpStatus < 300:
		fmt.Println("success")
	case httpStatus >= 400 && httpStatus < 500:
		fmt.Println("client error")
	case httpStatus >= 500:
		fmt.Println("server error")
	}
}
