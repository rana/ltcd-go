package main

import "testing"

// TreeNode100 represents a binary tree node
type TreeNode100 struct {
	Val   int
	Left  *TreeNode100
	Right *TreeNode100
}

// Time complexity: O(min(n,m)), n is the number of nodes in tree p. m is the number of nodes in tree q. We stop visiting nodes at the first mismatch. Worst case (identical trees), we visit all nodes in both trees.
// Space complexity: O(max(h1, h2)), h1 is the height of tree p. h2 is the height of tree q. Best case, O(logn) for balanced trees. Worst case, O(n) for skewed trees.
// https://claude.ai/chat/f29d0f10-76f9-412a-aafa-ff05c1592ba6
func isSameTree(p *TreeNode100, q *TreeNode100) bool {
	// Same Tree
	// Given the roots of two binary trees p an q.
	// Determine whether the two trees are identical.
	// Return true if conditions met.
	// Use a depth-first search (DFS) for simplicity.

	// Check base cases.
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	// Check value equality.
	// Check equality of left and right sides.
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

// Test cases
func TestIsSameTree(t *testing.T) {
	cases := []struct {
		name string
		p    *TreeNode100
		q    *TreeNode100
		want bool
	}{
		{
			name: "Example 1: Same trees [1,2,3]",
			p: &TreeNode100{
				Val:   1,
				Left:  &TreeNode100{Val: 2},
				Right: &TreeNode100{Val: 3},
			},
			q: &TreeNode100{
				Val:   1,
				Left:  &TreeNode100{Val: 2},
				Right: &TreeNode100{Val: 3},
			},
			want: true,
		},
		{
			name: "Example 2: Different structure [1,2] vs [1,null,2]",
			p: &TreeNode100{
				Val:  1,
				Left: &TreeNode100{Val: 2},
			},
			q: &TreeNode100{
				Val:   1,
				Right: &TreeNode100{Val: 2},
			},
			want: false,
		},
		{
			name: "Example 3: Different values [1,2,1] vs [1,1,2]",
			p: &TreeNode100{
				Val:   1,
				Left:  &TreeNode100{Val: 2},
				Right: &TreeNode100{Val: 1},
			},
			q: &TreeNode100{
				Val:   1,
				Left:  &TreeNode100{Val: 1},
				Right: &TreeNode100{Val: 2},
			},
			want: false,
		},
		{
			name: "Both nil trees",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "One nil tree",
			p: &TreeNode100{
				Val: 1,
			},
			q:    nil,
			want: false,
		},
		{
			name: "Single node same value",
			p:    &TreeNode100{Val: 1},
			q:    &TreeNode100{Val: 1},
			want: true,
		},
		{
			name: "Single node different values",
			p:    &TreeNode100{Val: 1},
			q:    &TreeNode100{Val: 2},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSameTree(tc.p, tc.q)
			if got != tc.want {
				t.Errorf("isSameTree() = %v, want %v", got, tc.want)
			}
		})
	}
}
