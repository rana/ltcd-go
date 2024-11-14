package main

import (
	"strconv"
	"testing"
)

// Time complexity: O(n * l * 4), n is the number of genes in the bank. l is the length of a gene. 4 is the number of nucleotides.
// Space complexity: O(n), for maps and queue.
// https://claude.ai/chat/57cffd0c-1611-40f7-b9d8-b976ec840da7
func minMutation433(strGen string, endGen string, bank []string) int {
	// Minimum Genetic Mutation
	// Given two strings strGen and endGen.
	// Given a string array bank of genes.
	// Determine the minimum number mutations to go
	// from strGen to endGen.
	// Return the minimum number of mutations; or, -1.
	// Conditions:
	//	* Four mutation choice of 'A', 'C', 'G', 'T'
	//	* Gene is eight characters long.
	//	* Mutation must exist in bank.
	// Model as graph traversal.
	// Breadth-first search guarantees minimum path.
	// Store gene and step count in queue.
	// Use map of bank for O(1) mutation check.

	// Transform bank to map for O(1) lookup.
	vld := make(map[string]bool)
	for _, gen := range bank {
		vld[gen] = true
	}

	// Check input edge case.
	if !vld[endGen] {
		return -1
	}

	// Nucleotide range.
	nuc := []byte{'A', 'C', 'G', 'T'}

	// Initialize visit array.
	vst := make(map[string]bool)
	vst[strGen] = true

	// Initialize que with gene and step count.
	que := [][2]string{{strGen, "0"}}

	// Breadth-first search.
	for len(que) > 0 {
		// Get and remove first gene and step count from que.
		cur := que[0][0]
		stp, _ := strconv.Atoi(que[0][1])
		que = que[1:]

		// Check whether we found end gene.
		if cur == endGen {
			return stp
		}

		// Iterate through mutations.
		buf := []byte(cur)
		for pos := 0; pos < 8; pos++ {
			org := buf[pos]
			for _, nxt := range nuc {
				if nxt != org {
					// Mutate
					buf[pos] = nxt
					nxtGen := string(buf)

					// Check valid and unvisited.
					if vld[nxtGen] && !vst[nxtGen] {
						vst[nxtGen] = true
						que = append(que, [2]string{nxtGen, strconv.Itoa(stp + 1)})
					}
				}
			}
			buf[pos] = org
		}
	}

	return -1
}

func TestMinMutation433(t *testing.T) {
	tst := []struct {
		name   string
		strGen string
		endGen string
		bank   []string
		want   int
	}{
		{
			name:   "Single mutation",
			strGen: "AACCGGTT",
			endGen: "AACCGGTA",
			bank:   []string{"AACCGGTA"},
			want:   1,
		},
		{
			name:   "Two mutations",
			strGen: "AACCGGTT",
			endGen: "AAACGGTA",
			bank:   []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			want:   2,
		},
		{
			name:   "No solution",
			strGen: "AAAAACCC",
			endGen: "AACCCCCC",
			bank:   []string{"AAAAAACC"},
			want:   -1,
		},
	}

	for _, tc := range tst {
		t.Run(tc.name, func(t *testing.T) {
			got := minMutation433(tc.strGen, tc.endGen, tc.bank)
			if got != tc.want {
				t.Errorf("minMutation() = %v, want %v", got, tc.want)
			}
		})
	}
}
