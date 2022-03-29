package benchmarking

import "testing"

func TestLengthOfFile(t *testing.T) {
	testResult, err := LengthOfFile("sample.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if testResult != 500 {
		t.Error("Expected 'Just Random Text', got ", testResult)
	}
}
