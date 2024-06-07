package tabletests_test

import (
	"testing"

	tabletests "github.com/aliyev-vladislav/learning-go/core/table_tests"
)

func TestDoMath(t *testing.T) {
	result, err := tabletests.DoMath(2, 2, "+")
	if result != 4 {
		t.Error("Expected 4, got", result)
	}
	if err != nil {
		t.Error("Expected nil error, got", err)
	}
	result2, err2 := tabletests.DoMath(2, 2, "-")
	if result2 != 0 {
		t.Error("Expected 0, got", result2)
	}
	if err2 != nil {
		t.Error("Expected nil error, got", err2)
	}
	result3, err3 := tabletests.DoMath(2, 2, "*")
	if result3 != 4 {
		t.Error("Expected 4, got", result3)
	}
	if err3 != nil {
		t.Error("Expected nil error, got", err3)
	}
	result4, err4 := tabletests.DoMath(2, 2, "/")
	if result4 != 1 {
		t.Error("Expected 1, got", result4)
	}
	if err4 != nil {
		t.Error("Expected nil error, got", err4)
	}
	result5, err5 := tabletests.DoMath(2, 0, "/")
	if result5 != 0 {
		t.Error("Expected 0, got", result5)
	}
	if err5.Error() != `division by zero` {
		t.Error("Expected error division by zero, got", err5)
	}

}

func TestDoMathTable(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"substraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_divison", 2, 0, "/", 0, `division by zero`},
		{"bad_op", 2, 2, "?", 0, `unknown operation ?`},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := tabletests.DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected: %d, got: %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message: `%s`, got: `%s`",
					d.errMsg, errMsg)
			}
		})
	}
}
