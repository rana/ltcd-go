## ltcd-go

LeetCode solutions in Go.

### LeetCode

[Top Interview 150](https://leetcode.com/studyplan/top-interview-150/)

## Spreadsheet

[Progress Tracking](https://docs.google.com/spreadsheets/d/18OZ1SwBK0OG4Wl3-_FqsO487ZhylPrbwA9HGuq3OHYs/edit?usp=sharing)

###

Example: words = ["lc", "lc", "cl", "cl", "cl"]
num = 2 (frequency of "lc")
rvs = 3 (frequency of "cl")

┌────────────────────────────────┐
│ Word Pairs Analysis: "lc"/"cl" │
└────────────────────────────────┘

Frequencies:
"lc": ┌─────┬─────┐
      │ lc  │ lc  │
      └─────┴─────┘

"cl": ┌─────┬─────┬─────┐
      │ cl  │ cl  │ cl  │
      └─────┴─────┴─────┘

1. Usable pairs: min(num, rvs) = min(2, 3) = 2 pairs

2. Result contribution: min(num, rvs) * 2
   Each word contributes 2 characters
   Two pairs contribute: "lc" + "cl" = "lccl" (4 chars)
   res += 2 * 2 = 4 (two words used in each pair)
```

### Combined Example
```
words = ["lc", "cl", "gg"]

┌────────────────────────────────┐
│ Complete Palindrome Formation  │
└────────────────────────────────┘

1. Process "gg" (palindrome):
   - Pairs: 1/2 = 0 pairs
   - Center: 1%2 = 1 (can use in center)

2. Process "lc"/"cl" (pairs):
   - min(1,1) = 1 pair
   - contributes "lc" + "cl"

Final palindrome: "lc" + "gg" + "cl" = "lcggcl"
Length = 6
