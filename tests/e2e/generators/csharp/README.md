# C# PowerPoint Generator for E2E Visual Testing

This is a .NET 9.0 console application that generates PowerPoint presentations from JSON test case definitions. It uses Microsoft's Open-XML-SDK to create PPTX files for visual comparison testing.

## Features

### Supported Chart Types
- **Bar Charts**: Clustered, stacked, percent stacked, column
- **Line Charts**: Standard, smooth, with markers
- **Pie Charts**: Standard pie
- **Doughnut Charts**: With configurable hole size
- **Area Charts**: Standard, stacked
- **Scatter Charts**: Standard, smooth, with lines
- **Bubble Charts**: X/Y/size visualization

### Supported Shapes
- **Preset Geometries**: Rectangle, rounded rectangle, ellipse, triangle, diamond, pentagon, hexagon, octagon, star, arrows, callouts
- **Fills**: Solid, linear gradient, radial gradient, pattern, none
- **Strokes**: Configurable width, color, dash styles (solid, dot, dash, dash-dot, long dash)
- **Effects**: Outer shadow, inner shadow, glow, reflection, soft edge

### Text Formatting
- Text boxes and shape text
- Paragraph formatting (alignment, line spacing, indentation, bullets)
- Character formatting (font family, size, color)
- Text styles (bold, italic, underline, strikethrough)

## Requirements

- .NET 9.0 SDK
- NuGet packages (auto-restored):
  - DocumentFormat.OpenXml 3.2.0
  - System.CommandLine 2.0.0-beta4.22272.1
  - Newtonsoft.Json 13.0.3

## Building

```bash
# Via Nix (recommended)
nix develop --command dotnet build

# Or directly with .NET SDK
dotnet build
```

## Usage

```bash
# Basic usage
dotnet run -- --input test_case.json --output presentation.pptx

# Via Nix
nix develop --command dotnet run --project PptxGenerator.csproj -- \
  --input path/to/test_case.json \
  --output path/to/output.pptx
```

## Project Structure

```
.
├── PptxGenerator.csproj    # .NET project file
├── Program.cs              # CLI entry point and main logic
├── Models/
│   └── TestCase.cs        # Data model classes (matches Go framework)
└── Generators/
    ├── ChartGenerator.cs  # Chart generation logic
    ├── ShapeGenerator.cs  # Shape generation logic
    ├── TextGenerator.cs   # Text generation logic
    └── Helpers.cs         # Utility classes (fill, color, EMU conversion)
```

## Input Format

The generator expects JSON files matching the test case schema defined in `tests/e2e/framework/testcase.go`. See example test cases in `tests/e2e/testcases/`.

### Minimal Example

```json
{
  "id": "example_chart",
  "name": "Example Chart",
  "description": "A simple bar chart",
  "category": "chart",
  "spec": {
    "slide_count": 1,
    "slide_size": {
      "width": 9144000,
      "height": 6858000
    },
    "slides": [
      {
        "index": 0,
        "layout": "Blank",
        "elements": [
          {
            "type": "chart",
            "position": { "x": 914400, "y": 914400 },
            "size": { "width": 7315200, "height": 5029200 },
            "chart": {
              "type": "bar_clustered",
              "title": "Sales by Quarter",
              "data": {
                "categories": ["Q1", "Q2", "Q3", "Q4"],
                "series": [
                  {
                    "name": "Revenue",
                    "values": [100, 150, 120, 180],
                    "color": "4472C4"
                  }
                ]
              },
              "style": {
                "font_size": 12,
                "font_family": "Calibri"
              }
            }
          }
        ]
      }
    ]
  }
}
```

## EMU Units

PowerPoint uses English Metric Units (EMUs) for positioning and sizing:
- 1 inch = 914,400 EMUs
- 1 cm = 360,000 EMUs
- 1 point = 12,700 EMUs

The `EMUConverter` helper class in `Helpers.cs` provides conversion utilities.

## Implementation Details

### Chart Generation
Charts are created using the Open-XML-SDK's chart API. Each chart type has specialized generation logic that creates the appropriate chart XML structure with:
- Series data and categories
- Titles and legends
- Axes configuration
- Styling and colors

### Shape Generation
Shapes use DrawingML's preset geometries and support:
- Multiple fill types (solid, gradient, pattern)
- Stroke/outline styling
- Visual effects (shadows, glow, reflection, soft edges)
- Text content within shapes

### Text Generation
Text is rendered using DrawingML text bodies with:
- Paragraph-level formatting
- Character-level formatting (runs)
- Font properties
- Alignment and spacing

## Testing

The generator has been tested with:
- ✅ Bar chart generation
- ✅ Shape rendering (rectangles, ellipses, etc.)
- ✅ Text formatting (bold, italic, colors, sizes)
- ✅ Multi-element slides
- ✅ Gradient fills
- ✅ Effects (shadows, glow)

## Known Limitations

1. Some advanced chart features not yet implemented:
   - Error bars
   - Trendlines
   - Data table display

2. Pattern fills use approximations (Open-XML-SDK has limited preset patterns)

3. Minor nullable reference warnings in compilation (non-critical)

## Compatibility

- **Output Format**: PowerPoint 2007+ (.pptx)
- **Platforms**: Cross-platform (Linux, macOS, Windows)
- **Framework**: .NET 9.0
- **Tested With**: LibreOffice Impress, Microsoft PowerPoint

## Related Components

This generator is part of the E2E visual testing framework:
- Go generator: `tests/e2e/generators/go/`
- Framework: `tests/e2e/framework/`
- Test cases: `tests/e2e/testcases/`

Both generators (Go and C#) produce presentations from the same JSON format, enabling cross-implementation comparison.
