Let me visualize how the `hits` function checks for collisions step by step:

```
Example:
Figure:           Field:
[0 1 1]          [0 0 0 0 0]
[1 0 0]          [0 0 0 0 0]
[0 1 0]          [1 1 0 0 0]
                 [1 1 0 0 0]

Let's check if figure hits at position (row=1, col=1):

Step-by-step check (figRow, figCol):
(0,0): fig[0][0]=0 & fld[1][1]=0 → No hit
(0,1): fig[0][1]=1 & fld[1][2]=0 → No hit
(0,2): fig[0][2]=1 & fld[1][3]=0 → No hit

(1,0): fig[1][0]=1 & fld[2][1]=1 → HIT! Return true
       ↑               ↑
       figure         field
```

Visual overlay of the check:
```
Figure overlay on field position (row=1, col=1):
[0 0 0 0 0]
[0 0 1 1 0]    ← row+figRow
[0 1 0 0 0]    * Collision at (2,1)
[1 1 0 0 0]
   ↑
   col+figCol

Formula for field position:
- Row = row + figRow
- Col = col + figCol
```

The function returns true when it finds ANY position where:
```
1. Figure has a 1 (fig[figRow][figCol] == 1)
   AND
2. Field has a 1 (fld[row+figRow][col+figCol] == 1)
```

Memory aid:
```
Think of it as overlaying the 3x3 figure on the field:
- If any figure's 1 overlaps with field's 1
- We have a hit (collision)
- Return true immediately
```

---
So `row+figRow` is a test of a "drop" with translation of the figure over the field?

---
Exactly! Let me visualize this "dropping" translation concept:

```
Figure:     Initial Field:    Testing drops at col=1:
[0 1 1]    [0 0 0 0 0]      row=0: [0 0 0 0 0]    No hit
[1 0 0]    [0 0 0 0 0]             [0 0 1 1 0]
[0 1 0]    [1 1 0 0 0]             [0 1 0 0 0]
           [1 1 0 0 0]             [1 1 0 0 0]

                            row=1: [0 0 0 0 0]  HIT!
                                   [0 0 1 1 0]
                                   [0 1 0 0 0]  ← Hits here
                                   [1 1 0 0 0]
```

The translation formula:
```
field[row+figRow][col+figCol]
      ↑   ↑       ↑   ↑
      |   |       |   |
      |   |       |   +-- Figure's column offset (0,1,2)
      |   |       +------ Starting column to check
      |   +-------------- Figure's row offset (0,1,2)
      +------------------ Starting row to check
```

Think of it like placing a transparent overlay:
```
row=0:          row=1:          row=2:
[0 1 1]         [0 1 1]         [0 1 1]
[1 0 0]   →     [1 0 0]   →     [1 0 0]
[0 1 0]         [0 1 0]         [0 1 0]
   ↓               ↓               ↓
[0 0 0 0 0]    [0 0 0 0 0]    [0 0 0 0 0]
[0 0 0 0 0]    [0 0 0 0 0]    [0 0 0 0 0]
[1 1 0 0 0]    [1 1 0 0 0]    [1 1 0 0 0]
[1 1 0 0 0]    [1 1 0 0 0]    [1 1 0 0 0]
```

The `hits` function is essentially:
1. Taking a 3x3 snapshot of the field starting at (row,col)
2. Checking if any 1s in the figure overlap with 1s in that snapshot
3. This simulates "dropping" the figure to test if it can fall further

---
```go
rowFull := func(fld, fig [][]int, row, col int) bool {
		for pos := range fldCols {
			if fld[row][pos] == 0 && (pos < col || pos >= col+3 || fig[0][pos-col] == 0) {
				return false
			}
		}
		return true
	}
```
`rowFull` checks for the negative case at the relevant row by:
* if the field column is zero outside of the figure (before the figure or after the figure)
* if the figure column is zero, `pos-col` acts as a translation of the figure over the field.

Would you explain why fig[0] is ok.

---
Ah, good catch! Let me explain why `fig[0]` is incorrect and should be `fig[row-row]` or more clearly, `figRow`:

```go
rowFull := func(fld, fig [][]int, row, col int) bool {
    // We need to know which row of the figure we're checking!
    figRow := row - row  // If checking first row where figure lands
    for pos := range fldCols {
        if fld[row][pos] == 0 && (pos < col || pos >= col+3 || fig[figRow][pos-col] == 0) {
            return false
        }
    }
    return true
}
```

Let's visualize why this matters:

```
Figure:        Field at row=row:
[0 1 1]        [0 0 0 0 0]
[1 0 0]   →    Check this row for fullness
[0 1 0]

When checking first row (figRow=0):
pos:     0  1  2  3  4
field:  [0  0  0  0  0]
         ↑  ↑  ↑  ↑  ↑
fig[0]: x [0  1  1] x   (pos-col translates to figure position)

When checking second row (figRow=1):
pos:     0  1  2  3  4
field:  [0  0  0  0  0]
         ↑  ↑  ↑  ↑  ↑
fig[1]: x [1  0  0] x   (need different figure row!)

When checking third row (figRow=2):
pos:     0  1  2  3  4
field:  [0  0  0  0  0]
         ↑  ↑  ↑  ↑  ↑
fig[2]: x [0  1  0] x   (need different figure row!)
```

The fix should be:
```go
rowFull := func(fld, fig [][]int, row, col int) bool {
    figRow := row - row  // Calculate which row of figure we're checking
    for pos := range fldCols {
        if fld[row][pos] == 0 && (pos < col || pos >= col+3 || fig[figRow][pos-col] == 0) {
            return false
        }
    }
    return true
}
```

Using just `fig[0]` would only ever check the first row of the figure, which is incorrect when checking if the second or third rows create a full row in the field.

The translation works in two dimensions:
1. `pos-col` translates horizontal position (which column of figure)
2. `figRow` should translate vertical position (which row of figure)

---

Naming pattern:
```
Pattern:    [source][dimension]
Examples:   fldRow, fldCol, figRow, figCol

Variables:
old → new        | Reason
----------------+----------------------------------
pos → curCol    | Scanning current column position
row → fldRow    | Field's row being checked
col → fldCol    | Field's starting column for figure
```

Visual explanation:
```
Field:                     Figure:
[. . . . .]               [0 1 1]
[. . . . .]               [1 0 0]
[. x x x .] fldRow   →    [0 1 0]
     ↑                      ↑
   curCol                 figRow
     ↑
   fldCol
```

This naming system:
1. Always identifies source (fld/fig)
2. Always identifies dimension (Row/Col)
3. Uses 'cur' for scanning position
4. Makes translations clearer:
   - `curCol-fldCol` maps field to figure position
   - `figRow` shows which figure row we're using
