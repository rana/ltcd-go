package main

import (
	"reflect"
	"testing"
)

type TreeNode173b struct {
	Val   int
	Left  *TreeNode173b
	Right *TreeNode173b
}

// Binary Search Tree Iterator
// Implement.
// Conditions:
//   - next, hasNext O(1) amortized time complexity.
//   - next, hasNext O(h) space complexity.
//   - Inorder traversal.
//
// Use a stack.
// Push all left subtrees for intialization.
// Push all left subtrees or right subtree for Next.
// Smallest items on top of stack.
type BSTIterator173b struct {
	stk []*TreeNode173b
}

func Constructor173b(root *TreeNode173b) BSTIterator173b {
	itr := BSTIterator173b{
		stk: make([]*TreeNode173b, 0),
	}

	// Push all left nodes.
	cur := root
	for cur != nil {
		itr.stk = append(itr.stk, cur)
		cur = cur.Left
	}

	return itr
}

// Conditions:
//   - next, hasNext O(1) amortized time complexity.
//   - next, hasNext O(h) space complexity.
//   - Moves pointer and return the number at pointer.
func (itr *BSTIterator173b) Next() int {
	// Get node and pop.
	nod := itr.stk[len(itr.stk)-1]
	itr.stk = itr.stk[:len(itr.stk)-1]

	// Push all left nodes or right subtree.
	cur := nod.Right
	for cur != nil {
		itr.stk = append(itr.stk, cur)
		cur = cur.Left
	}

	return nod.Val
}

// Conditions:
//   - next, hasNext O(1) amortized time complexity.
//   - next, hasNext O(h) space complexity.
//   - Return true if more numbers.
func (itr *BSTIterator173b) HasNext() bool {
	return len(itr.stk) > 0
}

func buildTree173b(nodes []int) *TreeNode173b {
	if len(nodes) == 0 {
		return nil
	}

	root := &TreeNode173b{Val: nodes[0]}
	queue := []*TreeNode173b{root}
	idx := 1

	for len(queue) > 0 && idx < len(nodes) {
		nod := queue[0]
		queue = queue[1:]

		// Add left child if exists
		if idx < len(nodes) && nodes[idx] != -1 {
			nod.Left = &TreeNode173b{Val: nodes[idx]}
			queue = append(queue, nod.Left)
		}
		idx++

		// Add right child if exists
		if idx < len(nodes) && nodes[idx] != -1 {
			nod.Right = &TreeNode173b{Val: nodes[idx]}
			queue = append(queue, nod.Right)
		}
		idx++
	}

	return root
}

// For testing purposes
func TestBSTIterator173b(t *testing.T) {
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
			root := buildTree173b(tc.nodes)

			// Create iterator
			itr := Constructor173b(root)

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
