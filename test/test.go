package main

import (
	"fmt"
)

func main() {
	var (
		operator string
		number1  float64
		number2  float64
	)
	fmt.Print("Enter operator (+, -, /, *): ")
	fmt.Scan(&operator)

	number("Enter first number: ", &number1)
	number("Enter second number: ", &number2)

	result := process(operator, number1, number2)
	fmt.Print("Total: ", result)

}

func number(msg string, numValue *float64) {
	fmt.Print(msg)
	fmt.Scan(numValue)

}
func process(operator string, number1, number2 float64) float64 {
	var result float64
	fmt.Println("opperator", operator, number1, number2)
	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		fmt.Println("Unknown operator")
	}
	return result
}
