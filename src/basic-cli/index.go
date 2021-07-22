package main

import (
	"fmt"
	"strings"
)

func print (param interface{}) {
	fmt.Print(param)
	fmt.Printf("\n")
}

func toLowerCase (text string) string {
	var textFormatted = strings.ToLower(text)
	return textFormatted
}

func main() {
	print("Insert any of the following commands:")

	var cmds = []string{"help", "add"}
	for _, cmd := range cmds {
		fmt.Printf("%v ", cmd)
	}
	print(":")

	var command string
	fmt.Scan(&command)
	command = toLowerCase(command)

	if (command == "help") {
		print("Prints this help message")
		help()	
	}
}
