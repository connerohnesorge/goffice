# Change: Comprehensive Document Metadata & Properties Support

## Why
Documents require robust metadata (core, extended, and custom properties) for indexing, searching, and compatibility with Office applications. Open-XML-SDK provides full W3C Dublin Core, Office Custom Properties, and Variable storage. This capability is essential for real-world document workflows.

## What Changes
- Core document properties (title, subject, creator, date created/modified, category, keywords, description, lastModifiedBy, revision)
- Extended properties (application, appVersion, totalTime, pages, words, characters, lines, paragraphs, slides, notes, hiddenSlides, mmClips, scalesCrop, headingPairs, titlesOfParts, company, manager, documentSecurity)
- Custom properties with support for all property types (text, date, integer, floating-point, boolean)
- Automatic metadata generation (created date, modified date, save count)
- Property validation against ECMA-376 constraints

## Impact
- Affected specs: wordprocessing-document, spreadsheet-document, presentation-document
- Affected code: packaging, openxml/features
- Breaking changes: None
- New APIs: MetadataService, PropertyType enums

## Effort Estimate
- Implementation: 3-4 days
- Testing: 2 days
- Documentation: 1 day
