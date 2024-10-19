package main

import "testing"

// Time complexity: O(n), n is the length of array nums. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67141a8d-0ae4-8002-a543-87e59e0782de
func majorityElement(nums []int) int {
	// Majority Element
	// Given and array nums.
	// Determine the majority element.
	// Return the majority element.
	// Use the Boyer-Moore Voting algorithm.

	// Initialize a candidate and count.
	can, cnt := nums[0], 0

	// Iterate through each element of nums.
	for _, num := range nums {
		// if cnt is zero, set candidate to current num.
		if cnt == 0 {
			can = num
		}

		// Increment or decrement count based on candidate match.
		if can == num {
			cnt++
		} else {
			cnt--
		}
	}

	// Candidate is now majority element.
	// Return majority element.
	return can
}

// TestMajorityElement_Case1 tests the function with input [3, 2, 3].
func TestMajorityElement_Case1(t *testing.T) {
	nums := []int{3, 2, 3}
	want := 3
	got := majorityElement(nums)
	if got != want {
		t.Errorf("TestMajorityElement_Case1 failed: got %d, want %d", got, want)
	}
}

// TestMajorityElement_Case2 tests the function with input [2, 2, 1, 1, 1, 2, 2].
func TestMajorityElement_Case2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	want := 2
	got := majorityElement(nums)
	if got != want {
		t.Errorf("TestMajorityElement_Case2 failed: got %d, want %d", got, want)
	}
}

// TestMajorityElement_Case3 tests the function with input [1].
func TestMajorityElement_Case3(t *testing.T) {
	nums := []int{1}
	want := 1
	got := majorityElement(nums)
	if got != want {
		t.Errorf("TestMajorityElement_Case3 failed: got %d, want %d", got, want)
	}
}

// TestMajorityElement_Case4 tests the function with input [6, 5, 5].
func TestMajorityElement_Case4(t *testing.T) {
	nums := []int{6, 5, 5}
	want := 5
	got := majorityElement(nums)
	if got != want {
		t.Errorf("TestMajorityElement_Case4 failed: got %d, want %d", got, want)
	}
}

// TestMajorityElement_Case5 tests the function with input [1, 1, 1, 2, 2].
func TestMajorityElement_Case5(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2}
	want := 1
	got := majorityElement(nums)
	if got != want {
		t.Errorf("TestMajorityElement_Case5 failed: got %d, want %d", got, want)
	}
}
