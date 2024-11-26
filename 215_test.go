package main

import (
	"container/heap"
	"testing"
)

// Time complexity: O(n log k), n is the lengths of array nums. We visit each element, and build a k-sized min heap.
// Space complexity: O(k), for the k-sized min heap.
// https://claude.ai/chat/4cf5dd34-97af-4d7b-a20b-8677ff107c2c
func findKthLargest(nums []int, k int) int {
	// Kth Largest Element in an Array
	// Given an integer array nums.
	// Given an integer k.
	// Find the Kth largest element in nums.
	// Return the Kth largest element.
	// Use a min-heap, and Go's container heap interface.

	// Initialize variables.
	hep := &minHep{}

	// Initialize heap.
	heap.Init(hep)

	// Build a k-sized min-heap.
	for _, num := range nums {
		heap.Push(hep, num)
		if hep.Len() > k {
			heap.Pop(hep)
		}
	}

	// First index of heap is Kth element.
	return (*hep)[0]
}

type minHep []int

// Define Heap interface functions.
func (h minHep) Len() int           { return len(h) }
func (h minHep) Less(i, j int) bool { return h[i] < h[j] }
func (h minHep) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHep) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *minHep) Pop() interface{} {
	prv := *h
	val := prv[len(prv)-1]
	*h = prv[:len(prv)-1]
	return val
}

// Tests
func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "Single element",
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			name: "All same elements",
			nums: []int{2, 2, 2, 2},
			k:    2,
			want: 2,
		},
		{
			name: "K equals array length",
			nums: []int{5, 4, 3, 2, 1},
			k:    5,
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findKthLargest(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tc.want)
			}
		})
	}
}
