package main

import "fmt"

func main() {
	d := 1 + 2i
	var e = 2 + 5.2i

	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)
	fmt.Println(real(e))
	fmt.Println(imag(e))
	
}