package main

import "fmt"

func main() {
	// Main Types
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint 32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name = "Brad"
	var age int32 = 37

	fmt.Println(name, age)
	fmt.Printf("%T\n%T\n", name, age)
}