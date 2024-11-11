package main

import (
	"reflect"
	"strconv"
	"testing"
)

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

// Time complexity: O(n), n is the number of nodes in the tree.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/96221765-9122-4c67-b411-dc1f1ad63952
func connect117(root *Node117) *Node117 {
	// Populating Next Right Pointers in Each Node II
	// Given the root of a binary tree.
	// Each node has a `Next` pointer.
	// Set each node's `Next` pointer to the node on the right.
	// Conditions:
	//	* O(1) space complexity
	// Use a head and tail node to connect a below level, and navigate to that level.

	// Check base case.
	if root == nil {
		return nil
	}

	// Iterate through each level of the tree.
	cur := root
	for cur != nil {
		// Track head and tail node to connect level below.
		// Head is used to navigate to level below.
		var hed, tal *Node117

		// Iterate current level, connect level below.
		for cur != nil {
			// Handle left node.
			if cur.Left != nil {
				if tal != nil {
					tal.Next = cur.Left
				} else {
					hed = cur.Left
				}
				tal = cur.Left
			}

			// Handle right node.
			if cur.Right != nil {
				if tal != nil {
					tal.Next = cur.Right
				} else {
					hed = cur.Right
				}
				tal = cur.Right
			}

			// Move to the next node on current level.
			cur = cur.Next
		}

		// Move to first node on level below.
		cur = hed
	}

	return root
}

func TestConnect117(t *testing.T) {
	cases := []struct {
		name     string
		tree     *Node117
		expected []string // Level order traversal with "#" marking level ends
	}{
		{
			name: "LeetCode Example 1",
			tree: &Node117{
				Val: 1,
				Left: &Node117{
					Val:   2,
					Left:  &Node117{Val: 4},
					Right: &Node117{Val: 5},
				},
				Right: &Node117{
					Val:   3,
					Right: &Node117{Val: 7},
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
			tree:     &Node117{Val: 1},
			expected: []string{"1", "#"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := connect117(tc.tree)
			output := verifyAndGetLevelOrder117(result)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, output)
			}
		})
	}
}

// Helper function to verify connections and get level order traversal
func verifyAndGetLevelOrder117(root *Node117) []string {
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
		start = getNextLevelStart117(start)
	}

	return res
}

func getNextLevelStart117(node *Node117) *Node117 {
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
