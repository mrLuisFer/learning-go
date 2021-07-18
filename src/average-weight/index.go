package main

import (
	"fmt"
)

func main() {
	// Sets the maximum value of values that can be searched
	var array [5]int

	var total, count int
 	fmt.Print("How many numbers you want to enter: ")
 	fmt.Scanln(&count)
 	
	 for i := 0; i < count; i++ {  
  	fmt.Print("Enter value: ")

  	fmt.Scanln(&array[i])
  	total += array[i]
 	}

	average := total / count
	fmt.Printf("Average is  %d", average)
}
