package main

import (
	"reflect"
	"testing"
)

type ListNode61 struct {
	Val  int
	Next *ListNode61
}

// Time complexity: O(n), n is the length of the list. We traverse up to two times.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/0e44c755-b461-4f3a-ac1d-0e0de548f5ce
func rotateRight61(hed *ListNode61, k int) *ListNode61 {
	// Rotate List
	// Given the head of a linked list.
	// Rotate the list to the right by k.
	// Return the head of the processed list.

	// Handle edge cases
	if hed == nil || hed.Next == nil || k == 0 {
		return hed
	}

	// Find length and last node
	len := 1
	cur := hed
	for cur.Next != nil {
		len++
		cur = cur.Next
	}

	// Connect last node to head (make a circle)
	cur.Next = hed

	// Calculate rotation position and find new break point
	k = k % len
	stepsToBreak := len - k

	// Move to the node before the new head
	for range stepsToBreak {
		cur = cur.Next
	}

	// Break the circle and return new head
	newHed := cur.Next
	cur.Next = nil

	return newHed
}

func TestRotateRight61(t *testing.T) {
	sliceToList := func(vals []int) *ListNode61 {
		if len(vals) == 0 {
			return nil
		}
		hed := &ListNode61{Val: vals[0]}
		cur := hed
		for i := 1; i < len(vals); i++ {
			cur.Next = &ListNode61{Val: vals[i]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode61) []int {
		var res []int
		for cur := hed; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}

	tests := []struct {
		name string
		vals []int
		k    int
		want []int
	}{
		{
			name: "Example 1: Five nodes, rotate 2",
			vals: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "Example 2: Three nodes, rotate 4",
			vals: []int{0, 1, 2},
			k:    4,
			want: []int{2, 0, 1},
		},
		{
			name: "Empty list",
			vals: []int{},
			k:    1,
			want: []int{},
		},
		{
			name: "Single node",
			vals: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "Zero rotation",
			vals: []int{1, 2, 3},
			k:    0,
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := listToSlice(rotateRight61(sliceToList(tc.vals), tc.k))
			if len(got) == 0 && len(tc.want) == 0 {
				return // Both slices are empty, test passes
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tc.want)
			}
		})
	}
}
