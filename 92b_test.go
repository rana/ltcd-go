package main

import "testing"

// ListNode92b represents a node in a linked list
type ListNode92b struct {
	Val  int
	Next *ListNode92b
}

// Time complexity:
// Space complexity:
func reverseBetween92b(hed *ListNode92b, lft, rht int) *ListNode92b {
	return nil
}

func TestReverseBetween92b(t *testing.T) {
	createList := func(vals []int) *ListNode92b {
		if len(vals) == 0 {
			return nil
		}
		hed := &ListNode92b{Val: vals[0]}
		cur := hed
		for i := 1; i < len(vals); i++ {
			cur.Next = &ListNode92b{Val: vals[i]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode92b) []int {
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
		hed  *ListNode92b
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
			got := reverseBetween92b(tt.hed, tt.lft, tt.rht)
			gotVals := listToSlice(got)
			if !sliceEqual(gotVals, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", gotVals, tt.want)
			}
		})
	}
}
