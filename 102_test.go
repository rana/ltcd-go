package main

import (
	"reflect"
	"testing"
)

// TreeNode102 represents a node in a binary tree
type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}

// Time complexity: O(n), n is the number of nodes in the tree.
// Space complexity: O(w), w is the maximum width of a tree level.
// https://claude.ai/chat/f8602b82-c3a5-437d-8fb0-ec03a4d72d8d
func levelOrder102(root *TreeNode102) [][]int {
	// Binary Tree Level Order Traversal
	// Given the root of a binary tree.
	// Traverse the tree in level-order left to right.
	// Gather node values for each level.
	// Return an array of level values.
	// Use level-order BFS with a queue.

	// Check base case.
	if root == nil {
		return [][]int{}
	}

	// Initialize variables.
	var res [][]int
	que := []*TreeNode102{root}

	// Iterate through the tree level by level.
	for len(que) > 0 {
		// Track level length for level processing.
		lvlLen := len(que)
		// Initialize array for level values.
		lvl := make([]int, 0, lvlLen)

		// Iterate through the current level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove first node from que.
			cur := que[0]
			que = que[1:]

			// Add current node value to level values.
			lvl = append(lvl, cur.Val)

			// Build next level from left to right.
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		// Store current level values.
		res = append(res, lvl)
	}

	return res
}

func TestLevelOrder102(t *testing.T) {
	// Helper function to create tree from array
	var arrayToTree func([]any, int) *TreeNode102
	arrayToTree = func(arr []any, idx int) *TreeNode102 {
		if idx >= len(arr) || arr[idx] == nil {
			return nil
		}

		val, ok := arr[idx].(int)
		if !ok {
			return nil
		}

		root := &TreeNode102{Val: val}
		root.Left = arrayToTree(arr, 2*idx+1)
		root.Right = arrayToTree(arr, 2*idx+2)
		return root
	}

	cases := []struct {
		name     string
		input    []any
		expected [][]int
	}{
		{
			name:     "Example 1: Standard tree",
			input:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Example 2: Single node",
			input:    []any{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: Empty tree",
			input:    []any{},
			expected: [][]int{},
		},
		{
			name:     "Complete binary tree",
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Left-skewed tree",
			input:    []any{1, 2, nil, 3, nil, nil, nil, 4},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := arrayToTree(tc.input, 0)
			result := levelOrder102(root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}
