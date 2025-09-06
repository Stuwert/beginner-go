package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	var name string = "Stuart"

	fmt.Printf("this is my name %s\n", name)

	age := 34

	// Passing %s for an int will print the int along with the type
	fmt.Printf("this is my age %d\n", age)

	var city string

	city = "Denver"
	fmt.Printf("this is my city %s\n", city)

	var state, continent string = "Colorado", "North America"

	fmt.Printf("%s is in %s %s\n", city, state, continent)

	var (
		isEmployed bool   = true
		salary     int    = 10000
		title      string = "developer"
	)

	fmt.Printf("Is Employed? %t with a salary %d and position %s\n", isEmployed, salary, title)

	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("Let's learn about some defaults int %d float %f string %s boolean %t\n", defaultInt, defaultFloat, defaultString, defaultBool)
}
