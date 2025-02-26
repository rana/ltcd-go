package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func evalRPN150b(tkns []string) int {
	return 0
}

// TestEvalRPN contains test cases for the evalRPN function
func TestEvalRPN150b(t *testing.T) {
	tests := []struct {
		name string
		tkns []string
		want int
	}{
		{
			name: "Simple addition and multiplication",
			tkns: []string{"2", "1", "+", "3", "*"},
			want: 9,
		},
		{
			name: "Division and addition",
			tkns: []string{"4", "13", "5", "/", "+"},
			want: 6,
		},
		{
			name: "Complex expression",
			tkns: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want: 22,
		},
		{
			name: "Single number",
			tkns: []string{"42"},
			want: 42,
		},
		{
			name: "Simple subtraction",
			tkns: []string{"5", "3", "-"},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN150b(tt.tkns); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
