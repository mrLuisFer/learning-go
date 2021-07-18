package main

import (
	"fmt"
)

func evenOrOdd(n int) bool {
	fmt.Scanln(&n)

	if n % 2 == 0 {
		println("Is even || pair \n")
		return true
	} else {
		println("Is odd \n")
		return false
	}
}

func main() {
	fmt.Print("Insert number: ")
	var n int


	var isEvenOrOdd bool = evenOrOdd(n)
	fmt.Println(isEvenOrOdd)
}