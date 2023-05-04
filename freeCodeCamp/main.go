package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 2 == 1
	fmt.Printf("%v, %T\n", m, n)

}