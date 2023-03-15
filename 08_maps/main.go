package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
}