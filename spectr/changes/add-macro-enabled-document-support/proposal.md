# Add Macro-Enabled Document Support

## Overview
Implement support for macro-enabled document formats (.docm, .xlsm, .pptm) including VBA project preservation, macro storage, and digital signatures. Note: preservation only, not execution.

## Motivation
Many enterprise users have macro-enabled documents. Current goffice may not properly handle VBA project parts. Users need to read/write macro-enabled documents without breaking macros.

## Goals
- Support macro-enabled format types (.docm, .dotm, .xlsm, .xltm, .xlam, .pptm, .potm, .ppam)
- Implement VBA project part handling (preserve vbaProject.bin, maintain VBA relationships, support extraction/replacement)
- Handle digital signatures on macros (preserve existing, remove when modified, validate presence not crypto)
- Support macro settings preservation

## Dependencies
- Depends on: packaging core, all document types
- Related: add-encryption-and-protection

## Estimated Effort
2 weeks

## Priority
P2
