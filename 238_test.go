package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of array nums. We traverse the array twice.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/671577be-d428-8002-8153-388aa27e3226
func productExceptSelf(nums []int) []int {
	// Product of Array Except Self
	// Given an integer array.
	// Compute the product of adjacent elements.
	//	- Exclude self element
	//	- No division operation
	//	- Time complexity O(1)
	// Use a two-pass approach.
	// Pass one, compute left prefixes.
	// Pass two, compute right suffixes.

	// Initialize variables.
	res := make([]int, len(nums))
	res[0] = 1

	// lft to rht: compute left prefix products and store.
	for idx := 1; idx < len(nums); idx++ {
		res[idx] = res[idx-1] * nums[idx-1]
	}

	// rht to lft: compute right suffix products and store.
	rht_prd := 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		res[idx] *= rht_prd
		rht_prd *= nums[idx]
	}

	return res
}

func TestProductExceptSelf_Case1(t *testing.T) {
	nms := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	result := productExceptSelf(nms)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Case 1 Failed: expected %v, got %v", expected, result)
	}
}

func TestProductExceptSelf_Case2(t *testing.T) {
	nms := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}
	result := productExceptSelf(nms)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Case 2 Failed: expected %v, got %v", expected, result)
	}
}

func TestProductExceptSelf_AllZeros(t *testing.T) {
	nms := []int{0, 0, 0}
	expected := []int{0, 0, 0}
	result := productExceptSelf(nms)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Case AllZeros Failed: expected %v, got %v", expected, result)
	}
}

func TestProductExceptSelf_NegativeNumbers(t *testing.T) {
	nms := []int{-1, -2, -3, -4}
	expected := []int{-24, -12, -8, -6}
	result := productExceptSelf(nms)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test Case NegativeNumbers Failed: expected %v, got %v", expected, result)
	}
}
