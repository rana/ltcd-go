package main

import "testing"

// TreeNode230 represents a node in a binary search tree
type TreeNode230 struct {
	Val   int
	Left  *TreeNode230
	Right *TreeNode230
}

// Time complexity: O(h + k), h is the height of the tree. We visits nodes and stop at k.
// Space complexity: O(h), for the recursion call stack.
// https://claude.ai/chat/f10b48a8-6d10-42f5-8949-b843a50352b3
func kthSmallest230(root *TreeNode230, k int) int {
	// Kth Smallest Element in a BST
	// Given the root of a binary search tree.
	// Given an integer k.
	// Find the Kth smallest value (k is 1-indexed).
	// Return the Kth smallest value.
	// Key insights:
	//	* DFS inorder traversal of BST is ascending order.
	//	* Count visited nodes and stop at k.
	// Use depth-first search (DFS) with recursion.
	// Use a closure for recursion and state management.

	// Initialize variables.
	var res int
	var cur int

	// Create a closure for inorder traversal.
	var inorder func(nod *TreeNode230)
	inorder = func(nod *TreeNode230) {
		// Check base cases.
		if nod == nil || cur > k {
			return
		}

		// Inorder traversal: left, current, right

		// Process left subtree.
		inorder(nod.Left)

		// Process current node.
		cur++
		if cur == k {
			res = nod.Val
			return
		}

		// Process right subtree.
		inorder(nod.Right)
	}

	// Process root node.
	inorder(root)

	return res
}

// iterative solution using a stack for illustration of alternative approach
func kthSmallestIterative230(root *TreeNode230, k int) int {
	// stk stores nodes to visit
	stk := make([]*TreeNode230, 0)
	cur := root
	cnt := 0

	for len(stk) > 0 || cur != nil {
		// Traverse left subtree
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		}

		// Process current node
		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		cnt++
		if cnt == k {
			return cur.Val
		}

		// Move to right subtree
		cur = cur.Right
	}
	return -1 // Not found
}

// Test cases
func TestKthSmallest230(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode230
		k    int
		want int
	}{
		{
			name: "Example 1: k=1",
			root: &TreeNode230{
				Val: 3,
				Left: &TreeNode230{
					Val:   1,
					Right: &TreeNode230{Val: 2},
				},
				Right: &TreeNode230{Val: 4},
			},
			k:    1,
			want: 1,
		},
		{
			name: "Example 2: k=3",
			root: &TreeNode230{
				Val: 5,
				Left: &TreeNode230{
					Val: 3,
					Left: &TreeNode230{
						Val:  2,
						Left: &TreeNode230{Val: 1},
					},
					Right: &TreeNode230{Val: 4},
				},
				Right: &TreeNode230{Val: 6},
			},
			k:    3,
			want: 3,
		},
		{
			name: "Single node",
			root: &TreeNode230{Val: 1},
			k:    1,
			want: 1,
		},
		{
			name: "Leftmost path",
			root: &TreeNode230{
				Val: 3,
				Left: &TreeNode230{
					Val:  2,
					Left: &TreeNode230{Val: 1},
				},
			},
			k:    1,
			want: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := kthSmallest230(tc.root, tc.k)
			if got != tc.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tc.want)
			}

			// Test iterative solution as well
			got = kthSmallestIterative230(tc.root, tc.k)
			if got != tc.want {
				t.Errorf("kthSmallestIterative() = %v, want %v", got, tc.want)
			}
		})
	}
}
