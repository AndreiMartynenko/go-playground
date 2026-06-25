package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}
type Cat struct {
	Name string
}

type Cow struct {
	Name string
}

func (dog Dog) Speak() string {
	return "Woof!"
}

func (cat Cat) Speak() string {
	return "Meow!"
}
func (cow Cow) Speak() string {
	return "Moo!"
}

func MakeSpeak(a Animal) {
	fmt.Println(a.Speak())
}

//func main() {
//	dog := Dog{Name: "Bob"}
//	cat := Cat{Name: "Kitty"}
//	cow := Cow{Name: "Bella"}
//
//	MakeSpeak(dog)
//	MakeSpeak(cat)
//	MakeSpeak(cow)
//
//}
