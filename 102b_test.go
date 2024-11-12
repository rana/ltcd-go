package main

import (
	"reflect"
	"testing"
)

// TreeNode102b represents a node in a binary tree
type TreeNode102b struct {
	Val   int
	Left  *TreeNode102b
	Right *TreeNode102b
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(w), w is the maximum width of a tree level.
func levelOrder102b(root *TreeNode102b) [][]int {
	// Binary Tree Level Order Traversal
	// Given the root of a binary tree.
	// Gather level-order traversal of nodes (left to right).
	// Return array of level-order traversals.
	// Use breadth-first search (BFS), and a queue.

	// Check base case.
	if root == nil {
		return [][]int{}
	}

	// Initialize variables.
	var res [][]int
	que := []*TreeNode102b{root}

	// Traverse the tree level by level.
	for len(que) > 0 {
		// Track the level length for level processing.
		lvlLen := len(que)
		// Initialize array for level values.
		lvlVals := make([]int, lvlLen)

		// Iterate through each node in the level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove the first node from que.
			cur := que[0]
			que = que[1:]

			// Store current node value.
			lvlVals[idx] = cur.Val

			// Add left to next level que.
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			// Add right to next level que.
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		// Store level values.
		res = append(res, lvlVals)
	}

	return res
}

func TestLevelOrder102b(t *testing.T) {
	// Helper function to create tree from array
	var arrayToTree func([]any, int) *TreeNode102b
	arrayToTree = func(arr []any, idx int) *TreeNode102b {
		if idx >= len(arr) || arr[idx] == nil {
			return nil
		}

		val, ok := arr[idx].(int)
		if !ok {
			return nil
		}

		root := &TreeNode102b{Val: val}
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
			result := levelOrder102b(root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}
