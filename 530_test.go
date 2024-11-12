package main

import "testing"

type TreeNode530 struct {
	Val   int
	Left  *TreeNode530
	Right *TreeNode530
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree.
// https://claude.ai/chat/4262a76a-06a8-4087-b0f8-3b868920b87b
func getMinimumDifference530(root *TreeNode530) int {
	// Minimum Absolute Difference in BST
	// Given the root of a Binary Search Tree.
	// Determine the absolute difference between
	// any two nodes.
	// Return the difference.
	// Key insights:
	//	* DFS inorder traversal visits with ascending order.
	//	* Minimum difference occurs between adjacent nodes.
	//	* Track previous value for comparison to current value.
	// Use DFS inorder traversal with recursion.
	// Use a closure for recursion and pointer reference.

	minDif := 100001
	var prv *int

	// Create closure for inorder traversal
	var inorder func(nod *TreeNode530)
	inorder = func(nod *TreeNode530) {
		// Check base case.
		if nod == nil {
			return
		}

		// Inorder traversal: left, current, right

		// Process left subtree.
		inorder(nod.Left)

		// Process current node.
		if prv != nil {
			// Calculate difference.
			// Know current value great than previous due to asc order.
			dif := nod.Val - *prv
			// Check for minimum difference.
			if dif < minDif {
				minDif = dif
			}
		}

		// Update previous value.
		prv = &nod.Val

		// Process right subtree.
		inorder(nod.Right)
	}

	// Process root node.
	inorder(root)

	return minDif
}

// Test cases
func TestGetMinimumDifference530(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode530
		want int
	}{
		{
			name: "Example 1: [4,2,6,1,3]",
			root: &TreeNode530{
				Val: 4,
				Left: &TreeNode530{
					Val:   2,
					Left:  &TreeNode530{Val: 1},
					Right: &TreeNode530{Val: 3},
				},
				Right: &TreeNode530{Val: 6},
			},
			want: 1,
		},
		{
			name: "Example 2: [1,0,48,null,null,12,49]",
			root: &TreeNode530{
				Val:  1,
				Left: &TreeNode530{Val: 0},
				Right: &TreeNode530{
					Val:   48,
					Left:  &TreeNode530{Val: 12},
					Right: &TreeNode530{Val: 49},
				},
			},
			want: 1,
		},
		{
			name: "Two nodes only",
			root: &TreeNode530{
				Val:  5,
				Left: &TreeNode530{Val: 2},
			},
			want: 3,
		},
		{
			name: "All values with max difference",
			root: &TreeNode530{
				Val:   100000,
				Left:  &TreeNode530{Val: 0},
				Right: &TreeNode530{Val: 100001},
			},
			want: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := getMinimumDifference530(tc.root)
			if got != tc.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tc.want)
			}
		})
	}
}
