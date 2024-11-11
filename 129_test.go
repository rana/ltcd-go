package main

import "testing"

// TreeNode129 Definition
type TreeNode129 struct {
	Val   int
	Left  *TreeNode129
	Right *TreeNode129
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree.
// https://claude.ai/chat/18ebadc6-bceb-4a54-9751-e0678b290228
func sumNumbers129(root *TreeNode129) int {
	// Summ Root to Leaf Numbers
	// Given the root of a binary tree.
	// Each node contains a digit 0-9.
	// Each root-to-leaf path forms a multi-digit number.
	// Sum all root-to-leaf multi-digit numbers.
	// Use depth-first search (DFS).
	// Use a dfs helper function.

	return dfs129(root, 0)
}

func dfs129(nod *TreeNode129, cur int) int {
	// Check base case.
	if nod == nil {
		return 0
	}

	// Calculate currently formed number.
	cur = cur*10 + nod.Val

	// Check leaf node condition.
	if nod.Left == nil && nod.Right == nil {
		return cur
	}

	// Recursively sum left and right subtrees.
	return dfs129(nod.Left, cur) + dfs129(nod.Right, cur)
}

// Test cases
func TestSumNumbers129(t *testing.T) {
	cases := []struct {
		name     string
		root     *TreeNode129
		expected int
	}{
		{
			name: "Example 1: Two leaves",
			root: &TreeNode129{
				Val:   1,
				Left:  &TreeNode129{Val: 2},
				Right: &TreeNode129{Val: 3},
			},
			expected: 25,
		},
		{
			name: "Example 2: Three paths",
			root: &TreeNode129{
				Val: 4,
				Left: &TreeNode129{
					Val:   9,
					Left:  &TreeNode129{Val: 5},
					Right: &TreeNode129{Val: 1},
				},
				Right: &TreeNode129{Val: 0},
			},
			expected: 1026,
		},
		{
			name:     "Single node",
			root:     &TreeNode129{Val: 5},
			expected: 5,
		},
		{
			name: "All zeros",
			root: &TreeNode129{
				Val:   0,
				Left:  &TreeNode129{Val: 0},
				Right: &TreeNode129{Val: 0},
			},
			expected: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumNumbers129(tc.root)
			if got != tc.expected {
				t.Errorf("got %v; want %v", got, tc.expected)
			}
		})
	}
}
