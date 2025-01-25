package main

import (
	"testing"
)

func TestQuery(t *testing.T) {
	query := ParseQueryInput("4 79 13 80 56")

	if query == nil {
		t.Errorf("Query is nil")
	}

	if len(query.fields) != 5 {
		t.Errorf("Expected 5 fields, got %d", len(query.fields))
	}
	expectedFields := []int{4, 79, 13, 80, 56}

	for i, field := range query.fields {
		if field != expectedFields[i] {
			t.Errorf("Expected %d, got %d", expectedFields[i], field)
		}
	}

}

func TestQuerySelectedBitset(t *testing.T) {
	query := ParseQueryInput("4 79 13 80 56")
	filename := "test.txt"
	infoFile := ReadFileInformation(filename)
	bitset := makeBitsetAndPopulate(infoFile)
	query.SelectedBitset(bitset)

	// Check if the selected bitset is correct
	expectedSelectedBitset := [][]uint64{
		{7},
		{7},
		{7},
		{7},
		{7},
	}
	for i, selected := range query.selectedBitset {
		if selected[0] != expectedSelectedBitset[i][0] {
			t.Errorf("Expected %b, got %b", expectedSelectedBitset[i], selected)
		}
	}

}
