package main

import "testing"

// ListNode86b represents a node in a linked list
type ListNode86b struct {
	Val  int
	Next *ListNode86b
}

// Time complexity:
// Space complexity:
func partition86b(hed *ListNode86b, x int) *ListNode86b {
	return nil
}

func TestPartition86b(t *testing.T) {
	createList := func(vals []int) *ListNode86b {
		if len(vals) == 0 {
			return nil
		}
		hed := &ListNode86b{Val: vals[0]}
		cur := hed
		for idx := 1; idx < len(vals); idx++ {
			cur.Next = &ListNode86b{Val: vals[idx]}
			cur = cur.Next
		}
		return hed
	}
	listToSlice := func(hed *ListNode86b) []int {
		var res []int
		cur := hed
		for cur != nil {
			res = append(res, cur.Val)
			cur = cur.Next
		}
		return res
	}

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
			list := createList(tt.list)
			got := partition86b(list, tt.x)
			gotSlice := listToSlice(got)

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
