package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}
func main() {
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Samantha", "Smith", "Boston", "f", 25}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
}