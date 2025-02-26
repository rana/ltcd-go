package main

import (
	"reflect"
	"testing"
)

type ListNode61b struct {
	Val  int
	Next *ListNode61b
}

// Time complexity:
// Space complexity:
func rotateRight61b(hed *ListNode61b, k int) *ListNode61b {
	return nil
}

func TestRotateRight61b(t *testing.T) {
	sliceToList := func(vals []int) *ListNode61b {
		if len(vals) == 0 {
			return nil
		}
		hed := &ListNode61b{Val: vals[0]}
		cur := hed
		for i := 1; i < len(vals); i++ {
			cur.Next = &ListNode61b{Val: vals[i]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode61b) []int {
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
			got := listToSlice(rotateRight61b(sliceToList(tc.vals), tc.k))
			if len(got) == 0 && len(tc.want) == 0 {
				return // Both slices are empty, test passes
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tc.want)
			}
		})
	}
}
