package main

import (
	"reflect"
	"testing"
)

// https://chatgpt.com/c/671433e1-611c-8002-b721-a9ab2468fb0e
func rotate(nums []int, k int) {
	// Rotate Array
	// Given an integer array nums.
	// Given a segment length k.
	// Rotate the array to the right by k.
	// The array wraps around.
	// Use three calls to a reverse function.
	// Reverse whole array, first segment, second segment.

	// Adjust k in case larger than nums.
	n := len(nums)
	k = k % n

	// Reverse whole array.
	reverse(nums, 0, n-1)

	// Reverse first segment
	reverse(nums, 0, k-1)

	// Reverse second segment
	reverse(nums, k, n-1)
}

func reverse(nums []int, lft, rht int) {
	// Reverse from left to right.
	for lft < rht {
		nums[lft], nums[rht] = nums[rht], nums[lft]
		lft++
		rht--
	}
}

func TestRotateExample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	expected := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(nums, k)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("TestRotateExample1 failed. Expected %v, got %v", expected, nums)
	}
}

func TestRotateExample2(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	expected := []int{3, 99, -1, -100}

	rotate(nums, k)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("TestRotateExample2 failed. Expected %v, got %v", expected, nums)
	}
}

func TestRotateKZero(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 0
	expected := []int{1, 2, 3, 4, 5}

	rotate(nums, k)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("TestRotateKZero failed. Expected %v, got %v", expected, nums)
	}
}

func TestRotateKEqualLength(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 5
	expected := []int{1, 2, 3, 4, 5}

	rotate(nums, k)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("TestRotateKEqualLength failed. Expected %v, got %v", expected, nums)
	}
}

func TestRotateLargeK(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 12
	expected := []int{4, 5, 1, 2, 3}

	rotate(nums, k)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("TestRotateLargeK failed. Expected %v, got %v", expected, nums)
	}
}
