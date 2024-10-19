package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the array nums. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/671400b2-bf88-8002-9a26-2d2fa7e3acc2
func removeElement(nums []int, val int) int {
	// Remove Element
	// Given an integer array nums.
	// Given and integer value val.
	// Remove all occurrences of val from nums in-place.
	// Return the count of non-val elements.
	// Use a two-pointer technique.

	// lft is the index of the last allowable element.
	lft := 0

	// Iterate through each element in nums.
	for rht := 0; rht < len(nums); rht++ {
		// Check for current not equal to val.
		// Skip over elments equal to val.
		if nums[rht] != val {
			// Move valid element left.
			nums[lft] = nums[rht]
			lft++
		}
	}

	return lft
}

// Unit tests
func TestRemoveElement_Case1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expLen := 2
	expNums := []int{2, 2}

	cpyNums := make([]int, len(nums))
	copy(cpyNums, nums)

	gotLen := removeElement(nums, val)

	if gotLen != expLen {
		t.Errorf("removeElement(%v, %d) = %d; want %d", cpyNums, val, gotLen, expLen)
	}

	if !reflect.DeepEqual(nums[:gotLen], expNums) {
		t.Errorf("After removeElement(%v, %d), nums = %v; want %v", cpyNums, val, nums[:gotLen], expNums)
	}
}

func TestRemoveElement_Case2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expLen := 5
	expNums := []int{0, 1, 3, 0, 4}

	cpyNums := make([]int, len(nums))
	copy(cpyNums, nums)

	gotLen := removeElement(nums, val)

	if gotLen != expLen {
		t.Errorf("removeElement(%v, %d) = %d; want %d", cpyNums, val, gotLen, expLen)
	}

	if !reflect.DeepEqual(nums[:gotLen], expNums) {
		t.Errorf("After removeElement(%v, %d), nums = %v; want %v", cpyNums, val, nums[:gotLen], expNums)
	}
}

func TestRemoveElement_EmptyArray(t *testing.T) {
	nums := []int{}
	val := 0
	expLen := 0
	expNums := []int{}

	cpyNums := make([]int, len(nums))
	copy(cpyNums, nums)

	gotLen := removeElement(nums, val)

	if gotLen != expLen {
		t.Errorf("removeElement(%v, %d) = %d; want %d", cpyNums, val, gotLen, expLen)
	}

	if !reflect.DeepEqual(nums[:gotLen], expNums) {
		t.Errorf("After removeElement(%v, %d), nums = %v; want %v", cpyNums, val, nums[:gotLen], expNums)
	}
}

func TestRemoveElement_NoMatch(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	val := 5
	expLen := 4
	expNums := []int{1, 2, 3, 4}

	cpyNums := make([]int, len(nums))
	copy(cpyNums, nums)

	gotLen := removeElement(nums, val)

	if gotLen != expLen {
		t.Errorf("removeElement(%v, %d) = %d; want %d", cpyNums, val, gotLen, expLen)
	}

	if !reflect.DeepEqual(nums[:gotLen], expNums) {
		t.Errorf("After removeElement(%v, %d), nums = %v; want %v", cpyNums, val, nums[:gotLen], expNums)
	}
}

func TestRemoveElement_AllMatch(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	val := 1
	expLen := 0
	expNums := []int{}

	cpyNums := make([]int, len(nums))
	copy(cpyNums, nums)

	gotLen := removeElement(nums, val)

	if gotLen != expLen {
		t.Errorf("removeElement(%v, %d) = %d; want %d", cpyNums, val, gotLen, expLen)
	}

	if !reflect.DeepEqual(nums[:gotLen], expNums) {
		t.Errorf("After removeElement(%v, %d), nums = %v; want %v", cpyNums, val, nums[:gotLen], expNums)
	}
}
