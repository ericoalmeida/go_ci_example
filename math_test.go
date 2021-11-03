package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(10, 10)

	if total != 20 {
		t.Errorf("Result of sum is invalid. Expected: %d  Received: %d", 20, total)
	}
}
