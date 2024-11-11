package main

import "testing"

// TreeNode105b represents a node in a binary tree
type TreeNode105b struct {
	Val   int
	Left  *TreeNode105b
	Right *TreeNode105b
}

func buildTree105b(preorder, inorder []int) *TreeNode105b {
	// Construct Binary Tree from
	// Preorder and Inorder Traversal
	// Given two integer arrays preorder and inorder.
	// Arrays are serialized binary trees.
	// Build a binary tree from the two arrays.
	// Return the root of the binary tree.
	// Use root from preorder.
	// Use left and right from inorder.
	// Use a helper closure `build`.
	// Call build recursively.
	// Use a value-index map of inorder.

	// Create a value-index of inorder for left-right finding.
	inMap := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		inMap[val] = idx
	}

	// Initialize preIdx
	preIdx := 0

	// Create closure `build` for recurive building.
	// lft, rht are indexes for inorder.
	var build func(lft, rht int) *TreeNode105b
	build = func(lft, rht int) *TreeNode105b {
		// Check out-of-bounds condition.
		if lft > rht {
			return nil
		}

		// Create current node with root value.
		val := preorder[preIdx]
		nod := &TreeNode105b{Val: val}
		preIdx++

		// Get inorder mid index.
		mid := inMap[val]

		// Recursively build subtrees.
		// Preorder: set left then right.
		nod.Left = build(lft, mid-1)
		nod.Right = build(mid+1, rht)

		return nod
	}

	return build(0, len(inorder)-1)
}

func TestBuildTree105b(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode105b
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &TreeNode105b{
				Val:  3,
				Left: &TreeNode105b{Val: 9},
				Right: &TreeNode105b{
					Val:   20,
					Left:  &TreeNode105b{Val: 15},
					Right: &TreeNode105b{Val: 7},
				},
			},
		},
		{
			name:     "Single Node",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     &TreeNode105b{Val: -1},
		},
		{
			name:     "Left Skewed Tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			want: &TreeNode105b{
				Val: 1,
				Left: &TreeNode105b{
					Val:  2,
					Left: &TreeNode105b{Val: 3},
				},
			},
		},
		{
			name:     "Right Skewed Tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want: &TreeNode105b{
				Val: 1,
				Right: &TreeNode105b{
					Val:   2,
					Right: &TreeNode105b{Val: 3},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := buildTree105b(tc.preorder, tc.inorder)
			if !isSameTree105b(got, tc.want) {
				t.Errorf("buildTree() = %v, want %v", got, tc.want)
			}
		})
	}
}

// Helper function to compare two binary trees
func isSameTree105b(p, q *TreeNode105b) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree105b(p.Left, q.Left) &&
		isSameTree105b(p.Right, q.Right)
}
