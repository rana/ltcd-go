### Solution Development

Provide the simplest, most straightforward implementation sufficient for interviews and most practical uses.

1. Question Fundamental Assumptions
   - What's the minimal information needed?
   - Are we storing/tracking more than necessary?
   - Can we work in a different direction?
   - Can we eliminate any data structures?

2. Start with Simplest Possible Approach
   - Begin with most direct solution
   - Add complexity only when necessary
   - Question each additional component
   - Start with a naive solution to establish baseline
   - Progress through optimizations systematically
   - Document trade-offs between different approaches

3. Compare Alternative Approaches
   - Consider reverse/forward processing
   - Evaluate direct vs indirect validation
   - Look for opportunities to combine steps
   - Progress through optimizations systematically
   - Document trade-offs between different approaches
   - Consider both time and space complexity at each step

### Verification Process
- Validate solution against edge cases
- Verify time and space complexity claims
- Check for potential optimizations
- Review code style and idiomaticity
- Ensure consistent naming conventions
- Confirm test coverage adequacy

### Solution Validation and Refinement
- Question solution fundamentals:
  * Is this the simplest possible approach?
  * What assumptions are we making?
  * Could we eliminate any data structures?
  * Are we solving the problem forwards when backwards might be simpler (or vice versa)?

- Verify correctness and efficiency:
  * Does it handle all edge cases?
  * Are we storing or processing more than necessary?
  * Could any operations be combined or eliminated?
  * Is the time/space complexity optimal?
  * Review code style and idiomaticity
  * Ensure consistent naming conventions
  * Confirm test coverage adequacy

- Improve code quality:
  * Remove unnecessary code
  * Make it more idiomatic for the language
  * Ensure consistent naming (using three-letter acronyms)
  * Maintain clarity over cleverness

### Common Pitfalls
1. Implementation Errors
   - Off-by-one in loops
   - Array bounds checking
   - Integer overflow scenarios
   - Null/nil pointer handling

2. Edge Cases
   - Empty input
   - Single element
   - Duplicate elements
   - Maximum/minimum values
   - Negative numbers
   - Power of 2 vs non-power of 2

3. Boundary Conditions
   - Start/end of sequences
   - First/last elements
   - State transitions
   - Cycle detection endpoints

4. State Management Pitfalls
   - Redundant state tracking
   - Manual counters mirroring collection sizes
   - Complex state synchronization logic
   - Unnecessary defensive checks
   - Over-engineering state validation

## Solution Presentation Approach
### Primary Focus
- Present the optimized, readable solution first
- Explain the intuition and approach thoroughly
- Demonstrate clean, idiomatic implementation
- Show comprehensive test coverage

### Additional Approaches
- Discuss alternative solutions after main implementation
- Compare trade-offs and use cases
- Explain when different approaches might be preferred
- Connect to similar problems and patterns

## Implementation Guidelines

### Programming Language (Go)
- Write all code in the Go programming language
- Ensure the code is idiomatic for Go
- Include clear and easily understandable comments for the source code
- Write `"package main"` and include `"import"` statements for required modules
- For structs and collections with zero initialization. Prefer *struct* idiom `var res strings.Builder` instead of `res := strings.Builder{}`
- For primitive value types. Prefer *primitive* idiom `hayLen := len(hay)` instead of `var hayLen = len(hay)`
- Prefer idiomatic `for` with presence of both index and value:
   ```go
   for idx, num := range nums {
      numIdx[num] = idx
   }
   ```
   Avoid:
   ```go
   for idx := 0; idx < len(nums); idx++ {
		numIdx[nums[idx]] = idx
	}
   ```
- For function parameter with identical types, prefer idiomatic identifer separated by comma `func buildTree(lft, rht int)` instead of `func buildTree(lft int, rht int)`.

#### Integer Range Iteration
- For iterating over integer ranges, prefer:
  ```go
  // Prefer: Simple range integer
  for idx := range numCrs {
      // Use idx
  }
  ```
  Over:
  ```go
  // Avoid: Traditional C-style loop
  for idx := 0; idx < numCrs; idx++ {
      // Use idx
  }
  ```

- For descending ranges, still prefer range with post-decrement:
  ```go
  // Prefer: Range with post-processing
  for idx := range numCrs {
      cur := numCrs - 1 - idx
      // Use cur
  }
  ```
  Over:
  ```go
  // Avoid: Traditional descending loop
  for idx := numCrs - 1; idx >= 0; idx-- {
      // Use idx
  }
  ```

- Exception: When stride is not 1
  ```go
  // Acceptable: Non-unit stride requires traditional loop
  for idx := 0; idx < numCrs; idx += 2 {
      // Use idx
  }
  ```

### Go-Specific Practices
- Follow Go idioms and conventions
- Use package organization best practices
- Implement proper error handling
- Utilize Go's concurrency patterns when appropriate


### Code Structure
- Group related functions together
- Separate implementation from tests
- Use meaningful package names
- Include comprehensive comments
- Implement proper error handling

### Identifier Naming Guidelines

#### Core Principles
- Use three-letter acronyms (TLAs) consistently
- Maintain phonetic relationship to original word
- Prefer consonants when creating new acronyms
- Choose first three characters when they preserve meaning
- Ensure acronyms are immediately recognizable
- When multiple words could describe same concept, prefer the word that:
  * Creates a more distinctive acronym (e.g., 'find/fnd' over 'search/src' to avoid collision with 'source')
  * Better preserves meaning in its three-letter form
  * Maintains clearer consonant structure
  * Examples:
    - find/fnd over search/src (avoids source collision)
    - remove/rem over delete/del (avoids delivery collision)
    - verify/vrf over check/chk (stronger consonants)

#### Standard Acronyms by Category

##### Data Structures
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| array        | arr     | Collections |
| matrix       | mtx     | 2D arrays |
| stack        | stk     | Stack DS |
| queue        | que     | Queue DS |
| heap         | hep     | Heap DS |
| tree         | tre     | Tree DS |
| graph        | gph     | Graph DS |
| node         | nod     | Trees, graphs |
| edge         | edg     | Graphs |
| vertex       | vtx     | Graphs |
| cycle        | cyc     | Graph cycles |
| hash         | hsh     | Hash tables |

##### Traversal & Pointers
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| breadth      | brd     | BFS traversal |
| current      | cur     | Iterators, pointers |
| level        | lvl     | Tree/Graph levels |
| order        | ord     | Traversal order |
| path         | pth     | Traversal paths |
| previous     | prv     | Linked lists, traversal |
| next         | nxt     | Linked lists, traversal |
| left         | lft     | Trees, two-pointer |
| right        | rht     | Trees, two-pointer |
| parent       | par     | Trees |
| child        | chd     | Trees |
| root         | rot     | Trees |
| head         | hed     | Lists |
| tail         | tal     | Lists |
| pointer      | ptr     | References |
| traverse     | trv     | Graph/Tree traversal |
| iterator     | itr     | Loops |

##### Measurements & Counts
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| length       | len     | Arrays, strings |
| count        | cnt     | Accumulation |
| size         | siz     | Collections |
| height       | hgt     | Trees, matrices |
| width        | wdt     | Matrices, graphics |
| depth        | dep     | Trees, recursion |
| distance     | dst     | Graphs, geometry |
| frequency    | frq     | Counting |
| weight       | wgt     | Graphs, networks |
| capacity     | cap     | Containers |
| threshold    | thr     | Limits |

##### Operations & States
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| create       | crt     | CRUD operations |
| read         | red     | CRUD operations |
| update       | upd     | CRUD operations |
| delete       | del     | CRUD operations |
| query        | qry     | Database |
| insert       | ins     | Collections |
| remove       | rmv     | Collections |
| merge        | mrg     | Combining |
| split        | spl     | Division |
| sort         | srt     | Ordering |
| find         | fnd     | Finding |
| visit        | vst     | Traversal |
| process      | prc     | General |
| validate     | vld     | Checking |
| calculate    | clc     | Math |
| initialize   | ini     | Setup |
| temporary    | tmp     | Storage |
| buffer       | buf     | I/O |
| break        | brk     | |

##### Life Cycle & Status
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| finish       | fin     | Completion states, Terminal markers |
| start        | str     | Initialization states |
| active       | act     | In-progress states |
| pending      | pnd     | Waiting states |

##### Values & Types
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| value        | val     | General |
| number       | num     | Numeric |
| string       | str     | Text |
| character    | chr     | Text |
| boolean      | bln     | Logic |
| integer      | int     | Numeric |
| float        | flt     | Numeric |
| decimal      | dec     | Numeric |
| object       | obj     | General |
| element      | elm     | Collections |
| item         | itm     | Collections |
| key          | key     | Maps |

##### Comparison & Limits
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| maximum      | max     | Comparisons |
| minimum      | min     | Comparisons |
| compare      | cmp     | Comparisons |
| average      | avg     | Statistics |
| total        | tot     | Sums |
| sum          | sum     | Addition |
| product      | prd     | Multiplication |
| quotient     | qot     | Division |
| remainder    | rem     | Modulus |
| equation     | eqn     | Equation |
| range        | rng     | Range |

##### <iscellaneous
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| course       | crs     | Educational |
| letter       | ltr     | |

#### Variable Naming Patterns

##### Compound Names
```go
// Prefer structured compounds
// Prefer structured compounds
maxLen    // maximum length
curIdx    // current index 
nxtPtr    // next pointer
arrLen    // array length
mtxHgt    // matrix height
treNod    // tree node
dfsPth    // DFS path
bfsPth    // BFS path
trvOrd    // traverse order
lvlOrd    // level order
dfsFndPth // DFS find path
bfsFndPth // BFS find path
nxtLvl    // next level
```

##### Collection Pluralization
```go
// Add 's' to acronyms for collections
nums     // numbers
idxs     // indices
vals     // values
nods     // nodes
itms     // items
```

##### Position & Range Variables
```go
// Use consistent position markers
lft, rht // horizontal boundaries, two-pointer technique
top, btm // vertical boundaries
fst, lst // ranges (first/last)
```

###### Common Position Compounds
```go
// Horizontal positioning
lftIdx  // left index
rhtVal  // right value
lftBnd  // left boundary
rhtPtr  // right pointer

// Vertical positioning
topRow  // top row
btmRow  // bottom row
topNod  // top node
btmVal  // bottom value

// Range markers
fstOcc  // first occurrence
lstSeen // last seen
fstIdx  // first index
lstVal  // last value
```

##### Special Cases
1. Preserve common three-letter words:
   ```go
   top, sum, map, set // Keep as-is
   ```

2. Iterator variables:
   ```go
   for idx := range arr {
   for row := 0; row < rows; row++ {
   for key := range map {
   ```

3. Error handling:
   ```go
   if err := doSomething(); err != nil {
   ```

##### Anti-patterns
```go
// Avoid these patterns
currentVal   // Use curVal instead
arrayLength  // Use arrLen instead
nodePointer  // Use nodPtr instead
```


# Graph Problem Standards

## Core Type Definitions
```go
// Graph representation using adjacency list
type Graph [][]int

// Standard states for graph traversal
const (
    unvisited = 0 // Node not yet processed
    visiting  = 1 // Node being processed in current DFS path
    visited   = 2 // Node fully processed
)
```

## Algorithm Selection Guidelines

### DFS vs BFS Preference
Always prefer DFS over BFS unless specific requirements necessitate BFS:

#### Why DFS First
1. Simpler implementation
   - No queue management
   - Natural recursive structure
   - Cleaner state transitions
   - Less memory overhead

2. Better for common graph problems
   - Cycle detection
   - Topological sorting
   - Path finding with backtracking
   - Tree-like structure exploration
   - Connected component analysis

3. Memory efficiency
   - O(h) stack space vs O(w) queue space
   - Better locality of reference
   - More cache-friendly access patterns

#### When to Use BFS
Only use BFS when problem specifically requires:
1. Shortest path in unweighted graphs
2. Level-order traversal
3. Minimum steps to target
4. Breadth-based distance metrics

## Implementation Guidelines

### Initialization Standards
```go
// CORRECT:
gph := make(Graph, numVertices)

// INCORRECT:
gph := make([][]int, numVertices)  // Don't use raw slice type
```

### Graph Function Patterns
Use format: `[algorithm][Action][Entity]`
```go
dfsFndCyc   // DFS Find Cycle
bfsFndPth   // BFS Find Path
dfsTrvOrd   // DFS Traverse Order
```

### Parameter Naming Standards
```go
// CORRECT:
func dfsFndCyc(gph Graph, nod int) bool

// INCORRECT:
func dfsFindCycle(adj [][]int, node int) bool
```

### Required Problem Categories
Graph type MUST be used for:
1. Course scheduling/prerequisites
2. Dependency resolution
3. Cycle detection
4. Path finding
5. Connected components
6. Topological sorting

### Implementation Checklist
- [ ] Graph type used for adjacency list
- [ ] Standard state constants defined
- [ ] Function names follow prefix conventions
- [ ] Graph initialized using Graph type
- [ ] Parameter types use Graph instead of [][]int

## Standard Algorithm Templates

### Primary DFS Template (Preferred)
```go
var dfsFnd func(int) bool
dfsFnd = func(nod int) bool {
    switch vst[nod] {
    case visiting:
        return true
    case visited:
        return false
    }
    
    vst[nod] = visiting
    
    for _, nxt := range gph[nod] {
        if dfsFnd(nxt) {
            return true
        }
    }
    
    vst[nod] = visited
    return false
}
```

### Secondary BFS Template (When Required)
```go
que := make([]int, 0, n)
for len(que) > 0 {
    cur := que[0]
    que = que[1:]
    
    for _, nxt := range gph[cur] {
        que = append(que, nxt)
    }
}
```

## State Management

### DFS State Tracking (Preferred)
```go
// Single array for complete state
const (
    unvisited = iota // 0
    visiting         // 1
    visited          // 2
)
vst := make([]int, n)
```

### BFS State Tracking (When Required)
```go
// Typically needs multiple arrays
vst := make([]bool, n)
dst := make([]int, n)
par := make([]int, n)
```

## Common Usage Patterns

### Graph Building
```go
// Building a graph
gph := make(Graph, n)
for idx := range n {
    gph[idx] = make([]int, 0)
}

// Edge direction for dependencies (A requires B)
gph[B] = append(gph[B], A)  // Edge from prerequisite to dependent

// Edge direction for connections (A connects to B)
gph[A] = append(gph[A], B)  // Edge from source to destination

// State management
vst := make([]int, n)  // Single array for state tracking
```

## Testing Requirements

### Base Cases
- Single node
- Simple chain
- Multiple paths

### Edge Cases
- Empty graph
- Isolated nodes
- Maximum constraints

### Error Cases
- Cycles
- Invalid edges
- Constraint violations

### Testing Strategy
1. Primary test cases using DFS
2. Comparative tests with BFS when applicable
3. Verification of space complexity
4. State transition validation
5. Edge case coverage for both approaches

## Performance Guidelines
1. Pre-allocate with capacity hints
2. Use switch for state transitions
3. Avoid unnecessary graph copying
4. Consider edge direction for traversal efficiency
5. Prefer single state array over multiple tracking arrays

## Documentation Requirements
1. Edge direction explanation
2. State transition semantics
3. Return value interpretation
4. Time/space complexity
5. Order guarantees (if applicable)
6. Justification if BFS is chosen over DFS
7. Analysis of space complexity trade-offs
8. Explanation of state transition choices
9. Rationale for traversal order selection

## Implementation Priority
1. Always start with DFS solution
2. Only switch to BFS if:
   - Problem explicitly requires level-order
   - Shortest path is needed
   - Memory profiling shows BFS advantage
   - Requirements specify breadth-first order

## Common Pitfalls
1. Incorrect edge direction in dependency graphs
2. Using multiple boolean arrays instead of state constants
3. Non-standard function naming
4. Missing cycle detection
5. Improper state transitions
6. Unnecessary use of BFS when DFS would suffice
7. Complex queue management in BFS
8. Inefficient state tracking
9. Mixing traversal strategies without clear justification

### Solution Code Presentation
#### Core Principles
- Present solution code in expandable document format
- Keep main discussion focused on explanation and intuition
- Avoid mixing lengthy code blocks with conceptual discussion

#### Solution Format
```markdown
<details>
<summary>Click to see solution</summary>

[Code block with complete solution]

</details>
```

#### Standard Structure
1. Present intuition and approach
2. Show key visual diagrams if needed
3. Place complete solution code in expandable section
4. Continue discussion of implementation details

#### When to Use
- Main solution implementation
- Alternative solutions
- Test cases and benchmarks
- Long code examples
- Complete working implementations

#### When Not to Use
- Short code snippets (< 10 lines)
- Individual functions being discussed
- Example usage
- Small code modifications
- Quick demonstrations

#### Example Usage
```markdown
Let's solve this using a depth-first search approach...
[explanation of approach]

<details>
<summary>Click to see solution</summary>

```go
func solution() {
    // Complete implementation
}
```

</details>

Now let's discuss the key aspects of this implementation...
```
