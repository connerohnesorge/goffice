using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Drawing;
using PptxGenerator.Models;
using A = DocumentFormat.OpenXml.Drawing;

namespace PptxGenerator.Generators;

public class FillGenerator
{
    public OpenXmlElement GenerateFill(FillSpec fillSpec)
    {
        return fillSpec.Type switch
        {
            "solid" => GenerateSolidFill(fillSpec),
            "gradient" => GenerateGradientFill(fillSpec),
            "pattern" => GeneratePatternFill(fillSpec),
            "none" => new A.NoFill(),
            _ => new A.NoFill()
        };
    }

    private A.SolidFill GenerateSolidFill(FillSpec fillSpec)
    {
        var solidFill = new A.SolidFill();

        if (!string.IsNullOrEmpty(fillSpec.Color))
        {
            solidFill.Append(new A.RgbColorModelHex { Val = fillSpec.Color });
        }

        return solidFill;
    }

    private A.GradientFill GenerateGradientFill(FillSpec fillSpec)
    {
        var gradientFill = new A.GradientFill();

        if (fillSpec.Gradient != null)
        {
            var gradientStopList = new A.GradientStopList();

            foreach (var stop in fillSpec.Gradient.Stops)
            {
                var gradientStop = new A.GradientStop
                {
                    Position = (int)(stop.Position * 100000)
                };
                gradientStop.Append(new A.RgbColorModelHex { Val = stop.Color });
                gradientStopList.Append(gradientStop);
            }

            gradientFill.Append(gradientStopList);

            // Linear or radial gradient
            if (fillSpec.Gradient.Type == "linear")
            {
                var linearGradientFill = new A.LinearGradientFill
                {
                    Angle = fillSpec.Gradient.Angle * 60000, // Convert degrees to 1/60000th of a degree
                    Scaled = true
                };
                gradientFill.Append(linearGradientFill);
            }
            else if (fillSpec.Gradient.Type == "radial")
            {
                var pathGradientFill = new A.PathGradientFill { Path = PathShadeValues.Circle };
                pathGradientFill.Append(new A.FillToRectangle
                {
                    Left = 50000,
                    Top = 50000,
                    Right = 50000,
                    Bottom = 50000
                });
                gradientFill.Append(pathGradientFill);
            }
        }

        return gradientFill;
    }

    private A.PatternFill GeneratePatternFill(FillSpec fillSpec)
    {
        var patternFill = new A.PatternFill();

        if (fillSpec.Pattern != null)
        {
            // Map pattern type
            var presetPattern = fillSpec.Pattern.Type switch
            {
                "dots" => PresetPatternValues.DottedDiamond,
                "grid" => PresetPatternValues.LargeGrid,
                "diagonalStripe" => PresetPatternValues.DiagonalBrick,
                "checkered" => PresetPatternValues.LargeConfetti,
                _ => PresetPatternValues.Cross
            };
            patternFill.Preset = presetPattern;

            // Foreground color
            if (!string.IsNullOrEmpty(fillSpec.Pattern.Foreground))
            {
                patternFill.Append(new A.ForegroundColor(
                    new A.RgbColorModelHex { Val = fillSpec.Pattern.Foreground }));
            }

            // Background color
            if (!string.IsNullOrEmpty(fillSpec.Pattern.Background))
            {
                patternFill.Append(new A.BackgroundColor(
                    new A.RgbColorModelHex { Val = fillSpec.Pattern.Background }));
            }
        }

        return patternFill;
    }
}

public static class ColorHelper
{
    public static string HexToRgb(string hex)
    {
        // Remove '#' if present
        hex = hex.TrimStart('#');

        // Ensure 6 characters
        if (hex.Length == 3)
        {
            hex = $"{hex[0]}{hex[0]}{hex[1]}{hex[1]}{hex[2]}{hex[2]}";
        }

        return hex.ToUpper();
    }

    public static A.RgbColorModelHex CreateColor(string hexColor)
    {
        return new A.RgbColorModelHex { Val = HexToRgb(hexColor) };
    }

    public static A.RgbColorModelHex CreateColorWithAlpha(string hexColor, double alpha)
    {
        var color = CreateColor(hexColor);
        if (alpha < 1.0)
        {
            color.Append(new A.Alpha { Val = (int)(alpha * 100000) });
        }
        return color;
    }
}

public static class EMUConverter
{
    private const double EMU_PER_INCH = 914400.0;
    private const double EMU_PER_CM = 360000.0;
    private const double POINTS_PER_INCH = 72.0;

    public static long InchesToEMU(double inches)
    {
        return (long)(inches * EMU_PER_INCH);
    }

    public static long CMToEMU(double cm)
    {
        return (long)(cm * EMU_PER_CM);
    }

    public static long PointsToEMU(double points)
    {
        return (long)((points / POINTS_PER_INCH) * EMU_PER_INCH);
    }

    public static double EMUToInches(long emus)
    {
        return emus / EMU_PER_INCH;
    }

    public static double EMUToCM(long emus)
    {
        return emus / EMU_PER_CM;
    }

    public static double EMUToPoints(long emus)
    {
        return (emus / EMU_PER_INCH) * POINTS_PER_INCH;
    }
}

public static class ValidationHelper
{
    public static void ValidateTestCase(TestCase testCase)
    {
        if (string.IsNullOrEmpty(testCase.Id))
        {
            throw new ArgumentException("Test case ID cannot be empty");
        }

        if (string.IsNullOrEmpty(testCase.Name))
        {
            throw new ArgumentException("Test case name cannot be empty");
        }

        if (testCase.Spec.SlideSize.Width <= 0 || testCase.Spec.SlideSize.Height <= 0)
        {
            throw new ArgumentException("Invalid slide size");
        }

        if (testCase.Spec.Slides.Count == 0)
        {
            throw new ArgumentException("Test case must have at least one slide");
        }
    }

    public static void ValidateChartData(ChartData chartData)
    {
        if (chartData.Series.Count == 0)
        {
            throw new ArgumentException("Chart must have at least one series");
        }

        foreach (var series in chartData.Series)
        {
            if (series.Values.Count == 0)
            {
                throw new ArgumentException($"Series '{series.Name}' has no values");
            }
        }
    }
}
