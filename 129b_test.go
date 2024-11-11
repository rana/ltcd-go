package main

import "testing"

// TreeNode129b Definition
type TreeNode129b struct {
	Val   int
	Left  *TreeNode129b
	Right *TreeNode129b
}

func sumNumbers129b(root *TreeNode129b) int {
	// Sum Root to Leaf Numbers
	// Given a root node of a binary tree.
	// Each node has a digit 0-9.
	// Root-to-leaf path is a multi-digit number.
	// Sum all root-to-leaf multi-digit numbers.
	// Return the sum.
	// Use depth-first search (DFS) for simplicity and clarity.
	// Use a `dfs` helper function.

	return dfs129b(root, 0)
}

func dfs129b(root *TreeNode129b, cur int) int {

	// Check base case.
	if root == nil {
		return 0
	}

	// Calculate running number.
	cur = cur*10 + root.Val

	// Check leaf node condition.
	if root.Left == nil && root.Right == nil {
		return cur
	}

	// Recursively calcuate subtrees.
	return dfs129b(root.Left, cur) + dfs129b(root.Right, cur)
}

func TestSumNumbers129b(t *testing.T) {
	cases := []struct {
		name     string
		root     *TreeNode129b
		expected int
	}{
		{
			name: "Example 1: Two leaves",
			root: &TreeNode129b{
				Val:   1,
				Left:  &TreeNode129b{Val: 2},
				Right: &TreeNode129b{Val: 3},
			},
			expected: 25,
		},
		{
			name: "Example 2: Three paths",
			root: &TreeNode129b{
				Val: 4,
				Left: &TreeNode129b{
					Val:   9,
					Left:  &TreeNode129b{Val: 5},
					Right: &TreeNode129b{Val: 1},
				},
				Right: &TreeNode129b{Val: 0},
			},
			expected: 1026,
		},
		{
			name:     "Single node",
			root:     &TreeNode129b{Val: 5},
			expected: 5,
		},
		{
			name: "All zeros",
			root: &TreeNode129b{
				Val:   0,
				Left:  &TreeNode129b{Val: 0},
				Right: &TreeNode129b{Val: 0},
			},
			expected: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumNumbers129b(tc.root)
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}
