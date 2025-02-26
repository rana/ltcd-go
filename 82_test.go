package main

import (
	"reflect"
	"testing"
)

// ListNode82 Definition for singly-linked list.
type ListNode82 struct {
	Val  int
	Next *ListNode82
}

// Time complexity: O(n), n is the length of the list. We traverse the list once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/54741edd-e21f-42e3-8e18-a7b5c444be10
func deleteDuplicates82(hed *ListNode82) *ListNode82 {
	// Remove Duplicates from Sorted List II
	// Given the head of a sorted list.
	// Remove all nodes with duplicate values.
	// Return the head of the linked list.

	// Check input min edge cases.
	if hed == nil || hed.Next == nil {
		return hed
	}

	// Initialize variables.
	dmy := &ListNode82{Next: hed}
	prv := dmy
	cur := hed

	// Iterate through list.
	for cur != nil {
		// Check for duplicates.
		hasDuplicate := false
		for cur.Next != nil && cur.Val == cur.Next.Val {
			hasDuplicate = true
			cur = cur.Next
		}

		if hasDuplicate {
			// Skip last duplicate.
			prv.Next = cur.Next
		} else {
			// Progress the prv pointer.
			prv = cur
		}
		cur = cur.Next
	}

	return dmy.Next
}

func TestDeleteDuplicates82(t *testing.T) {
	listToSlice := func(hed *ListNode82) []int {
		res := make([]int, 0) // Initialize with empty slice
		for cur := hed; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}
	createList := func(nums []int) *ListNode82 {
		if len(nums) == 0 {
			return nil
		}

		hed := &ListNode82{Val: nums[0]}
		cur := hed
		for idx := 1; idx < len(nums); idx++ {
			cur.Next = &ListNode82{Val: nums[idx]}
			cur = cur.Next
		}
		return hed
	}

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 3, 3, 4, 4, 5},
			want:  []int{1, 2, 5},
		},
		{
			name:  "Example 2",
			input: []int{1, 1, 1, 2, 3},
			want:  []int{2, 3},
		},
		{
			name:  "Empty list",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Single node",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "All duplicates",
			input: []int{1, 1, 1},
			want:  []int{},
		},
		{
			name:  "No duplicates",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hed := createList(tc.input)
			got := listToSlice(deleteDuplicates82(hed))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("input: %v, got: %v, want: %v", tc.input, got, tc.want)
			}
		})
	}
}
