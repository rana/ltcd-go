package main

import "testing"

// TreeNode230b represents a node in a binary search tree
type TreeNode230b struct {
	Val   int
	Left  *TreeNode230b
	Right *TreeNode230b
}

// Time complexity: O(h + k), h is the height of the tree. We stop early after k nodes.
// Space complexity: O(h), for recursion stack.
func kthSmallest230b(root *TreeNode230b, k int) int {
	// Kth Smallest Element in a BST
	// Given the root of a binary search tree.
	// Given an integer k.
	// Find the Kth smallest value (k is 1-based).
	// Return Kth smallest value.
	// Key insights:
	//	* DFS inorder traversal yields ascending order.
	//	* Count from the smallest node to k.
	//	* Stop early at k.
	// Use DFS inorder traversal with recursion.
	// Use a closure for recursion and state management.

	// Initialize variables.
	var res int
	var cur int

	// Create a closure for inorder traversal and recursion.
	var inorder func(nod *TreeNode230b)
	inorder = func(nod *TreeNode230b) {
		// Check base cases.
		if nod == nil || cur > k {
			return
		}

		// Inorder traversal: left, current right

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

func kthSmallestIterative230b(root *TreeNode230b, k int) int {
	// stk stores nodes to visit
	stk := make([]*TreeNode230b, 0)
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
func TestKthSmallest230b(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode230b
		k    int
		want int
	}{
		{
			name: "Example 1: k=1",
			root: &TreeNode230b{
				Val: 3,
				Left: &TreeNode230b{
					Val:   1,
					Right: &TreeNode230b{Val: 2},
				},
				Right: &TreeNode230b{Val: 4},
			},
			k:    1,
			want: 1,
		},
		{
			name: "Example 2: k=3",
			root: &TreeNode230b{
				Val: 5,
				Left: &TreeNode230b{
					Val: 3,
					Left: &TreeNode230b{
						Val:  2,
						Left: &TreeNode230b{Val: 1},
					},
					Right: &TreeNode230b{Val: 4},
				},
				Right: &TreeNode230b{Val: 6},
			},
			k:    3,
			want: 3,
		},
		{
			name: "Single node",
			root: &TreeNode230b{Val: 1},
			k:    1,
			want: 1,
		},
		{
			name: "Leftmost path",
			root: &TreeNode230b{
				Val: 3,
				Left: &TreeNode230b{
					Val:  2,
					Left: &TreeNode230b{Val: 1},
				},
			},
			k:    1,
			want: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := kthSmallest230b(tc.root, tc.k)
			if got != tc.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tc.want)
			}

			// Test iterative solution as well
			got = kthSmallestIterative230b(tc.root, tc.k)
			if got != tc.want {
				t.Errorf("kthSmallestIterative() = %v, want %v", got, tc.want)
			}
		})
	}
}
