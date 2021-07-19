package main

import (
	"fmt"
	"os"
)

func exists(filePath string) bool {
	var exists bool
	
	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("File %q exists! \n", filePath)
		exists = true
	} else {
		fmt.Printf("File doesn't %q exists! \n", filePath)
		exists = false
	}

	return exists
}

func main() {
	var exist bool = exists("../print-array/index.go")

	fmt.Printf("%v", exist)
}
