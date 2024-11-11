package main

import "testing"

// TreeNode114b definition matches the provided structure
type TreeNode114b struct {
	Val   int
	Left  *TreeNode114b
	Right *TreeNode114b
}

func flatten114b(root *TreeNode114b) {
	// Flatten Binary Tree to Linked List
	// Given the root of a binary tree.
	// Transform the tree into a linked list.
	// Conditions:
	//	* Use TreeNode.Right as the linked list pointer.
	//	* Linked list is in preorder traversal.
	//	* TreeNode.Left is nil
	//	* O(1) space complexity.
	// Move left pointer to right.
	// Use recursion.

	// Check base case.
	if root == nil {
		return
	}

	// Store original right.
	orgRht := root.Right

	// Move left to right.
	root.Right = root.Left
	root.Left = nil

	// Move to end of moved subtree.
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	// Attach original to current right.
	cur.Right = orgRht

	// Recursively flatten root right.
	flatten114b(root.Right)
}

func TestFlatten114b(t *testing.T) {
	// Test case 1: Standard tree
	test1 := &TreeNode114b{
		Val: 1,
		Left: &TreeNode114b{
			Val:   2,
			Left:  &TreeNode114b{Val: 3},
			Right: &TreeNode114b{Val: 4},
		},
		Right: &TreeNode114b{
			Val:   5,
			Right: &TreeNode114b{Val: 6},
		},
	}

	cases := []struct {
		name     string
		root     *TreeNode114b
		expected []int
	}{
		{
			name:     "Standard tree",
			root:     test1,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode114b{Val: 0},
			expected: []int{0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			flatten114b(tc.root)

			// Verify result
			var result []int
			cur := tc.root
			for cur != nil {
				// Verify left child is nil
				if cur.Left != nil {
					t.Error("Left child should be nil")
				}
				result = append(result, cur.Val)
				cur = cur.Right
			}

			// Compare values
			if len(result) != len(tc.expected) {
				t.Errorf("Expected length %d, got %d", len(tc.expected), len(result))
			}
			for i := 0; i < len(result) && i < len(tc.expected); i++ {
				if result[i] != tc.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tc.expected[i], result[i])
				}
			}
		})
	}
}
