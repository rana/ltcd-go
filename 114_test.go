package main

import "testing"

// TreeNode114 definition matches the provided structure
type TreeNode114 struct {
	Val   int
	Left  *TreeNode114
	Right *TreeNode114
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree. We use recursion.
// https://claude.ai/chat/a977069e-fc27-4ad2-97ac-a625d051776b
func flatten114(root *TreeNode114) {
	// Flatten Binary Tree to Linked List
	// Given the root of a binary tree.
	// Transform tree to a linked list.
	// Conditions:
	//	* Use TreeNode.Right as linked list pointer.
	//	* TreeNode.Left = nil.
	//	* O(1) space complexity.
	//	* Linked list is serialize as preorder.
	// Move left subtree to right.
	// Use recursion.

	// Check base case.
	if root == nil {
		return
	}

	// Store original right subtree.
	orgRht := root.Right

	// Move left to right.
	root.Right = root.Left
	root.Left = nil

	// Move to end of moved subtree.
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	// Attached original right to current right.
	cur.Right = orgRht

	// Recurse right.
	flatten114(root.Right)
}

// Testing implementation
func TestFlatten114(t *testing.T) {
	// Test case 1: Standard tree
	test1 := &TreeNode114{
		Val: 1,
		Left: &TreeNode114{
			Val:   2,
			Left:  &TreeNode114{Val: 3},
			Right: &TreeNode114{Val: 4},
		},
		Right: &TreeNode114{
			Val:   5,
			Right: &TreeNode114{Val: 6},
		},
	}

	cases := []struct {
		name     string
		root     *TreeNode114
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
			root:     &TreeNode114{Val: 0},
			expected: []int{0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			flatten114(tc.root)

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
