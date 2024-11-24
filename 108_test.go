package main

import "testing"

// TreeNode108 Definition
type TreeNode108 struct {
	Val   int
	Left  *TreeNode108
	Right *TreeNode108
}

// Time complexity: O(n), we visit each element of array nums.
// Space complexity: O(log n), for the recursive call stack.
// https://claude.ai/chat/fdbe5465-3786-48d4-8584-ebadaf9786e4
func sortedArrayToBST108(nums []int) *TreeNode108 {
	// Convert Sorted Array to Binary Search Tree
	// Given a sort-ascending integer array `nums`.
	// Create a height-balanced binary search tree.
	// Return the root of the BST.
	// Use a closure and recursion to build the BST.
	// Choose the middle element for a root node.

	// Check input min edge case.
	numLen := len(nums)
	if numLen == 0 {
		return nil
	}

	// Create a closure.
	var bld func(lft, rht int) *TreeNode108
	bld = func(lft, rht int) *TreeNode108 {
		// Check base case.
		if lft > rht {
			return nil
		}

		// Choose the middle element for the root node.
		mid := lft + (rht-lft)/2

		// Create root node with middle element.
		root := &TreeNode108{Val: nums[mid]}

		// Recursively create left and right nodes.
		root.Left = bld(lft, mid-1)
		root.Right = bld(mid+1, rht)

		return root
	}

	return bld(0, numLen-1)
}

// Test cases
func TestSortedArrayToBST108(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected *TreeNode108
		validate func(*TreeNode108) bool
	}{
		{
			name:  "Example 1: [-10,-3,0,5,9]",
			input: []int{-10, -3, 0, 5, 9},
			validate: func(root *TreeNode108) bool {
				return isBalanced108(root) && isBST108(root, nil, nil)
			},
		},
		{
			name:  "Example 2: [1,3]",
			input: []int{1, 3},
			validate: func(root *TreeNode108) bool {
				return isBalanced108(root) && isBST108(root, nil, nil)
			},
		},
		{
			name:  "Single element",
			input: []int{1},
			validate: func(root *TreeNode108) bool {
				return root.Val == 1 && root.Left == nil && root.Right == nil
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortedArrayToBST108(tc.input)
			if !tc.validate(got) {
				t.Errorf("sortedArrayToBST(%v) failed validation", tc.input)
			}
		})
	}
}

// Helper function to check if tree is height-balanced
func isBalanced108(root *TreeNode108) bool {
	var height func(*TreeNode108) int
	height = func(node *TreeNode108) int {
		if node == nil {
			return 0
		}

		lftHgt := height(node.Left)
		if lftHgt == -1 {
			return -1
		}

		rhtHgt := height(node.Right)
		if rhtHgt == -1 {
			return -1
		}

		if abs(lftHgt-rhtHgt) > 1 {
			return -1
		}

		return max(lftHgt, rhtHgt) + 1
	}

	return height(root) != -1
}

// Helper function to check if tree is a valid BST
func isBST108(root *TreeNode108, min, max *int) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Val <= *min) ||
		(max != nil && root.Val >= *max) {
		return false
	}

	return isBST108(root.Left, min, &root.Val) &&
		isBST108(root.Right, &root.Val, max)
}
