package main

import "testing"

// TreeNode222 represents a binary tree node
type TreeNode222 struct {
	Val   int
	Left  *TreeNode222
	Right *TreeNode222
}

// Time complexity: O(logn * logn), n is the number of nodes in the tree. logn recursive calls. logn work in each call.
// Space complexity: O(logn), logn call stack spaced used in recursive calls.
// https://claude.ai/chat/b08d721f-192c-4344-a342-19e6fdf67114
func countNodes222(root *TreeNode222) int {
	// Count Complete Tree Nodes
	// Given the root of a complete binary tree.
	// Determine the number of nodes in the tree.
	// Return the number of nodes in the tree.
	// Conditions:
	//	* Time complexity less than O(n)
	// Complete binary tree:
	//	* All nodes filled except possibly last level
	//	* Last level nodes fill from left.
	// Perfect binary tree:
	//	* Left height = Right height
	//	* Can calculate node count with formula 2^h-1.
	// Look for perfect trees and use formula.
	// Use recursion for non-perfect trees.

	// Check base case.
	if root == nil {
		return 0
	}

	// Calculate left subtree height.
	lftH := 0
	for cur := root; cur != nil; cur = cur.Left {
		lftH++
	}

	// Calculate right subtree height.
	rhtH := 0
	for cur := root; cur != nil; cur = cur.Right {
		rhtH++
	}

	// Check for perfect tree.
	if lftH == rhtH {
		return (1 << lftH) - 1 // 2^h-1
	}

	// Calculate recursively for non-perfect tree.
	return 1 + countNodes222(root.Left) + countNodes222(root.Right)
}

// Test cases
func TestCountNodes222(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode222
		want int
	}{
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "single node",
			root: &TreeNode222{Val: 1},
			want: 1,
		},
		{
			name: "perfect binary tree",
			root: &TreeNode222{
				Val: 1,
				Left: &TreeNode222{
					Val:   2,
					Left:  &TreeNode222{Val: 4},
					Right: &TreeNode222{Val: 5},
				},
				Right: &TreeNode222{
					Val:   3,
					Left:  &TreeNode222{Val: 6},
					Right: &TreeNode222{Val: 7},
				},
			},
			want: 7,
		},
		{
			name: "complete but not perfect",
			root: &TreeNode222{
				Val: 1,
				Left: &TreeNode222{
					Val:   2,
					Left:  &TreeNode222{Val: 4},
					Right: &TreeNode222{Val: 5},
				},
				Right: &TreeNode222{
					Val:  3,
					Left: &TreeNode222{Val: 6},
				},
			},
			want: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := countNodes222(tc.root); got != tc.want {
				t.Errorf("countNodes() = %v, want %v", got, tc.want)
			}
		})
	}
}
