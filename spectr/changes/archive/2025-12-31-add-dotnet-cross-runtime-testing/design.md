# Design Document: .NET Cross-Runtime Testing Infrastructure

**Related Documents**:
- High-level overview: See `proposal.md`
- Implementation tasks: See `tasks.md`
- E2E testing requirements: See `specs/e2e-testing/spec.md`
- Nix environment requirements: See `specs/nix-environment/spec.md`

## Table of Contents
1. [Architecture Overview](#architecture-overview)
2. [Detailed Component Design](#detailed-component-design)
3. [Code Patterns and Examples](#code-patterns-and-examples)
4. [Data Structures](#data-structures)
5. [Comparison Algorithms](#comparison-algorithms)
6. [Nix Integration](#nix-integration)
7. [Test Execution Flow](#test-execution-flow)
8. [API Mapping Patterns](#api-mapping-patterns)

---

## Architecture Overview

### System Diagram
```
┌─────────────────────────────────────────────────────────────┐
│                    Nix Development Shell                     │
│  ┌────────────┐         ┌────────────────────────────────┐ │
│  │ Go 1.25    │         │ .NET SDK 9                      │ │
│  │ Tools      │         │ (dotnet-sdk_9)                  │ │
│  └────────────┘         └────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────┐
│               E2E Test Framework (Go)                        │
│  ┌─────────────────────────────────────────────────────┐   │
│  │  Scenario Loader (YAML → TestScenario structs)      │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                             │
           ┌─────────────────┴─────────────────┐
           ▼                                   ▼
┌──────────────────────┐           ┌─────────────────────────┐
│  Go Bridge           │           │  C# Bridge              │
│  (goffice API)       │           │  (Open-XML-SDK API)     │
│                      │           │                         │
│  - Parses scenario   │           │  - Parses scenario      │
│  - Creates document  │           │  - Creates document     │
│    using goffice     │           │    using SDK            │
│  - Saves to disk     │           │  - Saves to disk        │
└──────────────────────┘           └─────────────────────────┘
           │                                   │
           │  go-output.docx                  │  dotnet-output.docx
           └─────────────────┬─────────────────┘
                             ▼
┌─────────────────────────────────────────────────────────────┐
│           Three-Level Comparison Pipeline                    │
│                                                              │
│  Level 1: XML Structure    (Fast)                           │
│  ├─ Parse XML from both documents                           │
│  ├─ Compare element trees                                   │
│  └─ Report: Element/attribute differences                   │
│                                                              │
│  Level 2: Binary Content   (Medium)                         │
│  ├─ Extract ZIP parts                                       │
│  ├─ Compare relationships, content types                    │
│  └─ Report: Package structure differences                   │
│                                                              │
│  Level 3: Visual Rendering (Slow)                           │
│  ├─ Convert to PDF (LibreOffice headless)                   │
│  ├─ Render to PNG (Ghostscript)                             │
│  ├─ Pixel diff (Euclidean distance)                         │
│  └─ Report: Visual differences with annotated images        │
└─────────────────────────────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────┐
│              Report Generator                                │
│  - HTML report (side-by-side comparison)                    │
│  - Diff images (highlighted differences)                    │
│  - Statistics (pass/fail, % match, pixel deltas)            │
│  - Logs (operation traces, error messages)                  │
└─────────────────────────────────────────────────────────────┘
```

### Directory Structure
```
tests/e2e/
├── README.md                          # Complete documentation
├── go.mod                             # Separate Go module for E2E tests (requires gopkg.in/yaml.v3 v3.0.1)
│                                      # Module path: github.com/connerohnesorge/goffice-e2e
│                                      # Uses replace directive: github.com/connerohnesorge/goffice => ../../
├── framework/                         # Test execution framework
│   ├── scenario.go                    # Scenario parsing and types
│   ├── executor.go                    # Test runner core
│   ├── comparison.go                  # Comparison orchestration
│   └── reporter.go                    # Report generation
├── bridges/                           # Language-specific generators
│   ├── go/
│   │   ├── bridge.go                  # Go→goffice translation
│   │   ├── wordprocessing.go          # Word document operations
│   │   ├── spreadsheet.go             # Excel document operations
│   │   └── presentation.go            # PowerPoint document operations
│   └── csharp/
│       ├── DocxBridge/                # C# project
│       │   ├── DocxBridge.csproj      # .NET 9 project with Open-XML-SDK and YamlDotNet
│       │   ├── Program.cs             # CLI entry point
│       │   ├── ScenarioParser.cs      # YAML → C# objects
│       │   ├── WordprocessingBridge.cs
│       │   ├── SpreadsheetBridge.cs
│       │   └── PresentationBridge.cs
│       └── build.sh                   # Build script
├── scenarios/                         # Test definitions
│   ├── wordprocessing/
│   │   ├── basic-paragraph.yaml
│   │   ├── formatted-text.yaml
│   │   ├── tables.yaml
│   │   ├── headers-footers.yaml
│   │   └── ...
│   ├── spreadsheet/
│   │   ├── basic-cells.yaml
│   │   ├── formulas.yaml
│   │   └── ...
│   └── presentation/
│       ├── basic-slide.yaml
│       └── ...
├── baselines/                         # Reference documents
│   ├── wordprocessing/
│   ├── spreadsheet/
│   └── presentation/
├── comparison/                        # Comparison logic
│   ├── xml.go                         # XML structure comparison
│   ├── binary.go                      # Binary content comparison
│   ├── visual.go                      # Visual rendering comparison
│   └── tolerance.go                   # Tolerance configuration
├── reports/                           # Generated reports (gitignored)
│   ├── html/
│   │   └── index.html                 # Main report
│   ├── diffs/
│   │   └── *.png                      # Diff images
│   └── logs/
│       └── *.log                      # Execution logs
├── output/                            # Generated test documents (gitignored)
│   ├── go/
│   └── dotnet/
├── scripts/                           # Utility scripts
│   ├── run-e2e.sh                     # Main test runner
│   ├── generate-baselines.sh          # Baseline generation
│   └── update-baselines.sh            # Baseline update workflow
├── cmd/                               # CLI tools
│   └── e2e-runner/
│       └── main.go                    # Test runner entry point
└── .gitignore                         # Git ignore rules:
                                       #   output/
                                       #   reports/
                                       #   bridges/csharp/DocxBridge/bin/
                                       #   bridges/csharp/DocxBridge/obj/
```

---

## Detailed Component Design

### 1. Test Scenario Definition

#### Scenario YAML Format
```yaml
# tests/e2e/scenarios/wordprocessing/formatted-text.yaml
name: formatted-text
description: Paragraph with bold, italic, and colored text
document_type: wordprocessing
category: basic
tags: [text-formatting, runs, properties]

# Comparison tolerance overrides
tolerance:
  xml_attribute_order_sensitive: false
  visual_pixel_tolerance: 2.0
  visual_diff_threshold: 0.005

# Document generation operations
operations:
  - action: create_document
    doc_type: document  # .docx

  - action: add_paragraph
    operations:
      - action: add_run
        text: "Normal text "

      - action: add_run
        text: "bold text "
        properties:
          bold: true

      - action: add_run
        text: "italic text "
        properties:
          italic: true

      - action: add_run
        text: "colored"
        properties:
          color: "FF0000"  # Red
          bold: true
          font_size: 12  # Points (converted to 24 half-points internally by Go API)
          font_name: "Calibri"
```

#### Scenario Go Struct
```go
// tests/e2e/framework/scenario.go

package framework

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"

    "gopkg.in/yaml.v3"
)

// TestScenario represents a complete test scenario
type TestScenario struct {
    Name         string             `yaml:"name"`
    Description  string             `yaml:"description"`
    DocumentType string             `yaml:"document_type"` // wordprocessing, spreadsheet, presentation
    Category     string             `yaml:"category"`      // basic, complex, edge-case
    Tags         []string           `yaml:"tags"`
    Tolerance    ToleranceConfig    `yaml:"tolerance"`
    Operations   []Operation        `yaml:"operations"`
}

// Operation represents a single document operation
type Operation struct {
    Action     string                 `yaml:"action"`  // create_document, add_paragraph, add_run, etc.
    Properties map[string]interface{} `yaml:",inline"` // Action-specific properties
}

// ToleranceConfig configures comparison tolerances
type ToleranceConfig struct {
    XMLAttributeOrderSensitive bool    `yaml:"xml_attribute_order_sensitive"`
    VisualPixelTolerance       float64 `yaml:"visual_pixel_tolerance"`
    VisualDiffThreshold        float64 `yaml:"visual_diff_threshold"`
}

// LoadScenario loads a scenario from YAML file
func LoadScenario(path string) (*TestScenario, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("open scenario file: %w", err)
    }
    defer file.Close()

    var scenario TestScenario
    decoder := yaml.NewDecoder(file)
    if err := decoder.Decode(&scenario); err != nil {
        return nil, fmt.Errorf("parse scenario YAML: %w", err)
    }

    // Set defaults
    if scenario.Tolerance.VisualPixelTolerance == 0 {
        scenario.Tolerance.VisualPixelTolerance = 2.0
    }
    if scenario.Tolerance.VisualDiffThreshold == 0 {
        scenario.Tolerance.VisualDiffThreshold = 0.005
    }

    return &scenario, nil
}

// LoadScenariosFromDirectory loads all scenarios in a directory
func LoadScenariosFromDirectory(dirPath string) ([]*TestScenario, error) {
    entries, err := os.ReadDir(dirPath)
    if err != nil {
        return nil, fmt.Errorf("read scenarios directory: %w", err)
    }

    var scenarios []*TestScenario
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        if !strings.HasSuffix(entry.Name(), ".yaml") && !strings.HasSuffix(entry.Name(), ".yml") {
            continue
        }

        scenarioPath := filepath.Join(dirPath, entry.Name())
        scenario, err := LoadScenario(scenarioPath)
        if err != nil {
            return nil, fmt.Errorf("load scenario %s: %w", entry.Name(), err)
        }
        scenarios = append(scenarios, scenario)
    }

    return scenarios, nil
}
```

### 2. Go Bridge (goffice API → Document)

```go
// tests/e2e/bridges/go/bridge.go

package gobridge

import (
    "fmt"

    "github.com/connerohnesorge/goffice/wordprocessing"
    "github.com/connerohnesorge/goffice/wordprocessing/elements"
    "github.com/connerohnesorge/goffice-e2e/framework"
)

// Bridge translates test scenarios into goffice API calls
type Bridge struct {
    scenario *framework.TestScenario
    doc      interface{} // Can be *wordprocessing.Document, *spreadsheet.Document, etc.
}

// NewBridge creates a new Go bridge
func NewBridge(scenario *framework.TestScenario) *Bridge {
    return &Bridge{scenario: scenario}
}

// Execute runs the scenario and produces a document
func (b *Bridge) Execute(outputPath string) error {
    switch b.scenario.DocumentType {
    case "wordprocessing":
        return b.executeWordprocessing(outputPath)
    case "spreadsheet":
        return b.executeSpreadsheet(outputPath)
    case "presentation":
        return b.executePresentation(outputPath)
    default:
        return fmt.Errorf("unknown document type: %s", b.scenario.DocumentType)
    }
}

// executeWordprocessing handles Word document scenarios
func (b *Bridge) executeWordprocessing(outputPath string) error {
    var doc *wordprocessing.Document
    var builder *wordprocessing.DocumentBuilder

    for _, op := range b.scenario.Operations {
        switch op.Action {
        case "create_document":
            docType, _ := op.Properties["doc_type"].(string)
            var dt wordprocessing.DocType
            switch docType {
            case "document":
                dt = wordprocessing.DocTypeDocument
            case "template":
                dt = wordprocessing.DocTypeTemplate
            default:
                dt = wordprocessing.DocTypeDocument
            }

            var err error
            doc, err = wordprocessing.New(outputPath, dt)
            if err != nil {
                return fmt.Errorf("create document: %w", err)
            }
            builder = wordprocessing.NewDocumentBuilder()

        case "add_paragraph":
            if err := b.addParagraph(builder, op); err != nil {
                return fmt.Errorf("add paragraph: %w", err)
            }

        case "add_table":
            if err := b.addTable(builder, op); err != nil {
                return fmt.Errorf("add table: %w", err)
            }

        default:
            return fmt.Errorf("unknown action: %s", op.Action)
        }
    }

    // Build final document element from builder and populate doc.MainPart().Document().Body()
    if builder != nil && doc != nil {
        builtDoc, err := builder.Build()
        if err != nil {
            return fmt.Errorf("build document: %w", err)
        }

        // Get the main part's document body and replace with builder's content
        mainPart := doc.MainPart()
        if mainPart == nil {
            return fmt.Errorf("no main part in document")
        }

        mainDoc := mainPart.Document()
        if mainDoc == nil {
            return fmt.Errorf("no document in main part")
        }

        // Replace body with builder's body
        mainDoc.SetBody(builtDoc.Body())
        b.doc = doc
    }

    // Save
    if doc != nil {
        if err := doc.Save(); err != nil {
            return fmt.Errorf("save document: %w", err)
        }
        return doc.Close()
    }

    return nil
}

// addParagraph translates add_paragraph operation
func (b *Bridge) addParagraph(builder *wordprocessing.DocumentBuilder, op framework.Operation) error {
    // Get paragraph text (if simple case)
    text, hasText := op.Properties["text"].(string)

    if hasText {
        // Simple case: single text string
        pb := builder.AddParagraph(text)

        // Apply paragraph-level properties
        if props, ok := op.Properties["properties"].(map[string]interface{}); ok {
            b.applyParagraphProperties(pb, props)
        }

        pb.Document() // Return to builder
        return nil
    }

    // Complex case: nested operations (runs)
    operations, hasOps := op.Properties["operations"].([]interface{})
    if !hasOps {
        return fmt.Errorf("paragraph must have either text or operations")
    }

    pb := builder.AddParagraph("")

    for _, opRaw := range operations {
        opMap, ok := opRaw.(map[string]interface{})
        if !ok {
            continue
        }

        action, _ := opMap["action"].(string)
        switch action {
        case "add_run":
            runText, _ := opMap["text"].(string)
            rb := pb.AddRun(runText)

            // Apply run properties
            if props, ok := opMap["properties"].(map[string]interface{}); ok {
                b.applyRunProperties(rb, props)
            }

            rb.Paragraph() // Back to paragraph builder
        }
    }

    pb.Document() // Return to document builder
    return nil
}

// applyParagraphProperties applies properties to paragraph builder
func (b *Bridge) applyParagraphProperties(pb *wordprocessing.ParagraphBuilder, props map[string]interface{}) {
    if align, ok := props["alignment"].(string); ok {
        switch align {
        case "left":
            pb.AlignLeft()
        case "center":
            pb.AlignCenter()
        case "right":
            pb.AlignRight()
        case "justify":
            pb.AlignJustify()
        }
    }

    if spacingAfter, ok := props["spacing_after"].(float64); ok {
        pb.SpacingAfter(int(spacingAfter))
    }

    if spacingBefore, ok := props["spacing_before"].(float64); ok {
        pb.SpacingBefore(int(spacingBefore))
    }

    if leftIndent, ok := props["left_indent"].(float64); ok {
        pb.LeftIndent(int(leftIndent))
    }
}

// applyRunProperties applies properties to run builder
func (b *Bridge) applyRunProperties(rb *wordprocessing.RunBuilder, props map[string]interface{}) {
    if bold, ok := props["bold"].(bool); ok && bold {
        rb.Bold()
    }

    if italic, ok := props["italic"].(bool); ok && italic {
        rb.Italic()
    }

    if underline, ok := props["underline"].(bool); ok && underline {
        rb.Underline()
    }

    if color, ok := props["color"].(string); ok {
        rb.Color(color)
    }

    if fontSize, ok := props["font_size"].(float64); ok {
        rb.FontSize(int(fontSize)) // Points (API converts to half-points internally)
    }

    if fontName, ok := props["font_name"].(string); ok {
        rb.Font(fontName)
    }
}

// addTable translates add_table operation
func (b *Bridge) addTable(builder *wordprocessing.DocumentBuilder, op framework.Operation) error {
    rows, _ := op.Properties["rows"].(float64)
    cols, _ := op.Properties["cols"].(float64)

    tb := builder.AddTable(int(rows), int(cols))

    // Set table properties
    if style, ok := op.Properties["style"].(string); ok {
        tb.SetStyle(style)
    }

    if widthPercent, ok := op.Properties["width_percent"].(float64); ok {
        tb.SetWidthPercent(int(widthPercent))
    }

    // Set cell values
    if cells, ok := op.Properties["cells"].([]interface{}); ok {
        for _, cellRaw := range cells {
            cellMap, ok := cellRaw.(map[string]interface{})
            if !ok {
                continue
            }

            row, _ := cellMap["row"].(float64)
            col, _ := cellMap["col"].(float64)
            text, _ := cellMap["text"].(string)

            tb.SetCellText(int(row), int(col), text)
        }
    }

    tb.Document() // Return to builder
    return nil
}
```

### 3. C# Bridge (Open-XML-SDK API → Document)

The C# bridge is a .NET 9 console application that uses the following NuGet packages:
- **DocumentFormat.OpenXml 3.2.0** - Microsoft's Open-XML-SDK for document manipulation
- **YamlDotNet 16.2.1** - YAML parsing library for reading test scenarios
- **System.CommandLine 2.0.0-beta4.22272.1** - Modern command-line argument parsing

```csharp
// tests/e2e/bridges/csharp/DocxBridge/Program.cs

using System;
using System.CommandLine;
using System.IO;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace DocxBridge
{
    class Program
    {
        static int Main(string[] args)
        {
            var scenarioOption = new Option<string>(
                "--scenario",
                description: "Path to scenario YAML file");
            scenarioOption.IsRequired = true;

            var outputOption = new Option<string>(
                "--output",
                description: "Output document path");
            outputOption.IsRequired = true;

            var rootCommand = new RootCommand("Generates Office documents from test scenarios")
            {
                scenarioOption,
                outputOption
            };

            rootCommand.SetHandler((string scenarioPath, string outputPath) =>
            {
                try
                {
                    var scenario = LoadScenario(scenarioPath);
                    GenerateDocument(scenario, outputPath);
                    Console.WriteLine($"SUCCESS: Generated {outputPath}");
                }
                catch (Exception ex)
                {
                    Console.Error.WriteLine($"ERROR: {ex.Message}");
                    Environment.Exit(1);
                }
            }, scenarioOption, outputOption);

            return rootCommand.Invoke(args);
        }

        static TestScenario LoadScenario(string path)
        {
            var yaml = File.ReadAllText(path);
            var deserializer = new DeserializerBuilder()
                .WithNamingConvention(UnderscoredNamingConvention.Instance)
                .Build();
            return deserializer.Deserialize<TestScenario>(yaml);
        }

        static void GenerateDocument(TestScenario scenario, string outputPath)
        {
            switch (scenario.DocumentType)
            {
                case "wordprocessing":
                    var wordBridge = new WordprocessingBridge(scenario);
                    wordBridge.Execute(outputPath);
                    break;
                case "spreadsheet":
                    var excelBridge = new SpreadsheetBridge(scenario);
                    excelBridge.Execute(outputPath);
                    break;
                case "presentation":
                    var pptBridge = new PresentationBridge(scenario);
                    pptBridge.Execute(outputPath);
                    break;
                default:
                    throw new ArgumentException($"Unknown document type: {scenario.DocumentType}");
            }
        }
    }
}

// tests/e2e/bridges/csharp/DocxBridge/WordprocessingBridge.cs

using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Packaging;
using DocumentFormat.OpenXml.Wordprocessing;
using System;
using System.IO;

namespace DocxBridge
{
    public class WordprocessingBridge
    {
        private readonly TestScenario _scenario;

        public WordprocessingBridge(TestScenario scenario)
        {
            _scenario = scenario;
        }

        public void Execute(string outputPath)
        {
            // Create directory if needed
            var directory = Path.GetDirectoryName(outputPath);
            if (!string.IsNullOrEmpty(directory) && !Directory.Exists(directory))
            {
                Directory.CreateDirectory(directory);
            }

            using (var doc = WordprocessingDocument.Create(outputPath, WordprocessingDocumentType.Document))
            {
                // Add main document part
                var mainPart = doc.AddMainDocumentPart();
                mainPart.Document = new Document();
                var body = new Body();
                mainPart.Document.Append(body);

                // Execute operations
                foreach (var op in _scenario.Operations)
                {
                    ExecuteOperation(doc, body, op);
                }

                mainPart.Document.Save();
            }
        }

        private void ExecuteOperation(WordprocessingDocument doc, Body body, Operation op)
        {
            switch (op.Action)
            {
                case "create_document":
                    // Already created in Execute()
                    break;

                case "add_paragraph":
                    AddParagraph(body, op);
                    break;

                case "add_table":
                    AddTable(body, op);
                    break;

                default:
                    throw new ArgumentException($"Unknown action: {op.Action}");
            }
        }

        private void AddParagraph(Body body, Operation op)
        {
            var paragraph = new Paragraph();

            // Simple case: direct text
            if (op.Properties.TryGetValue("text", out var textObj))
            {
                var run = new Run(new Text(textObj.ToString()));
                paragraph.Append(run);
            }
            // Complex case: nested operations (runs)
            else if (op.Properties.TryGetValue("operations", out var opsObj))
            {
                var operations = (opsObj as List<object>);
                foreach (var runOpObj in operations)
                {
                    var runOpMap = runOpObj as Dictionary<object, object>;
                    var action = runOpMap["action"].ToString();

                    if (action == "add_run")
                    {
                        var runText = runOpMap["text"].ToString();
                        var run = new Run(new Text(runText));

                        // Apply run properties
                        if (runOpMap.TryGetValue("properties", out var propsObj))
                        {
                            ApplyRunProperties(run, propsObj as Dictionary<object, object>);
                        }

                        paragraph.Append(run);
                    }
                }
            }

            // Apply paragraph properties
            if (op.Properties.TryGetValue("properties", out var paraPropsObj))
            {
                ApplyParagraphProperties(paragraph, paraPropsObj as Dictionary<object, object>);
            }

            body.Append(paragraph);
        }

        private void ApplyParagraphProperties(Paragraph para, Dictionary<object, object> props)
        {
            if (props == null) return;

            var paraProps = new ParagraphProperties();

            if (props.TryGetValue("alignment", out var alignObj))
            {
                var alignment = alignObj.ToString();
                JustificationValues justification = alignment switch
                {
                    "left" => JustificationValues.Left,
                    "center" => JustificationValues.Center,
                    "right" => JustificationValues.Right,
                    "justify" => JustificationValues.Both,
                    _ => JustificationValues.Left
                };
                paraProps.Append(new Justification() { Val = justification });
            }

            if (props.TryGetValue("spacing_after", out var spacingAfterObj))
            {
                var spacingAfter = Convert.ToInt32(spacingAfterObj);
                var spacing = paraProps.GetFirstChild<SpacingBetweenLines>() ?? new SpacingBetweenLines();
                spacing.After = (spacingAfter * 20).ToString(); // Convert to twips
                paraProps.Append(spacing);
            }

            para.PrependChild(paraProps);
        }

        private void ApplyRunProperties(Run run, Dictionary<object, object> props)
        {
            if (props == null) return;

            var runProps = new RunProperties();

            if (props.TryGetValue("bold", out var boldObj) && Convert.ToBoolean(boldObj))
            {
                runProps.Append(new Bold());
            }

            if (props.TryGetValue("italic", out var italicObj) && Convert.ToBoolean(italicObj))
            {
                runProps.Append(new Italic());
            }

            if (props.TryGetValue("underline", out var underlineObj) && Convert.ToBoolean(underlineObj))
            {
                runProps.Append(new Underline() { Val = UnderlineValues.Single });
            }

            if (props.TryGetValue("color", out var colorObj))
            {
                runProps.Append(new Color() { Val = colorObj.ToString() });
            }

            if (props.TryGetValue("font_size", out var fontSizeObj))
            {
                var fontSize = Convert.ToInt32(fontSizeObj);
                runProps.Append(new FontSize() { Val = fontSize.ToString() }); // Half-points
            }

            if (props.TryGetValue("font_name", out var fontNameObj))
            {
                runProps.Append(new RunFonts() { Ascii = fontNameObj.ToString(), HighAnsi = fontNameObj.ToString() });
            }

            run.PrependChild(runProps);
        }

        private void AddTable(Body body, Operation op)
        {
            var rows = Convert.ToInt32(op.Properties["rows"]);
            var cols = Convert.ToInt32(op.Properties["cols"]);

            var table = new Table();

            // Set table properties
            var tableProps = new TableProperties();

            if (op.Properties.TryGetValue("style", out var styleObj))
            {
                tableProps.Append(new TableStyle() { Val = styleObj.ToString() });
            }

            if (op.Properties.TryGetValue("width_percent", out var widthObj))
            {
                var widthPercent = Convert.ToInt32(widthObj);
                var width = new TableWidth()
                {
                    Type = TableWidthUnitValues.Pct,
                    Width = (widthPercent * 50).ToString() // 5000 = 100%
                };
                tableProps.Append(width);
            }

            table.Append(tableProps);

            // Create table grid
            var tableGrid = new TableGrid();
            for (int c = 0; c < cols; c++)
            {
                tableGrid.Append(new GridColumn());
            }
            table.Append(tableGrid);

            // Create rows
            for (int r = 0; r < rows; r++)
            {
                var row = new TableRow();
                for (int c = 0; c < cols; c++)
                {
                    var cell = new TableCell(new Paragraph(new Run(new Text(""))));
                    row.Append(cell);
                }
                table.Append(row);
            }

            // Set cell values
            if (op.Properties.TryGetValue("cells", out var cellsObj))
            {
                var cells = cellsObj as List<object>;
                foreach (var cellObj in cells)
                {
                    var cellMap = cellObj as Dictionary<object, object>;
                    var row = Convert.ToInt32(cellMap["row"]);
                    var col = Convert.ToInt32(cellMap["col"]);
                    var text = cellMap["text"].ToString();

                    var tableRow = table.Elements<TableRow>().ElementAt(row);
                    var tableCell = tableRow.Elements<TableCell>().ElementAt(col);
                    tableCell.RemoveAllChildren();
                    tableCell.Append(new Paragraph(new Run(new Text(text))));
                }
            }

            body.Append(table);
        }
    }
}

// tests/e2e/bridges/csharp/DocxBridge/TestScenario.cs

using System.Collections.Generic;

namespace DocxBridge
{
    public class TestScenario
    {
        public string Name { get; set; }
        public string Description { get; set; }
        public string DocumentType { get; set; }
        public string Category { get; set; }
        public List<string> Tags { get; set; }
        public ToleranceConfig Tolerance { get; set; }
        public List<Operation> Operations { get; set; }
    }

    public class ToleranceConfig
    {
        public bool XmlAttributeOrderSensitive { get; set; }
        public double VisualPixelTolerance { get; set; }
        public double VisualDiffThreshold { get; set; }
    }

    public class Operation
    {
        public string Action { get; set; }
        public Dictionary<object, object> Properties { get; set; }
    }
}

// tests/e2e/bridges/csharp/DocxBridge/DocxBridge.csproj

<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net9.0</TargetFramework>
    <LangVersion>latest</LangVersion>
    <Nullable>enable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="DocumentFormat.OpenXml" Version="3.2.0" />
    <PackageReference Include="System.CommandLine" Version="2.0.0-beta4.22272.1" />
    <PackageReference Include="YamlDotNet" Version="16.2.1" />
  </ItemGroup>

</Project>
```

### 4. Test Execution Framework

```go
// tests/e2e/framework/executor.go

package framework

import (
    "context"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

// Executor runs test scenarios through both bridges and compares results
type Executor struct {
    goBridge     string // Path to Go bridge executable (if needed)
    csharpBridge string // Path to C# bridge executable
    outputDir    string // Base output directory
    baselineDir  string // Baseline documents directory
    baselineMode bool   // If true, compare against baselines instead of running C# bridge
}

// NewExecutor creates a new test executor
func NewExecutor(csharpBridgePath, outputDir, baselineDir string) *Executor {
    return &Executor{
        csharpBridge: csharpBridgePath,
        outputDir:    outputDir,
        baselineDir:  baselineDir,
    }
}

// ExecuteScenario runs a scenario through both bridges and compares
func (e *Executor) ExecuteScenario(ctx context.Context, scenario *TestScenario) (*TestResult, error) {
    result := &TestResult{
        Scenario:  scenario.Name,
        StartTime: time.Now(),
    }

    // 1. Generate document with Go
    goOutputPath := filepath.Join(e.outputDir, "go", scenario.Name+".docx")
    if err := e.executeGoBridge(ctx, scenario, goOutputPath); err != nil {
        result.GoError = err.Error()
        result.Status = StatusFailed
        result.EndTime = time.Now()
        return result, nil // Don't fail whole test, just record error
    }
    result.GoOutputPath = goOutputPath

    // 2. Determine comparison target
    var comparisonTargetPath string
    if e.baselineMode {
        // Baseline mode: compare against pre-generated baseline
        comparisonTargetPath = filepath.Join(e.baselineDir, scenario.DocumentType, scenario.Name+".docx")
        if _, err := os.Stat(comparisonTargetPath); os.IsNotExist(err) {
            result.ComparisonError = fmt.Sprintf("baseline not found: %s", comparisonTargetPath)
            result.Status = StatusFailed
            result.EndTime = time.Now()
            return result, nil
        }
        result.DotNetOutputPath = comparisonTargetPath
    } else {
        // Normal mode: generate document with C# bridge
        dotnetOutputPath := filepath.Join(e.outputDir, "dotnet", scenario.Name+".docx")
        if err := e.executeCSharpBridge(ctx, scenario, dotnetOutputPath); err != nil {
            result.DotNetError = err.Error()
            result.Status = StatusFailed
            result.EndTime = time.Now()
            return result, nil
        }
        result.DotNetOutputPath = dotnetOutputPath
        comparisonTargetPath = dotnetOutputPath
    }

    // 3. Compare documents (three levels)
    comparisonResult, err := CompareDocuments(goOutputPath, comparisonTargetPath, scenario.Tolerance)
    if err != nil {
        result.ComparisonError = err.Error()
        result.Status = StatusFailed
        result.EndTime = time.Now()
        return result, nil
    }
    result.Comparison = comparisonResult

    // 4. Determine pass/fail
    if comparisonResult.XMLMatch && comparisonResult.BinaryMatch && comparisonResult.VisualMatch {
        result.Status = StatusPassed
    } else {
        result.Status = StatusFailed
    }

    result.EndTime = time.Now()
    return result, nil
}

// executeGoBridge generates document using goffice
func (e *Executor) executeGoBridge(ctx context.Context, scenario *TestScenario, outputPath string) error {
    // Ensure output directory exists
    if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
        return fmt.Errorf("create output directory: %w", err)
    }

    // Import the Go bridge directly (no subprocess needed)
    bridge := gobridge.NewBridge(scenario)
    if err := bridge.Execute(outputPath); err != nil {
        return fmt.Errorf("go bridge execution: %w", err)
    }

    return nil
}

// executeCSharpBridge generates document using Open-XML-SDK
func (e *Executor) executeCSharpBridge(ctx context.Context, scenario *TestScenario, outputPath string) error {
    // Add timeout context (default 30 seconds for C# bridge execution)
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    // Ensure output directory exists
    if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
        return fmt.Errorf("create output directory: %w", err)
    }

    // Write scenario to temp file for C# to read
    scenarioTempPath := filepath.Join(os.TempDir(), fmt.Sprintf("scenario_%s.yaml", scenario.Name))
    scenarioData, err := yaml.Marshal(scenario)
    if err != nil {
        return fmt.Errorf("marshal scenario: %w", err)
    }
    if err := os.WriteFile(scenarioTempPath, scenarioData, 0644); err != nil {
        return fmt.Errorf("write scenario temp file: %w", err)
    }
    defer os.Remove(scenarioTempPath)

    // Execute C# bridge with timeout
    cmd := exec.CommandContext(ctx, e.csharpBridge,
        "--scenario", scenarioTempPath,
        "--output", outputPath)

    output, err := cmd.CombinedOutput()
    if err != nil {
        // Check if timeout occurred
        if ctx.Err() == context.DeadlineExceeded {
            return fmt.Errorf("csharp bridge execution timeout exceeded (30s)")
        }
        return fmt.Errorf("csharp bridge execution: %w (output: %s)", err, string(output))
    }

    return nil
}

// TestResult represents the result of running a test scenario
type TestResult struct {
    Scenario         string
    Status           TestStatus
    StartTime        time.Time
    EndTime          time.Time
    GoOutputPath     string
    DotNetOutputPath string
    GoError          string
    DotNetError      string
    ComparisonError  string
    Comparison       *ComparisonResult
}

// TestStatus represents test execution status
type TestStatus string

const (
    StatusPassed TestStatus = "PASSED"
    StatusFailed TestStatus = "FAILED"
    StatusSkipped TestStatus = "SKIPPED"
)

// ComparisonResult contains results from three-level comparison
type ComparisonResult struct {
    XMLMatch    bool
    BinaryMatch bool
    VisualMatch bool

    XMLDiff     *XMLDiffResult
    BinaryDiff  *BinaryDiffResult
    VisualDiff  *VisualDiffResult
}
```

### 5. Three-Level Comparison Implementation

#### Level 1: XML Structure Comparison
```go
// tests/e2e/comparison/xml.go

package comparison

import (
    "archive/zip"
    "encoding/xml"
    "fmt"
    "io"
    "strings"
)

// XMLDiffResult contains XML comparison results
type XMLDiffResult struct {
    Match             bool
    ElementMismatches []ElementMismatch
    AttributeMismatches []AttributeMismatch
    ContentMismatches []ContentMismatch
}

// ElementMismatch represents a missing or extra element
type ElementMismatch struct {
    Path      string // XPath-style path
    Expected  string // Element name expected (from .NET)
    Actual    string // Element name actual (from Go)
}

// AttributeMismatch represents attribute value difference
type AttributeMismatch struct {
    Path      string // XPath to element
    Attribute string // Attribute name
    Expected  string // Value from .NET
    Actual    string // Value from Go
}

// ContentMismatch represents text content difference
type ContentMismatch struct {
    Path     string
    Expected string
    Actual   string
}

// CompareXMLStructure compares XML structure of two DOCX files
func CompareXMLStructure(goPath, dotnetPath string, tolerance ToleranceConfig) (*XMLDiffResult, error) {
    result := &XMLDiffResult{Match: true}

    // Open both ZIP archives
    goZip, err := zip.OpenReader(goPath)
    if err != nil {
        return nil, fmt.Errorf("open go docx: %w", err)
    }
    defer goZip.Close()

    dotnetZip, err := zip.OpenReader(dotnetPath)
    if err != nil {
        return nil, fmt.Errorf("open dotnet docx: %w", err)
    }
    defer dotnetZip.Close()

    // Compare document.xml (main document)
    goDocXML, err := extractXMLPart(goZip, "word/document.xml")
    if err != nil {
        return nil, fmt.Errorf("extract go document.xml: %w", err)
    }

    dotnetDocXML, err := extractXMLPart(dotnetZip, "word/document.xml")
    if err != nil {
        return nil, fmt.Errorf("extract dotnet document.xml: %w", err)
    }

    // Parse XML to trees
    goTree, err := parseXMLToTree(goDocXML)
    if err != nil {
        return nil, fmt.Errorf("parse go XML: %w", err)
    }

    dotnetTree, err := parseXMLToTree(dotnetDocXML)
    if err != nil {
        return nil, fmt.Errorf("parse dotnet XML: %w", err)
    }

    // Compare trees recursively
    compareNodes(goTree, dotnetTree, "", result, tolerance)

    if len(result.ElementMismatches) > 0 || len(result.AttributeMismatches) > 0 || len(result.ContentMismatches) > 0 {
        result.Match = false
    }

    return result, nil
}

// extractXMLPart extracts an XML part from DOCX ZIP
func extractXMLPart(zipReader *zip.ReadCloser, partName string) ([]byte, error) {
    for _, file := range zipReader.File {
        if file.Name == partName {
            rc, err := file.Open()
            if err != nil {
                return nil, err
            }
            defer rc.Close()

            return io.ReadAll(rc)
        }
    }
    return nil, fmt.Errorf("part not found: %s", partName)
}

// XMLNode represents a simplified XML tree node
type XMLNode struct {
    Name       xml.Name
    Attributes map[string]string
    Content    string
    Children   []*XMLNode
}

// parseXMLToTree parses XML into a tree structure
func parseXMLToTree(data []byte) (*XMLNode, error) {
    decoder := xml.NewDecoder(strings.NewReader(string(data)))

    var root *XMLNode
    var stack []*XMLNode

    for {
        token, err := decoder.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }

        switch t := token.(type) {
        case xml.StartElement:
            node := &XMLNode{
                Name:       t.Name,
                Attributes: make(map[string]string),
            }

            for _, attr := range t.Attr {
                node.Attributes[attr.Name.Local] = attr.Value
            }

            if len(stack) > 0 {
                parent := stack[len(stack)-1]
                parent.Children = append(parent.Children, node)
            } else {
                root = node
            }

            stack = append(stack, node)

        case xml.EndElement:
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }

        case xml.CharData:
            if len(stack) > 0 {
                node := stack[len(stack)-1]
                content := strings.TrimSpace(string(t))
                if content != "" {
                    node.Content += content
                }
            }
        }
    }

    return root, nil
}

// compareNodes recursively compares XML nodes
func compareNodes(goNode, dotnetNode *XMLNode, path string, result *XMLDiffResult, tolerance ToleranceConfig) {
    if goNode == nil && dotnetNode == nil {
        return
    }

    currentPath := path + "/" + dotnetNode.Name.Local

    // Check element names match
    if goNode == nil {
        result.ElementMismatches = append(result.ElementMismatches, ElementMismatch{
            Path:     currentPath,
            Expected: dotnetNode.Name.Local,
            Actual:   "(missing)",
        })
        return
    }

    if goNode.Name.Local != dotnetNode.Name.Local {
        result.ElementMismatches = append(result.ElementMismatches, ElementMismatch{
            Path:     currentPath,
            Expected: dotnetNode.Name.Local,
            Actual:   goNode.Name.Local,
        })
        return
    }

    // Check attributes
    for attrName, dotnetValue := range dotnetNode.Attributes {
        goValue, exists := goNode.Attributes[attrName]
        if !exists {
            result.AttributeMismatches = append(result.AttributeMismatches, AttributeMismatch{
                Path:      currentPath,
                Attribute: attrName,
                Expected:  dotnetValue,
                Actual:    "(missing)",
            })
        } else if goValue != dotnetValue {
            // Skip if attribute order doesn't matter and value is equivalent
            if tolerance.XMLAttributeOrderSensitive || goValue != dotnetValue {
                result.AttributeMismatches = append(result.AttributeMismatches, AttributeMismatch{
                    Path:      currentPath,
                    Attribute: attrName,
                    Expected:  dotnetValue,
                    Actual:    goValue,
                })
            }
        }
    }

    // Check for extra attributes in Go
    for attrName := range goNode.Attributes {
        if _, exists := dotnetNode.Attributes[attrName]; !exists {
            result.AttributeMismatches = append(result.AttributeMismatches, AttributeMismatch{
                Path:      currentPath,
                Attribute: attrName,
                Expected:  "(not present)",
                Actual:    goNode.Attributes[attrName],
            })
        }
    }

    // Check text content
    if goNode.Content != dotnetNode.Content {
        result.ContentMismatches = append(result.ContentMismatches, ContentMismatch{
            Path:     currentPath,
            Expected: dotnetNode.Content,
            Actual:   goNode.Content,
        })
    }

    // Compare children
    if len(goNode.Children) != len(dotnetNode.Children) {
        result.ElementMismatches = append(result.ElementMismatches, ElementMismatch{
            Path:     currentPath,
            Expected: fmt.Sprintf("%d children", len(dotnetNode.Children)),
            Actual:   fmt.Sprintf("%d children", len(goNode.Children)),
        })
        return
    }

    for i := 0; i < len(dotnetNode.Children); i++ {
        compareNodes(goNode.Children[i], dotnetNode.Children[i], currentPath, result, tolerance)
    }
}
```

#### Level 2: Binary Content Comparison
```go
// tests/e2e/comparison/binary.go

package comparison

import (
    "archive/zip"
    "bytes"
    "fmt"
    "io"
)

// BinaryDiffResult contains binary comparison results
type BinaryDiffResult struct {
    Match             bool
    MissingParts      []string // Parts in .NET but not in Go
    ExtraParts        []string // Parts in Go but not in .NET
    DifferentParts    []PartDifference
}

// PartDifference represents a part that exists in both but differs
type PartDifference struct {
    PartName      string
    SizeDiff      int64 // Go size - .NET size
    BytesDifferent int64
}

// CompareBinaryContent compares ZIP structure and binary parts
func CompareBinaryContent(goPath, dotnetPath string) (*BinaryDiffResult, error) {
    result := &BinaryDiffResult{Match: true}

    goZip, err := zip.OpenReader(goPath)
    if err != nil {
        return nil, fmt.Errorf("open go docx: %w", err)
    }
    defer goZip.Close()

    dotnetZip, err := zip.OpenReader(dotnetPath)
    if err != nil {
        return nil, fmt.Errorf("open dotnet docx: %w", err)
    }
    defer dotnetZip.Close()

    // Build maps of parts
    goParts := make(map[string]*zip.File)
    for _, file := range goZip.File {
        goParts[file.Name] = file
    }

    dotnetParts := make(map[string]*zip.File)
    for _, file := range dotnetZip.File {
        dotnetParts[file.Name] = file
    }

    // Find missing parts (in .NET but not in Go)
    for partName := range dotnetParts {
        if _, exists := goParts[partName]; !exists {
            result.MissingParts = append(result.MissingParts, partName)
            result.Match = false
        }
    }

    // Find extra parts (in Go but not in .NET)
    for partName := range goParts {
        if _, exists := dotnetParts[partName]; !exists {
            result.ExtraParts = append(result.ExtraParts, partName)
            result.Match = false
        }
    }

    // Compare common parts
    for partName, dotnetFile := range dotnetParts {
        goFile, exists := goParts[partName]
        if !exists {
            continue // Already recorded as missing
        }

        // Compare sizes
        if goFile.UncompressedSize64 != dotnetFile.UncompressedSize64 {
            diff := PartDifference{
                PartName: partName,
                SizeDiff: int64(goFile.UncompressedSize64) - int64(dotnetFile.UncompressedSize64),
            }

            // If binary part (not XML), compare bytes
            if !isXMLPart(partName) {
                bytesDiff, err := compareBinaryBytes(goFile, dotnetFile)
                if err != nil {
                    return nil, fmt.Errorf("compare bytes for %s: %w", partName, err)
                }
                diff.BytesDifferent = bytesDiff
            }

            result.DifferentParts = append(result.DifferentParts, diff)
            result.Match = false
        }
    }

    return result, nil
}

// isXMLPart returns true if the part is an XML file
func isXMLPart(partName string) bool {
    return strings.HasSuffix(partName, ".xml") ||
           strings.HasSuffix(partName, ".rels")
}

// compareBinaryBytes compares byte content of two ZIP file entries
func compareBinaryBytes(goFile, dotnetFile *zip.File) (int64, error) {
    goReader, err := goFile.Open()
    if err != nil {
        return 0, err
    }
    defer goReader.Close()

    dotnetReader, err := dotnetFile.Open()
    if err != nil {
        return 0, err
    }
    defer dotnetReader.Close()

    goData, err := io.ReadAll(goReader)
    if err != nil {
        return 0, err
    }

    dotnetData, err := io.ReadAll(dotnetReader)
    if err != nil {
        return 0, err
    }

    if bytes.Equal(goData, dotnetData) {
        return 0, nil
    }

    // Count different bytes
    var diff int64
    minLen := len(goData)
    if len(dotnetData) < minLen {
        minLen = len(dotnetData)
    }

    for i := 0; i < minLen; i++ {
        if goData[i] != dotnetData[i] {
            diff++
        }
    }

    // Add bytes from length difference
    diff += int64(abs(len(goData) - len(dotnetData)))

    return diff, nil
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

#### Level 3: Visual Rendering Comparison
```go
// tests/e2e/comparison/visual.go

package comparison

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"

    "github.com/connerohnesorge/goffice/pdf/comparison" // Reuse existing
)

// VisualDiffResult contains visual comparison results
type VisualDiffResult struct {
    Match            bool
    DiffPercentage   float64 // % of pixels that differ
    MaxColorDelta    float64 // Max color distance
    AvgColorDelta    float64 // Avg color distance
    DiffImagePath    string  // Path to generated diff image
}

// CompareVisualRendering converts documents to images and compares pixels
func CompareVisualRendering(goPath, dotnetPath string, tolerance ToleranceConfig) (*VisualDiffResult, error) {
    result := &VisualDiffResult{Match: true}

    // 1. Convert both documents to PDF
    goPDFPath, err := convertOfficeToPDF(goPath)
    if err != nil {
        return nil, fmt.Errorf("convert go document to PDF: %w", err)
    }
    defer os.Remove(goPDFPath) // Cleanup temp PDF

    dotnetPDFPath, err := convertOfficeToPDF(dotnetPath)
    if err != nil {
        return nil, fmt.Errorf("convert dotnet document to PDF: %w", err)
    }
    defer os.Remove(dotnetPDFPath)

    // 2. Convert PDFs to PNG
    config := comparison.DefaultComparisonConfig()
    config.PixelTolerance = tolerance.VisualPixelTolerance
    config.ImageDiffThreshold = tolerance.VisualDiffThreshold
    config.DPI = 150 // Standard DPI for comparison

    goPNGPath := strings.TrimSuffix(goPDFPath, ".pdf") + ".png"
    _, err = comparison.ConvertPDFToPNG(goPDFPath, goPNGPath, config.DPI)
    if err != nil {
        return nil, fmt.Errorf("convert go PDF to PNG: %w", err)
    }
    defer os.Remove(goPNGPath)

    dotnetPNGPath := strings.TrimSuffix(dotnetPDFPath, ".pdf") + ".png"
    _, err = comparison.ConvertPDFToPNG(dotnetPDFPath, dotnetPNGPath, config.DPI)
    if err != nil {
        return nil, fmt.Errorf("convert dotnet PDF to PNG: %w", err)
    }
    defer os.Remove(dotnetPNGPath)

    // 3. Compare images

    compResult, err := comparison.CompareImages(goPNGPath, dotnetPNGPath, config)
    if err != nil {
        return nil, fmt.Errorf("compare images: %w", err)
    }

    result.DiffPercentage = compResult.DiffPercentage()
    result.MaxColorDelta = compResult.MaxColorDelta
    result.AvgColorDelta = compResult.AvgColorDelta
    result.Match = compResult.IsWithinTolerance()

    // 4. Generate diff image if mismatch
    if !result.Match {
        diffPath := filepath.Join("reports", "diffs", filepath.Base(goPath)+".diff.png")
        if err := comparison.AnnotateDiffImage(dotnetPNGPath, goPNGPath, diffPath, config); err != nil {
            return nil, fmt.Errorf("generate diff image: %w", err)
        }
        result.DiffImagePath = diffPath
    }

    return result, nil
}

// convertOfficeToPDF converts an Office document to PDF
// Uses LibreOffice headless or system Office installation
func convertOfficeToPDF(docPath string) (string, error) {
    pdfPath := strings.TrimSuffix(docPath, filepath.Ext(docPath)) + ".pdf"

    // Add timeout for LibreOffice conversion (60 seconds)
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()

    // Try LibreOffice headless conversion with timeout
    cmd := exec.CommandContext(ctx,
        "soffice",
        "--headless",
        "--convert-to", "pdf",
        "--outdir", filepath.Dir(pdfPath),
        docPath)

    output, err := cmd.CombinedOutput()
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            return "", fmt.Errorf("libreoffice conversion timeout exceeded (60s)")
        }
        return "", fmt.Errorf("libreoffice conversion failed: %w (output: %s)", err, string(output))
    }

    // Check PDF was created
    if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
        return "", fmt.Errorf("PDF not created: %s", pdfPath)
    }

    return pdfPath, nil
}
```

---

## Nix Integration

### Extended flake.nix
```nix
# Additions to existing flake.nix

{
  description = "A development shell for go + .NET cross-runtime testing";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    treefmt-nix.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    treefmt-nix,
    ...
  }: flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [
          (final: prev: {
            final.buildGoModule = prev.buildGo125Module;
            buildGoModule = prev.buildGo125Module;
          })
        ];
      };

      # ... existing rooted helper ...

      scripts = {
        # ... existing scripts (dx, lint) ...

        # New E2E test scripts
        test-go = {
          exec = rooted ''
            echo "Running Go tests..."
            gotestsum --format short-verbose "$REPO_ROOT"/... --timeout=5m
          '';
          description = "Run Go tests only";
          deps = [pkgs.gotestsum];
        };

        test-dotnet = {
          exec = rooted ''
            echo "Running .NET Open-XML-SDK tests..."
            cd "$REPO_ROOT"/Open-XML-SDK
            dotnet test --verbosity normal
          '';
          description = "Run .NET SDK tests only";
          deps = [pkgs.dotnet-sdk_9];
        };

        test-e2e = {
          exec = rooted ''
            echo "Running E2E cross-runtime tests..."
            cd "$REPO_ROOT"/tests/e2e

            # Build C# bridge
            echo "Building C# bridge..."
            cd bridges/csharp/DocxBridge
            dotnet build --configuration Release
            cd ../../../

            # Run E2E tests
            go test ./cmd/e2e-runner -v -timeout=15m
          '';
          description = "Run cross-runtime E2E tests";
          deps = [pkgs.go_1_25 pkgs.dotnet-sdk_9 pkgs.ghostscript pkgs.libreoffice];
        };

        test-all = {
          exec = rooted ''
            set -e
            echo "================================"
            echo "Running complete test suite..."
            echo "================================"

            echo ""
            echo "1. Go unit tests..."
            gotestsum --format short-verbose "$REPO_ROOT"/... --timeout=5m

            echo ""
            echo "2. .NET SDK tests..."
            cd "$REPO_ROOT"/Open-XML-SDK
            dotnet test --verbosity quiet

            echo ""
            echo "3. E2E cross-runtime tests..."
            cd "$REPO_ROOT"/tests/e2e
            cd bridges/csharp/DocxBridge && dotnet build --configuration Release && cd ../../../
            go test ./cmd/e2e-runner -v -timeout=15m

            echo ""
            echo "================================"
            echo "All tests passed!"
            echo "================================"
          '';
          description = "Run all tests (Go + .NET + E2E)";
          deps = [pkgs.gotestsum pkgs.dotnet-sdk_9 pkgs.go_1_25 pkgs.ghostscript pkgs.libreoffice];
        };

        update-baselines = {
          exec = rooted ''
            echo "Updating E2E test baselines from .NET SDK..."
            cd "$REPO_ROOT"/tests/e2e

            # Build C# bridge
            cd bridges/csharp/DocxBridge
            dotnet build --configuration Release
            cd ../../../

            # Run baseline generation
            go run ./cmd/e2e-runner --generate-baselines

            echo ""
            echo "Baselines updated. Review changes with:"
            echo "  git diff tests/e2e/baselines/"
          '';
          description = "Regenerate test baselines from .NET SDK";
          deps = [pkgs.go_1_25 pkgs.dotnet-sdk_9];
        };

        # ... existing tests script ...
      };

      # ... existing scriptPackages ...

      # ... existing treefmtModule ...

    in {
      devShells.default = pkgs.mkShell {
        name = "dev";

        packages = with pkgs; [
          # ... existing packages (Go, Nix tools, etc.) ...

          # .NET SDK for cross-runtime testing
          dotnet-sdk_9               # Latest .NET SDK

          # Office document conversion tools
          libreoffice                # For Office→PDF conversion
          ghostscript                # For PDF→PNG conversion (already used)

          # ... existing script packages ...
        ]
        ++ builtins.attrValues scriptPackages;
      };

      # ... existing packages, formatter ...
    });
}
```

---

## Test Execution Flow

### CLI Entry Point
```go
// tests/e2e/cmd/e2e-runner/main.go

package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/connerohnesorge/goffice-e2e/framework"
)

func main() {
    var (
        scenariosDir    = flag.String("scenarios", "scenarios/wordprocessing", "Scenarios directory")
        outputDir       = flag.String("output", "output", "Output directory")
        baselineDir     = flag.String("baselines", "baselines", "Baselines directory")
        csharpBridge    = flag.String("csharp-bridge", "bridges/csharp/DocxBridge/bin/Release/net9.0/DocxBridge", "C# bridge executable")
        generateBaselines = flag.Bool("generate-baselines", false, "Generate new baselines from .NET SDK")
        htmlReport      = flag.Bool("html-report", true, "Generate HTML report")
    )
    flag.Parse()

    ctx := context.Background()

    // Load scenarios
    scenarios, err := framework.LoadScenariosFromDirectory(*scenariosDir)
    if err != nil {
        log.Fatalf("Failed to load scenarios: %v", err)
    }

    log.Printf("Loaded %d scenarios from %s", len(scenarios), *scenariosDir)

    // Create executor
    executor := framework.NewExecutor(*csharpBridge, *outputDir, *baselineDir)

    // Execute all scenarios
    var results []*framework.TestResult
    for _, scenario := range scenarios {
        log.Printf("Executing scenario: %s", scenario.Name)

        result, err := executor.ExecuteScenario(ctx, scenario)
        if err != nil {
            log.Printf("  ERROR: %v", err)
            continue
        }

        log.Printf("  Status: %s", result.Status)
        results = append(results, result)
    }

    // Generate report
    if *htmlReport {
        reportPath := filepath.Join("reports", "html", "index.html")
        if err := framework.GenerateHTMLReport(results, reportPath); err != nil {
            log.Fatalf("Failed to generate HTML report: %v", err)
        }
        log.Printf("HTML report generated: %s", reportPath)
    }

    // Print summary
    printSummary(results)

    // Exit with error if any tests failed
    for _, result := range results {
        if result.Status == framework.StatusFailed {
            os.Exit(1)
        }
    }
}

func printSummary(results []*framework.TestResult) {
    var passed, failed, skipped int
    for _, result := range results {
        switch result.Status {
        case framework.StatusPassed:
            passed++
        case framework.StatusFailed:
            failed++
        case framework.StatusSkipped:
            skipped++
        }
    }

    fmt.Println("\n================================")
    fmt.Printf("Test Results Summary\n")
    fmt.Println("================================")
    fmt.Printf("Total:   %d\n", len(results))
    fmt.Printf("Passed:  %d\n", passed)
    fmt.Printf("Failed:  %d\n", failed)
    fmt.Printf("Skipped: %d\n", skipped)
    fmt.Println("================================")
}
```

---

## API Mapping Patterns

### Go API → .NET SDK Mapping Reference

This section provides exact mappings between goffice and Open-XML-SDK APIs for common operations:

#### Document Creation
```go
// Go (goffice)
doc, err := wordprocessing.New("document.docx", wordprocessing.DocTypeDocument)
```
```csharp
// C# (Open-XML-SDK)
using var doc = WordprocessingDocument.Create("document.docx", WordprocessingDocumentType.Document);
```

#### Add Paragraph with Text
```go
// Go
builder := wordprocessing.NewDocumentBuilder()
builder.AddParagraph("Hello World")
```
```csharp
// C#
var para = new Paragraph(new Run(new Text("Hello World")));
body.Append(para);
```

#### Text Formatting (Bold, Italic, Color)
```go
// Go
pb := builder.AddParagraph("")
rb := pb.AddRun("Bold text")
rb.Bold().Color("FF0000").FontSize(24)
```
```csharp
// C#
var run = new Run(new Text("Bold text"));
var runProps = new RunProperties(
    new Bold(),
    new Color() { Val = "FF0000" },
    new FontSize() { Val = "24" }
);
run.PrependChild(runProps);
```

#### Tables
```go
// Go
tb := builder.AddTable(2, 3)
tb.SetCellText(0, 0, "Cell 1")
tb.SetStyle("TableGrid")
```
```csharp
// C#
var table = new Table();
var tableProps = new TableProperties(
    new TableStyle() { Val = "TableGrid" }
);
table.Append(tableProps);

var row = new TableRow();
var cell = new TableCell(new Paragraph(new Run(new Text("Cell 1"))));
row.Append(cell);
table.Append(row);
```

---

This design document provides the foundation for implementing the complete cross-runtime testing infrastructure. All code patterns are based on the actual APIs from both goffice and Open-XML-SDK as researched.
