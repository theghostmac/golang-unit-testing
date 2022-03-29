package benchmarking

import "testing"

func TestLengthOfFile(t *testing.T) {
	testResult, err := LengthOfFile("sample.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if testResult != 500 {
		t.Error("Expected 500, got ", testResult)
	}
}

var benchResult int

func BenchmarkLengthOfFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchTestResult, err := LengthOfFile("sample.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		benchResult = benchTestResult
	}
}g
