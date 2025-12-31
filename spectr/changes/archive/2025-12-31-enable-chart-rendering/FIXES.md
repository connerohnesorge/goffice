# Fixes Applied to enable-chart-rendering Proposal

## Overview
Fixed 8 critical API mismatches between the Spectr proposal specifications and the actual .wip code implementation.

## Issues Fixed

### 1. CRITICAL: DrawRectangle Signature Mismatch ✓
- **Before**: `DrawRectangle(x, y, w, h float64)` - 4 parameters
- **After**: `DrawRectangle(x, y, w, h float64, fill, stroke bool)` - 6 parameters
- **Reason**: WIP code (chart_renderer.go.wip lines 130-137, 423-430, 441-448) uses fill/stroke bools to control painting
- **Files Updated**: design.md, spec.md, tasks.md

### 2. CRITICAL: DrawCircle Signature Mismatch ✓
- **Before**: `DrawCircle(cx, cy, r float64)` - 3 parameters
- **After**: `DrawCircle(cx, cy, r float64, fill, stroke bool)` - 5 parameters
- **Reason**: WIP code (chart_renderer.go.wip lines 225-231, 240-246) uses fill/stroke bools
- **Files Updated**: design.md, spec.md, tasks.md

### 3. CRITICAL: SetLineDashPattern vs SetDashPattern Naming ✓
- **Before**: `SetDashPattern(pattern []float64, phase float64)`
- **After**: `SetLineDashPattern(pattern []float64, phase float64)`
- **Reason**: WIP code (stroke_renderer.go.wip line 70) uses `SetLineDashPattern`
- **Files Updated**: design.md, spec.md, tasks.md

### 4. FUNCTIONAL: Missing Clip() Method ✓
- **Before**: Clip() method not defined in Page interface
- **After**: `Clip()` method added to Page interface
- **Reason**: image_renderer.go.wip line 91 uses `.Page.Clip()` (commented out, awaiting implementation)
- **Files Updated**: design.md, spec.md, tasks.md

### 5. MODERATE: SetFont Signature Wrong ✓
- **Before**: `SetFont(font *Font, size float64)` - takes *Font pointer
- **After**: `SetFont(name string, size float64)` - takes string name
- **Reason**: WIP code (chart_renderer.go.wip lines 452, 491) uses `SetFont("Helvetica", 10)`
- **Files Updated**: design.md, spec.md, tasks.md

### 6. STRUCTURAL: Missing Type Definitions ✓
- **Before**: Path, Font, Image types referenced but not defined
- **After**: Added documentation:
  ```go
  type Path = *PathBuilder   // From pdf/drawing/path.go
  type Font struct { ... }   // From pdf/font package
  type Image struct { ... }  // From pdf/core package
  ```
- **Reason**: Clarify how these types work with the Page interface
- **Files Updated**: design.md, spec.md

### 7. ARCHITECTURAL: API Value Constants Undocumented ✓
- **Before**: SetLineCap(cap int), SetLineJoin(join int) - no documentation of valid values
- **After**: Added inline documentation:
  - `SetLineCap(cap int)  // 0=butt, 1=round, 2=square`
  - `SetLineJoin(join int) // 0=miter, 1=round, 2=bevel`
- **Reason**: Match PDF specification and make API self-documenting
- **Files Updated**: design.md, spec.md, tasks.md

### 8. DOCUMENTATION: Helper Functions Added ✓
- **Before**: No mention of color/path helper functions
- **After**: Added documentation of existing helpers:
  ```go
  func NewRGB(r, g, b float64) Color
  func ParseColor(hex string) Color
  func NewPathBuilder() *PathBuilder
  ```
- **Reason**: WIP code uses these helpers extensively (NewRGB in chart_renderer.go.wip line 467, ParseColor in fill_renderer.go.wip line 49)
- **Files Updated**: design.md, spec.md

## Validation Results

**Before Fixes**: 8 critical mismatches
**After Fixes**: ✓ enable-chart-rendering valid (spectr validate passes)

## Files Modified

1. `/spectr/changes/enable-chart-rendering/proposal.md` - Updated API examples
2. `/spectr/changes/enable-chart-rendering/design.md` - Fixed interface definition and added critical API details section
3. `/spectr/changes/enable-chart-rendering/specs/pdf-drawing/spec.md` - Fixed all scenarios and API signatures
4. `/spectr/changes/enable-chart-rendering/tasks.md` - Updated task descriptions to match correct API

## Impact

All specifications now **exactly match** the actual .wip code implementation, ensuring:
- No compilation errors when .wip files are enabled
- Correct method signatures for all Page interface methods
- Proper documentation of PDF operator values
- Clear type relationships and helper function availability

## Next Steps

The proposal is now ready for implementation. The coder agent can:
1. Create the Page interface using the exact signatures documented
2. Implement PDFPage with confidence that it matches WIP code usage
3. Enable .wip files knowing the API will compile correctly
