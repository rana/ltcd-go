package main

import (
	"reflect"
	"strconv"
	"testing"
)

type Node117b struct {
	Val   int
	Left  *Node117b
	Right *Node117b
	Next  *Node117b
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(h), h is the height of the tree. Best case, O(logn) for balanced tree. Worst case, O(n) for skewed tree.
func connect117b(root *Node117b) *Node117b {
	// Populating Next Right Pointers in Each Node II
	// Given the root node of a binary tree.
	// Each node has a `Next` pointer.
	// Set each `Next` pointer to the node to the right.
	// Return the root of the processed tree.
	// Use a head node to set the start of the level below.
	// Use a tail node to set Next pointer on the level below.
	// Condition:
	//	* O(1) space complexity.

	// Check base case.
	if root == nil {
		return nil
	}

	// Iterate through each level of the tree.
	cur := root
	for cur != nil {
		// Track head and tail for level below.
		var hed, tal *Node117b

		// Iterate through each node on current level.
		for cur != nil {
			// Handle left child.
			if cur.Left != nil {
				if tal != nil {
					tal.Next = cur.Left
				} else {
					hed = cur.Left
				}
				tal = cur.Left
			}

			// Handle right child.
			if cur.Right != nil {
				if tal != nil {
					tal.Next = cur.Right
				} else {
					hed = cur.Right
				}
				tal = cur.Right
			}

			// Move to next node on current level.
			cur = cur.Next
		}

		// Move to first node on next level.
		cur = hed
	}

	return root
}

func TestConnect117b(t *testing.T) {
	cases := []struct {
		name     string
		tree     *Node117b
		expected []string // Level order traversal with "#" marking level ends
	}{
		{
			name: "LeetCode Example 1",
			tree: &Node117b{
				Val: 1,
				Left: &Node117b{
					Val:   2,
					Left:  &Node117b{Val: 4},
					Right: &Node117b{Val: 5},
				},
				Right: &Node117b{
					Val:   3,
					Right: &Node117b{Val: 7},
				},
			},
			expected: []string{"1", "#", "2", "3", "#", "4", "5", "7", "#"},
		},
		{
			name:     "Empty Tree",
			tree:     nil,
			expected: nil,
		},
		{
			name:     "Single Node",
			tree:     &Node117b{Val: 1},
			expected: []string{"1", "#"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := connect117b(tc.tree)
			output := verifyAndGetLevelOrder117b(result)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, output)
			}
		})
	}
}

// Helper function to verify connections and get level order traversal
func verifyAndGetLevelOrder117b(root *Node117b) []string {
	if root == nil {
		return nil
	}

	var res []string
	// Start with the root level
	for start := root; start != nil; {
		// Process current level
		for cur := start; cur != nil; cur = cur.Next {
			res = append(res, strconv.Itoa(cur.Val))
		}
		// Add level separator
		res = append(res, "#")

		// Move to next level's start
		start = getNextLevelStart117b(start)
	}

	return res
}

func getNextLevelStart117b(node *Node117b) *Node117b {
	// Find first node with child in next level
	for cur := node; cur != nil; cur = cur.Next {
		if cur.Left != nil {
			return cur.Left
		}
		if cur.Right != nil {
			return cur.Right
		}
	}
	return nil
}
