package main

import (
	"reflect"
	"testing"
)

// ListNode definition is provided in the problem
type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n log n), O(n) for dividing each level. O(n log n) for recursive stack calls. O(n) for merging each level.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/5108dfca-fd36-44d1-b606-a5b537ab4f42
func sortList(hed *ListNode) *ListNode {
	// Given the head of a linked list.
	// Sort the list in ascending order.
	// Return the head of the list.
	// Conditions:
	//	* Time complexity: O(n log n)
	// 	* Space complexity: O(n)
	// Use merge sort with recursion.
	// Use fast/slow point technique to identify list middle without counting.

	// Check input min edge cases.
	if hed == nil || hed.Next == nil {
		return hed
	}

	// Find the middle node in the list without counting.
	// Use the fast/slow pointer technique.
	var prv *ListNode
	slw, fst := hed, hed
	for fst != nil && fst.Next != nil {
		prv = slw
		slw = slw.Next
		fst = fst.Next.Next
	}

	// Split list in half.
	prv.Next = nil

	// Recursively sort left and right halves.
	lft := sortList(hed)
	rht := sortList(slw)

	// Return merged list.
	return mrgLst(lft, rht)
}

// mrgLst merges two sort-ascending linked lists.
func mrgLst(lft, rht *ListNode) *ListNode {
	dmy := &ListNode{}
	cur := dmy

	// Merge lists in sort order.
	for lft != nil && rht != nil {
		if lft.Val <= rht.Val {
			cur.Next = lft
			lft = lft.Next
		} else {
			cur.Next = rht
			rht = rht.Next
		}
		cur = cur.Next
	}

	// Attach remaining nodes.
	if lft != nil {
		cur.Next = lft
	}
	if rht != nil {
		cur.Next = rht
	}

	return dmy.Next
}

func TestSortList(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1 from description",
			input:    []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2 from description",
			input:    []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Reverse sorted",
			input:    []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Duplicate values",
			input:    []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Create input linked list
			head := createList(tc.input)

			// Sort the list
			result := sortList(head)

			// Convert result back to slice for comparison
			got := listToSlice(result)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}

// Helper function to create linked list from slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

// Helper function to convert linked list to slice
func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
