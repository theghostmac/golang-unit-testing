package adder

import "fmt"

func addTwoNumbers(x, y int) int {
	return x + y
}

func main() {
	var x, y int
	solution := addTwoNumbers(x, y)
	fmt.Println(solution)
}
