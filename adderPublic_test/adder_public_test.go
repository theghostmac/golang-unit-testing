package adderPublic_test

import (
	"github.com/theghostmac/golang-unit-testing/adder"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	testResult := adder.AddTwoNumbers(10, 15)
	if testResult != 5 {
		t.Error("Result is incorrect: expected 5, got ", testResult)
	}
}
