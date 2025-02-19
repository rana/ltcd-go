package main

import (
	"reflect"
	"testing"
)

// Module 1 – Basic Coding
// Given an array a, your task is to output an array b of the same length by applying the following transformation:
// • For each i from 0 to a.length - 1 inclusive, b[i] = a[i - 1] + a[i] + a[i + 1]
// • If an element in the sum a[i - 1] + a[i] + a[i + 1] does not exist, use 0 in its place
// • For instance, b[0] = 0 + a[0] + a[1]
// Example
// For a = [4, 0, 1, -2, 3]:
// • b[0] = 0 + a[0] + a[1] = 0 + 4 + 0 = 4
// • b[1] = a[0] + a[1] + a[2] = 4 + 0 + 1 = 5
// • b[2] = a[1] + a[2] + a[3] = 0 + 1 + (-2) = -1
// • b[3] = a[2] + a[3] + a[4] = 1 + (-2) + 3 = 2
// • b[4] = a[3] + a[4] + 0 = (-2) + 3 + 0 = 1
// So, the output should be solution(a) = [4, 5, -1, 2, 1].

// Time complexity: O(n)
// Space complexity: O(1)
// https://console.anthropic.com/workbench/454e7a54-eb0d-4356-a480-f3dcb2bf400f
func solutionGCA1(arr []int) []int {
	// Get array length
	arrLen := len(arr)

	// Handle empty array case
	if arrLen == 0 {
		return []int{}
	}

	// Initialize result array
	res := make([]int, arrLen)

	// Process each position
	for idx := range arr {
		// Get previous element or 0 if out of bounds
		prv := 0
		if idx > 0 {
			prv = arr[idx-1]
		}

		// Get next element or 0 if out of bounds
		nxt := 0
		if idx < arrLen-1 {
			nxt = arr[idx+1]
		}

		// Calculate sum for current position
		res[idx] = prv + arr[idx] + nxt
	}

	return res
}

func TestSolutionGCA1(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		exp  []int
	}{
		{
			name: "Example case",
			arr:  []int{4, 0, 1, -2, 3},
			exp:  []int{4, 5, -1, 2, 1},
		},
		{
			name: "Empty array",
			arr:  []int{},
			exp:  []int{},
		},
		{
			name: "Single element",
			arr:  []int{1},
			exp:  []int{1},
		},
		{
			name: "All zeros",
			arr:  []int{0, 0, 0},
			exp:  []int{0, 0, 0},
		},
		{
			name: "Mixed values",
			arr:  []int{1, -1, 1, -1},
			exp:  []int{0, 1, -1, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := solutionGCA1(tc.arr)
			if !reflect.DeepEqual(got, tc.exp) {
				t.Errorf("got %v, want %v", got, tc.exp)
			}
		})
	}
}
