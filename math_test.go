package main

import (
	"testing"
)

func TestGenerateCombinations(t *testing.T) {
	combinations := generateCombinations(5, 3)
	if len(combinations) != 10 {
		t.Errorf("Expected 10 combinations, got %d", len(combinations))
	}
	combinations = generateCombinations(5, 2)
	if len(combinations) != 10 {
		t.Errorf("Expected 10 combinations, got %d", len(combinations))
	}

	combinations = generateCombinations(5, 4)
  if len(combinations) != 5 {
    t.Errorf("Expected 5 combinations, got %d", len(combinations))
  }
}
