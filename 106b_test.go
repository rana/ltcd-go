package main

import "testing"

// TreeNode106b Definition for a binary tree node.
type TreeNode106b struct {
	Val   int
	Left  *TreeNode106b
	Right *TreeNode106b
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree.
func buildTree106b(inorder, postorder []int) *TreeNode106b {
	// Construct Binary Tree from Inorder and PostOrder Traversal
	// Given two integer arrays postorder and inorder.
	// Each array is a serialized binary tree traversal.
	// Build a binary tree from the two arrays.
	// Return the root of the binary tree.
	// Use root from postorder.
	// Use left and right from inorder.
	// Use a value-index map for inorder.
	// User a helper closure `build` to support recursion.
	// Recursively build.
	// Postorder: Build right node then left node.

	// Create value-index map of inorder for left-right finding.
	inMap := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		inMap[val] = idx
	}

	// Initialize postorder index.
	pstIdx := len(postorder) - 1

	// Create `build` closure for recursive building.
	// lft and rht are indexes for inorder.
	var build func(lft, rht int) *TreeNode106b
	build = func(lft, rht int) *TreeNode106b {
		// Check out of bounds.
		if lft > rht {
			return nil
		}

		// Create a node with postorder root value.
		val := postorder[pstIdx]
		nod := &TreeNode106b{Val: val}
		pstIdx--

		// Get inorder mid index.
		mid := inMap[val]

		// Recursively build subtrees.
		// Postorder: build right then left.
		nod.Right = build(mid+1, rht)
		nod.Left = build(lft, mid-1)

		return nod
	}

	return build(0, len(inorder)-1)
}

func isSameTree106b(p *TreeNode106b, q *TreeNode106b) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree106b(p.Left, q.Left) &&
		isSameTree106b(p.Right, q.Right)
}

func TestBuildTree106b(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		want      *TreeNode106b
	}{
		{
			name:      "Example 1",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode106b{
				Val:  3,
				Left: &TreeNode106b{Val: 9},
				Right: &TreeNode106b{
					Val:   20,
					Left:  &TreeNode106b{Val: 15},
					Right: &TreeNode106b{Val: 7},
				},
			},
		},
		{
			name:      "Single node",
			inorder:   []int{-1},
			postorder: []int{-1},
			want:      &TreeNode106b{Val: -1},
		},
		{
			name:      "Left-skewed tree",
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
			want: &TreeNode106b{
				Val: 1,
				Left: &TreeNode106b{
					Val:  2,
					Left: &TreeNode106b{Val: 3},
				},
			},
		},
		{
			name:      "Right-skewed tree",
			inorder:   []int{1, 2, 3},
			postorder: []int{3, 2, 1},
			want: &TreeNode106b{
				Val: 1,
				Right: &TreeNode106b{
					Val:   2,
					Right: &TreeNode106b{Val: 3},
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
			got := buildTree106b(tc.inorder, tc.postorder)
			if !isSameTree106b(got, tc.want) {
				t.Errorf("buildTree() = %v, want %v", got, tc.want)
			}
		})
	}
}
