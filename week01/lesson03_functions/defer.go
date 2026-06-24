package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars result"

}

// defer is to delay the func execution

func main() {
	defer fmt.Println("After work. FIRST")
	//arguments of the func determine when defer was invoke? allocate not when the func was invoked
	//defer fmt.Println(getSomeVars())

	// often anonym func can be used
	defer func() {
		fmt.Println(getSomeVars())

	}()

	fmt.Println("Some useful work")
}

// defer is also very useful as a recovery after panic
