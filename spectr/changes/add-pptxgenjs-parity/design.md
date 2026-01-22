# Design: Achieving pptxgenjs Parity

## Overview

This design document outlines the architectural approach for adding pptxgenjs-equivalent features to goffice.

## Core Architectural Decisions

### 1. Media Support Architecture
- Implement via existing part system
- Create MediaPart interface extending packaging.Part
- Format validation and streaming for large files

### 2. Master Slides System
- SlideMasterPart managing master definitions
- Inheritance: Slide → Layout → Master → Theme
- Override system for customizations

### 3. RTL Text Support
- Integrate with golang.org/x/text for BiDi
- Use bidi algorithm for character reordering
- Text runs with directionality metadata

### 4. Fluent API Design
- Fluent methods return receiver
- Builder pattern for complex objects
- Helper methods for common patterns
- Non-breaking, additive approach

## Integration Points
- PDF rendering for all new features
- Validation system extensions
- Theme system coordination

## Phased Implementation
1. Core media, RTL, Fluent APIs
2. Charts, shapes, rich text
3. Masters, notes, HTML conversion
4. Animations, transitions, comments
5. Polish, documentation, examples
