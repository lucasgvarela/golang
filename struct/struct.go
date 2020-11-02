package main

import (
	"fmt"
)

// Car defines a car
type Car struct {
	CarName string
	CarYear int
}

func main() {
	car1 := Car{
		CarName: "Tesla Model X",
		CarYear: 2020,
	}

	fmt.Println(car1)
}
