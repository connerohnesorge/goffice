# Chart Formatting Enhancement Specification

## MODIFIED

#### Scenario: Enhanced data label formatting
- **GIVEN** a chart with data labels
- **WHEN** configuring label appearance
- **THEN** support custom label text and number formats
- **AND** allow label positioning (center, inside end, outside end)
- **AND** provide leader lines for distant labels
- **AND** support label boxes with fill and border
- **AND** enable per-series and per-point label customization

#### Scenario: Multiple axis support
- **GIVEN** a chart requiring different scales
- **WHEN** adding secondary axes
- **THEN** support primary and secondary Y axes
- **AND** provide secondary X axes for complex charts
- **AND** allow custom axis scales and intervals
- **AND** support axis crossing at specific values

#### Scenario: Enhanced axis formatting
- **GIVEN** chart axes requiring customization
- **WHEN** formatting axis appearance
- **THEN** support logarithmic scales and intervals
- **AND** provide custom axis label formatting (dates, currency, percentages)
- **AND** allow axis title positioning and rotation
- **AND** support axis gridline customization
- **AND** provide minor and major tick mark control

#### Scenario: Chart legend enhancements
- **GIVEN** a chart with multiple data series
- **WHEN** configuring legend appearance
- **THEN** support legend positioning (top, bottom, left, right, corner)
- **AND** allow custom legend item ordering
- **AND** provide legend text formatting
- **AND** support legend with multiple columns
- **AND** enable legend overlay on chart area

#### Scenario: Chart series formatting
- **GIVEN** chart data series
- **WHEN** applying series formatting
- **THEN** support gradient fills for series elements
- **AND** provide pattern fills (hatch, dots, stripes)
- **AND** allow custom marker shapes and sizes
- **AND** support line styles (solid, dashed, dotted)
- **AND** enable series overlap and gap width control

#### Scenario: Chart error bars
- **GIVEN** statistical or scientific data
- **WHEN** adding error bars to chart series
- **THEN** support fixed value, percentage, and standard deviation errors
- **AND** provide custom error amount specification
- **AND** allow error bar direction (both, plus, minus)
- **AND** support error bar end cap styles
- **AND** enable separate customization for X/Y errors
