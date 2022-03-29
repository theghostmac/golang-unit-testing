package main

import "testing"

func TestCalculate(t *testing.T) {
	outputVariables := []struct {
		operationName       string
		num1                int
		num2                int
		arithmeticOperation string
		expectedValue       int
		errorMessage        string
	}{
		{"addition", 4, 6, "+", 10, ""},
		{"subtraction", 6, 4, "-", 2, ""},
		{"multiplication", 5, 4, "*", 20, ""},
		{"division", 10, 5, "/", 2, ""},
		{"undefined", 5, 0, "/", 0, "division by zero is undefined"},
	}
	for _, valueOf := range outputVariables {
		t.Run(valueOf.operationName, func(t *testing.T) {
			testResult, err := Calculate(valueOf.num1, valueOf.num2, valueOf.arithmeticOperation)
			if testResult != valueOf.expectedValue {
				t.Errorf("Expected %d: got %d", valueOf.expectedValue, testResult)
			}
			var errorMessage string
			if err != nil {
				errorMessage = err.Error()
			}
			if errorMessage != valueOf.errorMessage {
				t.Errorf("Expected error message: `%s`, got `%s` ", valueOf.errorMessage, errorMessage)
			}
		})
	}
}
