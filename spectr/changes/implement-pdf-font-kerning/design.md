# Proposal: Implement Full Class-Based Font Kerning Support

## Summary

Implement complete class-based kerning support for TrueType fonts in the PDF rendering engine.

## Background

The PDF font handling code (`pdf/font/truetype.go`) has a TODO indicating that full class-based kerning support needs to be implemented:

**Reference TODO:** `pdf/font/truetype.go:1386`
```go
// Format 2 is more complex and less common
// For now, we skip detailed parsing
// TODO: Implement full class-based kerning support
```

## Motivation

Kerning improves text appearance by adjusting space between specific character pairs. Class-based kerning:
- Groups characters into classes (e.g., all uppercase letters, all vowels)
- Defines kerning between classes rather than individual pairs
- More efficient than pair-based kerning for large fonts
- Used by professional typefaces

## Technical Design

### TrueType Kerning Formats

TrueType fonts support multiple kerning formats:
- **Format 0**: Simple pair-based kerning (already implemented)
- **Format 2**: Class-based kerning (needs implementation)
- **Format 3**: Additional features (not priority)

### Class-Based Kerning Structure

Format 2 kerning subtable structure:
```
Kerning Subtable (Format 2):
- Coverage: Which glyphs are covered
- Class Definition Tables:
  - Left Class: Maps glyphs to left classes
  - Right Class: Maps glyphs to right classes
- Class Kerning Array: Kerning values for class pairs
```

### Implementation

```go
// ClassBasedKerning represents Format 2 kerning data
type ClassBasedKerning struct {
    LeftClassDef  ClassDefTable  // Maps glyphs to left classes
    RightClassDef ClassDefTable  // Maps glyphs to right classes
    KerningArray  [][]int16      // Kerning values [leftClass][rightClass]
}

// ClassDefTable maps glyph IDs to class IDs
type ClassDefTable struct {
    Format      uint16
    Classes     map[uint16]uint16  // glyphID -> classID
    DefaultClass uint16
}

func (f *Font) parseKerningFormat2(data []byte, offset uint32) (*ClassBasedKerning, error) {
    r := bytes.NewReader(data[offset:])
    
    // Read header
    var leftClassOffset, rightClassOffset, kerningArrayOffset uint16
    binary.Read(r, binary.BigEndian, &leftClassOffset)
    binary.Read(r, binary.BigEndian, &rightClassOffset)
    binary.Read(r, binary.BigEndian, &kerningArrayOffset)
    
    // Parse left class definition table
    leftClass, err := f.parseClassDefTable(data, offset+uint32(leftClassOffset))
    if err != nil {
        return nil, fmt.Errorf("failed to parse left class: %w", err)
    }
    
    // Parse right class definition table
    rightClass, err := f.parseClassDefTable(data, offset+uint32(rightClassOffset))
    if err != nil {
        return nil, fmt.Errorf("failed to parse right class: %w", err)
    }
    
    // Parse kerning array
    kerningArray, err := f.parseKerningArray(data, offset+uint32(kerningArrayOffset))
    if err != nil {
        return nil, fmt.Errorf("failed to parse kerning array: %w", err)
    }
    
    return &ClassBasedKerning{
        LeftClassDef:  leftClass,
        RightClassDef: rightClass,
        KerningArray:  kerningArray,
    }, nil
}

func (f *Font) getClassBasedKerning(left, right rune) int16 {
    if f.ClassKerning == nil {
        return 0
    }
    
    // Get glyph IDs
    leftGlyph := f.GetGlyphID(left)
    rightGlyph := f.GetGlyphID(right)
    
    // Get class IDs
    leftClass := f.ClassKerning.LeftClassDef.GetClass(leftGlyph)
    rightClass := f.ClassKerning.RightClassDef.GetClass(rightGlyph)
    
    // Look up kerning
    return f.ClassKerning.KerningArray[leftClass][rightClass]
}
```

### Integration with Existing Kerning

```go
func (f *Font) GetKerningForRunes(left, right rune) int16 {
    // Try pair-based kerning first (Format 0)
    if kerning := f.getPairKerning(left, right); kerning != 0 {
        return kerning
    }
    
    // Fall back to class-based kerning (Format 2)
    return f.getClassBasedKerning(left, right)
}
```

## Requirements

### SHALL Requirements

#### Requirement: Format 2 Parsing
The font parser SHALL correctly parse TrueType kerning format 2 subtables.

##### Scenario: Parse Class Definition Table
Given a TrueType font with class-based kerning
When the font is loaded
Then the class definition tables SHALL be parsed correctly

##### Scenario: Parse Kerning Array
Given a font with class-based kerning
When the font is loaded
Then the kerning values array SHALL be parsed correctly

#### Requirement: Class-Based Kerning Lookup
The font SHALL return correct kerning values using class-based lookup.

##### Scenario: Same Class Pair
Given two characters in the same class
When kerning is requested
Then the correct class-pair kerning SHALL be returned

##### Scenario: Different Class Pair
Given two characters in different classes
When kerning is requested
Then the correct class-pair kerning SHALL be returned

#### Requirement: Fallback to Pair Kerning
The implementation SHALL check pair-based kerning before class-based kerning.

##### Scenario: Pair Override
Given a font with both pair and class kerning
When a specific pair has kerning defined
Then the pair kerning SHALL take precedence

### SHOULD Requirements

#### Requirement: Performance
Class-based kerning lookup SHOULD be efficient (O(1) after initial class lookup).

#### Requirement: Memory Efficiency
The implementation SHOULD use memory-efficient data structures for class tables.

## Testing Strategy

### Unit Tests
- Test format 2 subtable parsing
- Test class definition table parsing
- Test kerning array parsing
- Test class-based lookup
- Test integration with pair-based kerning

### Integration Tests
- Test with real fonts using class-based kerning
- Compare kerning output with expected values

### Visual Tests
- Compare text rendering with/without class kerning
- Verify improved spacing in kerned text

## Implementation Plan

1. Define ClassBasedKerning data structures
2. Implement class definition table parsing
3. Implement kerning array parsing
4. Implement class-based kerning lookup
5. Integrate with existing kerning system
6. Add unit tests
7. Test with real fonts

## Related Changes

- `pdf/font/truetype.go` - Main implementation
- `pdf/font/kerning.go` - May split out kerning-specific code

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Complex table formats | Medium | Thorough testing; reference implementation |
| Font variations | Low | Test with diverse font files |
| Performance impact | Low | Optimize lookup tables |

## Acceptance Criteria

- [ ] Format 2 subtables parsed correctly
- [ ] Class definition tables work correctly
- [ ] Kerning values returned correctly
- [ ] Integration with pair kerning works
- [ ] Unit tests pass with >90% coverage
- [ ] Works with real fonts using class kerning
