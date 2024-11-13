package main

import "testing"

// Graph207 is a directed graph as an adjacency list (course-prerequisite).
type Graph207 [][]int

// Time complexity: O(v + e), v is the number of course. e is the number of prerequisites.
// Space complexity: O(e), graph, recursion, arrays.
// https://claude.ai/chat/60f35838-1433-41dd-ad7d-c4e917b4ff32
func canFinish207(numCrs int, deps [][]int) bool {
	// Course Schedule
	// Given an integer number of courses numCrs.
	// Given an integer pair array of pre-requisites.
	// Pair idx1 depends on idx2.
	// Determine if all course can be finished.
	// Return true if condition met.
	// Key insights:
	//	* A graph cycle detection challenge
	//	* Model a directed graph with an adjacency list.
	// Use depth-first search with two visited arrays.
	// One array for current dfs path.
	// One array for fully explored nodes.
	// Cycle exists if we revist a node in current dfs path.

	// Initialize graph.
	gph := make(Graph207, numCrs)
	for crs := range numCrs {
		gph[crs] = make([]int, 0)
	}

	// Populate graph with directed edges course-prerequisite.
	for _, dep := range deps {
		// Get course and prerequisite.
		crs, pre := dep[0], dep[1]

		// Populate graph.
		gph[crs] = append(gph[crs], pre)
	}

	// Initialize arrays.
	// dfsVst tracks state in dfs path
	// fin tracks fully explored nodes
	dfsVst := make([]bool, numCrs)
	fin := make([]bool, numCrs)

	// Iterate through unfinished courses looking for graph cycles.
	for crs := range numCrs {
		if !fin[crs] && dfsFndCyc207(crs, gph, dfsVst, fin) {
			return false
		}
	}

	return true
}

// dfsFndCyc207 performs a depth-first search for a graph cycle.
// Return true if a graph cycle is found.
func dfsFndCyc207(crs int, gph Graph207, dfsVst, fin []bool) bool {
	// Check whether there was a dfs visited node.
	if dfsVst[crs] {
		return true
	}

	// Skip fully explored courses.
	if fin[crs] {
		return false
	}

	// Mark current course as dfs visited.
	dfsVst[crs] = true

	// Check prerequisites for cycles.
	for _, pre := range gph[crs] {
		if dfsFndCyc207(pre, gph, dfsVst, fin) {
			return true
		}
	}

	// Mark course as finished.
	fin[crs] = true
	// Remove course from dfs visit.
	dfsVst[crs] = false

	return false
}

// Test cases
func TestCanFinish207(t *testing.T) {
	tests := []struct {
		name     string
		numCrs   int
		preReqs  [][]int
		expected bool
	}{
		{
			name:     "Example 1 - Simple valid path",
			numCrs:   2,
			preReqs:  [][]int{{1, 0}},
			expected: true,
		},
		{
			name:     "Example 2 - Simple cycle",
			numCrs:   2,
			preReqs:  [][]int{{1, 0}, {0, 1}},
			expected: false,
		},
		{
			name:     "No prerequisites",
			numCrs:   3,
			preReqs:  [][]int{},
			expected: true,
		},
		{
			name:     "Complex valid path",
			numCrs:   4,
			preReqs:  [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected: true,
		},
		{
			name:     "Complex cycle",
			numCrs:   4,
			preReqs:  [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}},
			expected: false,
		},
		{
			name:     "Single course",
			numCrs:   1,
			preReqs:  [][]int{},
			expected: true,
		},
		{
			name:     "Multiple prerequisites for one course",
			numCrs:   3,
			preReqs:  [][]int{{2, 0}, {2, 1}},
			expected: true,
		},
		{
			name:     "Self loop",
			numCrs:   2,
			preReqs:  [][]int{{1, 1}},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := canFinish207(tc.numCrs, tc.preReqs)
			if got != tc.expected {
				t.Errorf("canFinish(%v, %v) = %v, want %v",
					tc.numCrs, tc.preReqs, got, tc.expected)
			}
		})
	}
}
