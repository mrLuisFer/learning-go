package main

import (
	"os"
)

func createFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Texto escrito con golang")
}

func main() {
	var file = "cities.txt"
	createFile(file)
}
