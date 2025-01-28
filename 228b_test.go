package main

import (
	"testing"
)

func summaryRanges228b(nums []int) []string {
	return nil
}

func TestSummaryRanges228b(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "Example 2",
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: []string{},
		},
		{
			name: "Single element",
			nums: []int{1},
			want: []string{"1"},
		},
		{
			name: "All continuous",
			nums: []int{1, 2, 3, 4, 5},
			want: []string{"1->5"},
		},
		{
			name: "All separate",
			nums: []int{1, 3, 5, 7, 9},
			want: []string{"1", "3", "5", "7", "9"},
		},
		{
			name: "Negative numbers",
			nums: []int{-3, -2, -1, 1, 2},
			want: []string{"-3->-1", "1->2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summaryRanges228b(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
