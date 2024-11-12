package main

import "testing"

// TreeNode236 represents a node in a binary tree
type TreeNode236 struct {
	Val   int
	Left  *TreeNode236
	Right *TreeNode236
}

// Time complexity: O(n), n is the number of nodes in the tree.
// Space complexity: O(h), h is the height of the tree.
// https://claude.ai/chat/abf141f5-9856-4347-b2fb-0866a046aac8
func lowestCommonAncestor236(root, p, q *TreeNode236) *TreeNode236 {
	// Lowest Common Ancestor of Binary Tree
	// Given the root of a binary tree.
	// Given two nodes p and q in the tree.
	// Find the lowest common ancestor.
	// Lowest common ancestor is closest to leaf.
	// Conditions:
	//	* Self node can be lowest common ancestor.
	// Use recursion and post-order traversal.
	// Post-order traversal enables knowing about
	// descedant nodes prior to ancestor nodes.

	// Check base cases.
	if root == nil || root == p || root == q {
		return root
	}

	// Post-order traversal:
	lft := lowestCommonAncestor236(root.Left, p, q)
	rht := lowestCommonAncestor236(root.Right, p, q)

	// Check success condition.
	if lft != nil && rht != nil {
		return root
	}

	// Return non-nil node.
	if lft != nil {
		return lft
	}
	return rht
}

// Helper function to create a tree node
func createNode236(val int) *TreeNode236 {
	return &TreeNode236{Val: val}
}

func TestLowestCommonAncestor236(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // Array representation of tree with -1 for null
		pVal     int   // Value of node p
		qVal     int   // Value of node q
		expected int   // Expected LCA value
	}{
		{
			name:     "Example 1: LCA is root",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     5,
			qVal:     1,
			expected: 3,
		},
		{
			name:     "Example 2: Node is ancestor of itself",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     5,
			qVal:     4,
			expected: 5,
		},
		{
			name:     "Example 3: Simple tree",
			treeVals: []int{1, 2},
			pVal:     1,
			qVal:     2,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Build tree from values
			if len(tc.treeVals) == 0 {
				t.Fatal("Empty tree values")
			}

			// Create nodes for the tree
			nodes := make([]*TreeNode236, len(tc.treeVals))
			for i, val := range tc.treeVals {
				if val != -1 {
					nodes[i] = createNode236(val)
				}
			}

			// Connect nodes to form the tree
			for i := 0; i < len(nodes); i++ {
				if nodes[i] != nil {
					lftIdx := 2*i + 1
					rhtIdx := 2*i + 2
					if lftIdx < len(nodes) {
						nodes[i].Left = nodes[lftIdx]
					}
					if rhtIdx < len(nodes) {
						nodes[i].Right = nodes[rhtIdx]
					}
				}
			}

			// Find nodes p and q
			var p, q *TreeNode236
			for _, node := range nodes {
				if node != nil {
					if node.Val == tc.pVal {
						p = node
					}
					if node.Val == tc.qVal {
						q = node
					}
				}
			}

			if p == nil || q == nil {
				t.Fatal("Could not find p or q nodes")
			}

			// Find LCA
			res := lowestCommonAncestor236(nodes[0], p, q)

			// Verify result
			if res.Val != tc.expected {
				t.Errorf("Expected LCA value %d, got %d", tc.expected, res.Val)
			}
		})
	}
}
