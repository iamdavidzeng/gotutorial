package utils

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	if GetSum(1, 1) != 2 {
		t.Error("Expected 1 + 1 to euqal 2")
	}
}
