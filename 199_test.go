package main

import (
	"reflect"
	"testing"
)

type TreeNode199 struct {
	Val   int
	Left  *TreeNode199
	Right *TreeNode199
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(w), w is the maximum width of a tree level. We store a tree level at a time. A perfect binary tree has n/2 nodes.
// https://claude.ai/chat/0dd2282e-de47-4719-bc4f-f6e77c9b767c
func rightSideView199(root *TreeNode199) []int {
	// Binary Tree Right Side View
	// Given the root of a binary tree.
	// Determine the right-side view.
	// Return an array of the node values right-side view.
	// Array is ordered top to bottom.
	// Use level-order BFS traversal with a queue.

	// Check base case.
	if root == nil {
		return nil
	}

	// Initialize variables.
	var res []int
	que := []*TreeNode199{root}

	// Process tree level by level until que is empty.
	for len(que) > 0 {
		// Track the current level length for level processing.
		lvlLen := len(que)

		// Iterate through the current level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove the first node from the que.
			nod := que[0]
			que = que[1:]

			// Check for right-most node.
			// idx == 0 due to right inserted prior to left.
			if idx == 0 {
				res = append(res, nod.Val)
			}

			// Insert right node first.
			// Allows right node to appear at front of level.
			if nod.Right != nil {
				que = append(que, nod.Right)
			}

			// Insert left node last.
			// Allows left node to appear at end of level.
			if nod.Left != nil {
				que = append(que, nod.Left)
			}
		}
	}

	return res
}

func TestRightSideView199(t *testing.T) {
	cases := []struct {
		name     string
		root     *TreeNode199
		expected []int
	}{
		{
			name: "Example 1: Complete tree with right node",
			root: &TreeNode199{
				Val: 1,
				Left: &TreeNode199{
					Val:   2,
					Right: &TreeNode199{Val: 5},
				},
				Right: &TreeNode199{
					Val:   3,
					Right: &TreeNode199{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "Example 2: Right-only tree",
			root: &TreeNode199{
				Val:   1,
				Right: &TreeNode199{Val: 3},
			},
			expected: []int{1, 3},
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: nil,
		},
		{
			name: "Left side deeper than right",
			root: &TreeNode199{
				Val: 1,
				Left: &TreeNode199{
					Val:  2,
					Left: &TreeNode199{Val: 4},
				},
				Right: &TreeNode199{Val: 3},
			},
			expected: []int{1, 3, 4},
		},
		{
			name:     "Single node",
			root:     &TreeNode199{Val: 1},
			expected: []int{1},
		},
		{
			name: "Full binary tree",
			root: &TreeNode199{
				Val: 1,
				Left: &TreeNode199{
					Val:   2,
					Left:  &TreeNode199{Val: 4},
					Right: &TreeNode199{Val: 5},
				},
				Right: &TreeNode199{
					Val:   3,
					Left:  &TreeNode199{Val: 6},
					Right: &TreeNode199{Val: 7},
				},
			},
			expected: []int{1, 3, 7},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := rightSideView199(tc.root)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("rightSideView() = %v, want %v", got, tc.expected)
			}
		})
	}
}
