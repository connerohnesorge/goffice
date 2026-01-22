## ADDED Requirements

### Requirement: Media Embedding Support
The system MUST support embedding video and audio content in slides.

#### Scenario: Embed video content in slides
- **WHEN** adding a video to a slide with valid video file path
- **THEN** a video part is created with proper content types
- **AND** the video is added to the slide with specified position and size
- **AND** playback options can be configured (auto-play, loop, fullscreen)

#### Scenario: Embed audio content in slides
- **WHEN** adding audio to a slide with valid audio file
- **THEN** an audio part is created with proper content types
- **AND** the audio supports play-across-slides functionality
- **AND** looping and volume settings are configurable

### Requirement: Speaker Notes Support
The system MUST support adding presentation notes to slides.

#### Scenario: Add speaker notes to slides
- **WHEN** adding speaker notes to a slide
- **THEN** a notes slide part is created and linked to the parent slide
- **AND** notes support basic text formatting (bold, italic)
- **AND** multiple paragraphs are supported

### Requirement: Slide Master Templates
The system MUST support creating and applying slide master templates.

#### Scenario: Create slide master template
- **WHEN** creating a slide master
- **THEN** layout positions for placeholders are defined
- **AND** backgrounds (colors, gradients, images) can be set
- **AND** corporate branding elements are supported