package main

import (
	"reflect"
	"testing"
)

// Graph210 is an adjacency list.
type Graph210 [][]int

// Visit states
const (
	unvisited210 = iota
	visiting210
	visited210
)

// Time complexity: O(v + e), v is the number of courses. e is the number of requisites.
// Space complexity: O(e), for visit states, and recursion call stack.
// https://claude.ai/chat/ce810e86-3de9-44df-8b45-418db858c725
func findOrder210(numCrs int, deps [][]int) []int {
	// Course Schedule II
	// Given an integer number of courses.
	// Given an array of [course, prerequisite].
	// Determine the order of course to take
	// with prerequisite taken before course.
	// Return the course in the order need to complete them.
	// Return nil if cycle present.
	// Use topological sort to determine order.
	// Use depth-first search, recursion, and post-order traversal for topological sort.
	// Use function dfsFndCyc for recursion and state management.
	// Use three visit states: unvisited, visiting, visited.

	// Check input min edge case.
	if len(deps) == 0 {
		res := make([]int, numCrs)
		for idx := range numCrs {
			res[idx] = idx
		}
		return res
	}

	// Initialize graph.
	gph := make(Graph210, numCrs)
	for idx := range numCrs {
		gph[idx] = make([]int, 0)
	}

	// Build graph: prerequisite -> course
	for _, dep := range deps {
		crs, pre := dep[0], dep[1]
		gph[pre] = append(gph[pre], crs)
	}

	// Initialize state.
	vst := make([]int, numCrs)
	res := make([]int, 0)

	// Create closure for dfs cycle detection, and topological sort.
	var dfsFndCyc func(cur int) bool
	dfsFndCyc = func(cur int) bool {
		// Check current visit state.
		switch vst[cur] {
		case visiting210:
			return true // Detected cycle
		case visited210:
			return false // Already explored
		}

		// Update visiting.
		vst[cur] = visiting210

		// Recurse through neighbors.
		for _, nxt := range gph[cur] {
			if dfsFndCyc(nxt) {
				return true // propogate cycle
			}
		}

		// Post-order traversal.
		vst[cur] = visited210
		res = append(res, cur)

		return false
	}

	// Traverse each course.
	for crs := range numCrs {
		if dfsFndCyc(crs) {
			return nil // no courses when cycle
		}
	}

	// Reverse results for final topological sort.
	for lft, rht := 0, len(res)-1; lft < rht; {
		res[lft], res[rht] = res[rht], res[lft]
		lft++
		rht--
	}

	return res
}

func TestFindOrder210(t *testing.T) {
	tests := []struct {
		name    string
		numCrs  int
		deps    [][]int
		want    []int
		wantAlt []int // Alternative valid order if applicable
	}{
		{
			name:   "Example 1",
			numCrs: 2,
			deps:   [][]int{{1, 0}},
			want:   []int{0, 1},
		},
		{
			name:    "Example 2",
			numCrs:  4,
			deps:    [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:    []int{0, 1, 2, 3},
			wantAlt: []int{0, 2, 1, 3},
		},
		{
			name:   "Example 3",
			numCrs: 1,
			deps:   [][]int{},
			want:   []int{0},
		},
		{
			name:   "Cycle detection",
			numCrs: 2,
			deps:   [][]int{{1, 0}, {0, 1}},
			want:   nil,
		},
		{
			name:   "Empty dependencies",
			numCrs: 3,
			deps:   [][]int{},
			want:   []int{0, 1, 2},
		},
		{
			name:   "Linear chain",
			numCrs: 4,
			deps:   [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:   []int{0, 1, 2, 3},
		},
		{
			name:    "Multiple prerequisites",
			numCrs:  3,
			deps:    [][]int{{2, 0}, {2, 1}},
			want:    []int{0, 1, 2},
			wantAlt: []int{1, 0, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findOrder210(tc.numCrs, tc.deps)

			// Check if result matches either wanted order
			if !reflect.DeepEqual(got, tc.want) &&
				(tc.wantAlt == nil || !reflect.DeepEqual(got, tc.wantAlt)) {
				t.Errorf("findOrder() = %v, want %v", got, tc.want)
				if tc.wantAlt != nil {
					t.Errorf("or want alternative %v", tc.wantAlt)
				}
			}
		})
	}
}
