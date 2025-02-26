package main

import "testing"

// ListNode92 represents a node in a linked list
type ListNode92 struct {
	Val  int
	Next *ListNode92
}

// Time complexity: O(n), n is the length of the list. We traverse the list once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/caa88a7d-5ce0-4d4a-959b-86a70927084e
func reverseBetween92(hed *ListNode92, lft, rht int) *ListNode92 {
	// Reverse Linked List II
	// Given the head of a linked list.
	// Given two integers lft and rht: lft <= rht.
	// Reverse nodes from postions left to right.
	// Return the head of the processed list.

	// Check input min edge case.
	if hed == nil || lft == rht {
		return hed
	}

	// Initialize variables.
	dmy := &ListNode92{Next: hed}
	prv := dmy

	// Move to node prior to reversal.
	for idx := 1; idx < lft; idx++ {
		prv = prv.Next
	}

	// Current is first reversal node.
	cur := prv.Next

	// Reverse nodes in the range.
	for idx := lft; idx < rht; idx++ {
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = prv.Next
		prv.Next = nxt
	}

	return dmy.Next
}

func TestReverseBetween92(t *testing.T) {
	createList := func(vals []int) *ListNode92 {
		if len(vals) == 0 {
			return nil
		}
		hed := &ListNode92{Val: vals[0]}
		cur := hed
		for i := 1; i < len(vals); i++ {
			cur.Next = &ListNode92{Val: vals[i]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode92) []int {
		var res []int
		cur := hed
		for cur != nil {
			res = append(res, cur.Val)
			cur = cur.Next
		}
		return res
	}
	sliceEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	tests := []struct {
		name string
		hed  *ListNode92
		lft  int
		rht  int
		want []int
	}{
		{
			name: "Example 1: reverse middle portion",
			hed:  createList([]int{1, 2, 3, 4, 5}),
			lft:  2,
			rht:  4,
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "Example 2: single node",
			hed:  createList([]int{5}),
			lft:  1,
			rht:  1,
			want: []int{5},
		},
		{
			name: "Test 3: reverse entire list",
			hed:  createList([]int{1, 2, 3}),
			lft:  1,
			rht:  3,
			want: []int{3, 2, 1},
		},
		{
			name: "Test 4: reverse first two nodes",
			hed:  createList([]int{1, 2, 3, 4}),
			lft:  1,
			rht:  2,
			want: []int{2, 1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween92(tt.hed, tt.lft, tt.rht)
			gotVals := listToSlice(got)
			if !sliceEqual(gotVals, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", gotVals, tt.want)
			}
		})
	}
}
