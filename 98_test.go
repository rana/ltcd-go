package main

import (
	"math"
	"testing"
)

// TreeNode98 represents a binary tree node
type TreeNode98 struct {
	Val   int
	Left  *TreeNode98
	Right *TreeNode98
}

// Time complexity: O(n), n is the number of nodes.
// Space complexity: O(h), h is the height of the tree.
func isValidBST98(root *TreeNode98) bool {
	// Validate Binary Search Tree
	// Given the root of a binary search tree.
	// Determine whether the tree is valid.
	// Return true if conditions met.
	// Conditions:
	//	* Left subtree values less than current.
	//	* Right subtree values greater than current.
	//	* Subtrees are binary search trees.
	// Use depth-first search (DFS) with recursion.
	// Use a helper function validateBST with min/max bounds.

	return validateBST98(root, math.MinInt64, math.MaxInt64)
}

func validateBST98(nod *TreeNode98, min, max int64) bool {
	// Check base case.
	if nod == nil {
		// Nil is valid BST property.
		return true
	}

	// Check current node bounds.
	cur := int64(nod.Val)
	if cur <= min || cur >= max {
		return false
	}

	// Recursively validate left and right.
	return validateBST98(nod.Left, min, cur) &&
		validateBST98(nod.Right, cur, max)
}

// Test cases
func TestIsValidBST98(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode98
		want bool
	}{
		{
			name: "Example 1: Valid BST [2,1,3]",
			root: &TreeNode98{
				Val:   2,
				Left:  &TreeNode98{Val: 1},
				Right: &TreeNode98{Val: 3},
			},
			want: true,
		},
		{
			name: "Example 2: Invalid BST [5,1,4,null,null,3,6]",
			root: &TreeNode98{
				Val:  5,
				Left: &TreeNode98{Val: 1},
				Right: &TreeNode98{
					Val:   4,
					Left:  &TreeNode98{Val: 3},
					Right: &TreeNode98{Val: 6},
				},
			},
			want: false,
		},
		{
			name: "Single node",
			root: &TreeNode98{Val: 1},
			want: true,
		},
		{
			name: "Handle integer overflow",
			root: &TreeNode98{
				Val:   math.MaxInt32,
				Right: &TreeNode98{Val: math.MaxInt32},
			},
			want: false,
		},
		{
			name: "Duplicate values invalid",
			root: &TreeNode98{
				Val:  5,
				Left: &TreeNode98{Val: 5},
			},
			want: false,
		},
		{
			name: "Deep tree valid",
			root: &TreeNode98{
				Val: 5,
				Left: &TreeNode98{
					Val:   3,
					Left:  &TreeNode98{Val: 1},
					Right: &TreeNode98{Val: 4},
				},
				Right: &TreeNode98{
					Val:   7,
					Left:  &TreeNode98{Val: 6},
					Right: &TreeNode98{Val: 8},
				},
			},
			want: true,
		},
		{
			name: "Valid range but duplicate in path",
			root: &TreeNode98{
				Val: 10,
				Left: &TreeNode98{
					Val:   5,
					Right: &TreeNode98{Val: 10}, // Invalid: equal to ancestor
				},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidBST98(tc.root)
			if got != tc.want {
				t.Errorf("isValidBST() = %v, want %v", got, tc.want)
			}
		})
	}
}
