package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of the array jmps.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/671482a5-b5c4-8002-9a7c-a259de69231e
func jump(jmps []int) int {
	return 0
}

// Test cases for the jump function.
func TestJumpCase1(t *testing.T) {
	nms := []int{2, 3, 1, 1, 4}
	exp := 2
	res := jump(nms)
	if res != exp {
		t.Errorf("TestJumpCase1 failed: expected %d, got %d", exp, res)
	}
}

func TestJumpCase2(t *testing.T) {
	nms := []int{2, 3, 0, 1, 4}
	exp := 2
	res := jump(nms)
	if res != exp {
		t.Errorf("TestJumpCase2 failed: expected %d, got %d", exp, res)
	}
}

func TestJumpCase3(t *testing.T) {
	nms := []int{0}
	exp := 0
	res := jump(nms)
	if res != exp {
		t.Errorf("TestJumpCase3 failed: expected %d, got %d", exp, res)
	}
}

func TestJumpCase4(t *testing.T) {
	nms := []int{1, 2, 1, 1, 1}
	exp := 3
	res := jump(nms)
	if res != exp {
		t.Errorf("TestJumpCase4 failed: expected %d, got %d", exp, res)
	}
}
