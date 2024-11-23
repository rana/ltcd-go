package main

import (
	"reflect"
	"sort"
	"testing"
)

// Time complexity: O(n choose k), for all combinations.
// Space complexity: O(k), for recursive stack.
func combine77b(n int, k int) [][]int {
	// Combinations
	// Given two integers n and k.
	// Create all combinations of k numbers from a range [1,n].
	// Return the combinations.
	// Use backtracking, recursion, and early pruning.
	// Use a closure for recursion and state management.

	// Initialize variables.
	var res [][]int
	cur := make([]int, 0, k)

	// Create a closure.
	var bckCmb func(pos int)
	bckCmb = func(pos int) {
		// Check base case.
		// Check combination length k reached.
		if len(cur) == k {
			// Make a copy of cur.
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		// Check early pruning.
		// Check remaining length less than needed length.
		remLen := n - pos
		nedLen := k - len(cur)
		if remLen < nedLen {
			return
		}

		// Recurse from current position.
		for idx := pos; idx < n; idx++ {
			// Add current number; 1-based.
			cur = append(cur, idx+1)
			// Recurse for next index.
			bckCmb(idx + 1)
			// Backtrack from current number.
			cur = cur[:len(cur)-1]
		}
	}

	// Recurse from index 0.
	bckCmb(0)

	return res
}

// TestCombine77b validates the combine function with various test cases
func TestCombine77b(t *testing.T) {
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
			got := combine77b(tt.n, tt.k)

			// Sort both got and want for consistent comparison
			sortCombinations77b(got)
			sortCombinations77b(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

// sortCombinations77b sorts each combination and the slice of combinations
// for consistent comparison in tests
func sortCombinations77b(combs [][]int) {
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
