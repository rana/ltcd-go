package main

import "testing"

// Time complexity: O(n), n is the length of the array prcs. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/6714797e-7598-8002-b8ea-f625fcd45f64
func maxProfit2(prcs []int) int {
	// Best Time to Buy and Sell Stock II
	// Given and integer array prcs.
	// Buy and sell multiple times.
	// Hold one share at a time.
	// Determine the maximum profit.
	// Return the maximum profit.
	// Use a local optimization "greedy" algorithm.

	// Initialize variable.
	max_pft := 0

	// Iterate through each element of the prcs array.
	for idx := 1; idx < len(prcs); idx++ {
		// Check whether there is a daily profit.
		if prcs[idx] > prcs[idx-1] {
			// Accumulate daily profit.
			max_pft += prcs[idx] - prcs[idx-1]
		}
	}

	return max_pft
}

// TestMaxProfitAscendingPrices tests increasing prices.
func TestMaxProfitAscendingPrices(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	expected := 4
	res := maxProfit2(prices)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

// TestMaxProfitDescendingPrices tests decreasing prices.
func TestMaxProfitDescendingPrices(t *testing.T) {
	prices := []int{5, 4, 3, 2, 1}
	expected := 0
	res := maxProfit2(prices)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

// TestMaxProfitMixedPrices tests mixed price changes.
func TestMaxProfitMixedPrices(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 7
	res := maxProfit2(prices)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

// TestMaxProfitSingleDay tests a single-day price.
func TestMaxProfitSingleDay(t *testing.T) {
	prices := []int{5}
	expected := 0
	res := maxProfit2(prices)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

// TestMaxProfitEmptyPrices tests an empty price list.
func TestMaxProfitEmptyPrices(t *testing.T) {
	prices := []int{}
	expected := 0
	res := maxProfit2(prices)
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
