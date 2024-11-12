package main

import "testing"

// TreeNode222b represents a binary tree node
type TreeNode222b struct {
	Val   int
	Left  *TreeNode222b
	Right *TreeNode222b
}

func countNodes222b(root *TreeNode222b) int {
	// Count Complete Tree Nodes
	// Given the root of a complete binary tree.
	// Count the number nodes in the tree.
	// Conditions:
	//	* Time complexity less than O(n)
	// Key insights:
	//	* Perfect binary tree has 2^h-1 nodes.
	// 	* Perfect binary tree has equal left and right height.
	// Use perfect binary tree knowledge for < O(n) time complexity.
	// Use a depth-first search DFS and recursion.

	// Check base case.
	if root == nil {
		return 0
	}

	// Measure the height of the left subtree.
	lftH := 0
	for cur := root; cur != nil; cur = cur.Left {
		lftH++
	}

	// Measure the height of the right subtree.
	rhtH := 0
	for cur := root; cur != nil; cur = cur.Right {
		rhtH++
	}

	// Check for perfect tree.
	if lftH == rhtH {
		return (1 << lftH) - 1 // 2^h-1
	}

	// Measure complete tree height with recursion.
	return 1 + countNodes222b(root.Left) + countNodes222b(root.Right)
}

func TestCountNodes222b(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode222b
		want int
	}{
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "single node",
			root: &TreeNode222b{Val: 1},
			want: 1,
		},
		{
			name: "perfect binary tree",
			root: &TreeNode222b{
				Val: 1,
				Left: &TreeNode222b{
					Val:   2,
					Left:  &TreeNode222b{Val: 4},
					Right: &TreeNode222b{Val: 5},
				},
				Right: &TreeNode222b{
					Val:   3,
					Left:  &TreeNode222b{Val: 6},
					Right: &TreeNode222b{Val: 7},
				},
			},
			want: 7,
		},
		{
			name: "complete but not perfect",
			root: &TreeNode222b{
				Val: 1,
				Left: &TreeNode222b{
					Val:   2,
					Left:  &TreeNode222b{Val: 4},
					Right: &TreeNode222b{Val: 5},
				},
				Right: &TreeNode222b{
					Val:  3,
					Left: &TreeNode222b{Val: 6},
				},
			},
			want: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := countNodes222b(tc.root); got != tc.want {
				t.Errorf("countNodes() = %v, want %v", got, tc.want)
			}
		})
	}
}
