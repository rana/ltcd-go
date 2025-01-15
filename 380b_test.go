package main

import (
	"testing"
)

type RandomizedSet struct {
}

func Constructor380b() RandomizedSet {
	return RandomizedSet{}
}

func (rs *RandomizedSet) Insert(val int) bool {
	return false
}

func (rs *RandomizedSet) Remove(val int) bool {
	return false
}

func (rs *RandomizedSet) GetRandom() int {
	return 0
}

func TestRandomizedSet380b(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		args     []int
		expected []bool
	}{
		{
			name:     "Test Insert",
			ops:      []string{"Insert", "Insert"},
			args:     []int{1, 1},
			expected: []bool{true, false},
		},
		{
			name:     "Test Remove",
			ops:      []string{"Insert", "Remove", "Remove"},
			args:     []int{1, 1, 1},
			expected: []bool{true, true, false},
		},
		{
			name:     "Test Insert Remove Sequence",
			ops:      []string{"Insert", "Remove", "Insert"},
			args:     []int{1, 1, 1},
			expected: []bool{true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := Constructor380b()

			for i, op := range tt.ops {
				var result bool

				switch op {
				case "Insert":
					result = rs.Insert(tt.args[i])
				case "Remove":
					result = rs.Remove(tt.args[i])
				}

				if result != tt.expected[i] {
					t.Errorf("%s: operation %d: got %v, want %v",
						tt.name, i, result, tt.expected[i])
				}
			}
		})
	}
}

func TestRandomizedSet_Random380b(t *testing.T) {
	tests := []struct {
		name     string
		inserts  []int
		validSet map[int]bool
	}{
		{
			name:     "Test GetRandom with 3 numbers",
			inserts:  []int{1, 2, 3},
			validSet: map[int]bool{1: true, 2: true, 3: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := Constructor380b()
			for _, num := range tt.inserts {
				rs.Insert(num)
			}

			rnd := rs.GetRandom()
			if !tt.validSet[rnd] {
				t.Errorf("%s: GetRandom() = %d, want one of %v",
					tt.name, rnd, tt.validSet)
			}
		})
	}
}
