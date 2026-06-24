package main

import (
	"fmt"
)

type Car struct {
	Brand   string
	Model   string
	Year    int
	Mileage int
}

func (c Car) Info() string {
	return fmt.Sprintf("%s %s (%d)\nMileage: %d km", c.Brand, c.Model, c.Year, c.Mileage)
}

func (c *Car) Drive(km int) {
	c.Mileage += km

}

func (c Car) Age(currentYear int) int {
	return currentYear - c.Year
}

func main() {
	car := Car{
		Brand:   "Bentley",
		Model:   "GTC",
		Year:    2025,
		Mileage: 2675,
	}

	fmt.Println(car.Info())
	car.Drive(30)

	fmt.Println(car.Info())
	fmt.Println("Age:", car.Age(2026))

}
