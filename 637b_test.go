package main

import (
	"math"
	"testing"
)

// TreeNode637b represents a binary tree node
type TreeNode637b struct {
	Val   int
	Left  *TreeNode637b
	Right *TreeNode637b
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(w), w is the maximum width of a tree level.
func averageOfLevels637b(root *TreeNode637b) []float64 {
	// Average of Levels in Binary Tree
	// Given the root of a binary tree.
	// Calculate the average level value of each level.
	// Return an array of level averages.
	// Use breadth-first search (BFS) and a queue.

	// Check base case.
	if root == nil {
		return []float64{}
	}

	// Initialize variables.
	var res []float64
	que := []*TreeNode637b{root}

	// Traverse tree level by level.
	for len(que) > 0 {
		// Track level length to process level.
		lvlLen := len(que)
		lvlSum := 0

		// Iterate through each node in level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove first node in que.
			cur := que[0]
			que = que[1:]

			// Add current value to level sum.
			lvlSum += cur.Val

			// Add left to next level.
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			// Add right to next level.
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		// Calculate level average.
		lvlAvg := float64(lvlSum) / float64(lvlLen)

		// Store level average.
		res = append(res, lvlAvg)
	}

	return res
}

func TestAverageOfLevels637b(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode637b
		expected []float64
	}{
		{
			name: "Example 1: Complete tree",
			root: &TreeNode637b{
				Val:  3,
				Left: &TreeNode637b{Val: 9},
				Right: &TreeNode637b{
					Val:   20,
					Left:  &TreeNode637b{Val: 15},
					Right: &TreeNode637b{Val: 7},
				},
			},
			expected: []float64{3.0, 14.5, 11.0},
		},
		{
			name:     "Example 2: Single node",
			root:     &TreeNode637b{Val: 1},
			expected: []float64{1.0},
		},
		{
			name: "Example 3: Large values",
			root: &TreeNode637b{
				Val:   math.MaxInt32,
				Left:  &TreeNode637b{Val: math.MaxInt32},
				Right: &TreeNode637b{Val: math.MaxInt32},
			},
			expected: []float64{float64(math.MaxInt32), float64(math.MaxInt32)},
		},
		{
			name: "Example 4: Negative values",
			root: &TreeNode637b{
				Val:   -1,
				Left:  &TreeNode637b{Val: -2},
				Right: &TreeNode637b{Val: -3},
			},
			expected: []float64{-1.0, -2.5},
		},
		{
			name:     "Example 5: Empty tree",
			root:     nil,
			expected: []float64{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := averageOfLevels637b(tc.root)
			if len(got) != len(tc.expected) {
				t.Errorf("length mismatch: got %v, want %v", got, tc.expected)
				return
			}

			// Check each value within 10^-5 precision
			for i := range got {
				if math.Abs(got[i]-tc.expected[i]) > 1e-5 {
					t.Errorf("value mismatch at index %d: got %v, want %v", i, got, tc.expected)
					return
				}
			}
		})
	}
}
