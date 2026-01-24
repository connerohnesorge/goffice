# Add Spreadsheet Pivot Table Support

## Overview
Implement Excel pivot table creation, modification, and refresh functionality including pivot table configuration, data source management, field arrangement, calculations, and styling.

## Motivation
Pivot tables are one of Excel most powerful features for data analysis. Current goffice preserves existing pivot tables but cannot create or modify them.

## Goals
- Implement pivot table creation (define data source range or table, configure row/column/value/filter fields, set aggregation functions, define calculated fields)
- Support pivot table operations (refresh from source, add/remove/rearrange fields, change aggregations, apply filters and slicers)
- Implement pivot table styling (built-in styles, custom formatting, subtotal/grand total formatting)
- Support pivot table features (grouping dates/numbers, sorting/filtering, Top 10 filtering, value field settings show-as, custom field names)
- Enable pivot chart creation from pivot tables

## Dependencies
- Depends on: spreadsheet core, add-spreadsheet-formula-evaluation
- Related: add-spreadsheet-slicer-support

## Estimated Effort
16 weeks

## Priority
P1
