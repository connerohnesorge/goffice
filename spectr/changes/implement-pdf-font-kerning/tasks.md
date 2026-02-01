# Tasks: Implement Full Class-Based Font Kerning Support

## Data Structures

- [ ] 1.1 Define ClassBasedKerning struct
- [ ] 1.2 Define ClassDefTable struct
- [ ] 1.3 Define ClassDefFormat1 struct (glyph-by-glyph)
- [ ] 1.4 Define ClassDefFormat2 struct (range-based)
- [ ] 1.5 Add ClassBasedKerning field to Font struct

## Class Definition Table Parsing

- [ ] 2.1 Implement parseClassDefTable method
- [ ] 2.2 Implement parseClassDefFormat1 (glyph array)
- [ ] 2.3 Implement parseClassDefFormat2 (range records)
- [ ] 2.4 Add GetClass method for glyph-to-class lookup
- [ ] 2.5 Handle default class assignment

## Kerning Array Parsing

- [ ] 3.1 Implement parseKerningArray method
- [ ] 3.2 Parse array dimensions (left/right class counts)
- [ ] 3.3 Parse 2D kerning value array
- [ ] 3.4 Handle signed 16-bit kerning values
- [ ] 3.5 Validate array bounds

## Format 2 Subtable Parsing

- [ ] 4.1 Implement parseKerningFormat2 method
- [ ] 4.2 Parse subtable header (offsets)
- [ ] 4.3 Call class definition parsing
- [ ] 4.4 Call kerning array parsing
- [ ] 4.5 Integrate into main kerning table parsing

## Kerning Lookup

- [ ] 5.1 Implement getClassBasedKerning method
- [ ] 5.2 Convert runes to glyph IDs
- [ ] 5.3 Look up left class
- [ ] 5.4 Look up right class
- [ ] 5.5 Retrieve kerning from array
- [ ] 5.6 Handle missing class assignments

## Integration

- [ ] 6.1 Update GetKerningForRunes to check class-based kerning
- [ ] 6.2 Ensure pair kerning takes precedence
- [ ] 6.3 Handle fonts with multiple kerning subtables
- [ ] 6.4 Update font loading to detect format 2

## Unit Tests

- [ ] 7.1 Create test: Parse class definition format 1
- [ ] 7.2 Create test: Parse class definition format 2
- [ ] 7.3 Create test: Parse kerning array
- [ ] 7.4 Create test: Class-based lookup (same class)
- [ ] 7.5 Create test: Class-based lookup (different classes)
- [ ] 7.6 Create test: Integration with pair kerning
- [ ] 7.7 Create test: Unknown glyph handling
- [ ] 7.8 Create test: Default class handling

## Integration Tests

- [ ] 8.1 Test with font using class-based kerning
- [ ] 8.2 Verify kerning values match expected
- [ ] 8.3 Test with font using both pair and class kerning
- [ ] 8.4 Test with font having no class kerning

## Test Fonts

- [ ] 9.1 Identify fonts with class-based kerning
- [ ] 9.2 Create minimal test font with known class kerning
- [ ] 9.3 Document font characteristics for testing

## Visual Tests

- [ ] 10.1 Compare text rendering with/without class kerning
- [ ] 10.2 Verify improved letter spacing
- [ ] 10.3 Test with various text samples

## Documentation

- [ ] 11.1 Update pdf/font/AGENTS.md with kerning details
- [ ] 11.2 Document class-based kerning algorithm
- [ ] 11.3 Add code comments to parsing methods
- [ ] 11.4 Document supported kerning formats

## Verification

- [ ] 12.1 Run all font tests - ensure no regressions
- [ ] 12.2 Verify test coverage >90% for changed code
- [ ] 12.3 Test with real-world documents
- [ ] 12.4 Performance benchmark shows acceptable lookup speed
