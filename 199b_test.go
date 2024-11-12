package main

import (
	"reflect"
	"testing"
)

type TreeNode199b struct {
	Val   int
	Left  *TreeNode199b
	Right *TreeNode199b
}

func rightSideView199b(root *TreeNode199b) []int {
	// Binary Tree Right Side View
	// Given the root of a binary tree.
	// Gather values of the nodes on the right side.
	// Return the right-side values.
	// Use breadth-first search (BFS) and a queue.

	// Check base case.
	if root == nil {
		return nil
	}

	// Initialize variables.
	var res []int
	que := []*TreeNode199b{root}

	// Traverse the tree level by level.
	for len(que) > 0 {
		// Track the level length for level processing.
		lvlLen := len(que)

		// Iterate through the current level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove the first node from the que.
			cur := que[0]
			que = que[1:]

			// Check for the end node of the level.
			if idx == lvlLen-1 {
				// Store right-side node value.
				res = append(res, cur.Val)
			}

			// Process next level from left to right.
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}
	}

	return res
}

func TestRightSideView199b(t *testing.T) {
	cases := []struct {
		name     string
		root     *TreeNode199b
		expected []int
	}{
		{
			name: "Example 1: Complete tree with right node",
			root: &TreeNode199b{
				Val: 1,
				Left: &TreeNode199b{
					Val:   2,
					Right: &TreeNode199b{Val: 5},
				},
				Right: &TreeNode199b{
					Val:   3,
					Right: &TreeNode199b{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "Example 2: Right-only tree",
			root: &TreeNode199b{
				Val:   1,
				Right: &TreeNode199b{Val: 3},
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
			root: &TreeNode199b{
				Val: 1,
				Left: &TreeNode199b{
					Val:  2,
					Left: &TreeNode199b{Val: 4},
				},
				Right: &TreeNode199b{Val: 3},
			},
			expected: []int{1, 3, 4},
		},
		{
			name:     "Single node",
			root:     &TreeNode199b{Val: 1},
			expected: []int{1},
		},
		{
			name: "Full binary tree",
			root: &TreeNode199b{
				Val: 1,
				Left: &TreeNode199b{
					Val:   2,
					Left:  &TreeNode199b{Val: 4},
					Right: &TreeNode199b{Val: 5},
				},
				Right: &TreeNode199b{
					Val:   3,
					Left:  &TreeNode199b{Val: 6},
					Right: &TreeNode199b{Val: 7},
				},
			},
			expected: []int{1, 3, 7},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := rightSideView199b(tc.root)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("rightSideView() = %v, want %v", got, tc.expected)
			}
		})
	}
}
