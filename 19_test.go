package main

import (
	"reflect"
	"testing"
)

type ListNode19 struct {
	Val  int
	Next *ListNode19
}

// Time complexity: O(n), n is the length of the list. We traverse the list once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/21591150-5d1e-4c66-9158-766cd58d7890
func removeNthFromEnd19(hed *ListNode19, n int) *ListNode19 {
	// Remove Nth Node From End of List
	// Given the head of a linked list.
	// Remove the Nth node from the end.
	// Return the head of the linked list.
	// Use a two-pointer approach.
	// Move the fast pointer n nodes.
	// Move the slow pointer prior to the removal nodes.

	// Check input min edge cases.
	if hed == nil {
		return hed
	}

	// Initialize variables.
	dmy := &ListNode19{Next: hed}
	slw := dmy
	fst := dmy

	// Move the fast pointer forward n nodes.
	for range n {
		// Check whether list is smaller than n.
		if fst.Next == nil {
			return hed
		}
		fst = fst.Next
	}

	// Move fast and slow pointers simultaneously.
	// Stop when fast pointer reaches the end.
	for fst.Next != nil {
		slw = slw.Next
		fst = fst.Next
	}

	// Remove the Nth node.
	// slw is just before the removal node.
	slw.Next = slw.Next.Next

	return dmy.Next
}

func TestRemoveNthFromEnd19(t *testing.T) {
	createList := func(nums []int) *ListNode19 {
		if len(nums) == 0 {
			return nil
		}

		hed := &ListNode19{Val: nums[0]}
		cur := hed
		for i := 1; i < len(nums); i++ {
			cur.Next = &ListNode19{Val: nums[i]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode19) []int {
		res := []int{}
		for cur := hed; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}

	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1: Remove 2nd from end",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Example 2: Single node",
			input:    []int{1},
			n:        1,
			expected: []int{},
		},
		{
			name:     "Example 3: Remove last node",
			input:    []int{1, 2},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "Remove first node",
			input:    []int{1, 2, 3},
			n:        3,
			expected: []int{2, 3},
		},
		{
			name:     "Empty list",
			input:    []int{},
			n:        1,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := createList(tc.input)
			result := removeNthFromEnd19(input, tc.n)
			got := listToSlice(result)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
