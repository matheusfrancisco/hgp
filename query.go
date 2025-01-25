package main

import (
	"strconv"
	"strings"
)

type Query struct {
	fields         []int
	selectedBitset [][]uint64
}

func ParseQueryInput(input string) *Query {
	numbers := strings.Split(input, " ")
	if len(numbers) != 5 {
		return nil
	}

	winnerNumbers := make([]int, 5)
	for i := range winnerNumbers {
		selectedNumber, err := strconv.Atoi(numbers[i])
		if err != nil || selectedNumber < 1 || selectedNumber > 90 {
			return nil
		}
		winnerNumbers[i] = selectedNumber
	}
	return NewQuery(winnerNumbers)
}

func NewQuery(fields []int) *Query {
	selectedBitset := make([][]uint64, 5)
	return &Query{fields: fields, selectedBitset: selectedBitset}
}

func (q *Query) SelectedBitset(bitset *Bitset) *Query {
	selectedBitset := make([][]uint64, 5)
	for i, number := range q.fields {
		selectedBitset[i] = bitset.numbersBitset[number]
	}
	q.selectedBitset = selectedBitset
	return q
}

func (q *Query) countWinnerIntersection(k int) int {
	indices := generateCombinations(len(q.selectedBitset), k)

	count := 0
	for _, combination := range indices {
		intersection := q.selectedBitset[combination[0]]
		for _, idx := range combination[1:] {
			intersection = bitwiseAND(intersection, q.selectedBitset[idx])
		}
		count += countSetBits(intersection)
	}

	return count
}

