package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range numbers {
		var evenOrOdd string

		if i % 2 == 0 {
			evenOrOdd = "even"
		} else{
			evenOrOdd = "odd"
		}

		fmt.Printf("%v is %v\n", i, evenOrOdd)
	}
}