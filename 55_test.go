package main

import "testing"

// Time complexity: O(n), n is the length of array jmps. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67147d95-3944-8002-8b7a-aa0a63347efe
func canJump(jmps []int) bool {
	return false
}

// Unit tests
func TestCanJumpTrue(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	if !canJump(nums) {
		t.Errorf("Expected true, got false")
	}
}

func TestCanJumpFalse(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	if canJump(nums) {
		t.Errorf("Expected false, got true")
	}
}

func TestCanJumpSingleElement(t *testing.T) {
	nums := []int{0}
	if !canJump(nums) {
		t.Errorf("Expected true, got false")
	}
}

func TestCanJumpAllZeros(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	if canJump(nums) {
		t.Errorf("Expected false, got true")
	}
}

func TestCanJumpMaxValue(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	if !canJump(nums) {
		t.Errorf("Expected true, got false")
	}
}

func TestCanJumpLargeInput(t *testing.T) {
	nums := make([]int, 10000)
	for idx := 0; idx < 10000; idx++ {
		nums[idx] = 1
	}
	if !canJump(nums) {
		t.Errorf("Expected true, got false")
	}
}
