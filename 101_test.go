package main

import "testing"

type TreeNode101 struct {
	Val   int
	Left  *TreeNode101
	Right *TreeNode101
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree. Best case, O(logn) for a balanced tree. Worst case, O(n) for a skewed tree.
// https://claude.ai/chat/fbc760d3-c76b-4788-8442-efc9b9afdc62
func isSymmetric101(root *TreeNode101) bool {
	// Symmetric Tree
	// Given the root of a binary tree.
	// Determine whether the tree is a mirror of itself.
	// Return true if conditions met.
	// Conditions:
	//	* lft val == rht val
	//	* lft.Left == rht.Right
	//	* lft.Right == rht.Left
	// Use a depth-first search (DFS) for simplicity and clarity.
	// Use a helper function isMirror.

	// Check base case.
	if root == nil {
		return true
	}
	return isMirror101(root.Left, root.Right)
}

func isMirror101(lft, rht *TreeNode101) bool {
	// Check whther both lft and rht are nil.
	if lft == nil && rht == nil {
		return true
	}
	// Check whether one side is nil.
	if lft == nil || rht == nil {
		return false
	}
	// Check whether lft/rht values and subtrees are equal.
	return lft.Val == rht.Val &&
		isMirror101(lft.Left, rht.Right) &&
		isMirror101(lft.Right, rht.Left)
}

// Test cases
func TestIsSymmetric101(t *testing.T) {
	cases := []struct {
		name     string
		root     *TreeNode101
		expected bool
	}{
		{
			name: "symmetric tree",
			root: &TreeNode101{
				Val: 1,
				Left: &TreeNode101{
					Val:   2,
					Left:  &TreeNode101{Val: 3},
					Right: &TreeNode101{Val: 4},
				},
				Right: &TreeNode101{
					Val:   2,
					Left:  &TreeNode101{Val: 4},
					Right: &TreeNode101{Val: 3},
				},
			},
			expected: true,
		},
		{
			name: "asymmetric tree",
			root: &TreeNode101{
				Val: 1,
				Left: &TreeNode101{
					Val:   2,
					Right: &TreeNode101{Val: 3},
				},
				Right: &TreeNode101{
					Val:   2,
					Right: &TreeNode101{Val: 3},
				},
			},
			expected: false,
		},
		{
			name:     "nil root",
			root:     nil,
			expected: true,
		},
		{
			name:     "single node",
			root:     &TreeNode101{Val: 1},
			expected: true,
		},
		{
			name: "different values",
			root: &TreeNode101{
				Val:   1,
				Left:  &TreeNode101{Val: 2},
				Right: &TreeNode101{Val: 3},
			},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSymmetric101(tc.root)
			if got != tc.expected {
				t.Errorf("isSymmetric() = %v, want %v", got, tc.expected)
			}
		})
	}
}
