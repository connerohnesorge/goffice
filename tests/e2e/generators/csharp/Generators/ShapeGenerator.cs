using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Drawing;
using DocumentFormat.OpenXml.Presentation;
using PptxGenerator.Models;
using A = DocumentFormat.OpenXml.Drawing;
using P = DocumentFormat.OpenXml.Presentation;

namespace PptxGenerator.Generators;

public class ShapeGenerator
{
    public P.Shape GenerateShape(
        ShapeSpec shapeSpec,
        Models.Position position,
        Models.Size size,
        uint shapeId)
    {
        var shape = new P.Shape();

        // Non-visual shape properties
        shape.Append(new P.NonVisualShapeProperties(
            new P.NonVisualDrawingProperties { Id = shapeId, Name = $"Shape {shapeId}" },
            new P.NonVisualShapeDrawingProperties(),
            new ApplicationNonVisualDrawingProperties()));

        // Shape properties
        var spPr = new P.ShapeProperties();

        // Transform
        spPr.Append(new A.Transform2D(
            new A.Offset { X = position.X, Y = position.Y },
            new A.Extents { Cx = size.Width, Cy = size.Height }));

        // Preset geometry
        var presetGeometry = GetPresetGeometry(shapeSpec.Type);
        spPr.Append(presetGeometry);

        // Fill
        var fillGenerator = new FillGenerator();
        var fill = fillGenerator.GenerateFill(shapeSpec.Fill);
        spPr.Append(fill);

        // Stroke/Outline
        if (shapeSpec.Stroke.Width > 0)
        {
            var outline = new A.Outline { Width = (int)shapeSpec.Stroke.Width };

            if (!string.IsNullOrEmpty(shapeSpec.Stroke.Color))
            {
                outline.Append(new A.SolidFill(new A.RgbColorModelHex { Val = shapeSpec.Stroke.Color }));
            }

            // Dash style
            if (shapeSpec.Stroke.DashStyle != "solid")
            {
                var presetDash = GetDashStyle(shapeSpec.Stroke.DashStyle);
                outline.Append(new A.PresetDash { Val = presetDash });
            }

            spPr.Append(outline);
        }

        // Effects
        if (shapeSpec.Effects.Count > 0)
        {
            var effectList = new A.EffectList();
            foreach (var effectSpec in shapeSpec.Effects)
            {
                AddEffect(effectList, effectSpec);
            }
            spPr.Append(effectList);
        }

        shape.Append(spPr);

        // Text body (if shape contains text)
        if (shapeSpec.Text != null)
        {
            var textGen = new TextGenerator();
            var textBody = textGen.GenerateTextBody(shapeSpec.Text);
            shape.Append(textBody);
        }
        else
        {
            // Empty text body
            shape.Append(new P.TextBody(
                new A.BodyProperties(),
                new A.ListStyle()));
        }

        return shape;
    }

    private A.PresetGeometry GetPresetGeometry(string shapeType)
    {
        var shapeTypeValue = shapeType switch
        {
            "rect" => ShapeTypeValues.Rectangle,
            "roundRect" => ShapeTypeValues.RoundRectangle,
            "ellipse" => ShapeTypeValues.Ellipse,
            "triangle" => ShapeTypeValues.Triangle,
            "diamond" => ShapeTypeValues.Diamond,
            "pentagon" => ShapeTypeValues.Pentagon,
            "hexagon" => ShapeTypeValues.Hexagon,
            "octagon" => ShapeTypeValues.Octagon,
            "star5" => ShapeTypeValues.Star5,
            "rightArrow" => ShapeTypeValues.RightArrow,
            "wedgeRectCallout" => ShapeTypeValues.WedgeRectangleCallout,
            _ => ShapeTypeValues.Rectangle
        };

        return new A.PresetGeometry(new A.AdjustValueList())
        {
            Preset = shapeTypeValue
        };
    }

    private A.PresetLineDashValues GetDashStyle(string dashStyle)
    {
        return dashStyle switch
        {
            "dot" => A.PresetLineDashValues.Dot,
            "dash" => A.PresetLineDashValues.Dash,
            "dashDot" => A.PresetLineDashValues.DashDot,
            "longDash" => A.PresetLineDashValues.LargeDash,
            "longDashDot" => A.PresetLineDashValues.LargeDashDot,
            "longDashDotDot" => A.PresetLineDashValues.LargeDashDotDot,
            _ => A.PresetLineDashValues.Solid
        };
    }

    private void AddEffect(A.EffectList effectList, EffectSpec effectSpec)
    {
        switch (effectSpec.Type)
        {
            case "shadow":
                if (effectSpec.Shadow != null)
                {
                    AddShadow(effectList, effectSpec.Shadow);
                }
                break;
            case "glow":
                if (effectSpec.Glow != null)
                {
                    AddGlow(effectList, effectSpec.Glow);
                }
                break;
            case "reflection":
                if (effectSpec.Reflection != null)
                {
                    AddReflection(effectList, effectSpec.Reflection);
                }
                break;
            case "soft_edge":
                if (effectSpec.SoftEdge != null)
                {
                    AddSoftEdge(effectList, effectSpec.SoftEdge);
                }
                break;
        }
    }

    private void AddShadow(A.EffectList effectList, ShadowSpec shadowSpec)
    {
        if (shadowSpec.Type == "outer")
        {
            var shadow = new A.OuterShadow
            {
                BlurRadius = shadowSpec.BlurRadius,
                Distance = shadowSpec.Distance,
                Direction = shadowSpec.Angle * 60000 // Convert degrees to 1/60000th of a degree
            };

            if (!string.IsNullOrEmpty(shadowSpec.Color))
            {
                var color = new A.RgbColorModelHex { Val = shadowSpec.Color };
                if (shadowSpec.Transparency > 0)
                {
                    color.Append(new A.Alpha { Val = (int)((1.0 - shadowSpec.Transparency) * 100000) });
                }
                shadow.Append(color);
            }

            effectList.Append(shadow);
        }
        else if (shadowSpec.Type == "inner")
        {
            var shadow = new A.InnerShadow
            {
                BlurRadius = shadowSpec.BlurRadius,
                Distance = shadowSpec.Distance,
                Direction = shadowSpec.Angle * 60000
            };

            if (!string.IsNullOrEmpty(shadowSpec.Color))
            {
                var color = new A.RgbColorModelHex { Val = shadowSpec.Color };
                if (shadowSpec.Transparency > 0)
                {
                    color.Append(new A.Alpha { Val = (int)((1.0 - shadowSpec.Transparency) * 100000) });
                }
                shadow.Append(color);
            }

            effectList.Append(shadow);
        }
    }

    private void AddGlow(A.EffectList effectList, GlowSpec glowSpec)
    {
        var glow = new A.Glow { Radius = glowSpec.Radius };

        if (!string.IsNullOrEmpty(glowSpec.Color))
        {
            glow.Append(new A.RgbColorModelHex { Val = glowSpec.Color });
        }

        effectList.Append(glow);
    }

    private void AddReflection(A.EffectList effectList, ReflectionSpec reflectionSpec)
    {
        var reflection = new A.Reflection
        {
            BlurRadius = reflectionSpec.BlurRadius,
            StartOpacity = (int)(reflectionSpec.StartOpacity * 100000),
            EndAlpha = (int)(reflectionSpec.EndOpacity * 100000),
            Distance = reflectionSpec.Distance,
            Direction = reflectionSpec.Direction * 60000,
            FadeDirection = reflectionSpec.FadeDirection * 60000,
            StartPosition = (int)(reflectionSpec.StartPosition * 100000),
            EndPosition = (int)(reflectionSpec.EndPosition * 100000)
        };

        effectList.Append(reflection);
    }

    private void AddSoftEdge(A.EffectList effectList, SoftEdgeSpec softEdgeSpec)
    {
        var softEdge = new A.SoftEdge { Radius = softEdgeSpec.Radius };
        effectList.Append(softEdge);
    }
}
