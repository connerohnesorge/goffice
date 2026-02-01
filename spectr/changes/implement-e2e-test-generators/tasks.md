# Tasks: Implement E2E Test Generator Suite

## Background Application Generator

- [ ] 1.1 Create background_gen.go file
- [ ] 1.2 Implement solid color background generation
- [ ] 1.3 Implement linear gradient background generation
- [ ] 1.4 Implement radial gradient background generation
- [ ] 1.5 Implement pattern background generation
- [ ] 1.6 Add tests for each background type

## Image Generator

- [ ] 2.1 Create image_gen.go file
- [ ] 2.2 Implement PNG image test generation (with transparency)
- [ ] 2.3 Implement JPEG image test generation
- [ ] 2.4 Implement GIF image test generation
- [ ] 2.5 Implement various image sizes (small, medium, large)
- [ ] 2.6 Implement inline vs floating placement
- [ ] 2.7 Add test images to testdata/

## Table Generator

- [ ] 3.1 Create table_gen.go file
- [ ] 3.2 Implement simple table generation (2x2, 3x3, 5x5)
- [ ] 3.3 Implement table with borders
- [ ] 3.4 Implement table with cell shading
- [ ] 3.5 Implement table with header rows
- [ ] 3.6 Implement horizontal cell merging
- [ ] 3.7 Implement vertical cell merging
- [ ] 3.8 Implement nested tables

## Baseline Generation Mode

- [ ] 4.1 Add --generate-baselines flag to e2e-runner
- [ ] 4.2 Implement baseline generation logic
- [ ] 4.3 Implement baseline file naming convention
- [ ] 4.4 Add baseline output directory (testdata/baselines/)
- [ ] 4.5 Implement skip comparison when generating baselines

## Test Document Generation

- [ ] 5.1 Generate Word documents with backgrounds
- [ ] 5.2 Generate Word documents with images
- [ ] 5.3 Generate Word documents with tables
- [ ] 5.4 Generate Excel documents with backgrounds
- [ ] 5.5 Generate Excel documents with images
- [ ] 5.6 Generate PowerPoint documents with backgrounds

## PDF Rendering for Baselines

- [ ] 6.1 Integrate PDF rendering into test generation
- [ ] 6.2 Ensure consistent PDF output format
- [ ] 6.3 Handle PDF rendering errors gracefully

## Unit Tests

- [ ] 7.1 Create test: Background generator produces valid documents
- [ ] 7.2 Create test: Image generator produces valid documents
- [ ] 7.3 Create test: Table generator produces valid documents
- [ ] 7.4 Create test: Baseline generation mode works correctly
- [ ] 7.5 Create test: Generated documents are structurally valid

## Integration Tests

- [ ] 8.1 Create test: End-to-end background test generation
- [ ] 8.2 Create test: End-to-end image test generation
- [ ] 8.3 Create test: End-to-end table test generation
- [ ] 8.4 Create test: Full baseline generation pipeline

## Initial Baselines

- [ ] 9.1 Generate baseline PDFs for background tests
- [ ] 9.2 Generate baseline PDFs for image tests
- [ ] 9.3 Generate baseline PDFs for table tests
- [ ] 9.4 Verify baselines are valid PDFs
- [ ] 9.5 Commit baselines to repository

## Documentation

- [ ] 10.1 Update e2e/README.md with generator usage
- [ ] 10.2 Document how to add new test generators
- [ ] 10.3 Document baseline regeneration process
- [ ] 10.4 Add code comments to generator functions

## Verification

- [ ] 11.1 Run all e2e generator tests
- [ ] 11.2 Verify test coverage >80%
- [ ] 11.3 Verify baselines can be regenerated deterministically
- [ ] 11.4 Manual verification of generated test documents
