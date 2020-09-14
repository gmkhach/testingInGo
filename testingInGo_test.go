// Test files must be in the same package as the file that's being tested.
// The test file must end with "_test.go"
package main

import (
	"testing"
)

// Test must be in a func that have a signuature like "func Test____(t *testing.T)"
// Run test with "go test"
func TestSum(t *testing.T) {
	sum := Sum(42, 1985)
	if sum != 2027 {
		t.Error("Expected: 2027, Got:", sum)
	}
}
