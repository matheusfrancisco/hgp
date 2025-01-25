package main

import (
	"math/bits"
)

func generateCombinations(n, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	}
	if n == 0 {
		return nil
	}

	var result [][]int
	// Include the first element in the combination
	for _, combination := range generateCombinations(n-1, k-1) {
		result = append(result, append([]int{n - 1}, combination...))
	}
	// Exclude the first element from the combination
	result = append(result, generateCombinations(n-1, k)...)

	return result
}

func bitwiseAND(a, b []uint64) []uint64 {
	result := make([]uint64, len(a))
	for i := range a {
		result[i] = a[i] & b[i]
	}
	return result
}

func countSetBits(bitset []uint64) int {
	count := 0
	for _, word := range bitset {
		count += bits.OnesCount64(word)
	}
	return count
}
