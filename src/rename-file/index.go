package main

import (
	"fmt"
	"os"
)

func main() {
	src := "./src/rename-file/test.txt"
	dst := "./src/rename-file/test2.txt"

	fmt.Println("Renaming the file...")
	err := os.Rename(src, dst)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File renamed successfully")
	}
}
