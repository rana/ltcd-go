package main

import (
	"sort"
	"testing"
)

// Time complexity: O(nlogn), n is the length of array cts. We sort the array once. We traverse the array for h-index determination once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67153f25-7d64-8002-abb6-192eab61846f
func hIndex(cts []int) int {
	// H-Index
	// Given an integer array cts.
	// Determine h-index.
	// h-index condition: element value = 1-based index
	// Return the h-index.

	// Sort array descending.
	sort.Sort(sort.Reverse(sort.IntSlice(cts)))

	// Iterate through each element of cts.
	for idx, ct := range cts {
		// Check for h-index condition.
		// h-index condition: element value == 1-based index
		if ct < idx+1 {
			return idx
		}
	}

	// All elements meet h-index condition.
	return len(cts)
}

// TestHIndexCase1 tests the standard case with mixed citation counts.
func TestHIndexCase1(t *testing.T) {
	cts := []int{3, 0, 6, 1, 5}
	exp := 3
	res := hIndex(cts)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

// TestHIndexCase2 tests the case with multiple low citation counts.
func TestHIndexCase2(t *testing.T) {
	cts := []int{1, 3, 1}
	exp := 1
	res := hIndex(cts)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

// TestHIndexCase3 tests the case with zero citations.
func TestHIndexCase3(t *testing.T) {
	cts := []int{0}
	exp := 0
	res := hIndex(cts)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

// TestHIndexCase4 tests the case with a single high citation count.
func TestHIndexCase4(t *testing.T) {
	cts := []int{100}
	exp := 1
	res := hIndex(cts)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

// TestHIndexCase5 tests the case with all zero citations in a larger array.
func TestHIndexCase5(t *testing.T) {
	cts := []int{0, 0, 0}
	exp := 0
	res := hIndex(cts)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
