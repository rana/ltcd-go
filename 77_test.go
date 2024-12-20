package main

import (
	"reflect"
	"sort"
	"testing"
)

// Time complexity: O(n choose k), to create all combinations.
// Space complexity: O(k), for recursive stack depth.
// https://claude.ai/chat/43d50773-b075-4a30-8328-b1c44d5febdb
func combine77(n int, k int) [][]int {
	// Combinations
	// Given two integers n and k.
	// Create k combinations from numbers in the range [1,n].
	// Return the combinations.
	// Use backtracking with early pruning and recursion.
	// Use a closure for recursion and state capture.

	// Initialize variables.
	var res [][]int
	cur := make([]int, 0, k)

	// Create a closure for recursive backtracking.
	var bckCmb func(fstIdx int)
	bckCmb = func(fstIdx int) {
		// Check base case.
		// Check whether whole combination of length k created.
		if len(cur) == k {
			// Copy combination as cur is shared.
			cpy := make([]int, k)
			copy(cpy, cur)
			res = append(res, cpy)
			return
		}

		// Check early pruning.
		// Check whether remaining length is less than needed length.
		rmnLen := n - fstIdx
		nedLen := k - len(cur)
		if rmnLen < nedLen {
			return
		}

		// Recurse from first index.
		for idx := fstIdx; idx < n; idx++ {
			// Add to combination. Number is 1-based.
			cur = append(cur, idx+1)
			// Recurse next index.
			bckCmb(idx + 1)
			// Backtrack.
			cur = cur[:len(cur)-1]
		}
	}

	// Recurse from index 0.
	bckCmb(0)

	return res
}

// TestCombine77 validates the combine function with various test cases
func TestCombine77(t *testing.T) {
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
			got := combine77(tt.n, tt.k)

			// Sort both got and want for consistent comparison
			sortCombinations77(got)
			sortCombinations77(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

// sortCombinations77 sorts each combination and the slice of combinations
// for consistent comparison in tests
func sortCombinations77(combs [][]int) {
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
