# Tasks: Implement PDF Image Renderer Features

## Clip Operation Implementation

- [ ] 1.1 Add clip rectangle parameter to RenderClippedPicture
- [ ] 1.2 Implement PDF clip path definition
- [ ] 1.3 Implement state push/pop for clip operations
- [ ] 1.4 Add rectangular clip support
- [ ] 1.5 Test clip with various rectangle sizes
- [ ] 1.6 Test clip with offset rectangles

## Tiling Pattern Implementation

- [ ] 2.1 Implement PDF tiling pattern creation
- [ ] 2.2 Add tile calculation (number of tiles in X/Y)
- [ ] 2.3 Implement alignment options (top-left, center, etc.)
- [ ] 2.4 Implement flip options (horizontal, vertical, both)
- [ ] 2.5 Add scale transformation support
- [ ] 2.6 Implement bounds clipping for tiles

## Relationship Loading Implementation

- [ ] 3.1 Add loadImageFromRelationship method
- [ ] 3.2 Implement relationship resolution from document
- [ ] 3.3 Add image part loading from package
- [ ] 3.4 Implement image format detection
- [ ] 3.5 Add image caching mechanism
- [ ] 3.6 Update GetImageDimensions to use relationships

## Image Format Support

- [ ] 4.1 Add PNG parsing for dimensions
- [ ] 4.2 Add JPEG parsing for dimensions
- [ ] 4.3 Add GIF parsing for dimensions
- [ ] 4.4 Add BMP parsing for dimensions
- [ ] 4.5 Add format validation

## Unit Tests

- [ ] 5.1 Create test: Clip with full image bounds
- [ ] 5.2 Create test: Clip with partial image bounds
- [ ] 5.3 Create test: Clip with offset rectangle
- [ ] 5.4 Create test: Basic tiling 2x2
- [ ] 5.5 Create test: Tiling with alignment
- [ ] 5.6 Create test: Tiling with flip
- [ ] 5.7 Create test: Load image from valid relationship
- [ ] 5.8 Create test: Handle missing relationship
- [ ] 5.9 Create test: Image caching
- [ ] 5.10 Create test: Parse PNG dimensions
- [ ] 5.11 Create test: Parse JPEG dimensions

## Integration Tests

- [ ] 6.1 Create test: End-to-end clip rendering
- [ ] 6.2 Create test: End-to-end tile rendering
- [ ] 6.3 Create test: End-to-end with relationship image
- [ ] 6.4 Create test: Combined clip and tile (if applicable)

## Visual/Regression Tests

- [ ] 7.1 Create visual test: Clipped image vs reference
- [ ] 7.2 Create visual test: Tiled pattern vs reference
- [ ] 7.3 Create visual test: Various image formats
- [ ] 7.4 Compare output with Office PDF export

## Test Fixtures

- [ ] 8.1 Create test image: PNG with transparency
- [ ] 8.2 Create test image: JPEG photo
- [ ] 8.3 Create test image: Small pattern for tiling
- [ ] 8.4 Create test document: With clipped image
- [ ] 8.5 Create test document: With tiled background

## Documentation

- [ ] 9.1 Update pdf/AGENTS.md with image rendering details
- [ ] 9.2 Add code comments for clip operations
- [ ] 9.3 Add code comments for tiling operations
- [ ] 9.4 Add code comments for relationship loading
- [ ] 9.5 Document supported image formats

## Verification

- [ ] 10.1 Run all pdf drawing tests - ensure no regressions
- [ ] 10.2 Verify test coverage >90% for changed code
- [ ] 10.3 Visual comparison with Excel/Word PDF export
- [ ] 10.4 Performance test with large images
