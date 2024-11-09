package main

import "testing"

// ListNode21 definition for singly-linked list
type ListNode21 struct {
	Val  int
	Next *ListNode21
}

// Time complexity: O(n + m), n and m are the lengths of the lists. We process each node once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/4d8be773-f034-4545-9530-421137a65f51
func mergeTwoLists(lst1 *ListNode21, lst2 *ListNode21) *ListNode21 {
	// Merge Two Sorted Lists
	// Given the heads of two sort-ascending linked lists.
	// Merge the two into one sort-ascending linked list.
	// Return the head of the merged linked list.
	// Use a dummy node to simplify edge cases.

	// Initialize variables.
	var dmy ListNode21
	cur := &dmy

	// Merge while both lists have nodes.
	for lst1 != nil && lst2 != nil {
		if lst1.Val <= lst2.Val {
			cur.Next = lst1
			lst1 = lst1.Next
		} else {
			cur.Next = lst2
			lst2 = lst2.Next
		}
		cur = cur.Next
	}

	// Attached remaining from wither list.
	if lst1 != nil {
		cur.Next = lst1
	}
	if lst2 != nil {
		cur.Next = lst2
	}

	return dmy.Next
}

// Helper function to create linked list from slice
func createList21(nums []int) *ListNode21 {
	if len(nums) == 0 {
		return nil
	}

	var dmy ListNode21
	cur := &dmy

	for _, num := range nums {
		cur.Next = &ListNode21{Val: num}
		cur = cur.Next
	}

	return dmy.Next
}

// Helper function to convert linked list to slice
func listToSlice21(hed *ListNode21) []int {
	var res []int

	for cur := hed; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}

	return res
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		lst1 []int
		lst2 []int
		want []int
	}{
		{
			name: "Example 1: Regular merge",
			lst1: []int{1, 2, 4},
			lst2: []int{1, 3, 4},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "Example 2: Both empty lists",
			lst1: []int{},
			lst2: []int{},
			want: []int{},
		},
		{
			name: "Example 3: One empty list",
			lst1: []int{},
			lst2: []int{0},
			want: []int{0},
		},
		{
			name: "First list longer",
			lst1: []int{1, 2, 3, 4, 5},
			lst2: []int{1, 2},
			want: []int{1, 1, 2, 2, 3, 4, 5},
		},
		{
			name: "Second list longer",
			lst1: []int{1, 2},
			lst2: []int{1, 2, 3, 4, 5},
			want: []int{1, 1, 2, 2, 3, 4, 5},
		},
		{
			name: "Negative numbers",
			lst1: []int{-3, -1, 0},
			lst2: []int{-2, 0, 2},
			want: []int{-3, -2, -1, 0, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst1 := createList21(tt.lst1)
			lst2 := createList21(tt.lst2)
			got := listToSlice21(mergeTwoLists(lst1, lst2))

			if len(got) != len(tt.want) {
				t.Errorf("got len %v, want len %v", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
