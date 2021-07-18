package main

import (
	"fmt"
)

func sayHello(name string)  {
	msg := fmt.Sprintf("Hello %v, to learning Go! :D", name)
	fmt.Println(msg)
}


func basicMath(n int, n2 int, operationSymbol string) {
	
	var result int

	if operationSymbol == "+" {
		result = n + n2
	} else if operationSymbol == "-" {
		result = n - n2
	} else if operationSymbol == "/" {
		result = n / n2
	} else if operationSymbol == "*" {
		result = n * n2
	}

	// print the result
	fmt.Println(result)
}


func mathPrompt() {
	var (
		num1 int
		num2 int
		operationSymbol string
	)

	fmt.Print("Insert number:")
	fmt.Scanln(&num1)

	fmt.Print("Insert number:")
	fmt.Scanln(&num2)

	fmt.Print("Insert operation:")
	fmt.Scanln(&operationSymbol)

		if operationSymbol == "+" || operationSymbol == "add" {
		fmt.Printf("\n Result: %d",num1 + num2)
	} else if operationSymbol == "-" || operationSymbol == "sub" {
		fmt.Printf("\n Result: %d",num1 - num2)
	} else if operationSymbol == "/" || operationSymbol == "div" {
		fmt.Printf("\n Result: %d",num1 / num2)
	} else if operationSymbol == "*" || operationSymbol == "mult" {
		fmt.Printf("\n Result: %d",num1 * num2)
	}
}

func main() {
		sayHello("Luis")
		basicMath(3,4, "+")
		basicMath(5,9, "-")

		mathPrompt()
}
