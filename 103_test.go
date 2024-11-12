package main

import (
	"reflect"
	"testing"
)

// TreeNode103 represents a node in a binary tree
type TreeNode103 struct {
	Val   int
	Left  *TreeNode103
	Right *TreeNode103
}

// Time complexity: O(n), n is the number of nodes in the tree.
// Space complexity: O(w), w is the maximum width of a tree level.
// https://claude.ai/chat/0de0fc9f-4c55-4bd9-bdc5-7b8b222b0e58
func zigzagLevelOrder103(root *TreeNode103) [][]int {
	// Binary Tree Zigzag Level Order Traversal
	// Given the root of a binary tree.
	// Traverse the tree in level-order.
	// Return an array of level values in zigzag order.
	// One level of values is left to right.
	// Next level of values are right to left.
	// Use level-order BFS traversal with a queue.
	// Use value array indexes to achieve zigzag ordering.

	// Check base case.
	if root == nil {
		return [][]int{}
	}

	// Initialize variables.
	var res [][]int
	que := []*TreeNode103{root}
	lvlDep := 0

	// Iterate through the tree level by level.
	for len(que) > 0 {
		// Track level length for level processing.
		lvlLen := len(que)
		// Initialize array of level values.
		lvl := make([]int, lvlLen)

		// Iterate through the current level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove first node from que.
			cur := que[0]
			que = que[1:]

			// Add current node value.
			// Choose zigzag order based on level depth.
			if lvlDep%2 == 0 {
				lvl[idx] = cur.Val
			} else {
				lvl[lvlLen-idx-1] = cur.Val
			}

			// Add nodes to que for level below.
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		// Add level values to result.
		res = append(res, lvl)

		// Increment level depth.
		lvlDep++
	}

	return res
}

func TestZigzagLevelOrder103(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode103
		want [][]int
	}{
		{
			name: "Example 1: Three levels",
			root: &TreeNode103{
				Val:  3,
				Left: &TreeNode103{Val: 9},
				Right: &TreeNode103{
					Val:   20,
					Left:  &TreeNode103{Val: 15},
					Right: &TreeNode103{Val: 7},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "Example 2: Single node",
			root: &TreeNode103{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "Example 3: Empty tree",
			root: nil,
			want: [][]int{},
		},
		{
			name: "Complete binary tree",
			root: &TreeNode103{
				Val: 1,
				Left: &TreeNode103{
					Val:   2,
					Left:  &TreeNode103{Val: 4},
					Right: &TreeNode103{Val: 5},
				},
				Right: &TreeNode103{
					Val:   3,
					Left:  &TreeNode103{Val: 6},
					Right: &TreeNode103{Val: 7},
				},
			},
			want: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := zigzagLevelOrder103(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tc.want)
			}
		})
	}
}
