### Solution Development
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
#### Acronym Usage
- Use three-letter acronyms whenever possible
- Prefer consonants
- Choose the first three characters of a word when they reflect its meaning
- Ensure acronyms maintain a close relation to the word's meaning and phonetic sound

#### Standard Acronyms
| Full Word    | Acronym | Usage Context |
|--------------|---------|---------------|
| current      | cur     | Iterators, pointers |
| length       | len     | Arrays, strings |
| right        | rht     | Trees, pointers |
| left         | lft     | Trees, pointers |
| previous     | prv     | Linked lists |
| next         | nxt     | Linked lists |
| word         | wrd     | Strings |
| count        | cnt     | Accumulation |
| index        | idx     | Arrays, loops |
| matrix       | mtx     | 2D arrays |
| first        | fst     | General |
| last         | lst     | General |
| char         | chr     | Characters |
| old          | prv     | General |
| new          | nxt     | General |
| result       | res     | Return values |
| response     | res     | APIs |
| request      | req     | APIs |
| merge        | mrg     | Combining |
| original     | org     | Source data |
| copy         | cpy     | Duplicates |
| node         | nod     | Trees, graphs |
| head         | hed     | Lists |
| source       | src     | Source data |
| destination  | dst     | Target data |
| number       | num     | Numeric data |
| product      | prd     | Multiplication |
| maximum      | max     | Comparisons |
| minimum      | min     | Comparisons |
| total        | tot     | Sums |
| bottom       | btm     | Position |
| error        | err     | Error handling |
| haystack     | hay     | Search context |
| needle       | ndl     | Search target |
| frequency    | frq     | Counting |
| item         | itm     | General |
| value        | val     | General |
| user         | usr     | General |
| create       | crt     | Data operations (DB, File, Cache, API) |
| insert       | ins     | Data operations (DB, File, Cache, API) |
| update       | upd     | Data operations (DB, File, Cache, API) |
| delete       | del     | Data operations (DB, File, Cache, API) |

#### Compound Names
- Join TLAs for compound concepts
  * `maxLen` (maximum length)
  * `curIdx` (current index)
  * `nxtPtr` (next pointer)
- Maintain consistency in ordering

#### Special Techniques
- In the two-pointer technique, use `lft` for the left pointer and `rht` for the right pointer

#### Preferred Naming Conventions
- Use `idx` or `n` as index identifiers
- Pluralize three-letter acronyms by adding an `s` (e.g., `num` becomes `nums`, `idx` becomes `idxs`)
- Prefer compound words with three-letter acronyms over reducing compound words to three letters (e.g., `currentLength = curLen` instead of further abbreviating)
- Use three-letter acronyms within compound acronyms (e.g., `maximumRight = maxRht`)
- Use similar naming conventions in `for` variables. For example, `for row := 0; row < rows; row++ {` instead of `for row := 0; row < mtxH; row++ {`

#### Exceptions
- Preserve non-acronym three-letter words. For example, `top = top`