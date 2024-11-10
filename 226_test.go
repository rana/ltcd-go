package main

import "testing"

// TreeNode226 Definition for a binary tree node.
type TreeNode226 struct {
	Val   int
	Left  *TreeNode226
	Right *TreeNode226
}

// Time complexity: O(n), n is the number of nodes. We vist each node once.
// Space complexity: O(h), h is the height of the tree. Space is due to recursive call stack. Best case, O(logn) for a balanced tree. Worst case, O(n) for a skewed tree.
// https://claude.ai/chat/2272913d-6ba9-4a52-b6e7-2fc4383c83e6
func invertTree(root *TreeNode226) *TreeNode226 {
	// Invert Binary Tree
	// Given the root of a binary tree.
	// Invert the tree.
	// Return the root of the processed tree.
	// Inversion is swapping left and right nodes.
	// Use depth-first search (DFS) for simplicity and clarity.

	// Check base case.
	if root == nil {
		return nil
	}

	// Swap left and right nodes.
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert left and right subtrees.
	invertTree(root.Left)
	invertTree(root.Right)

	// Return root
	return root
}

// Testing Implementation
func TestInvertTree(t *testing.T) {
	// Helper function to create a tree from array representation
	createTree := func(vals []int) *TreeNode226 {
		if len(vals) == 0 {
			return nil
		}

		root := &TreeNode226{Val: vals[0]}
		queue := []*TreeNode226{root}
		i := 1

		for len(queue) > 0 && i < len(vals) {
			nod := queue[0]
			queue = queue[1:]

			if i < len(vals) {
				if vals[i] != -101 { // Using -101 to represent nil
					nod.Left = &TreeNode226{Val: vals[i]}
					queue = append(queue, nod.Left)
				}
				i++
			}

			if i < len(vals) {
				if vals[i] != -101 {
					nod.Right = &TreeNode226{Val: vals[i]}
					queue = append(queue, nod.Right)
				}
				i++
			}
		}
		return root
	}

	// Helper function to convert tree to array for comparison
	treeToArray := func(root *TreeNode226) []int {
		if root == nil {
			return []int{}
		}

		res := []int{}
		queue := []*TreeNode226{root}

		for len(queue) > 0 {
			nod := queue[0]
			queue = queue[1:]

			if nod == nil {
				res = append(res, -101)
				continue
			}

			res = append(res, nod.Val)

			if nod.Left != nil || nod.Right != nil {
				queue = append(queue, nod.Left, nod.Right)
			}
		}

		// Trim trailing nulls
		for len(res) > 0 && res[len(res)-1] == -101 {
			res = res[:len(res)-1]
		}
		return res
	}

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Example 2",
			input:    []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "Empty Tree",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single Node",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Unbalanced Tree",
			input:    []int{1, 2, -101, 3},
			expected: []int{1, -101, 2, -101, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.input)
			inverted := invertTree(root)
			result := treeToArray(inverted)

			if len(result) != len(tc.expected) {
				t.Errorf("got length %v, want %v", len(result), len(tc.expected))
				return
			}

			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("got %v, want %v", result, tc.expected)
					return
				}
			}
		})
	}
}
