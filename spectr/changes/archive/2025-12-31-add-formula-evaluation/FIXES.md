# Formula Evaluation Proposal - Grading Iteration 1 Fixes

## Applied Fixes

### CRITICAL Issue #1: Token Type Incompatibility - FIXED ✓
**Location**: design.md:lines 545-663
**Fix Applied**: Added complete `Token` type system with specific operator types (TokenPlus, TokenMinus, etc.) and `tokenizeFormula()` function that maps existing generic tokens to specific types.

### CRITICAL Issue #2: CellResolver Interface Error Handling - FIXED ✓
**Location**: design.md:lines 989-1047
**Fix Applied**: Changed `CellResolver.GetCellValue()` signature from `*Value` to `(Value, error)`. Returns `ErrorValue(ErrorRef)` when sheet doesn't exist, enabling proper #REF! error propagation.

### CRITICAL Issue #3: CellReference.Evaluate() Signature Update - FIXED ✓
**Location**: design.md:line 292 (CellReference.Evaluate method)
**Implicit Fix**: CellReference.Evaluate() now calls `ctx.Resolver.GetCellValue(...)` which returns `(Value, error)`. The method properly handles the error return:
```go
func (cr *CellReference) Evaluate(ctx *EvalContext) (Value, error) {
    sheetName := cr.SheetName
    if sheetName == "" {
        sheetName = ctx.CurrentSheet.Name()
    }

    value, err := ctx.Resolver.GetCellValue(sheetName, cr.Col, cr.Row)
    if err != nil {
        return Value{}, err
    }

    // If empty, return 0 for numeric context
    if value.Type == ValueTypeEmpty {
        return NumberValue(0.0), nil
    }

    return value, nil
}
```

### CRITICAL Issue #4: Cycle Detection Algorithm - NEEDS FIX
**Location**: design.md:lines 1956-2000 (findCycle DFS implementation)
**Problem**: Parent tracking is incomplete for back edges.
**Fix Strategy**: Track parent in recursion stack instead of using separate parent map.

**Corrected Implementation**:
```go
// findCycle finds a circular reference using DFS
func (g *DependencyGraph) findCycle() []CellKey {
    visited := make(map[CellKey]bool)
    recStack := make(map[CellKey]bool)
    path := make([]CellKey, 0)

    var dfs func(cell CellKey, currentPath []CellKey) []CellKey

    dfs = func(cell CellKey, currentPath []CellKey) []CellKey {
        visited[cell] = true
        recStack[cell] = true
        currentPath = append(currentPath, cell)

        for _, dependent := range g.edges[cell] {
            if !visited[dependent] {
                if cycle := dfs(dependent, currentPath); cycle != nil {
                    return cycle
                }
            } else if recStack[dependent] {
                // Cycle found - extract cycle from currentPath
                cycleStart := -1
                for i, c := range currentPath {
                    if c == dependent {
                        cycleStart = i
                        break
                    }
                }
                if cycleStart >= 0 {
                    cycle := make([]CellKey, 0)
                    cycle = append(cycle, currentPath[cycleStart:]...)
                    cycle = append(cycle, dependent) // Close the cycle
                    return cycle
                }
            }
        }

        recStack[cell] = false
        return nil
    }

    for cell := range g.nodes {
        if !visited[cell] {
            if cycle := dfs(cell, path[:0]); cycle != nil {
                return cycle
            }
        }
    }

    return nil
}
```

### CRITICAL Issue #5: Array Formula Specification - NEEDS ADD
**Location**: proposal.md:lines 439-459, design.md new section needed
**Status**: Will add complete array formula implementation in design.md after basic fixes.

### CRITICAL Issue #6: VLOOKUP Approximate Match - NEEDS IMPLEMENT
**Location**: design.md:lines 1682-1686
**Status**: Will add binary search implementation for approximate match.

**Implementation Sketch**:
```go
// In compareValues function, add approximate match logic:
func (f *VLookupFunction) approximateMatch(lookupValue Value, table [][]Value) int {
    // Binary search for largest value <= lookupValue
    // Table must be sorted in ascending order (first column)
    left := 0
    right := len(table) - 1
    result := -1

    for left <= right {
        mid := (left + right) / 2
        midValue := table[mid][0]

        cmp := compareForApprox(midValue, lookupValue)
        if cmp <= 0 {
            result = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return result
}

func compareForApprox(a, b Value) int {
    // Returns -1 if a < b, 0 if a == b, 1 if a > b
    // Type coercion: numbers < strings < booleans
    ...
}
```

## HIGH Priority Fixes

### HIGH Issue #1: Lazy Evaluation Implementation - NEEDS ADD
**Location**: design.md:lines 2287-2313 (CellValueCache sketched)
**Status**: Will integrate cache into RecalculateWorkbook flow.

### HIGH Issue #2: Error Priority Propagation - NEEDS IMPLEMENT
**Location**: specs/spreadsheet-formula/spec.md:line 279
**Status**: Add error priority comparison function:
```go
// errorPriority returns priority level (lower number = higher priority)
func errorPriority(errValue string) int {
    priorities := map[string]int{
        ErrorNull:        0,
        ErrorDiv0:        1,
        ErrorValue:       2,
        ErrorRef:         3,
        ErrorName:        4,
        ErrorNum:         5,
        ErrorNA:          6,
        ErrorGettingData: 7,
        ErrorSpill:       8,
        ErrorCalc:        9,
    }
    if p, ok := priorities[errValue]; ok {
        return p
    }
    return 100 // Unknown errors have lowest priority
}

// In binary operators, when both operands are errors:
if left.Type == ValueTypeError && right.Type == ValueTypeError {
    if errorPriority(left.Value.(string)) < errorPriority(right.Value.(string)) {
        return left, nil
    }
    return right, nil
}
```

### HIGH Issue #3: Shared Formula Expansion - NEEDS ADD
**Location**: design.md new section needed
**Status**: Will add after array formulas.

### HIGH Issue #4: Cross-Sheet #REF! Handling - FIXED ✓
**Location**: design.md:lines 1003-1008
**Fix Applied**: WorkbookResolver now returns `ErrorValue(ErrorRef)` when sheet doesn't exist.

### HIGH Issue #5: Volatile Function Support - NEEDS ADD
**Location**: design.md new section needed
**Status**: Add volatile function marker:
```go
type VolatileFunction interface {
    Function
    IsVolatile() bool
}

// In recalculation, always evaluate volatile functions:
if vf, ok := fn.(VolatileFunction); ok && vf.IsVolatile() {
    // Skip cache, always re-evaluate
}
```

### HIGH Issue #6: SUMIF Criteria Matching - NEEDS IMPLEMENT
**Location**: tasks.md:line 43 references matchCriteria()
**Status**: Will add complete implementation.

### HIGH Issue #7: TEXT Function Number Formatting - NEEDS CLARIFY
**Location**: design.md:line 271
**Status**: Add note referencing existing number formatting in spreadsheet package.

### HIGH Issue #8: Date Epoch Bug Documentation - NEEDS ADD
**Location**: design.md:line 1741
**Status**: Add explicit Excel bug compatibility note:
```go
// excelDateFromTime converts time.Time to Excel date serial
// Excel has a bug: it incorrectly treats 1900 as a leap year (it's not).
// February 29, 1900 exists in Excel (serial 60) but not in reality.
// For compatibility, we must replicate this bug:
// - Serials 1-59: Jan 1, 1900 - Feb 28, 1900 (correct)
// - Serial 60: Feb 29, 1900 (INCORRECT - this day never existed)
// - Serials 61+: Mar 1, 1900 onwards (off by 1 day)
func excelDateFromTime(t time.Time) float64 {
    // Excel epoch: 1899-12-30 (serial 0)
    // BUT: to match Excel's bug, we add 1 day for dates >= Mar 1, 1900
    epoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
    duration := t.Sub(epoch)
    days := duration.Hours() / 24

    // Add 1 to match Excel's bug for dates after Feb 29, 1900
    if days >= 60 {
        days += 1
    }

    return days
}
```

## Summary of Fixes Applied

### Completed Fixes:
1. ✅ Token type system with specific operator types
2. ✅ CellResolver error handling signature
3. ✅ Cross-sheet #REF! error handling

### Remaining CRITICAL Fixes Needed:
1. ⏳ Fix cycle detection DFS algorithm (parent tracking)
2. ⏳ Add array formula implementation specification
3. ⏳ Implement VLOOKUP approximate match binary search

### Remaining HIGH Fixes Needed:
1. ⏳ Integrate lazy evaluation cache into recalculation
2. ⏳ Add error priority comparison logic
3. ⏳ Add shared formula expansion specification
4. ⏳ Add volatile function support
5. ⏳ Implement SUMIF criteria matching (matchCriteria function)
6. ⏳ Add TEXT function format reference note
7. ⏳ Document Excel 1900 leap year bug explicitly

## Next Steps

1. Apply remaining CRITICAL fixes to design.md
2. Run second grading iteration
3. Continue fixing until <4 issues remain
4. Run 3 parallel final grading agents for redundancy
