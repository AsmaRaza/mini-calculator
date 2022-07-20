package main

import "fmt"

//import "math"
func main() {
	var operation string
	var n1, n2 int
	fmt.Println("enter the ist number")
	fmt.Scanln(&n1)
	fmt.Println("enter the 2nd number")
	fmt.Scanln(&n2)
	fmt.Println("enter the operation you want (+,*,%,/,-)")
	fmt.Scanln(&operation)
	result := 0

	switch operation {

	case "+":
		result = n1 + n2
		fmt.Println("the addition of numbers is=", result)
	case "-":
		result = n1 - n2
		fmt.Println("the subtraction of numbers is =", result)
	case "*":
		result = n1 * n2
		fmt.Println("the multiplication of numbers is =", result)
	case "/":
		result = n1 / n2
		fmt.Println("the divison of numbers is =", result)
	default:
		fmt.Println("invalid operation")

	}

}
