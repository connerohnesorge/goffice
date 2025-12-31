using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Packaging;
using DocumentFormat.OpenXml.Wordprocessing;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

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
                var run = new Run(new Text(textObj.ToString() ?? string.Empty));
                paragraph.Append(run);
            }
            // Complex case: nested operations (runs)
            else if (op.Properties.TryGetValue("operations", out var opsObj))
            {
                var operations = (opsObj as List<object>);
                if (operations != null)
                {
                    foreach (var runOpObj in operations)
                    {
                        var runOpMap = runOpObj as Dictionary<object, object>;
                        if (runOpMap == null) continue;

                        var action = runOpMap["action"].ToString();

                        if (action == "add_run")
                        {
                            var runText = runOpMap["text"].ToString() ?? string.Empty;
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
            }

            // Apply paragraph properties
            if (op.Properties.TryGetValue("properties", out var paraPropsObj))
            {
                ApplyParagraphProperties(paragraph, paraPropsObj as Dictionary<object, object>);
            }

            body.Append(paragraph);
        }

        private void ApplyParagraphProperties(Paragraph para, Dictionary<object, object>? props)
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
                if (paraProps.GetFirstChild<SpacingBetweenLines>() == null)
                {
                    paraProps.Append(spacing);
                }
            }

            if (props.TryGetValue("spacing_before", out var spacingBeforeObj))
            {
                var spacingBefore = Convert.ToInt32(spacingBeforeObj);
                var spacing = paraProps.GetFirstChild<SpacingBetweenLines>() ?? new SpacingBetweenLines();
                spacing.Before = (spacingBefore * 20).ToString(); // Convert to twips
                if (paraProps.GetFirstChild<SpacingBetweenLines>() == null)
                {
                    paraProps.Append(spacing);
                }
            }

            if (props.TryGetValue("left_indent", out var leftIndentObj))
            {
                var leftIndent = Convert.ToInt32(leftIndentObj);
                var indentation = new Indentation() { Left = (leftIndent * 20).ToString() }; // Convert to twips
                paraProps.Append(indentation);
            }

            if (paraProps.HasChildren)
            {
                para.PrependChild(paraProps);
            }
        }

        private void ApplyRunProperties(Run run, Dictionary<object, object>? props)
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
                // Font size in YAML is in points, Open-XML-SDK expects half-points
                runProps.Append(new FontSize() { Val = (fontSize * 2).ToString() });
            }

            if (props.TryGetValue("font_name", out var fontNameObj))
            {
                runProps.Append(new RunFonts() { Ascii = fontNameObj.ToString(), HighAnsi = fontNameObj.ToString() });
            }

            if (runProps.HasChildren)
            {
                run.PrependChild(runProps);
            }
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

            if (tableProps.HasChildren)
            {
                table.Append(tableProps);
            }

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
                if (cells != null)
                {
                    foreach (var cellObj in cells)
                    {
                        var cellMap = cellObj as Dictionary<object, object>;
                        if (cellMap == null) continue;

                        var row = Convert.ToInt32(cellMap["row"]);
                        var col = Convert.ToInt32(cellMap["col"]);
                        var text = cellMap["text"].ToString() ?? string.Empty;

                        var tableRow = table.Elements<TableRow>().ElementAt(row);
                        var tableCell = tableRow.Elements<TableCell>().ElementAt(col);
                        tableCell.RemoveAllChildren();
                        tableCell.Append(new Paragraph(new Run(new Text(text))));
                    }
                }
            }

            body.Append(table);
        }
    }
}
