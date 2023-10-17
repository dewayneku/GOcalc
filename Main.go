package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Simple Go Calculator")

	for {
		fmt.Print("Enter an expression (e.g., 2 + 2), or type 'exit' to quit: ")

		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

func evaluateExpression(input string) (float64, error) {
	tokens := parseExpression(input)
	if len(tokens) != 3 {
		return 0, fmt.Errorf("Invalid input. Please use the format 'number operator number'.")
	}

	num1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid first number.")
	}

	num2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid second number.")
	}

	operator := tokens[1]

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Division by zero is not allowed.")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Invalid operator. Please use +, -, *, or /.")
	}
}

func parseExpression(input string) []string {
	return []string{input}
}
