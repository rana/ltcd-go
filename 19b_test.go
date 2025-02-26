package main

import (
	"reflect"
	"testing"
)

type ListNode19b struct {
	Val  int
	Next *ListNode19b
}

// Time complexity:
// Space complexity:
func removeNthFromEnd19b(hed *ListNode19b, n int) *ListNode19b {
	return nil
}

func TestRemoveNthFromEnd19b(t *testing.T) {
	createList := func(nums []int) *ListNode19b {
		if len(nums) == 0 {
			return nil
		}

		hed := &ListNode19b{Val: nums[0]}
		cur := hed
		for i := 1; i < len(nums); i++ {
			cur.Next = &ListNode19b{Val: nums[i]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode19b) []int {
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
			result := removeNthFromEnd19b(input, tc.n)
			got := listToSlice(result)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
