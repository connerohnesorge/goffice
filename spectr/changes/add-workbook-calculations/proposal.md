# Change: Spreadsheet Workbook Calculations & Formula Engine

## Why
Excel calculations are core to spreadsheet functionality. Open-XML-SDK provides cell evaluation, formula parsing, circular reference detection, and calculation modes. Complete formula support enables proper spreadsheet simulation and validation.

## What Changes
- Formula parsing and normalization
- Cell evaluation and calculation engine
- Circular reference detection and reporting
- Calculation mode support (automatic, manual, automatic-except-tables)
- Dependent cell tracking
- Volatile function handling
- Array formula support
- Named range formula resolution
- External link reference handling
- Calculation chain optimization

## Impact
- Affected specs: spreadsheet-formulas, spreadsheet-cells
- Affected code: spreadsheet package
- Breaking changes: None
- New APIs: FormulaParser, CalculationEngine, CircularReferenceDetector

## Effort Estimate
- Implementation: 5-6 days
- Testing: 3-4 days
- Documentation: 1 day
