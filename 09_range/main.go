package main

import "fmt"

func main() {
	ids := []int{33,76,54,23,11,2}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("%d - ID: %d\n", id, id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
		fmt.Printf("%d - ID: %d\n", sum, id)
	}

	// Range with map
	emails := map[string]string{"bob":"bob@gmaail.com", "sharon":"sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}