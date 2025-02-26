package main

import (
	"reflect"
	"testing"
)

// ListNode2b definition for singly-linked list
type ListNode2b struct {
	Val  int
	Next *ListNode2b
}

// Time complexity:
// Space complexity:
func addTwoNumbers2b(l1 *ListNode2b, l2 *ListNode2b) *ListNode2b {
	return nil
}

// Test cases
func TestAddTwoNumbers2b(t *testing.T) {
	// Helper function to create linked list from slice
	createList := func(nums []int) *ListNode2b {
		var hed = &ListNode2b{}
		var cur = hed
		for _, num := range nums {
			cur.Next = &ListNode2b{Val: num}
			cur = cur.Next
		}
		return hed.Next
	}

	// Helper function to convert list to slice
	listToSlice := func(hed *ListNode2b) []int {
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
		got := listToSlice(addTwoNumbers2b(l1, l2))

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Test %d: got %v, want %v", idx, got, tt.want)
		}
	}
}
