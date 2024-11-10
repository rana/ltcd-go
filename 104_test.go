package main

import "testing"

// TreeNode104 represents a binary tree node
type TreeNode104 struct {
	Val   int
	Left  *TreeNode104
	Right *TreeNode104
}

// Time complexity: O(n), n is the number of the nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree. Best case, O(logn) for a balanced tree. Worst case, O(n) for a skewed tree.
// https://claude.ai/chat/164015ab-9eaa-4928-93fc-4027486417ae
func maxDepth(root *TreeNode104) int {
	// Maximum Depth of Binary Tree
	// Given the root of a binary tree.
	// Determine the maximum depth.
	// Return the maximum depth.
	// Use a depth-first approach (DFS) for simplicity.

	// Check base case.
	if root == nil {
		return 0
	}

	// Measure the depth of left and right nodes.
	lftDep := maxDepth(root.Left)
	rhtDep := maxDepth(root.Right)

	// Return the maximum depth of the right and left nodes
	// plus one for the current node.
	return max(lftDep, rhtDep) + 1
}

// Test cases
func TestMaxDepth(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode104
		want int
	}{
		{
			name: "Example 1: Balanced tree",
			root: &TreeNode104{
				Val:  3,
				Left: &TreeNode104{Val: 9},
				Right: &TreeNode104{
					Val:   20,
					Left:  &TreeNode104{Val: 15},
					Right: &TreeNode104{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "Example 2: Right skewed tree",
			root: &TreeNode104{
				Val:   1,
				Right: &TreeNode104{Val: 2},
			},
			want: 2,
		},
		{
			name: "Empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "Single node tree",
			root: &TreeNode104{Val: 1},
			want: 1,
		},
		{
			name: "Left skewed tree",
			root: &TreeNode104{
				Val: 1,
				Left: &TreeNode104{
					Val: 2,
					Left: &TreeNode104{
						Val: 3,
					},
				},
			},
			want: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxDepth(tc.root)
			if got != tc.want {
				t.Errorf("maxDepth() = %v, want %v", got, tc.want)
			}
		})
	}
}
