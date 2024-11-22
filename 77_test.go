package main

import (
	"reflect"
	"sort"
	"testing"
)

// https://claude.ai/chat/43d50773-b075-4a30-8328-b1c44d5febdb
func combine(n int, k int) [][]int {
	// Combinations
	// Given two integers
}

// TestCombine validates the combine function with various test cases
func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{
			name: "Example 1: n=4, k=2",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "Example 2: n=1, k=1",
			n:    1,
			k:    1,
			want: [][]int{{1}},
		},
		{
			name: "Minimum Valid: n=k",
			n:    3,
			k:    3,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "Three Choose Two",
			n:    3,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name: "Maximum Constraint Test",
			n:    5,
			k:    5,
			want: [][]int{{1, 2, 3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.n, tt.k)

			// Sort both got and want for consistent comparison
			sortCombinations(got)
			sortCombinations(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

// sortCombinations sorts each combination and the slice of combinations
// for consistent comparison in tests
func sortCombinations(combs [][]int) {
	for _, comb := range combs {
		sort.Ints(comb)
	}
	sort.Slice(combs, func(i, j int) bool {
		for k := 0; k < len(combs[i]) && k < len(combs[j]); k++ {
			if combs[i][k] != combs[j][k] {
				return combs[i][k] < combs[j][k]
			}
		}
		return len(combs[i]) < len(combs[j])
	})
}
