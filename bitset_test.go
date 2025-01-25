package main

import (
	"testing"
)

func TestNewBitset(t *testing.T) {
	// Test creating a Bitset
	rows := 90
	lengthPerRow := 10000000
	bitset := NewBitset(rows, lengthPerRow)

	// Check dimensions
	if len(bitset.numbersBitset) != rows {
		t.Errorf("Expected rows: %d, got: %d", rows, len(bitset.numbersBitset))
	}

	for i, row := range bitset.numbersBitset {
		expectedLength := (lengthPerRow + 63) / 64
		if len(row) != expectedLength {
			t.Errorf("Row %d: Expected length: %d, got: %d", i, expectedLength, len(row))
		}
	}
}

func TestBisetCreated(t *testing.T) {
	// Test creating a Bitset
	filename := "test.txt"
	infoFile := ReadFileInformation(filename)
	bitset := makeBitsetAndPopulate(infoFile)

	for i, item := range infoFile.Items {
		for _, j := range item {
			indexToSet := int(i / 64)
			offset := uint(i % 64)
			// Check if the bit is set
			if !bitset.IsSet(j, indexToSet, (1<<offset)) {
				t.Errorf("The bit for number %d in item %d is not set", j, i)
			}
		}
	}
}
