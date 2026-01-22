# Media Support Specification

## ADDED

#### Scenario: Embed video content in slides
- **GIVEN** a presentation with editable access
- **WHEN** adding a video to a slide with valid video file path or stream
- **THEN** the system creates a video part with proper content types
- **AND** adds the video to the slide with specified position and size
- **AND** handles multiple video formats (MP4, AVI, WMV, MOV)
- **AND** supports video from file path, io.Reader, or embedded data

#### Scenario: Configure video playback options
- **GIVEN** a video has been added to a slide
- **WHEN** configuring playback options
- **THEN** support setting auto-play, loop, mute, and fullscreen options
- **AND** allow setting poster frame (preview image)
- **AND** support bookmark points for navigation
- **AND** manage aspect ratio preservation

## MODIFIED

#### Scenario: Extend shape creation to support chaining
- **GIVEN** a shape creation flow
- **WHEN** adding multiple properties
- **THEN** support fluent API chaining
- **AND** return self from setter methods
- **AND** maintain current functionality

# RTL Text Support Specification

## ADDED

#### Scenario: Basic RTL text rendering
- **GIVEN** a slide with text shapes
- **WHEN** adding Hebrew, Arabic, or Persian text content
- **THEN** render text right-to-left by default for RTL languages
- **AND** correctly handle bidirectional text mixing with LTR content
- **AND** support RTL paragraph direction
- **AND** maintain proper character ordering and joining

# Chart Type Extension Specification

## ADDED

#### Scenario: Add bubble chart support
- **GIVEN** data with X values, Y values, and bubble sizes
- **WHEN** creating a bubble chart
- **THEN** render bubbles positioned by X/Y coordinates
- **AND** scale bubble sizes proportionally

#### Scenario: Add radar chart support
- **GIVEN** multi-dimensional data sets
- **WHEN** creating a radar chart
- **THEN** display data on circular grid with radial axes
- **AND** support multiple series as overlapping polygons
