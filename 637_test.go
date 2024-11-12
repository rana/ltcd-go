package main

import (
	"math"
	"testing"
)

// TreeNode637 represents a binary tree node
type TreeNode637 struct {
	Val   int
	Left  *TreeNode637
	Right *TreeNode637
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(w), w is the max width of a tree level. We enqueue one tree level at a time. Worst case, a perfect binary tree last level has n/2 nodes.
// https://claude.ai/chat/48c1298f-cb0a-48e1-86af-d50ade3303c0
func averageOfLevels637(root *TreeNode637) []float64 {
	// Average of Levels in Binary Tree
	// Given the root of a binary tree.
	// Determine the average value of nodes on each tree level.
	// Return a float array of average values.
	// Use level-order BFS with a queue.

	// Check base case.
	if root == nil {
		return nil
	}

	// Initialize variables.
	var res []float64
	que := []*TreeNode637{root}

	// Process tree level by level until que is empty.
	for len(que) > 0 {
		// Track level length for level processing.
		lvlLen := len(que)
		// Track current level sum for average calculation.
		lvlSum := 0

		// Process current level nodes.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove node from que.
			nod := que[0]
			que = que[1:]

			// Add current node value to level sum.
			lvlSum += nod.Val

			// Add left node to que.
			if nod.Left != nil {
				que = append(que, nod.Left)
			}
			// Add right node to que.
			if nod.Right != nil {
				que = append(que, nod.Right)
			}
		}

		// Calculate level average.
		lvlAvg := float64(lvlSum) / float64(lvlLen)
		// Store level average.
		res = append(res, lvlAvg)
	}

	return res
}

// Test cases
func TestAverageOfLevels637(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode637
		expected []float64
	}{
		{
			name: "Example 1: Complete tree",
			root: &TreeNode637{
				Val:  3,
				Left: &TreeNode637{Val: 9},
				Right: &TreeNode637{
					Val:   20,
					Left:  &TreeNode637{Val: 15},
					Right: &TreeNode637{Val: 7},
				},
			},
			expected: []float64{3.0, 14.5, 11.0},
		},
		{
			name:     "Example 2: Single node",
			root:     &TreeNode637{Val: 1},
			expected: []float64{1.0},
		},
		{
			name: "Example 3: Large values",
			root: &TreeNode637{
				Val:   math.MaxInt32,
				Left:  &TreeNode637{Val: math.MaxInt32},
				Right: &TreeNode637{Val: math.MaxInt32},
			},
			expected: []float64{float64(math.MaxInt32), float64(math.MaxInt32)},
		},
		{
			name: "Example 4: Negative values",
			root: &TreeNode637{
				Val:   -1,
				Left:  &TreeNode637{Val: -2},
				Right: &TreeNode637{Val: -3},
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
			got := averageOfLevels637(tc.root)
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
