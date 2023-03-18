package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	// Declare and assign
	fruitArr2 := [2]string{"Apple, Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])


}