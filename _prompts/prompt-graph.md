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
