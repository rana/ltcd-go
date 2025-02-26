package main

import "testing"

// ListNode21b definition for singly-linked list
type ListNode21b struct {
	Val  int
	Next *ListNode21b
}

// Time complexity:
// Space complexity:
func mergeTwoLists21b(l1 *ListNode21b, l2 *ListNode21b) *ListNode21b {
	return nil
}

func TestMergeTwoLists21b(t *testing.T) {
	createList21b := func(nums []int) *ListNode21b {
		if len(nums) == 0 {
			return nil
		}
		var dmy ListNode21b
		cur := &dmy
		for _, num := range nums {
			cur.Next = &ListNode21b{Val: num}
			cur = cur.Next
		}
		return dmy.Next
	}
	listToSlice21b := func(hed *ListNode21b) []int {
		var res []int
		for cur := hed; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}

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
			lst1 := createList21b(tt.lst1)
			lst2 := createList21b(tt.lst2)
			got := listToSlice21b(mergeTwoLists21b(lst1, lst2))

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
