package main

import "testing"

func TestCalculate(t *testing.T) {
	testResult, err := Calculate(6, 4, "*")
	if testResult != 24 {
		t.Error("Result not correct: expect 4, got ", testResult)
	}
	if err != nil {
		t.Error("Error not nil: got ", err)
	}
	secondTestResult, secondErr := Calculate(16, 4, "/")
	if secondTestResult != 4 {
		t.Error("Result not correct: expect 4, got ", secondTestResult)
	}
	if secondErr != nil {
		t.Error("Error not nil: got ", secondErr)
	}
	// and so forth...
}
