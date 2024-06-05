package pubadder_test

import (
	"testing"

	"github.com/aliyev-vladislav/learning-go/core/pubadder"
)

func TestAddNumbers(t *testing.T) {
	result := pubadder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)

	}
}
