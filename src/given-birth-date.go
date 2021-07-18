package main

import (
	"fmt"
)

func givenBirthDate(n int) {

	result := 2021 - n

	fmt.Printf("Your birth date is: %d", result)
}

func main() {
	var n int	
	fmt.Print("Write yout age: ")
	fmt.Scanln(&n)
	givenBirthDate(n)
}

