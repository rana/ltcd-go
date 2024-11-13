package main

import "testing"

// Graph399 is an adjacency list with directed edge a->b and weight w, where a/b = w.
type Graph399 map[string]map[string]float64

// Time complexity: O(v + e) per query: v is number of verticies. e is number of edges.
// Space complexity: O(v + e), for graph storage and visited set in DFS.
// https://claude.ai/chat/09f333d1-464b-4fc1-81c3-b7e45e64ce88
func calcEquation399(eqns [][]string, vals []float64, qrys [][]string) []float64 {
	// Evaluate Division
	// Given an array of equation variable pairs `rels`. Represents equations.
	// Given an array of float64 values. Represents the value of an equation.
	// Given an array of query variable pairs `qrys`.
	// A query is an equation variable pair with unknown value.
	// Return an array of query answers.
	// An unknowable answer receives -1.
	// Key insights:
	//	* Model as a weighted directed graph using an adjacency list.
	//	* Each equation a/b=k has two directed edges
	//		- Edge a->b with weight k
	//		- Edge b->a with weight 1/k
	// Use depth-first search (DFS) recursion for weighted path finding.
	// Use helper function dfsFndPath for recursion.

	// Initialize graph.
	gph := make(Graph399)

	// Build graph.
	for idx, eqn := range eqns {
		// Get equation variables and value.
		src, dst := eqn[0], eqn[1]
		val := vals[idx]

		// Initialize adjacency maps.
		if gph[src] == nil {
			gph[src] = make(map[string]float64)
		}
		if gph[dst] == nil {
			gph[dst] = make(map[string]float64)
		}

		// Set bidirectional weights.
		gph[src][dst] = val
		gph[dst][src] = 1.0 / val
	}

	// Process querys.
	res := make([]float64, len(qrys))
	for idx, qry := range qrys {
		// Get query variables.
		src, dst := qry[0], qry[1]

		// Check for special cases.
		if gph[src] == nil || gph[dst] == nil {
			res[idx] = -1.0
			continue
		}
		if src == dst {
			res[idx] = 1.0
			continue
		}

		// Evaluate query.
		vst := make(map[string]bool)
		if val := dfsFndPth399(gph, src, dst, vst); val != 0.0 {
			res[idx] = val
		} else {
			res[idx] = -1.0
		}
	}

	return res
}

// dfsFndPth399 performs a depth-first search for the path weight.
func dfsFndPth399(gph Graph399, cur, dst string, vst map[string]bool) float64 {
	// Mark current node as visited.
	vst[cur] = true

	// Iterate through current neighbors.
	for nxt, val := range gph[cur] {
		// Check whether we found weight.
		if nxt == dst {
			return val
		}

		// Skip visited node.
		if vst[nxt] {
			continue
		}

		// Recursively search from neighbor.
		if subVal := dfsFndPth399(gph, nxt, dst, vst); subVal != 0.0 {
			return val * subVal
		}
	}

	// No weight found.
	return 0.0
}

// Tests
func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			name:      "Example 1",
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			want:      []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name:      "Example 2",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			want:      []float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			name:      "Example 3",
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			want:      []float64{0.5, 2.0, -1.0, -1.0},
		},
		{
			name:      "Single node self query",
			equations: [][]string{{"a", "b"}},
			values:    []float64{2.0},
			queries:   [][]string{{"a", "a"}},
			want:      []float64{1.0},
		},
		{
			name:      "Circular references",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"c", "a"}},
			values:    []float64{2.0, 3.0, 0.166666},
			queries:   [][]string{{"a", "a"}, {"b", "b"}, {"c", "c"}},
			want:      []float64{1.0, 1.0, 1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEquation399(tt.equations, tt.values, tt.queries)
			if len(got) != len(tt.want) {
				t.Errorf("calcEquation() returned %v results, want %v", len(got), len(tt.want))
			}
			for i := range got {
				if !approximatelyEqual(got[i], tt.want[i], 0.00001) {
					t.Errorf("calcEquation()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

// Helper function to compare floating point values
func approximatelyEqual(a, b, epsilon float64) bool {
	if a == b {
		return true
	}
	// Handle -1.0 case separately as it's our error indicator
	if a == -1.0 && b == -1.0 {
		return true
	}
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}
