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

#### Scenario: Embed audio content in slides
- **GIVEN** a presentation with editable access
- **WHEN** adding audio to a slide with valid audio file
- **THEN** the system creates an audio part with proper content types
- **AND** adds the audio with position for playback controls
- **AND** handles multiple audio formats (MP3, WAV, WMA, M4A)
- **AND** supports background music across multiple slides

#### Scenario: Configure audio playback options
- **GIVEN** audio has been added to presentation
- **WHEN** configuring audio settings
- **THEN** support play across slides option
- **AND** allow loop, auto-play, and volume settings
- **AND** support fade in/out effects
- **AND** show/hide audio icon on slide

#### Scenario: Media validation and error handling
- **GIVEN** invalid or unsupported media files
- **WHEN** attempting to add them to slides
- **THEN** the system validates file format and codecs
- **AND** returns appropriate error messages
- **AND** provides fallback options for incompatible formats
