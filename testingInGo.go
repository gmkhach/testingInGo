// Test files must be in the same package as the file that's being tested.
// The test file must end with "_test.go"
// Test must be in a func that have a signuature like "func Test____(t *testing.T)"
// Run test with "go test ___"

package main

import (
    "fmt"
)

func main() {
	fmt.Println("Sum of numbers between 1 and 5:", Sum(1, 2, 3, 4, 5))
	fmt.Println("Sum of numbers between -5 and -1:", Sum(-1, -2, -3, -4, -5))
}

func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v	
	}	
	return sum
}