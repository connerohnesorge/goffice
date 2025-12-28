# ML Paper Results Excel Example

This example demonstrates creating a complex, realistic machine learning research paper results spreadsheet using goffice's spreadsheet capabilities.

## Features Demonstrated

### Data Organization
- **4 sheets** with different types of ML research data
- **Multiple data types**: strings, numbers, formulas
- **Professional headers** with consistent naming
- **Realistic ML benchmarks** for neural network models

### Sheet 1: Model Results
- 10 state-of-the-art deep learning models (ResNet, EfficientNet, Vision Transformer, YOLO, etc.)
- 8 performance metrics per model:
  - Parameters (M)
  - Training Time (hrs)
  - Inference Speed (imgs/sec)
  - ImageNet Top-1 Accuracy %
  - ImageNet Top-5 Accuracy %
  - COCO mAP (Mean Average Precision)
  - Memory (GB)

### Sheet 2: Analysis
- Computed columns using Excel formulas:
  - `Efficiency Score = Accuracy / Parameters`
  - `Params-per-Accuracy = Parameters / (Accuracy - 50)`
  - `Accuracy Gain vs ResNet-50` (baseline reference)
  - `Memory per Accuracy` (efficiency metric)
- Demonstrates formula-based derived metrics

### Sheet 3: Training Curve
- 9 epochs of training progress data
- Metrics tracked:
  - Training Loss and Validation Loss
  - Training Accuracy and Validation Accuracy
  - Learning Rate schedule
  - Overfitting Gap (formula: `Val Loss - Train Loss`)
- Realistic training dynamics showing convergence

### Sheet 4: Hyperparameter Search
- 8 hyperparameter tuning experiments
- Variables:
  - Learning Rate
  - Batch Size
  - Weight Decay
  - Warmup Steps
- Results:
  - Final Accuracy
  - Convergence Time (hrs)
  - Accuracy per Hour (formula: `Accuracy / Time`)

### Charts
- **Bar Chart**: Model accuracy comparison (ImageNet Top-1 %)
- **Line Chart**: Training vs validation accuracy over epochs
- **Scatter Plot**: Parameters vs accuracy trade-off analysis

## Running the Example

```bash
cd examples/ml_results
go run ml_results_example.go
```

This generates `ml_results_example.xlsx` with all data and charts in the current directory.

## Testing with LibreOffice

```bash
# View the spreadsheet in LibreOffice Calc
soffice --headless --calc ml_results_example.xlsx

# Convert to PDF (for sharing/archiving)
soffice --headless --convert-to pdf ml_results_example.xlsx
```

## Code Highlights

### Creating Sheets and Cells
```go
sheet, err := doc.AddSheet("Model Results")
sheet.Cell("A1").SetString("Model Name")
sheet.Cell("B1").SetString("Accuracy %")
sheet.Cell("A2").SetString("ResNet-50")
sheet.Cell("B2").SetNumber(76.1)
```

### Using Formulas
```go
// Compute efficiency score
row := 4
efficiencyFormula := fmt.Sprintf("=C%d/B%d", row, row)
sheet.Cell(fmt.Sprintf("D%d", row)).SetFormula(efficiencyFormula)
```

### Creating Charts
```go
chart, err := sheet.AddChart(
    spreadsheet.ChartTypeBar,
    "E5:E14",  // data range
    spreadsheet.CellRef{Row: 15, Col: 1},  // anchor
)
chart.SetTitle("Model Accuracy Comparison")
chart.AddSeries("Accuracy %", "E5:E14")
```

## Realistic Data

All data reflects actual or plausible performance metrics for real models:
- ResNet-50: Well-established baseline
- EfficientNet-B4: High accuracy with moderate parameters
- Vision Transformer: High accuracy, large model
- YOLO-v8: Fast inference, good accuracy
- Swin Transformer: Strong performance
- ConvNeXt: Modern efficient architecture

## File Output

- `ml_results_example.xlsx` (9.2 KB): Complete spreadsheet with 4 sheets, 10+ models, 40+ data points, multiple formulas, and 3 charts
- `ml_results_example.pdf` (90 KB): PDF export for sharing and archiving

## Library Capabilities Tested

✅ Multiple sheets  
✅ String and number values  
✅ Excel formulas (=, cell references, arithmetic)  
✅ Charts (bar, line, scatter)  
✅ Chart series with data ranges  
✅ Cell references (A1-style notation)  
✅ LibreOffice CLI compatibility  
✅ PDF export via LibreOffice  

## Notes

This example showcases goffice's ability to create complex, professional Excel documents programmatically without requiring Microsoft Office or OpenOffice installed on the development machine.
