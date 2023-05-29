package main

import "fmt"

func main() {
	a := 10 // 1010
	b := 3 // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100

	c := 8
	fmt.Println( c << 3) // 2^3 * 2^3
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0
}