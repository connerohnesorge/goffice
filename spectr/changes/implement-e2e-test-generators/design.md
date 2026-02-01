# Proposal: Implement E2E Test Generator Suite

## Summary

Implement comprehensive E2E test generators for background application, image generation, table generation, and baseline generation mode.

## Background

The E2E test framework has several TODOs indicating missing test generation capabilities:

**Reference TODOs:**
1. `tests/e2e/generators/go/main.go:265` - "TODO: Implement background application"
2. `tests/e2e/generators/go/main.go:457` - "TODO: Implement image generation (future phase)"
3. `tests/e2e/generators/go/main.go:474` - "TODO: Implement table generation (future phase)"
4. `tests/e2e/cmd/e2e-runner/main.go:52` - "TODO: Implement baseline generation mode when --generate-baselines is set"

## Motivation

Comprehensive E2E testing requires test documents covering:
- Background fills, colors, and patterns (for PDF rendering validation)
- Images of various types, sizes, and placements
- Tables with different structures and formatting
- Automated baseline generation for regression testing

## Technical Design

### 1. Background Application Generator

**Purpose:** Generate documents with various background types for PDF rendering tests.

**Background Types to Support:**
- Solid color fills
- Gradient fills (linear, radial)
- Pattern fills
- Picture/texture fills

**Implementation:**
```go
func generateBackgroundTests() []TestDocument {
    var tests []TestDocument
    
    // Solid colors
    tests = append(tests, generateSolidBackgroundTest("red", "FF0000"))
    tests = append(tests, generateSolidBackgroundTest("blue", "0000FF"))
    tests = append(tests, generateSolidBackgroundTest("gradient_bg", "linear_gradient"))
    
    // Gradients
    tests = append(tests, generateGradientBackgroundTest("linear_h", LinearHorizontal))
    tests = append(tests, generateGradientBackgroundTest("linear_v", LinearVertical))
    tests = append(tests, generateGradientBackgroundTest("radial", Radial))
    
    // Patterns
    tests = append(tests, generatePatternBackgroundTest("diagonal_stripes"))
    tests = append(tests, generatePatternBackgroundTest("crosshatch"))
    
    return tests
}
```

### 2. Image Generation

**Purpose:** Create test documents with various image configurations.

**Image Types to Support:**
- PNG (with and without transparency)
- JPEG (various quality levels)
- GIF (animated and static)
- BMP
- TIFF

**Configurations:**
- Various sizes (small icon to full page)
- Inline vs floating
- Different anchor types
- Behind text vs in front of text

**Implementation:**
```go
func generateImageTests() []TestDocument {
    var tests []TestDocument
    
    // By format
    tests = append(tests, generateImageTest("png_transparent", PNG, WithTransparency))
    tests = append(tests, generateImageTest("jpeg_photo", JPEG, HighQuality))
    tests = append(tests, generateImageTest("gif_static", GIF, Static))
    
    // By size
    tests = append(tests, generateImageTest("small_icon", PNG, Size(64, 64)))
    tests = append(tests, generateImageTest("medium_image", JPEG, Size(400, 300)))
    tests = append(tests, generateImageTest("large_photo", JPEG, Size(1600, 1200)))
    
    // By placement
    tests = append(tests, generateImageTest("inline", PNG, Inline))
    tests = append(tests, generateImageTest("floating", PNG, Floating))
    
    return tests
}
```

### 3. Table Generation

**Purpose:** Generate documents with tables for PDF table rendering validation.

**Table Types to Support:**
- Simple tables (uniform rows/columns)
- Complex tables (merged cells)
- Tables with formatting (borders, shading)
- Nested tables
- Tables with images

**Configurations:**
- Various sizes (2x2 to 10x10)
- Different border styles
- Cell shading/colors
- Header rows
- Merged cells (horizontal, vertical)

**Implementation:**
```go
func generateTableTests() []TestDocument {
    var tests []TestDocument
    
    // Simple tables
    tests = append(tests, generateTableTest("simple_2x2", 2, 2))
    tests = append(tests, generateTableTest("simple_5x5", 5, 5))
    
    // With formatting
    tests = append(tests, generateTableTest("with_borders", 3, 3, Borders(Single)))
    tests = append(tests, generateTableTest("with_shading", 3, 3, Shading(Gray)))
    
    // Complex
    tests = append(tests, generateTableTest("merged_cells", 4, 4, 
        HorizontalMerge(0, 0, 2),
        VerticalMerge(2, 0, 2)))
    
    return tests
}
```

### 4. Baseline Generation Mode

**Purpose:** Enable automatic generation of baseline/reference files for regression testing.

**Command-line Flag:**
```bash
e2e-runner --generate-baselines
```

**Behavior:**
- Run all test generators
- Generate PDF output for each test
- Save to `testdata/baselines/` directory
- Skip comparison phase

**Implementation:**
```go
func main() {
    generateBaselines := flag.Bool("generate-baselines", false, "Generate baseline files")
    flag.Parse()
    
    if *generateBaselines {
        runBaselineGeneration()
    } else {
        runComparisonTests()
    }
}

func runBaselineGeneration() {
    tests := collectAllTests()
    for _, test := range tests {
        doc := generateTestDocument(test)
        pdf := renderToPDF(doc)
        saveBaseline(test.Name, pdf)
    }
}
```

## Requirements

### SHALL Requirements

#### Requirement: Background Application Tests
The E2E generator SHALL create documents with various background types.

##### Scenario: Solid Color Background
Given a request for solid color background tests
When the generator runs
Then it SHALL create documents with solid fill backgrounds

##### Scenario: Gradient Background
Given a request for gradient background tests
When the generator runs
Then it SHALL create documents with linear and radial gradients

#### Requirement: Image Generation Tests
The E2E generator SHALL create documents with various image configurations.

##### Scenario: PNG Image Test
Given a request for PNG image tests
When the generator runs
Then it SHALL create documents with PNG images (with and without transparency)

##### Scenario: JPEG Image Test
Given a request for JPEG image tests
When the generator runs
Then it SHALL create documents with JPEG images at various qualities

##### Scenario: Multiple Image Sizes
Given a request for image size tests
When the generator runs
Then it SHALL create documents with images ranging from 64x64 to 1600x1200

#### Requirement: Table Generation Tests
The E2E generator SHALL create documents with various table configurations.

##### Scenario: Simple Table Test
Given a request for simple table tests
When the generator runs
Then it SHALL create documents with basic uniform tables

##### Scenario: Formatted Table Test
Given a request for formatted table tests
When the generator runs
Then it SHALL create documents with tables having borders and shading

##### Scenario: Merged Cell Table Test
Given a request for complex table tests
When the generator runs
Then it SHALL create documents with horizontally and vertically merged cells

#### Requirement: Baseline Generation Mode
The E2E runner SHALL support a --generate-baselines flag.

##### Scenario: Generate Baselines Flag
Given the --generate-baselines flag is provided
When the e2e-runner executes
Then it SHALL generate and save baseline PDFs without comparison

##### Scenario: Baseline Output Location
Given baseline generation is enabled
When the runner executes
Then baselines SHALL be saved to the testdata/baselines/ directory

### SHOULD Requirements

#### Requirement: Test Coverage
The generators SHOULD cover at least 80% of common document features.

#### Requirement: Performance
Test generation SHOULD complete within 60 seconds for a full suite.

## Testing Strategy

### Unit Tests
- Test each generator function independently
- Test document structure validation
- Test baseline file naming conventions

### Integration Tests
- End-to-end test generation pipeline
- Baseline generation and verification
- PDF output validation

### Regression Tests
- Compare generated PDFs against baselines
- Detect visual differences
- Report regressions

## Implementation Plan

1. Implement background application generator
2. Implement image generator
3. Implement table generator
4. Implement baseline generation mode
5. Add unit tests
6. Add integration tests
7. Create initial baseline files

## Related Changes

- `tests/e2e/generators/go/main.go` - Main generator orchestration
- `tests/e2e/generators/go/background_gen.go` - NEW: Background generator
- `tests/e2e/generators/go/image_gen.go` - NEW: Image generator
- `tests/e2e/generators/go/table_gen.go` - NEW: Table generator
- `tests/e2e/cmd/e2e-runner/main.go` - Runner with baseline mode

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Large baseline files | Medium | Compress baselines; store in git LFS if needed |
| Test generation time | Low | Optimize; parallelize generation |
| Flaky visual comparisons | Medium | Use tolerance thresholds; focus on critical areas |

## Acceptance Criteria

- [ ] Background generator creates solid, gradient, and pattern fills
- [ ] Image generator creates tests for PNG, JPEG, GIF
- [ ] Table generator creates simple, formatted, and merged tables
- [ ] Baseline generation mode works with --generate-baselines flag
- [ ] All generators have unit tests
- [ ] Integration tests pass
- [ ] Initial baselines created and verified
