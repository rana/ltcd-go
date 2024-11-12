package main

import "testing"

type TreeNode530b struct {
	Val   int
	Left  *TreeNode530b
	Right *TreeNode530b
}

// Time complexity: O(n), n is the number of nodes in the tree.
// Space complexity: O(h), h is the height of the tree.
func getMinimumDifference530b(root *TreeNode530b) int {
	// Minimum Absolute Difference in BST
	// Given the root of a Binary Search Tree.
	// Find the minimum absolute difference between
	// any two node values.
	// Return the minimum absolute difference.
	// Key insights:
	//	* Inorder traversal of BST yields ascending order.
	//	* Just compare adjacent values.
	//	* Track the previous value for comparison to current.
	// Use a depth-first search (DFS) with recursion.
	// Use a closure for recursion and state management.

	// Initialize variables.
	minDif := 100001 // max + 1
	var prv *int

	// Create closure for inorder DFS recursion.
	var inorder func(nod *TreeNode530b)
	inorder = func(nod *TreeNode530b) {
		// Check base case.
		if nod == nil {
			return
		}

		// Inorder traversal: left, current, right

		// Process left subtree.
		inorder(nod.Left)

		// Process current node.
		if prv != nil {
			// Calculate difference from previous.
			// Know current value greater than previous.
			dif := nod.Val - *prv
			// Check minimum difference.
			if dif < minDif {
				minDif = dif
			}
		}
		prv = &nod.Val

		// Process right subtree.
		inorder(nod.Right)
	}

	// Process root node.
	inorder(root)

	return minDif
}

func TestGetMinimumDifference530b(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode530b
		want int
	}{
		{
			name: "Example 1: [4,2,6,1,3]",
			root: &TreeNode530b{
				Val: 4,
				Left: &TreeNode530b{
					Val:   2,
					Left:  &TreeNode530b{Val: 1},
					Right: &TreeNode530b{Val: 3},
				},
				Right: &TreeNode530b{Val: 6},
			},
			want: 1,
		},
		{
			name: "Example 2: [1,0,48,null,null,12,49]",
			root: &TreeNode530b{
				Val:  1,
				Left: &TreeNode530b{Val: 0},
				Right: &TreeNode530b{
					Val:   48,
					Left:  &TreeNode530b{Val: 12},
					Right: &TreeNode530b{Val: 49},
				},
			},
			want: 1,
		},
		{
			name: "Two nodes only",
			root: &TreeNode530b{
				Val:  5,
				Left: &TreeNode530b{Val: 2},
			},
			want: 3,
		},
		{
			name: "All values with max difference",
			root: &TreeNode530b{
				Val:   100000,
				Left:  &TreeNode530b{Val: 0},
				Right: &TreeNode530b{Val: 100001},
			},
			want: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := getMinimumDifference530b(tc.root)
			if got != tc.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tc.want)
			}
		})
	}
}
