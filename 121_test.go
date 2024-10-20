package main

import (
	"math"
	"testing"
)

// Time complexity: O(n), n is the length of the array prcs. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67147424-275c-8002-a970-61aaec808a22
func maxProfit1(prcs []int) int {
	// Best Time to Buy and Sell Stock
	// Given an integer array prcs.
	// Buy once and sell once.
	// Determine the maximum profit.
	// Return the maximum profit.
	// Use a local optimization "greedy" algorithm.

	// Initialize variable.
	min_prc := math.MaxInt32
	max_pft := 0

	// Iterate through each element of array prcs.
	for _, prc := range prcs {
		// Check for minium price.
		if prc < min_prc {
			min_prc = prc
		} else {
			// Calculate profit.
			pft := prc - min_prc
			// Check for maximum profit.
			if pft > max_pft {
				max_pft = pft
			}
		}
	}

	return max_pft
}

// TestMaxProfitValidCases tests maxProfit with valid input cases.
func TestMaxProfitValidCases(t *testing.T) {
	// Test case 1: Typical case.
	prices1 := []int{7, 1, 5, 3, 6, 4}
	expected1 := 5
	result1 := maxProfit1(prices1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: No profit possible.
	prices2 := []int{7, 6, 4, 3, 1}
	expected2 := 0
	result2 := maxProfit1(prices2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Profit with increasing prices.
	prices3 := []int{1, 2, 3, 4, 5}
	expected3 := 4
	result3 := maxProfit1(prices3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}
}

// TestMaxProfitEdgeCases tests maxProfit with edge input cases.
func TestMaxProfitEdgeCases(t *testing.T) {
	// Test case 4: Single price.
	prices4 := []int{5}
	expected4 := 0
	result4 := maxProfit1(prices4)
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Prices with zeros.
	prices5 := []int{0, 0, 0, 0}
	expected5 := 0
	result5 := maxProfit1(prices5)
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}
}
