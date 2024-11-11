package main

import "testing"

type TreeNode101b struct {
	Val   int
	Left  *TreeNode101b
	Right *TreeNode101b
}

func isSymmetric101b(root *TreeNode101b) bool {
	// Symmetric Tree
	// Given the root of a binary tree.
	// Determine whether if is a mirror of itself.
	// Symmetric around center.
	// Return true if condition met.
	// Condition:
	// lft.val == rht.val
	// Use depth-first search (DFS) for simplicity and clarity.
	// Use a helper function.

	// Check base case.
	if root == nil {
		return true
	}

	return isMirror101b(root.Left, root.Right)
}

func isMirror101b(lft, rht *TreeNode101b) bool {
	// Check lft and rht both nil.
	if lft == nil && rht == nil {
		return true
	}
	// Check one side nil.
	if lft == nil || rht == nil {
		return false
	}

	// Check value equal.
	// Check symmetric.
	return lft.Val == rht.Val &&
		isMirror101b(lft.Left, rht.Right) &&
		isMirror101b(lft.Right, rht.Left)
}

func TestIsSymmetric(t *testing.T) {
	cases := []struct {
		name     string
		root     *TreeNode101b
		expected bool
	}{
		{
			name: "symmetric tree",
			root: &TreeNode101b{
				Val: 1,
				Left: &TreeNode101b{
					Val:   2,
					Left:  &TreeNode101b{Val: 3},
					Right: &TreeNode101b{Val: 4},
				},
				Right: &TreeNode101b{
					Val:   2,
					Left:  &TreeNode101b{Val: 4},
					Right: &TreeNode101b{Val: 3},
				},
			},
			expected: true,
		},
		{
			name: "asymmetric tree",
			root: &TreeNode101b{
				Val: 1,
				Left: &TreeNode101b{
					Val:   2,
					Right: &TreeNode101b{Val: 3},
				},
				Right: &TreeNode101b{
					Val:   2,
					Right: &TreeNode101b{Val: 3},
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
			root:     &TreeNode101b{Val: 1},
			expected: true,
		},
		{
			name: "different values",
			root: &TreeNode101b{
				Val:   1,
				Left:  &TreeNode101b{Val: 2},
				Right: &TreeNode101b{Val: 3},
			},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSymmetric101b(tc.root)
			if got != tc.expected {
				t.Errorf("isSymmetric() = %v, want %v", got, tc.expected)
			}
		})
	}
}
