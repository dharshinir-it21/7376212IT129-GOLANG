package main

import (
	"fmt"
	"golang/external_module"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println(external_module.Greet(name))

	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	fmt.Println("Sum:", external_module.Add(num1, num2))
}
