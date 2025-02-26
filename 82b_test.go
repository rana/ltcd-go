package main

import (
	"reflect"
	"testing"
)

// ListNode82b Definition for singly-linked list.
type ListNode82b struct {
	Val  int
	Next *ListNode82b
}

// Time complexity:
// Space complexity:
func deleteDuplicates82b(hed *ListNode82b) *ListNode82b {
	return nil
}

func TestDeleteDuplicates82b(t *testing.T) {
	listToSlice := func(hed *ListNode82b) []int {
		res := make([]int, 0) // Initialize with empty slice
		for cur := hed; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}
	createList := func(nums []int) *ListNode82b {
		if len(nums) == 0 {
			return nil
		}

		hed := &ListNode82b{Val: nums[0]}
		cur := hed
		for idx := 1; idx < len(nums); idx++ {
			cur.Next = &ListNode82b{Val: nums[idx]}
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
			got := listToSlice(deleteDuplicates82b(hed))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("input: %v, got: %v, want: %v", tc.input, got, tc.want)
			}
		})
	}
}
