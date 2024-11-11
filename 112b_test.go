package main

import (
	"testing"
)

type TreeNode112b struct {
	Val   int
	Left  *TreeNode112b
	Right *TreeNode112b
}

func hasPathSum112b(root *TreeNode112b, targetSum int) bool {
	// Path Sum
	// Given the root of a binary tree.
	// Given an integer targetSum.
	// Determine whether a root-to-leaf path sum equals targetSum.
	// Return true if condition met.
	// For current node, subtract current value from targetSum.
	// Use depth-first search (DFS) for simplicity and clarity.
	// Use recursion.

	// Check base case.
	if root == nil {
		return false
	}

	// Calculate remaining sum.
	rem := targetSum - root.Val

	// Check leaf node condition.
	if root.Left == nil && root.Right == nil {
		return rem == 0
	}

	// Recursively check left and right subtrees
	return hasPathSum112b(root.Left, rem) ||
		hasPathSum112b(root.Right, rem)
}

func createTree112b(values []int, idx int) *TreeNode112b {
	if idx >= len(values) || values[idx] == -1001 { // using -1001 to represent nil
		return nil
	}

	root := &TreeNode112b{Val: values[idx]}
	root.Left = createTree112b(values, 2*idx+1)
	root.Right = createTree112b(values, 2*idx+2)
	return root
}

func TestHasPathSum112b(t *testing.T) {
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
			root := createTree112b(tc.treeVals, 0)
			got := hasPathSum112b(root, tc.targetSum)
			if got != tc.expected {
				t.Errorf("hasPathSum(%v, %d) = %v; want %v", tc.treeVals, tc.targetSum, got, tc.expected)
			}
		})
	}
}
