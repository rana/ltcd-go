package main

import "testing"

// TreeNode105 represents a node in a binary tree
type TreeNode105 struct {
	Val   int
	Left  *TreeNode105
	Right *TreeNode105
}

// Time complexity: O(n), n is the number of nodes in the tree.
// Space complexity: O(n), we build a map of inorder.
// https://claude.ai/chat/6d219ff4-c375-4bb3-8010-d551b697d6f0
func buildTree105(preorder, inorder []int) *TreeNode105 {
	// Construct Binary Tree from
	// Preorder and Inorder Traversal
	// Given two integer array preorder and inorder.
	// The array are serialized binary trees.
	// Build a binary tree.
	// Return the root node of the tree.
	// Use the preorder to identify the root.
	// Use the inorder to identify the left an right.
	// Use a vale-index map of inorder.
	// Use a closure to support recursively building the tree.
	// Preorder: build the left node prior to the right node.

	// Create a value-index map of inorder.
	inMap := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		inMap[val] = idx
	}

	// Initialize preorder index.
	preIdx := 0

	// Create closure for recursive building.
	// lft and rht are inorder indexes.
	var build func(lft, rht int) *TreeNode105
	build = func(lft, rht int) *TreeNode105 {
		// Check for out of bounds.
		if lft > rht {
			return nil
		}

		// Create a root node with preorder value.
		val := preorder[preIdx]
		nod := &TreeNode105{Val: val}
		preIdx++

		// Get inorder mid index.
		mid := inMap[val]

		// Build subtrees recursively (left then right).
		nod.Left = build(lft, mid-1)
		nod.Right = build(mid+1, rht)

		return nod
	}

	return build(0, len(inorder)-1)
}

// Test cases
func TestBuildTree105(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode105
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &TreeNode105{
				Val:  3,
				Left: &TreeNode105{Val: 9},
				Right: &TreeNode105{
					Val:   20,
					Left:  &TreeNode105{Val: 15},
					Right: &TreeNode105{Val: 7},
				},
			},
		},
		{
			name:     "Single Node",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     &TreeNode105{Val: -1},
		},
		{
			name:     "Left Skewed Tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			want: &TreeNode105{
				Val: 1,
				Left: &TreeNode105{
					Val:  2,
					Left: &TreeNode105{Val: 3},
				},
			},
		},
		{
			name:     "Right Skewed Tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want: &TreeNode105{
				Val: 1,
				Right: &TreeNode105{
					Val:   2,
					Right: &TreeNode105{Val: 3},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := buildTree105(tc.preorder, tc.inorder)
			if !isSameTree105(got, tc.want) {
				t.Errorf("buildTree() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Helper function to compare two binary trees
func isSameTree105(p, q *TreeNode105) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree105(p.Left, q.Left) &&
		isSameTree105(p.Right, q.Right)
}
