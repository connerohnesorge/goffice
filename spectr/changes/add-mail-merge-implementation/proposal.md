# Add Mail Merge Implementation

## Overview
Implement Word mail merge functionality for generating multiple documents from template and data source. Includes merge field resolution, conditional blocks, and data source integration.

## Motivation
Mail merge is critical Word feature for document generation at scale (labels, letters, invoices, reports). Current goffice has basic field support but no merge execution engine.

## Goals
- Implement merge field resolution (MERGEFIELD, IF, NEXT, NEXTIF)
- Support multiple data source types (CSV, Excel, JSON, SQL)
- Generate single merged document or per-record documents
- Support nested fields and complex expressions
- Handle missing data gracefully
- Enable custom field formatters
- Support filtering and sorting data sources

## Dependencies
- Depends on: wordprocessing core, wordprocessing fields
- Related: add-package-clone-support

## Estimated Effort
8 weeks

## Priority
P1
