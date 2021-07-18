package main

import (
	"fmt"
)

func main() {
	var n float64

	fmt.Println("Write a number: ")
	fmt.Scanln(&n)

	if (n > 0) {
		result := n / 2

		fmt.Printf("Result: %f", result)
	} else {
		fmt.Printf("The number %f is less than or equal to 0", n)
	}
}
