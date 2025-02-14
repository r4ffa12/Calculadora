package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the operation (e.g., 12 + 34):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var num1, num2 int
	var operator string

	splitInput := strings.Fields(input)
	if len(splitInput) != 3 {
		fmt.Println("Invalid format. Use: number operator number")
		return
	}

	num1, err1 := strconv.Atoi(splitInput[0])
	num2, err2 := strconv.Atoi(splitInput[2])
	operator = splitInput[1]

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers")
		return
	}

	result, err := getResult(num1, num2, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
}

func getResult(num1, num2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Operator not valid")
	}
}
