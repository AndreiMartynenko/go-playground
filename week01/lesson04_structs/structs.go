package main

import "fmt"

// Person represents a person.
type Person struct {
	ID      int
	Name    string
	Address string
}

// Account demonstrates struct composition and embedding.
type Account struct {
	ID   int
	Name string

	Bartender func(string) string

	Owner Person // Regular field

	Person // Embedded field
}

func main() {

	// Full struct initialization using named fields.
	acc := Account{
		ID:   1,
		Name: "John Doe",

		Person: Person{
			Name:    "Andrew Mart",
			Address: "London",
		},
	}

	fmt.Printf("%#v\n", acc)

	// Short struct initialization using field order.
	acc.Owner = Person{2, "John Doe", "New York"}

	fmt.Printf("%#v\n", acc)

	// Accessing promoted fields from the embedded struct.
	fmt.Println(acc.Address)

	// Accessing Account.Name
	fmt.Println(acc.Name)

	// Accessing Person.Name explicitly
	fmt.Println(acc.Person.Name)
}
