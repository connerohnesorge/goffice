# Change: Complete Style Inheritance and Default Handling

## Why
Style resolution and inheritance chains are incomplete. Document styles, character styles, table styles, and list styles need proper cascade resolution and default value application matching Open-XML-SDK behavior.

## What Changes
- Style hierarchy resolution (based-on relationships)
- Default style determination
- Linked style support (paragraph to character)
- Style inheritance cascading for all element types
- Theme color and font resolution
- Style property precedence rules
- QuickStyle and hidden style support
- Automatic style generation (e.g., Normal, Heading 1-9)
- Style aliases and linked styles
- Style comparison and equality

## Impact
- Affected specs: wordprocessing-styles, spreadsheet-styles, validation, framework
- Affected code: wordprocessing/styles/*.go, spreadsheet/styles/*.go, openxml/styles/*.go (new)
- Breaking changes: None (changes internal resolution logic)
