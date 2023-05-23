package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 42
	fmt.Printf("%v, %T\n", m, m)

	var p uint16 = 42
	fmt.Printf("%v, %T\n", p, p)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c int = 10
	var d int8 = 3
	fmt.Println(c + int(d))

}