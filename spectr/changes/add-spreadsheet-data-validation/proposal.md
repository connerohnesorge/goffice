# Add Spreadsheet Data Validation Support

## Overview
Implement data validation rules for Excel cells including dropdowns, numeric ranges, date ranges, text length, and custom formulas for input validation and dropdown list creation.

## Motivation
Data validation is essential for data entry forms and template spreadsheets. Current goffice lacks support for creating or modifying validation rules.

## Goals
- Implement all data validation types (list/dropdown validation with source range or hardcoded values, whole number validation with min/max, decimal validation with min/max, date validation with date ranges, time validation with time ranges, text length validation, custom formula validation)
- Support input messages and error alerts (Information, Warning, Stop styles)
- Support named ranges as validation sources
- Enable copying validation rules
- Validate rule configurations
- Support validation in template generation

## Dependencies
- Depends on: spreadsheet core
- Related: add-spreadsheet-formula-evaluation

## Estimated Effort
3 weeks

## Priority
P1
