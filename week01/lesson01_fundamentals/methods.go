package main

import "fmt"

// A method is a function associated with a specific type.

// Methods with structs

type Person struct {
	ID   int
	Name string
}

// Value receiver.
// This method works with a copy of Person,
// so the original struct is not changed.
func (p Person) UpdateName(name string) {
	p.Name = name
}

// Pointer receiver.
// This method works with the original Person,
// so changes are saved.
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	ID   int
	Name string
	Person
}

func (a *Account) SetName(name string) {
	a.Name = name
}

// Methods can be used not only with structs.

type MySlice []int

func (s *MySlice) Add(value int) {
	*s = append(*s, value)
}

func (s MySlice) Count() int {
	return len(s)
}

func main() {
	pers := new(Person)

	pers.SetName("Andrew Mart")

	fmt.Printf("updated person: %+v\n", pers)

	acc := Account{
		ID:   1,
		Name: "AndyMart",
		Person: Person{
			ID:   2,
			Name: "Andrew Mart",
		},
	}

	acc.SetName("Andrew Mart")
	acc.Person.SetName("Test")

	fmt.Printf("updated account: %+v\n", acc)

	sl := MySlice{1, 2}

	sl.Add(5)

	fmt.Println(sl.Count(), sl)
}
