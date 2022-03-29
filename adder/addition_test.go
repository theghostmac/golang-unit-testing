package adder

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	test_result := addTwoNumbers(10, 15)
	if test_result != 25 {
		t.Error("Results are not correct: expected 25, got ", test_result)
	}
}
