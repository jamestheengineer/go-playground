package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	//firstName string
	//lastName string
	//city string
	//gender string
	//age int

	firstName, lastName, city, gender string
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Birthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// Change last name after marrried
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Samantha", "Smith", "Boston", "f", 25}
	person3 := Person{"Joe", "Smith", "Boston", "m", 25}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Doe")
	fmt.Println(person1.greet())

	person3.getMarried("Doe")
	fmt.Println(person3.greet())
}