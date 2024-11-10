package main

import "testing"

// ListNode86 represents a node in a linked list
type ListNode86 struct {
	Val  int
	Next *ListNode86
}

// Time complexity: O(n), n is the length of the list. We traverse the list once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/b00c4d5d-bb62-43b3-9502-ed75f80bbd9a
func partition(hed *ListNode86, x int) *ListNode86 {
	// Partition List
	// Given the head of a linked list.
	// Given an integer x.
	// Partition the list at x.
	// Conditions:
	//	* Nodes < x come before nodes >= x.
	//	* Preserve original list node order.
	// Use two intermediate lists to partition.

	// Check input min edge cases.
	if hed == nil || hed.Next == nil {
		return hed
	}

	// Initialize variables.
	smlHed := &ListNode86{}
	bigHed := &ListNode86{}
	smlCur := smlHed
	bigCur := bigHed

	// Iterate through original list and partition.
	cur := hed
	for cur != nil {
		if cur.Val < x {
			smlCur.Next = cur
			smlCur = cur
		} else {
			bigCur.Next = cur
			bigCur = cur
		}
		cur = cur.Next
	}

	// Connect two lists.
	smlCur.Next = bigHed.Next
	bigCur.Next = nil

	return smlHed.Next
}

// For testing
func createList86(vals []int) *ListNode86 {
	if len(vals) == 0 {
		return nil
	}
	hed := &ListNode86{Val: vals[0]}
	cur := hed
	for idx := 1; idx < len(vals); idx++ {
		cur.Next = &ListNode86{Val: vals[idx]}
		cur = cur.Next
	}
	return hed
}

func listToSlice86(hed *ListNode86) []int {
	var res []int
	cur := hed
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		list []int
		x    int
		want []int
	}{
		{
			name: "Example 1",
			list: []int{1, 4, 3, 2, 5, 2},
			x:    3,
			want: []int{1, 2, 2, 4, 3, 5},
		},
		{
			name: "Example 2",
			list: []int{2, 1},
			x:    2,
			want: []int{1, 2},
		},
		{
			name: "Empty List",
			list: []int{},
			x:    1,
			want: []int{},
		},
		{
			name: "Single Element",
			list: []int{1},
			x:    2,
			want: []int{1},
		},
		{
			name: "All Elements Less",
			list: []int{1, 2, 3},
			x:    4,
			want: []int{1, 2, 3},
		},
		{
			name: "All Elements Greater",
			list: []int{4, 5, 6},
			x:    4,
			want: []int{4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createList86(tt.list)
			got := partition(list, tt.x)
			gotSlice := listToSlice86(got)

			if len(gotSlice) != len(tt.want) {
				t.Errorf("partition() = %v, want %v", gotSlice, tt.want)
				return
			}

			for idx := range gotSlice {
				if gotSlice[idx] != tt.want[idx] {
					t.Errorf("partition() = %v, want %v", gotSlice, tt.want)
					return
				}
			}
		})
	}
}
