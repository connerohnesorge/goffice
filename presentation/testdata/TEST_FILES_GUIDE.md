# Test Files Guide

Quick reference guide for choosing the right test files for different testing scenarios.

## Basic Functionality Testing

### Simple Presentations (Good for Basic Read/Write Tests)
- `Presentation.pptx` - Standard presentation file (1002KB)
- `autosave.pptx` - Simple autosave test (34KB)
- `mcppt.pptx` - Minimal markup compatibility presentation (31KB)

### Templates
- `Presentation.potx` - PowerPoint template file for template testing (39KB)

## Advanced Feature Testing

### Animation & Transitions
```go
// Testing animation and effects
"animation.pptx"     // Animation features
"3dtestdash.pptx"    // 3D graphics with dashes
"3dtestdot.pptx"     // 3D graphics with dots
```

### Text & Formatting
```go
// Text alignment and formatting
"Algn_tab_TabAlignment.pptx"  // Tab alignment features
```

### Media Handling
```go
// Media references and embedded content
"mediareference.pptx"  // Media reference handling
```

### Security Features
```go
// Testing encrypted presentations
"encrypted_pptx.pptx"  // Encrypted presentation file
```

## Performance Testing

### Large & Complex Presentations
```go
// Performance and stress testing
"o09_Performance_typical.pptx"  // 1.5MB - Typical Office 2009 performance test
"Presentation.pptx"              // 1002KB - Full-featured presentation
```

## Office 2016 Features

### Modern PowerPoint Features
```go
// Office 2016 specific features
"Of16-01.pptx"  // Office 2016 features test 1
"Of16-02.pptx"  // Office 2016 features test 2
"Of16-03.pptx"  // Office 2016 features test 3
```

## Markup Compatibility Testing

```go
"mcppt.pptx"  // Markup compatibility presentation
```

## Example Test Usage

```go
// Basic read test
func TestBasicRead(t *testing.T) {
    pres, err := Open(filepath.Join("testdata", "Presentation.pptx"))
    // ...
}

// Template test
func TestTemplateRead(t *testing.T) {
    pres, err := Open(filepath.Join("testdata", "Presentation.potx"))
    // ...
}

// Animation test
func TestAnimations(t *testing.T) {
    pres, err := Open(filepath.Join("testdata", "animation.pptx"))
    // ...
}

// 3D graphics test
func Test3DGraphics(t *testing.T) {
    testCases := []string{
        "3dtestdash.pptx",
        "3dtestdot.pptx",
    }
    for _, file := range testCases {
        pres, err := Open(filepath.Join("testdata", file))
        // ...
    }
}

// Media reference test
func TestMediaReferences(t *testing.T) {
    pres, err := Open(filepath.Join("testdata", "mediareference.pptx"))
    // ...
}

// Encryption test
func TestEncryptedPresentation(t *testing.T) {
    pres, err := Open(filepath.Join("testdata", "encrypted_pptx.pptx"))
    // Should handle encrypted files appropriately
}

// Office version compatibility
func TestOffice2016Features(t *testing.T) {
    files := []string{
        "Of16-01.pptx",
        "Of16-02.pptx",
        "Of16-03.pptx",
    }
    for _, file := range files {
        pres, err := Open(filepath.Join("testdata", file))
        // ...
    }
}

// Performance test
func TestLargePresentation(t *testing.T) {
    pres, err := Open(filepath.Join("testdata", "o09_Performance_typical.pptx"))
    // Test performance with large file
}
```

## Recommended Test Progression

1. **Start Simple**: Use `mcppt.pptx` and `autosave.pptx`
2. **Add Features**: Test `animation.pptx`, `mediareference.pptx`
3. **Test 3D Graphics**: Use `3dtestdash.pptx`, `3dtestdot.pptx`
4. **Test Formatting**: Use `Algn_tab_TabAlignment.pptx`
5. **Go Complex**: Test with `Presentation.pptx`, `o09_Performance_typical.pptx`
6. **Modern Features**: Test Office 2016 files
7. **Security**: Test `encrypted_pptx.pptx`

## File Size Reference

| Category | Files | Size Range |
|----------|-------|------------|
| Small | mcppt, autosave, Algn_tab | 31-37KB |
| Medium | Presentation.potx, Of16 series | 39-115KB |
| Large | animation, mediareference | 98-272KB |
| Very Large | 3D tests, Presentation.pptx | 529KB-1002KB |
| Extra Large | o09_Performance_typical | 1.5MB |

## Quick Stats

- Total test files: 14
- .pptx files: 13
- .potx files: 1
- Total size: ~4.3MB
- Office versions covered: 2009, 2016
- Largest file: o09_Performance_typical.pptx (1.5MB)

## Feature Coverage Matrix

| Feature | Test Files |
|---------|------------|
| Animations | animation.pptx |
| 3D Graphics | 3dtestdash.pptx, 3dtestdot.pptx |
| Text Alignment | Algn_tab_TabAlignment.pptx |
| Media References | mediareference.pptx |
| Templates | Presentation.potx |
| Encryption | encrypted_pptx.pptx |
| Markup Compatibility | mcppt.pptx |
| Autosave | autosave.pptx |
| Office 2016 | Of16-01.pptx, Of16-02.pptx, Of16-03.pptx |
| Performance | o09_Performance_typical.pptx |

## Special Considerations

### Encrypted Files
- `encrypted_pptx.pptx` - May require special handling for password protection
- Test error handling and security features

### Performance Files
- `o09_Performance_typical.pptx` - Use for stress testing and performance benchmarks
- Good for testing memory usage and processing speed

### 3D Graphics
- `3dtestdash.pptx`, `3dtestdot.pptx` - Test 3D rendering and graphics features
- Large files (529-531KB) for comprehensive 3D testing
