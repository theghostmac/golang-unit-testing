package main

import (
	"errors"
	"fmt"
)

func Calculate(num1, num2 int, arithmeticOperations string) (int, error) {
	switch arithmeticOperations {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero is undefined")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operation %s ", arithmeticOperations)
	}
}

func main() {
	calculate, err := Calculate(2, 3, "*")
	if err != nil {
		return
	}
	fmt.Println(calculate)
}
