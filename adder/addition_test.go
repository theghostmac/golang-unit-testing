package adder

import "testing"

func Test_AddTwoNumbers(t *testing.T) {
	testResult := AddTwoNumbers(10, 15)
	if testResult != 25 {
		t.Error("Results are not correct: expected 25, got ", testResult)
	}
}
