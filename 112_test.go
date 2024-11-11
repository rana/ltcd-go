package main

import (
	"testing"
)

type TreeNode112 struct {
	Val   int
	Left  *TreeNode112
	Right *TreeNode112
}

// Time complexity: O(n), n is the number of nodes in the tree. We vist each node once.
// Space complexity: O(h), h is the height of the tree.
// https://claude.ai/chat/230e2ca6-1829-486a-8cc1-17de71d912f4
func hasPathSum112(root *TreeNode112, targetSum int) bool {
	// Path Sum
	// Given the root of a binary tree.
	// Given an integer targetSum.
	// Determine whether a root-to-leaf path equals targetSum.
	// Current root value subtracted from targetSum.
	// Use depth-first search (DFS).

	// Check base case.
	if root == nil {
		return false
	}

	// Calculate current remaining sum.
	rem := targetSum - root.Val

	// Check leaf node and remaining zero.
	if root.Left == nil && root.Right == nil {
		return rem == 0
	}

	// Recursively check both subtrees.
	return hasPathSum112(root.Left, rem) ||
		hasPathSum112(root.Right, rem)
}

// Helper function to create a binary tree from a slice
// nil values in the slice represent null nodes
func createTree112(values []int, idx int) *TreeNode112 {
	if idx >= len(values) || values[idx] == -1001 { // using -1001 to represent nil
		return nil
	}

	root := &TreeNode112{Val: values[idx]}
	root.Left = createTree112(values, 2*idx+1)
	root.Right = createTree112(values, 2*idx+2)
	return root
}

func TestHasPathSum112(t *testing.T) {
	cases := []struct {
		name      string
		treeVals  []int // -1001 represents nil
		targetSum int
		expected  bool
	}{
		{
			name:      "Example 1: Valid path sum exists",
			treeVals:  []int{5, 4, 8, 11, -1001, 13, 4, 7, 2, -1001, -1001, -1001, 1},
			targetSum: 22,
			expected:  true,
		},
		{
			name:      "Example 2: No valid path sum",
			treeVals:  []int{1, 2, 3},
			targetSum: 5,
			expected:  false,
		},
		{
			name:      "Example 3: Empty tree",
			treeVals:  []int{},
			targetSum: 0,
			expected:  false,
		},
		{
			name:      "Single node equals target",
			treeVals:  []int{1},
			targetSum: 1,
			expected:  true,
		},
		{
			name:      "Single node doesn't equal target",
			treeVals:  []int{1},
			targetSum: 2,
			expected:  false,
		},
		{
			name:      "Negative values in path",
			treeVals:  []int{1, -2, 3},
			targetSum: -1,
			expected:  true,
		},
		{
			name:      "All negative values",
			treeVals:  []int{-1, -2, -3},
			targetSum: -3,
			expected:  true,
		},
		{
			name:      "Long path with exact sum",
			treeVals:  []int{1, 2, 3, 4, 5, -1001, -1001, 8},
			targetSum: 15,
			expected:  true,
		},
		{
			name:      "Path must go to leaf",
			treeVals:  []int{1, 2, 3, 4},
			targetSum: 3,
			expected:  false, // Sum exists but not to leaf
		},
		{
			name:      "Zero sum path exists",
			treeVals:  []int{0, 1, 2, -1, -1001, -2},
			targetSum: 0,
			expected:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree112(tc.treeVals, 0)
			got := hasPathSum112(root, tc.targetSum)
			if got != tc.expected {
				t.Errorf("hasPathSum(%v, %d) = %v; want %v", tc.treeVals, tc.targetSum, got, tc.expected)
			}
		})
	}
}
