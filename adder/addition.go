package adder

import "fmt"

func AddTwoNumbers(x, y int) int {
	return x + y
}

func main() {
	var x, y int
	solution := AddTwoNumbers(x, y)
	fmt.Println(solution)
}
