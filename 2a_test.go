package main

import (
	"reflect"
	"testing"
)

// ListNode2 definition for singly-linked list
type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// Time complexity: O(max(m,n)), m and n are the lengths of the list. We traverse each list once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/703d6000-3fc2-4646-afc2-05dc51e03593
func addTwoNumbers2(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	// Add Two Numbers
	// Given two linked lists.
	// Each linked list has digits of an integer in reverse.
	// Sum the two integers together.
	// Return a linked list of the digits of the sum.
	// Use a dummy head node and a carry value.

	// Initialize variables.
	hed := &ListNode2{Val: 0}
	cur := hed
	cry := 0

	// Process while any node not nil or carry exists.
	for l1 != nil || l2 != nil || cry > 0 {
		// Initialize sum with carry.
		sum := cry
		// Sum list 1 node.
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		// Sum list 2 node.
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Store current digit and carry.
		cur.Next = &ListNode2{Val: sum % 10}
		cur = cur.Next
		cry = sum / 10
	}

	return hed.Next
}

// Test cases
func TestAddTwoNumbers2(t *testing.T) {
	// Helper function to create linked list from slice
	createList := func(nums []int) *ListNode2 {
		var hed = &ListNode2{}
		var cur = hed
		for _, num := range nums {
			cur.Next = &ListNode2{Val: num}
			cur = cur.Next
		}
		return hed.Next
	}

	// Helper function to convert list to slice
	listToSlice := func(hed *ListNode2) []int {
		var res []int
		for cur := hed; cur != nil; cur = cur.Next {
			res = append(res, cur.Val)
		}
		return res
	}

	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for idx, tt := range tests {
		l1 := createList(tt.l1)
		l2 := createList(tt.l2)
		got := listToSlice(addTwoNumbers2(l1, l2))

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Test %d: got %v, want %v", idx, got, tt.want)
		}
	}
}
