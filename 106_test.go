package main

import "testing"

// TreeNode106 Definition for a binary tree node.
type TreeNode106 struct {
	Val   int
	Left  *TreeNode106
	Right *TreeNode106
}

// Time complexity: O(n), n is the number of nodes.
// Space complexity: O(n), space for inorder map, and O(h) recursive call stack.
// https://claude.ai/chat/beb52abe-ae4b-4710-876f-40392079ee62
func buildTree106(inorder, postorder []int) *TreeNode106 {
	// Construct Binary Tree from
	// Inorder and Postorder Traversal
	// Given two integer arrays inorder and postorder.
	// Each array is a serialization of a binary tree.
	// Build a binary tree from the two arrays.
	// Use postorder to identify the root node.
	// Use the inorder to identify the left and right nodes.
	// Use a value-index map of inorder.
	// Use a closure to support recursively building the tree.
	// Postorder: build right node prior to left node.

	// Create a value-index map of inorder.
	inMap := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		inMap[val] = idx
	}

	// Initialize postorder index
	pstIdx := len(postorder) - 1

	// Create a closure for recursive building.
	var build func(lft, rht int) *TreeNode106
	build = func(lft, rht int) *TreeNode106 {
		// Check for out of bounds.
		if lft > rht {
			return nil
		}

		// Create a node with the postorder root value.
		val := postorder[pstIdx]
		nod := &TreeNode106{Val: val}
		pstIdx--

		// Get the inorder mid index.
		mid := inMap[val]

		// Build subtrees recursively.
		// Postorder: build right subtree before left subtree.
		nod.Right = build(mid+1, rht)
		nod.Left = build(lft, mid-1)

		return nod
	}

	return build(0, len(inorder)-1)
}

// Helper function to check if two trees are identical
func isSameTree106(p *TreeNode106, q *TreeNode106) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree106(p.Left, q.Left) &&
		isSameTree106(p.Right, q.Right)
}

func TestBuildTree106(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		want      *TreeNode106
	}{
		{
			name:      "Example 1",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode106{
				Val:  3,
				Left: &TreeNode106{Val: 9},
				Right: &TreeNode106{
					Val:   20,
					Left:  &TreeNode106{Val: 15},
					Right: &TreeNode106{Val: 7},
				},
			},
		},
		{
			name:      "Single node",
			inorder:   []int{-1},
			postorder: []int{-1},
			want:      &TreeNode106{Val: -1},
		},
		{
			name:      "Left-skewed tree",
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
			want: &TreeNode106{
				Val: 1,
				Left: &TreeNode106{
					Val:  2,
					Left: &TreeNode106{Val: 3},
				},
			},
		},
		{
			name:      "Right-skewed tree",
			inorder:   []int{1, 2, 3},
			postorder: []int{3, 2, 1},
			want: &TreeNode106{
				Val: 1,
				Right: &TreeNode106{
					Val:   2,
					Right: &TreeNode106{Val: 3},
				},
			},
		},
		{
			name:      "Empty tree",
			inorder:   []int{},
			postorder: []int{},
			want:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := buildTree106(tc.inorder, tc.postorder)
			if !isSameTree106(got, tc.want) {
				t.Errorf("buildTree() = %v, want %v", got, tc.want)
			}
		})
	}
}
