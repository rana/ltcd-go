package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(m + n), m is the length of array nums1. n is the length of array nums2. We traverse each array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/6713ef14-e2c4-8002-b858-76ac90fa6a4a
func merge88(nums1 []int, m int, nums2 []int, n int) {
	// Merge Sorted Array
	// Given two integer arrays sorted ascending.
	// Given two integer array lengths.
	// Merge the two arrays into the first array.
	// Maintain the sort-ascending.
	// Use a three-pointer approach.

	// Initialize variables.
	idx1 := m - 1
	idx2 := n - 1
	idx := m + n - 1

	// Iterate through each array from the end.
	// From the end due to the sort ascending.
	for idx1 >= 0 && idx2 >= 0 {
		if nums1[idx1] > nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			nums1[idx] = nums2[idx2]
			idx2--
		}
		idx--
	}

	// Copy any reamaining elements from nums2.
	for idx2 >= 0 {
		nums1[idx] = nums2[idx2]
		idx2--
		idx--
	}
}

func TestMrg88_BothNonEmpty(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge88(nums1, 3, nums2, 3)
	exp := []int{1, 2, 2, 3, 5, 6}
	if !reflect.DeepEqual(nums1, exp) {
		t.Errorf("Expected %v, got %v", exp, nums1)
	}
}

func TestMrg88_Nums2Empty(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{}
	merge88(nums1, 1, nums2, 0)
	exp := []int{1}
	if !reflect.DeepEqual(nums1, exp) {
		t.Errorf("Expected %v, got %v", exp, nums1)
	}
}

func TestMrg88_Nums1Empty(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	merge88(nums1, 0, nums2, 1)
	exp := []int{1}
	if !reflect.DeepEqual(nums1, exp) {
		t.Errorf("Expected %v, got %v", exp, nums1)
	}
}

func TestMrg88_DuplicateElements(t *testing.T) {
	nums1 := []int{2, 2, 3, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge88(nums1, 3, nums2, 3)
	exp := []int{1, 2, 2, 2, 3, 3}
	if !reflect.DeepEqual(nums1, exp) {
		t.Errorf("Expected %v, got %v", exp, nums1)
	}
}
