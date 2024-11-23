package main

import (
	"reflect"
	"sort"
	"testing"
)

// Time complexity: O(n^(t/m)), n is the number of candidates. t is the target value. m is min value in candidates.
// Space complexity: O(t/m), for recrusion stack.
// https://claude.ai/chat/cdc02c4a-b27a-472b-925f-8c27f569625d
func combinationSum39(cnds []int, tgt int) [][]int {
	// Combination Sum
	// Given a distinct integer array candidates.
	// Given an integer target.
	// Create combinations of candidates which sum to target.
	// Return the combinations.
	// Use backtracking, recursion, and early pruning.
	// Use a closure for recursion and state management.
	// Sort candidates to enable early pruning.

	// Sort candidates to enable early pruning.
	sort.Ints(cnds)

	// Initialize variables.
	var res [][]int
	var cur []int

	// Create closure for backtracking.
	var bckTrk func(pos, remTgt int)
	bckTrk = func(pos, remTgt int) {
		// Check base cases.
		switch {
		case remTgt < 0:
			// Early pruning.
			return
		case remTgt == 0:
			// Found combination.
			// Create copy of cur.
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		// Recurse through candidates from current position.
		for idx := pos; idx < len(cnds); idx++ {
			// Check early pruning.
			if cnds[idx] > remTgt {
				// All further candidates larger due to sorting.
				break
			}

			// Use current candidate.
			cur = append(cur, cnds[idx])
			// Recurse with current index and next remaining target.
			bckTrk(idx, remTgt-cnds[idx])
			// Backtrack current candidate.
			cur = cur[:len(cur)-1]
		}
	}

	// Recurse from index 0
	bckTrk(0, tgt)

	return res
}

func TestCombinationSum39(t *testing.T) {
	cases := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Example 3",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "Single number multiple times",
			candidates: []int{1},
			target:     3,
			expected:   [][]int{{1, 1, 1}},
		},
		{
			name:       "No solution exists",
			candidates: []int{4, 5, 6},
			target:     3,
			expected:   [][]int{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := combinationSum39(tc.candidates, tc.target)

			// Special handling for empty result cases
			if len(tc.expected) == 0 && len(got) == 0 {
				return // Both are empty, test passes
			}

			if len(got) != len(tc.expected) {
				t.Errorf("got %v combinations, want %v", len(got), len(tc.expected))
				return
			}

			// Sort both slices for comparison
			sortCombinations := func(combs [][]int) {
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

			sortCombinations(got)
			sortCombinations(tc.expected)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
