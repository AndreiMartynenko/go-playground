package main

import "fmt"

// A map is an unordered collection of key-value pairs.
// Maps provide fast lookups, inserts, updates, and deletes.

func main() {

	// -------------------------------------------------------
	// Creating a map
	// -------------------------------------------------------

	// Empty map
	ages := map[string]int{}

	// Using make
	scores := make(map[string]int)

	// Map literal
	countries := map[string]string{
		"UK":     "London",
		"France": "Paris",
		"Japan":  "Tokyo",
	}

	fmt.Println(ages)
	fmt.Println(scores)
	fmt.Println(countries)

	// -------------------------------------------------------
	// Adding elements
	// -------------------------------------------------------

	ages["John"] = 25
	ages["Alice"] = 31

	fmt.Println(ages)

	// -------------------------------------------------------
	// Updating elements
	// -------------------------------------------------------

	ages["John"] = 30

	fmt.Println(ages)

	// -------------------------------------------------------
	// Reading values
	// -------------------------------------------------------

	fmt.Println(ages["John"])

	// If the key doesn't exist,
	// the zero value is returned.

	fmt.Println(ages["Bob"]) // 0

	// -------------------------------------------------------
	// Checking whether a key exists
	// -------------------------------------------------------

	age, ok := ages["Bob"]

	fmt.Println(age)
	fmt.Println(ok)

	if ok {
		fmt.Println("Bob exists")
	} else {
		fmt.Println("Bob does not exist")
	}

	// -------------------------------------------------------
	// Deleting elements
	// -------------------------------------------------------

	delete(ages, "Alice")

	fmt.Println(ages)

	// delete() never panics
	delete(ages, "Unknown")

	// -------------------------------------------------------
	// Length
	// -------------------------------------------------------

	fmt.Println("length:", len(ages))

	// -------------------------------------------------------
	// Iterating over a map
	// -------------------------------------------------------

	for key, value := range countries {
		fmt.Println(key, "->", value)
	}

	// Keys only
	for key := range countries {
		fmt.Println(key)
	}

	// Values only
	for _, value := range countries {
		fmt.Println(value)
	}

	// -------------------------------------------------------
	// Map lookup
	// -------------------------------------------------------

	capital, ok := countries["Japan"]

	if ok {
		fmt.Println(capital)
	}

	// -------------------------------------------------------
	// Nested maps
	// -------------------------------------------------------

	users := map[string]map[string]string{
		"john": {
			"city": "London",
			"job":  "Developer",
		},
	}

	fmt.Println(users)

	// -------------------------------------------------------
	// Maps are reference types
	// -------------------------------------------------------

	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2 := m1

	m2["a"] = 100

	fmt.Println(m1)
	fmt.Println(m2)

	// -------------------------------------------------------
	// Copying a map
	// -------------------------------------------------------

	copyMap := make(map[string]int)

	for k, v := range m1 {
		copyMap[k] = v
	}

	copyMap["a"] = 999

	fmt.Println(m1)
	fmt.Println(copyMap)

	// -------------------------------------------------------
	// Nil map
	// -------------------------------------------------------

	var nilMap map[string]int

	fmt.Println(nilMap == nil)

	// Reading is allowed
	fmt.Println(nilMap["hello"])

	// Writing causes panic
	// nilMap["hello"] = 1

	// -------------------------------------------------------
	// Clear a map
	// -------------------------------------------------------

	clear(m1)

	fmt.Println(m1)
}
