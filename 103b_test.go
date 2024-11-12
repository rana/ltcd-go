package main

import (
	"reflect"
	"testing"
)

// TreeNode103b represents a node in a binary tree
type TreeNode103b struct {
	Val   int
	Left  *TreeNode103b
	Right *TreeNode103b
}

// Time complexity: O(n), n is the number of nodes in the tree. We visit each node once.
// Space complexity: O(w), w is the maximum width of a tree level.
func zigzagLevelOrder103b(root *TreeNode103b) [][]int {
	// Binary Tree Zigzag Level Order Traversal
	// Given the root of a binary tree.
	// Collect the zigzag level-order traversal.
	// Return an array pf zigzag level-order traversal.
	// Zigzag starts left to right for level,
	// then right to left at next level, and continues
	// alternation.
	// Use breadth-first search (BFS), and queue.
	// Use level depth counter to enable zigzag.

	// Check base case.
	if root == nil {
		return [][]int{}
	}

	// Initialize variables.
	var res [][]int
	que := []*TreeNode103b{root}
	lvlDep := 0

	// Traverse tree level by level.
	for len(que) > 0 {
		// Track level length for level processing.
		lvlLen := len(que)
		// Initialize level values array.
		lvlVals := make([]int, lvlLen)

		// Iterate through each node in the level.
		for idx := 0; idx < lvlLen; idx++ {
			// Get and remove first node from que.
			cur := que[0]
			que = que[1:]

			// Store current node value in array.
			// Determine index based on level depth for zigzag.
			if lvlDep%2 == 0 {
				lvlVals[idx] = cur.Val
			} else {
				lvlVals[lvlLen-idx-1] = cur.Val
			}

			// Add left to que.
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			// Add right to que.
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		// Store level values.
		res = append(res, lvlVals)

		// Increment level depth.
		lvlDep++
	}

	return res
}

func TestZigzagLevelOrder103b(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode103b
		want [][]int
	}{
		{
			name: "Example 1: Three levels",
			root: &TreeNode103b{
				Val:  3,
				Left: &TreeNode103b{Val: 9},
				Right: &TreeNode103b{
					Val:   20,
					Left:  &TreeNode103b{Val: 15},
					Right: &TreeNode103b{Val: 7},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "Example 2: Single node",
			root: &TreeNode103b{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "Example 3: Empty tree",
			root: nil,
			want: [][]int{},
		},
		{
			name: "Complete binary tree",
			root: &TreeNode103b{
				Val: 1,
				Left: &TreeNode103b{
					Val:   2,
					Left:  &TreeNode103b{Val: 4},
					Right: &TreeNode103b{Val: 5},
				},
				Right: &TreeNode103b{
					Val:   3,
					Left:  &TreeNode103b{Val: 6},
					Right: &TreeNode103b{Val: 7},
				},
			},
			want: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := zigzagLevelOrder103b(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tc.want)
			}
		})
	}
}
