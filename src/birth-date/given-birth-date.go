package main

import (
	"fmt"
	"time"
)

func givenBirthDate(n int) {
	var currentDate int = time.Now().Year()
	fmt.Println(currentDate)

	result := currentDate - n

	fmt.Printf("Your birth year is: %d", result)
}

func main() {
	var n int	
	fmt.Print("Write yout age: ")
	fmt.Scanln(&n)
	givenBirthDate(n)
}

