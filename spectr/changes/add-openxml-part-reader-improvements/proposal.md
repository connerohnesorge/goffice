# Add OpenXML Part Reader Improvements

## Overview
Enhance OpenXmlPartReader with additional capabilities from Open-XML-SDK including attribute reading, namespace resolution, element skipping, and read state management.

## Motivation
Current goffice OpenXmlPartReader provides basic streaming read. Open-XML-SDK version offers more advanced features like attribute enumeration, namespace context, efficient skipping, and detailed state tracking.

## Goals
- Enhance OpenXmlPartReader with missing methods (GetAttributes, GetAttribute, LoadCurrentElement, Skip, IsStartElement, IsEndElement, GetLineInfo for line/column error reporting)
- Implement namespace context tracking
- Support mixed read modes (streaming + DOM for selected elements)
- Provide better error messages with position information
- Enable efficient partial document processing

## Dependencies
- Depends on: openxml framework
- Related: add-openxml-dom-reader

## Estimated Effort
3 weeks

## Priority
P2
