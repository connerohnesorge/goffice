using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Drawing;
using DocumentFormat.OpenXml.Presentation;
using PptxGenerator.Models;
using A = DocumentFormat.OpenXml.Drawing;
using P = DocumentFormat.OpenXml.Presentation;

namespace PptxGenerator.Generators;

public class TextGenerator
{
    public P.Shape GenerateTextBox(
        TextSpec textSpec,
        Models.Position position,
        Models.Size size,
        uint shapeId)
    {
        var shape = new P.Shape();

        // Non-visual shape properties
        shape.Append(new P.NonVisualShapeProperties(
            new P.NonVisualDrawingProperties { Id = shapeId, Name = $"TextBox {shapeId}" },
            new P.NonVisualShapeDrawingProperties(new A.ShapeLocks { NoGrouping = true }),
            new ApplicationNonVisualDrawingProperties { PlaceholderShape = new PlaceholderShape() }));

        // Shape properties
        var spPr = new P.ShapeProperties();
        spPr.Append(new A.Transform2D(
            new A.Offset { X = position.X, Y = position.Y },
            new A.Extents { Cx = size.Width, Cy = size.Height }));
        spPr.Append(new A.PresetGeometry(new A.AdjustValueList()) { Preset = ShapeTypeValues.Rectangle });
        spPr.Append(new A.NoFill());
        shape.Append(spPr);

        // Text body
        var textBody = GenerateTextBody(textSpec);
        shape.Append(textBody);

        return shape;
    }

    public P.TextBody GenerateTextBody(TextSpec textSpec)
    {
        var textBody = new P.TextBody();
        textBody.Append(new A.BodyProperties());
        textBody.Append(new A.ListStyle());

        // Generate paragraphs
        if (textSpec.Paragraphs != null && textSpec.Paragraphs.Count > 0)
        {
            foreach (var paraSpec in textSpec.Paragraphs)
            {
                textBody.Append(GenerateParagraph(paraSpec));
            }
        }
        else if (!string.IsNullOrEmpty(textSpec.Content))
        {
            // Simple text content - create single paragraph with default formatting
            var para = new A.Paragraph();
            var run = GenerateRun(new RunSpec
            {
                Text = textSpec.Content,
                FontFamily = textSpec.DefaultRun.FontFamily,
                FontSize = textSpec.DefaultRun.FontSize,
                Color = textSpec.DefaultRun.Color,
                Bold = textSpec.DefaultRun.Bold,
                Italic = textSpec.DefaultRun.Italic,
                Underline = textSpec.DefaultRun.Underline,
                Strikethrough = textSpec.DefaultRun.Strikethrough
            });
            para.Append(run);
            textBody.Append(para);
        }
        else
        {
            // Empty paragraph
            textBody.Append(new A.Paragraph());
        }

        return textBody;
    }

    private A.Paragraph GenerateParagraph(ParagraphSpec paraSpec)
    {
        var para = new A.Paragraph();

        // Paragraph properties
        var pPr = new A.ParagraphProperties();

        // Alignment
        if (!string.IsNullOrEmpty(paraSpec.Alignment))
        {
            var alignment = paraSpec.Alignment switch
            {
                "left" => A.TextAlignmentTypeValues.Left,
                "center" => A.TextAlignmentTypeValues.Center,
                "right" => A.TextAlignmentTypeValues.Right,
                "justify" => A.TextAlignmentTypeValues.Justified,
                _ => A.TextAlignmentTypeValues.Left
            };
            pPr.Alignment = alignment;
        }

        // Line spacing
        if (paraSpec.LineSpacing > 0)
        {
            pPr.Append(new A.LineSpacing(
                new A.SpacingPercent { Val = (int)(paraSpec.LineSpacing * 100000) }));
        }

        // Space before/after
        if (paraSpec.SpaceBefore > 0)
        {
            pPr.Append(new A.SpaceBefore(new A.SpacingPoints { Val = (int)paraSpec.SpaceBefore }));
        }
        if (paraSpec.SpaceAfter > 0)
        {
            pPr.Append(new A.SpaceAfter(new A.SpacingPoints { Val = (int)paraSpec.SpaceAfter }));
        }

        // Indent
        if (paraSpec.Indent > 0)
        {
            pPr.Indent = (int)paraSpec.Indent;
        }

        // Bullet
        if (!string.IsNullOrEmpty(paraSpec.BulletChar))
        {
            pPr.Append(new A.CharacterBullet { Char = paraSpec.BulletChar });
        }

        if (pPr.HasAttributes || pPr.HasChildren)
        {
            para.Append(pPr);
        }

        // Text runs
        foreach (var runSpec in paraSpec.Runs)
        {
            para.Append(GenerateRun(runSpec));
        }

        return para;
    }

    private A.Run GenerateRun(RunSpec runSpec)
    {
        var run = new A.Run();

        // Run properties
        var rPr = new A.RunProperties { Language = "en-US" };

        // Font size (convert points to half-points * 100)
        if (runSpec.FontSize > 0)
        {
            rPr.FontSize = runSpec.FontSize * 100;
        }

        // Font family
        if (!string.IsNullOrEmpty(runSpec.FontFamily))
        {
            rPr.Append(new A.LatinFont { Typeface = runSpec.FontFamily });
        }

        // Color
        if (!string.IsNullOrEmpty(runSpec.Color))
        {
            rPr.Append(new A.SolidFill(new A.RgbColorModelHex { Val = runSpec.Color }));
        }

        // Bold
        if (runSpec.Bold)
        {
            rPr.Bold = true;
        }

        // Italic
        if (runSpec.Italic)
        {
            rPr.Italic = true;
        }

        // Underline
        if (runSpec.Underline != "none")
        {
            var underlineValue = runSpec.Underline switch
            {
                "single" => A.TextUnderlineValues.Single,
                "double" => A.TextUnderlineValues.Double,
                "dotted" => A.TextUnderlineValues.Dotted,
                "dash" => A.TextUnderlineValues.Dash,
                _ => A.TextUnderlineValues.None
            };
            rPr.Underline = underlineValue;
        }

        // Strikethrough
        if (runSpec.Strikethrough)
        {
            rPr.Strike = A.TextStrikeValues.SingleStrike;
        }

        run.Append(rPr);

        // Text content
        run.Append(new A.Text { Text = runSpec.Text });

        return run;
    }
}
