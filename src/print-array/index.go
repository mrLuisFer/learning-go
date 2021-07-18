package main

import (
	"fmt"
)

func main() {
	intArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i:])
	}
}
