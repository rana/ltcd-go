package main

import "testing"

type MinStack struct {
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		vals: []int{},
		mins: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.vals = append(s.vals, val)

	if len(s.mins) == 0 {
		s.mins = append(s.mins, val)
	} else {
		curMin := s.mins[len(s.mins)-1]
		if val < curMin {
			curMin = val
		}
		s.mins = append(s.mins, curMin)
	}
}

func (s *MinStack) Pop() {
	s.vals = s.vals[:len(s.vals)-1]
	s.mins = s.mins[:len(s.mins)-1]
}

func (s *MinStack) Top() int {
	return s.vals[len(s.vals)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

func TestMinStack(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		vals     []int
		expected []interface{}
	}{
		{
			name:     "LeetCode example",
			ops:      []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			vals:     []int{-2, 0, -3}, // Only include actual push values
			expected: []interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		},
		{
			name:     "Single element operations",
			ops:      []string{"MinStack", "push", "top", "getMin", "pop"},
			vals:     []int{5}, // Only one push value
			expected: []interface{}{nil, nil, 5, 5, nil},
		},
		{
			name:     "Multiple identical elements",
			ops:      []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			vals:     []int{3, 3, 3}, // Only push values
			expected: []interface{}{nil, nil, nil, nil, 3, nil, 3, nil, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var stk MinStack
			pushIdx := 0

			for i, op := range tc.ops {
				t.Logf("Operation %d: %s", i, op)
				switch op {
				case "MinStack":
					stk = Constructor()
				case "push":
					val := tc.vals[pushIdx]
					t.Logf("Pushing value: %d", val)
					stk.Push(val)
					pushIdx++
				case "pop":
					stk.Pop()
				case "top":
					got := stk.Top()
					exp, _ := tc.expected[i].(int)
					if got != exp {
						t.Errorf("Top() = %v, want %v", got, exp)
					}
				case "getMin":
					got := stk.GetMin()
					exp, _ := tc.expected[i].(int)
					if got != exp {
						t.Errorf("GetMin() = %v, want %v", got, exp)
					}
				}
			}
		})
	}
}
