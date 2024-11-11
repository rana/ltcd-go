package main

import (
	"reflect"
	"testing"
)

type TreeNode173 struct {
	Val   int
	Left  *TreeNode173
	Right *TreeNode173
}

// Binary Search Tree Iterator.
// Implement iterator with inorder traversal.
// Conditions:
//   - Left < Right
//   - next, hasNext O(1) average time complexity
//   - next, hasNext O(h) space complexity
//
// Use a stack.
// Initialize with left nodes of root.
// Next: Push left nodes of right subtree.
type BSTIterator struct {
	stk []*TreeNode173
}

func Constructor173(root *TreeNode173) BSTIterator {
	itr := BSTIterator{
		stk: make([]*TreeNode173, 0),
	}

	// Push all left nodes for smallest number at stack top.
	cur := root
	for cur != nil {
		itr.stk = append(itr.stk, cur)
		cur = cur.Left
	}

	return itr
}

func (itr *BSTIterator) Next() int {
	// Get and pop current node.
	nod := itr.stk[len(itr.stk)-1]
	itr.stk = itr.stk[:len(itr.stk)-1]

	// Push all left nodes of right subtree.
	cur := nod.Right
	for cur != nil {
		itr.stk = append(itr.stk, cur)
		cur = cur.Left
	}

	return nod.Val
}

func (itr *BSTIterator) HasNext() bool {
	return len(itr.stk) > 0
}

// buildTree173 creates a binary tree from an array representation
// where -1 represents null nodes
func buildTree173(nodes []int) *TreeNode173 {
	if len(nodes) == 0 {
		return nil
	}

	root := &TreeNode173{Val: nodes[0]}
	queue := []*TreeNode173{root}
	idx := 1

	for len(queue) > 0 && idx < len(nodes) {
		nod := queue[0]
		queue = queue[1:]

		// Add left child if exists
		if idx < len(nodes) && nodes[idx] != -1 {
			nod.Left = &TreeNode173{Val: nodes[idx]}
			queue = append(queue, nod.Left)
		}
		idx++

		// Add right child if exists
		if idx < len(nodes) && nodes[idx] != -1 {
			nod.Right = &TreeNode173{Val: nodes[idx]}
			queue = append(queue, nod.Right)
		}
		idx++
	}

	return root
}

// For testing purposes
func TestBSTIterator(t *testing.T) {
	tests := []struct {
		name     string
		nodes    []int
		expected []int
	}{
		{
			name:     "Example from problem",
			nodes:    []int{7, 3, 15, -1, -1, 9, 20}, // -1 represents null
			expected: []int{3, 7, 9, 15, 20},
		},
		{
			name:     "Single node",
			nodes:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Left skewed tree",
			nodes:    []int{3, 2, -1, 1},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Build tree from nodes
			root := buildTree173(tc.nodes)

			// Create iterator
			itr := Constructor173(root)

			// Get all values using iterator
			var got []int
			for itr.HasNext() {
				got = append(got, itr.Next())
			}

			// Compare results
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
