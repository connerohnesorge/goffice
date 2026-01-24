# Add OpenXML Element Equality Comparison

## Overview
Implement element equality comparison infrastructure matching Open-XML-SDK's Equality namespace. Enables deep comparison of OpenXML elements with configurable options for attribute ordering, namespace handling, and whitespace sensitivity.

## Motivation
Users need to compare documents for semantic equivalence, detect changes, and write tests that verify document content. Open-XML-SDK provides comprehensive equality comparison.

## Goals
- Implement OpenXmlElementEqualityComparer with full feature parity
- Support configurable comparison options (ignore whitespace, namespace-aware, etc.)  
- Provide helper methods for common comparison scenarios
- Enable use in testing frameworks (testify assertions)
- Support comparison of entire documents and individual elements

## Dependencies
- Depends on: openxml core framework
- Related: add-openxml-linq-support, add-validation-semantic-constraints
