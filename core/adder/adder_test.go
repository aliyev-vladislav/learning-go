package adder

import "testing"

func Test_addNumber(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Fatal("incorrect result: expected 5, got", result)
	}
}
