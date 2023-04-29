package main

import "fmt"

var a = 42
var (
	actorName string = "Joe"
	companion string = "Jack"
	doctorNumber int = 3
	season int = 11
)

func main() {
	var k int
	k = 42
	var i int = 42
	j := 42

	fmt.Println(i, j, k)
	fmt.Printf("%v, %T\n", j, j)
}