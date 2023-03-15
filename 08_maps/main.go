package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)
	emails2 := map[string]string{"bob":"bob@gmaail.com", "sharon":"sharon@gmail.com"}

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Bob")
	fmt.Println(emails)

	fmt.Println(emails2)
}