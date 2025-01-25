package main

import "testing"

func TestFileInformation(t *testing.T) {
	result := ReadFileInformation("test.txt")
	expected := 3

	if result.NumberOfLines != expected {
		t.Errorf("CountFileLines(test.txt) = %d; want %d", result.NumberOfLines, expected)
	}

	for i, item := range result.Items {
		if len(item) != 5 {
			t.Errorf("The item %d should have 5 numbers", i)
		}
	}

  for i, item := range result.Items {
    for j, number := range item {
      if number < 1 || number > 90 {
        t.Errorf("The number %d in the item %d is invalid", j, i)
      }
    }
  }
}

