# Presentation Test Data

This directory contains comprehensive test files for integration testing of the PowerPoint presentation processing functionality.

## Source

All files were copied from the Open-XML-SDK test assets:
- Source: `Open-XML-SDK/test/DocumentFormat.OpenXml.Tests.Assets/assets/TestFiles/`
- Files: 13 .pptx files and 1 .potx file

## File Categories

### Presentation Templates
- `Presentation.potx` - PowerPoint template file (39KB)

### Standard Presentations
- `Presentation.pptx` - Full-featured presentation file (1002KB)
- `autosave.pptx` - Autosave feature test (34KB)

### Animation & Effects
- `animation.pptx` - Presentation with animations (98KB)
- `3dtestdash.pptx` - 3D graphics with dashes (529KB)
- `3dtestdot.pptx` - 3D graphics with dots (531KB)

### Text Formatting
- `Algn_tab_TabAlignment.pptx` - Tab alignment testing (37KB)

### Media & References
- `mediareference.pptx` - Media reference handling (52KB)

### Performance Testing
- `o09_Performance_typical.pptx` - Typical Office 2009 performance test (1.5MB)

### Security
- `encrypted_pptx.pptx` - Encrypted presentation file (272KB)

### Office 2016 Feature Presentations
- `Of16-01.pptx` - Office 2016 features test 1 (58KB)
- `Of16-02.pptx` - Office 2016 features test 2 (51KB)
- `Of16-03.pptx` - Office 2016 features test 3 (115KB)

### Markup Compatibility
- `mcppt.pptx` - Markup compatibility presentation (31KB)

## Usage in Tests

These files are available for Go tests using the standard testdata directory convention. Access them in tests like:

```go
filepath.Join("testdata", "Presentation.pptx")
```

## File Statistics

- Total files: 14
- .pptx files: 13
- .potx files: 1
- Total size: ~4.3MB
- Largest file: o09_Performance_typical.pptx (1.5MB)
- Smallest file: mcppt.pptx (31KB)

## Feature Coverage

The test files cover:
- Basic presentation structure
- Slides and layouts
- Text formatting and alignment
- Animations and transitions
- 3D graphics and effects
- Media references
- Template files (.potx)
- Encrypted presentations
- Office 2016 features
- Markup compatibility
- Performance testing scenarios
