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

	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	const constAge int = 54

	fmt.Printf("Monday %d - Tuesday %d Wednesday %d\n", Monday, Tuesday, Wednesday)

	// Increments the next integer value for
	// subsequent declarations
	const (
		Jan int = iota + 1
		Feb
		Mar
		Apr
	)

	fmt.Printf("jan - %d feb - %d mar - %d apr - %d\n", Jan, Feb, Mar, Apr)

	addResult, multiplyResult := add(3, 4)

	fmt.Printf("This is the result %d, %d\n", addResult, multiplyResult)

	myLoop()
	myArrays()
	mySlices()
}

// Is Exportable
func Add() {

}

// Is Not Exportable
// The int within the parameters below covers
// the grounds of all of the parameters leading up to it

// This return value is not a tuple
// Multiple returns are a first class citizen
func add(a, b int) (int, int) {
	// Go can ignore return types
	// because errors are values
	return a + b, a * b
}

func myLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	counter := 0
	for counter < 3 {
		counter++
		fmt.Println("This is the counter", counter)
	}

	iterations := 0
	for {

		if iterations > 3 {
			break
		}
		iterations++
		fmt.Println("Num of iterations", iterations)
	}
}

func myArrays() {
	numbers := [5]int{10, 20, 30, 40, 50}
	// array length cannot change. You can change and reassign values but you can't reassign length
	fmt.Printf("this is our array, %v\n", numbers)
	fmt.Printf("Our length is %d\n", len(numbers))
	fmt.Printf("Last value is %d\n", numbers[len(numbers)-1])

	matrix := [2][3]int{
		{1, 2, 3},
		{5, 8, 123},
	}

	fmt.Printf("This is my matrix %v\n", matrix)

	// unboundedNumbers := [...]{2, 3, 4, 5}
}

func mySlices() {
	// A slice or a portion of an array
	numbers := [2]int{1, 2}

	allNumbers := numbers[:]

	first := numbers[0:1]

	fmt.Printf("These are all numbers %v\n", allNumbers)
	fmt.Printf("These are first %v\n", first)

	fruits := []string{"apple", "banana", "strawberry", "blueberry"}

	fmt.Printf("These are my fruits %v\n", fruits)
	fruits = append(fruits, "kiwi")
	fmt.Printf("The length of fruits is %d\n", len(fruits)) //5 So the memory abstraction of length is hidden underneath
	fmt.Printf("These are my fruits %v\n", fruits)
}
