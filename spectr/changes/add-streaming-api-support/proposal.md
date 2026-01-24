# Change: Streaming API Implementation for Large Document Processing

## Why
Processing very large documents (>50MB) exhausts memory with DOM-based API. Open-XML-SDK provides streaming reader/writer for memory-efficient processing of large files.

## What Changes
- StreamingOpenXmlPart reader for sequential part reading
- StreamingPartReader and StreamingPartWriter abstractions
- Streaming element iteration without full document loading
- Streaming validation for parts without loading full content
- Memory-efficient chunked processing
- Cancellation token support for long operations
- Progress reporting for streaming operations
- Partial document writing (append-only mode)
- Streaming find and replace operations

## Impact
- Affected specs: streaming-api, package, openxml/validation
- Affected code: openxml/streaming/*.go, packaging/*.go
- Breaking changes: None (new API surface)
