# Proposal: Implement PDF Image Renderer Features

## Summary

Implement three key features in the PDF image renderer: clip operations, tiling pattern support, and loading images from relationships.

## Background

The PDF image renderer (`pdf/drawing/image_renderer.go`) has three TODOs indicating missing functionality:

**Reference TODOs:**
1. `pdf/drawing/image_renderer.go:123` - "TODO: Implement clip operation when available in Page interface"
2. `pdf/drawing/image_renderer.go:149` - "TODO: Implement tiling pattern support"  
3. `pdf/drawing/image_renderer.go:160` - "TODO: Load image from relationship when implemented"

These features are needed for complete DrawingML image rendering support in PDF output.

## Motivation

Modern Office documents use various image fill modes:
- **Clip**: Shows only a portion of an image within a shape
- **Tile**: Repeats an image pattern to fill a shape
- **Relationships**: Images are often stored as separate parts referenced by relationships

Without these features, PDF rendering may not match the visual appearance of the original document.

## Technical Design

### 1. Clip Operation

**Purpose:** Clip an image to show only a portion within a defined rectangular area.

**Implementation:**
```go
func (r *ImageRenderer) RenderClippedPicture(
    blipFill *drawingml.BlipFill,
    bounds RenderBounds,
    clipRect *Rect, // NEW: Clip rectangle
) error {
    // Get image data
    imageData, err := r.loadImage(blipFill)
    if err != nil {
        return err
    }
    
    // Apply clip using PDF clip path
    r.page.PushState()
    r.page.DefineClipPath(clipRect)
    r.page.Clip()
    
    // Render image
    err = r.renderImageData(imageData, bounds)
    
    r.page.PopState()
    return err
}
```

**Requirements:**
- Define clip region in PDF content stream
- Apply clip before rendering image
- Restore state after rendering
- Support rectangular clips initially (more complex shapes in future)

### 2. Tiling Pattern Support

**Purpose:** Repeat an image to fill a shape (like wallpaper).

**Implementation:**
```go
func (r *ImageRenderer) RenderTiledPicture(
    blipFill *drawingml.BlipFill,
    bounds RenderBounds,
    tileOpts TileOptions, // UPDATED: Now used
) error {
    // Get image data and dimensions
    imageData, width, height, err := r.loadImageWithDimensions(blipFill)
    if err != nil {
        return err
    }
    
    // Calculate tiling parameters
    tilesX := int(math.Ceil(bounds.Width / width))
    tilesY := int(math.Ceil(bounds.Height / height))
    
    // Create tile pattern in PDF
    pattern := r.page.CreateTilingPattern(width, height)
    
    // Render image into pattern
    pattern.DrawImage(imageData, 0, 0, width, height)
    
    // Fill bounds with pattern
    r.page.FillWithPattern(bounds, pattern)
    
    return nil
}
```

**TileOptions fields:**
- `Alignment`: How tiles align (top-left, center, etc.)
- `Flip`: Whether to flip alternate tiles
- `Scale`: Scale factor for tiles

**Requirements:**
- Create PDF tiling pattern object
- Handle alignment options
- Support flipping (horizontal, vertical, both)
- Handle scale transformations
- Clip tiles to bounds

### 3. Load Image from Relationship

**Purpose:** Load image data from document relationships instead of embedded data.

**Implementation:**
```go
func (r *ImageRenderer) loadImageFromRelationship(
    relationshipID string,
) (imageData []byte, contentType string, err error) {
    // Get relationship from document
    rel, err := r.document.GetRelationship(relationshipID)
    if err != nil {
        return nil, "", fmt.Errorf("relationship not found: %w", err)
    }
    
    // Load image part
    imagePart, err := r.document.LoadPart(rel.TargetURI)
    if err != nil {
        return nil, "", fmt.Errorf("failed to load image part: %w", err)
    }
    
    // Read image data
    data, err := imagePart.ReadAll()
    if err != nil {
        return nil, "", fmt.Errorf("failed to read image data: %w", err)
    }
    
    return data, imagePart.ContentType(), nil
}

// Updated GetImageDimensions
func (r *ImageRenderer) GetImageDimensions(
    relationshipID string,
) (width, height float64, err error) {
    // Load from relationship
    data, contentType, err := r.loadImageFromRelationship(relationshipID)
    if err != nil {
        return 0, 0, err
    }
    
    // Parse image to get dimensions
    return r.parseImageDimensions(data, contentType)
}
```

**Requirements:**
- Resolve relationship ID to part
- Load image data from package part
- Support common image formats (PNG, JPEG, GIF, BMP)
- Handle missing relationships gracefully
- Cache loaded images for performance

## Requirements

### SHALL Requirements

#### Requirement: Clip Operation
The image renderer SHALL support clipping images to rectangular regions.

##### Scenario: Clip to Center Portion
Given an image and a clip rectangle defining the center 50%
When rendered with clip
Then only the center 50% SHALL be visible

##### Scenario: Clip with Offset
Given an image and a clip rectangle with X/Y offset
When rendered with clip
Then the image SHALL be offset within the clip region

#### Requirement: Tiling Pattern
The image renderer SHALL support tiling images to fill shapes.

##### Scenario: Basic Tiling
Given a small image (100x100) and large bounds (500x500)
When rendered with tile fill
Then the image SHALL repeat 5x5 to fill the bounds

##### Scenario: Tile Alignment
Given an image with tile alignment set to center
When rendered
Then tiles SHALL align according to the alignment setting

#### Requirement: Relationship Loading
The image renderer SHALL load images from document relationships.

##### Scenario: PNG from Relationship
Given a PNG image referenced by relationship ID "rId1"
When GetImageDimensions is called
Then the PNG SHALL be loaded and dimensions returned

##### Scenario: JPEG from Relationship
Given a JPEG image referenced by relationship ID "rId2"
When rendering
Then the JPEG SHALL be loaded and rendered

### SHOULD Requirements

#### Requirement: Image Caching
The implementation SHOULD cache loaded images to improve performance.

#### Requirement: Format Support
The implementation SHOULD support PNG, JPEG, GIF, and BMP formats.

## Testing Strategy

### Unit Tests
- Test clip operation with various rectangles
- Test tiling with different sizes and alignments
- Test relationship loading
- Test image format detection

### Integration Tests
- End-to-end rendering with clip
- End-to-end rendering with tile
- End-to-end with relationship-loaded images

### Visual Tests
- Compare clipped images with reference
- Compare tiled images with reference
- Verify image quality preservation

## Implementation Plan

1. Implement clip operation
2. Implement tiling pattern support
3. Implement relationship loading
4. Add image caching
5. Add unit tests
6. Add integration tests
7. Visual regression testing

## Related Changes

- `pdf/drawing/image_renderer.go` - Main implementation
- `pdf/core/page.go` - May need clip and pattern APIs
- `packaging/` - Relationship resolution

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Complex clip shapes | Medium | Start with rectangular clips |
| Large tiled patterns | Low | Optimize pattern creation |
| Missing relationships | Low | Graceful error handling |
| Unsupported image formats | Low | Clear error messages |

## Acceptance Criteria

- [ ] Clip operation working for rectangles
- [ ] Tiling pattern support with alignment
- [ ] Image loading from relationships
- [ ] Unit tests pass with >90% coverage
- [ ] Integration tests pass
- [ ] Visual fidelity matches Office
