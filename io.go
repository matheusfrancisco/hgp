package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FileType represents a file with a
// filename and items
type FileType struct {
	Filename      string
	Items         [][]int
	NumberOfLines int
}

func ReadFileInformation(fileName string) *FileType {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There is an error open the file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	validLines := 0
	items := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		info := strings.Fields(line)
		// validate the information should have 5 numbers
		if len(info) < 5 || len(info) > 5 {
			fmt.Fprintf(os.Stderr, "The information should have 5 numbers\n")
			lines++
			continue
		}
		invalidNumber := false
		numberSlice := make([]int, 5)
		for number := range numberSlice {
			selectedNumber, err := strconv.Atoi(info[number])
			if err != nil || selectedNumber < 1 || selectedNumber > 90 {
				lines++
				invalidNumber = true
				break
			}
			numberSlice[number] = selectedNumber
		}

		if invalidNumber {
			fmt.Fprintf(os.Stderr, "The game is invalid\n")
			lines++
			continue
		}

		items = append(items, numberSlice)
		validLines++
		lines++
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	return &FileType{
		Filename:      fileName,
		Items:         items,
		NumberOfLines: validLines,
	}
}
