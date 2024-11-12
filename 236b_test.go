package main

import "testing"

// TreeNode236b represents a node in a binary tree
type TreeNode236b struct {
	Val   int
	Left  *TreeNode236b
	Right *TreeNode236b
}

func lowestCommonAncestor236b(root, p, q *TreeNode236b) *TreeNode236b {
	// Lowest Common Ancestor of a Binary Tree
	// Given the root of a binary tree.
	// Given two nodes a and b in the tree.
	// Find the lowest common ancestor node.
	// Return the lowest common ancestor node.
	// Lowest means closest to the leaf.
	// A node may be it's own ancestor.
	// Use depth-first search (DFS) with postorder traversal.
	// Postorder left -> right -> current.
	// Postorder traversal visits all descendant nodes prior to ancestor nodes.
	// We would like the bottom-up approach since we're looking
	// for nodes closest to leaves.

	// Check base cases.
	if root == nil || root == p || root == q {
		return root
	}

	// Postorder traversal: left, right, current.

	// Process left subtree.
	lft := lowestCommonAncestor236b(root.Left, p, q)

	// Process right subtree.
	rht := lowestCommonAncestor236b(root.Right, p, q)

	// Check for success condition.
	if lft != nil && rht != nil {
		// Current node is lowest common ancestor.
		return root
	}

	// Return non-nil node.
	if lft != nil {
		return lft
	}
	return rht
}

func createNode236b(val int) *TreeNode236b {
	return &TreeNode236b{Val: val}
}

func TestLowestCommonAncestor236b(t *testing.T) {
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
			nodes := make([]*TreeNode236b, len(tc.treeVals))
			for i, val := range tc.treeVals {
				if val != -1 {
					nodes[i] = createNode236b(val)
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
			var p, q *TreeNode236b
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
			res := lowestCommonAncestor236b(nodes[0], p, q)

			// Verify result
			if res.Val != tc.expected {
				t.Errorf("Expected LCA value %d, got %d", tc.expected, res.Val)
			}
		})
	}
}
