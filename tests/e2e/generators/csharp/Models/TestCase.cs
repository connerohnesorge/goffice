using Newtonsoft.Json;

namespace PptxGenerator.Models;

// Test Case root structure
public class TestCase
{
    [JsonProperty("id")]
    public string Id { get; set; } = "";

    [JsonProperty("name")]
    public string Name { get; set; } = "";

    [JsonProperty("description")]
    public string Description { get; set; } = "";

    [JsonProperty("category")]
    public string Category { get; set; } = "";

    [JsonProperty("tags")]
    public List<string> Tags { get; set; } = new();

    [JsonProperty("created")]
    public DateTime Created { get; set; }

    [JsonProperty("author")]
    public string Author { get; set; } = "";

    [JsonProperty("spec")]
    public TestSpec Spec { get; set; } = new();

    [JsonProperty("config")]
    public TestConfig Config { get; set; } = new();

    [JsonProperty("metadata")]
    public Dictionary<string, string> Metadata { get; set; } = new();
}

// Test specification
public class TestSpec
{
    [JsonProperty("slide_count")]
    public int SlideCount { get; set; }

    [JsonProperty("slide_size")]
    public SlideSize SlideSize { get; set; } = new();

    [JsonProperty("slides")]
    public List<SlideSpec> Slides { get; set; } = new();

    [JsonProperty("theme")]
    public ThemeSpec? Theme { get; set; }

    [JsonProperty("fonts")]
    public List<FontSpec> Fonts { get; set; } = new();
}

// Slide size (EMUs)
public class SlideSize
{
    [JsonProperty("width")]
    public long Width { get; set; }

    [JsonProperty("height")]
    public long Height { get; set; }
}

// Slide specification
public class SlideSpec
{
    [JsonProperty("index")]
    public int Index { get; set; }

    [JsonProperty("layout")]
    public string Layout { get; set; } = "Blank";

    [JsonProperty("background")]
    public BackgroundSpec? Background { get; set; }

    [JsonProperty("elements")]
    public List<ElementSpec> Elements { get; set; } = new();
}

// Background specification
public class BackgroundSpec
{
    [JsonProperty("fill")]
    public FillSpec Fill { get; set; } = new();
}

// Element specification
public class ElementSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "";

    [JsonProperty("position")]
    public Position Position { get; set; } = new();

    [JsonProperty("size")]
    public Size Size { get; set; } = new();

    [JsonProperty("z_index")]
    public int ZIndex { get; set; }

    [JsonProperty("shape")]
    public ShapeSpec? Shape { get; set; }

    [JsonProperty("chart")]
    public ChartSpec? Chart { get; set; }

    [JsonProperty("text")]
    public TextSpec? Text { get; set; }

    [JsonProperty("image")]
    public ImageSpec? Image { get; set; }

    [JsonProperty("table")]
    public TableSpec? Table { get; set; }
}

// Position (EMUs)
public class Position
{
    [JsonProperty("x")]
    public long X { get; set; }

    [JsonProperty("y")]
    public long Y { get; set; }
}

// Size (EMUs)
public class Size
{
    [JsonProperty("width")]
    public long Width { get; set; }

    [JsonProperty("height")]
    public long Height { get; set; }
}

// Chart specification
public class ChartSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "";

    [JsonProperty("title")]
    public string Title { get; set; } = "";

    [JsonProperty("legend")]
    public LegendSpec? Legend { get; set; }

    [JsonProperty("data")]
    public ChartData Data { get; set; } = new();

    [JsonProperty("style")]
    public ChartStyle Style { get; set; } = new();

    [JsonProperty("axes")]
    public AxesSpec? Axes { get; set; }
}

// Chart data
public class ChartData
{
    [JsonProperty("categories")]
    public List<string> Categories { get; set; } = new();

    [JsonProperty("series")]
    public List<SeriesData> Series { get; set; } = new();
}

// Series data
public class SeriesData
{
    [JsonProperty("name")]
    public string Name { get; set; } = "";

    [JsonProperty("values")]
    public List<double> Values { get; set; } = new();

    [JsonProperty("x_values")]
    public List<double>? XValues { get; set; }

    [JsonProperty("color")]
    public string Color { get; set; } = "";

    [JsonProperty("marker_style")]
    public string? MarkerStyle { get; set; }

    [JsonProperty("smooth")]
    public bool Smooth { get; set; }
}

// Chart style
public class ChartStyle
{
    [JsonProperty("color_palette")]
    public List<string> ColorPalette { get; set; } = new();

    [JsonProperty("font_size")]
    public int FontSize { get; set; }

    [JsonProperty("font_family")]
    public string FontFamily { get; set; } = "";

    [JsonProperty("show_data_labels")]
    public bool ShowDataLabels { get; set; }

    [JsonProperty("grid_lines")]
    public GridLineSpec GridLines { get; set; } = new();
}

// Grid line specification
public class GridLineSpec
{
    [JsonProperty("major_horizontal")]
    public bool MajorHorizontal { get; set; }

    [JsonProperty("minor_horizontal")]
    public bool MinorHorizontal { get; set; }

    [JsonProperty("major_vertical")]
    public bool MajorVertical { get; set; }

    [JsonProperty("minor_vertical")]
    public bool MinorVertical { get; set; }
}

// Legend specification
public class LegendSpec
{
    [JsonProperty("position")]
    public string Position { get; set; } = "right";

    [JsonProperty("show_legend")]
    public bool ShowLegend { get; set; }
}

// Axes specification
public class AxesSpec
{
    [JsonProperty("category")]
    public AxisSpec Category { get; set; } = new();

    [JsonProperty("value")]
    public AxisSpec Value { get; set; } = new();
}

// Axis specification
public class AxisSpec
{
    [JsonProperty("title")]
    public string Title { get; set; } = "";

    [JsonProperty("min")]
    public double Min { get; set; }

    [JsonProperty("max")]
    public double Max { get; set; }

    [JsonProperty("major_unit")]
    public double MajorUnit { get; set; }

    [JsonProperty("minor_unit")]
    public double MinorUnit { get; set; }

    [JsonProperty("number_format")]
    public string NumberFormat { get; set; } = "";

    [JsonProperty("tick_label_rotation")]
    public int TickLabelRotation { get; set; }
}

// Shape specification
public class ShapeSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "";

    [JsonProperty("fill")]
    public FillSpec Fill { get; set; } = new();

    [JsonProperty("stroke")]
    public StrokeSpec Stroke { get; set; } = new();

    [JsonProperty("effects")]
    public List<EffectSpec> Effects { get; set; } = new();

    [JsonProperty("text")]
    public TextSpec? Text { get; set; }
}

// Fill specification
public class FillSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "solid";

    [JsonProperty("color")]
    public string? Color { get; set; }

    [JsonProperty("gradient")]
    public GradientSpec? Gradient { get; set; }

    [JsonProperty("pattern")]
    public PatternSpec? Pattern { get; set; }
}

// Gradient specification
public class GradientSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "linear";

    [JsonProperty("angle")]
    public int Angle { get; set; }

    [JsonProperty("stops")]
    public List<GradientStop> Stops { get; set; } = new();
}

// Gradient stop
public class GradientStop
{
    [JsonProperty("position")]
    public double Position { get; set; }

    [JsonProperty("color")]
    public string Color { get; set; } = "";
}

// Pattern specification
public class PatternSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "";

    [JsonProperty("foreground")]
    public string Foreground { get; set; } = "";

    [JsonProperty("background")]
    public string Background { get; set; } = "";
}

// Stroke specification
public class StrokeSpec
{
    [JsonProperty("width")]
    public long Width { get; set; }

    [JsonProperty("color")]
    public string Color { get; set; } = "";

    [JsonProperty("dash_style")]
    public string DashStyle { get; set; } = "solid";

    [JsonProperty("cap")]
    public string Cap { get; set; } = "flat";

    [JsonProperty("join")]
    public string Join { get; set; } = "miter";
}

// Effect specification
public class EffectSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "";

    [JsonProperty("shadow")]
    public ShadowSpec? Shadow { get; set; }

    [JsonProperty("glow")]
    public GlowSpec? Glow { get; set; }

    [JsonProperty("reflection")]
    public ReflectionSpec? Reflection { get; set; }

    [JsonProperty("soft_edge")]
    public SoftEdgeSpec? SoftEdge { get; set; }
}

// Shadow specification
public class ShadowSpec
{
    [JsonProperty("type")]
    public string Type { get; set; } = "outer";

    [JsonProperty("angle")]
    public int Angle { get; set; }

    [JsonProperty("distance")]
    public long Distance { get; set; }

    [JsonProperty("blur_radius")]
    public long BlurRadius { get; set; }

    [JsonProperty("color")]
    public string Color { get; set; } = "";

    [JsonProperty("transparency")]
    public double Transparency { get; set; }
}

// Glow specification
public class GlowSpec
{
    [JsonProperty("radius")]
    public long Radius { get; set; }

    [JsonProperty("color")]
    public string Color { get; set; } = "";
}

// Reflection specification
public class ReflectionSpec
{
    [JsonProperty("blur_radius")]
    public long BlurRadius { get; set; }

    [JsonProperty("start_opacity")]
    public double StartOpacity { get; set; }

    [JsonProperty("end_opacity")]
    public double EndOpacity { get; set; }

    [JsonProperty("distance")]
    public long Distance { get; set; }

    [JsonProperty("direction")]
    public int Direction { get; set; }

    [JsonProperty("fade_direction")]
    public int FadeDirection { get; set; }

    [JsonProperty("start_position")]
    public double StartPosition { get; set; }

    [JsonProperty("end_position")]
    public double EndPosition { get; set; }
}

// Soft edge specification
public class SoftEdgeSpec
{
    [JsonProperty("radius")]
    public long Radius { get; set; }
}

// Text specification
public class TextSpec
{
    [JsonProperty("content")]
    public string Content { get; set; } = "";

    [JsonProperty("paragraphs")]
    public List<ParagraphSpec> Paragraphs { get; set; } = new();

    [JsonProperty("default_run")]
    public RunSpec DefaultRun { get; set; } = new();
}

// Paragraph specification
public class ParagraphSpec
{
    [JsonProperty("runs")]
    public List<RunSpec> Runs { get; set; } = new();

    [JsonProperty("alignment")]
    public string Alignment { get; set; } = "left";

    [JsonProperty("line_spacing")]
    public double LineSpacing { get; set; }

    [JsonProperty("space_before")]
    public long SpaceBefore { get; set; }

    [JsonProperty("space_after")]
    public long SpaceAfter { get; set; }

    [JsonProperty("indent")]
    public long Indent { get; set; }

    [JsonProperty("bullet_char")]
    public string? BulletChar { get; set; }
}

// Run specification
public class RunSpec
{
    [JsonProperty("text")]
    public string Text { get; set; } = "";

    [JsonProperty("font_family")]
    public string FontFamily { get; set; } = "Arial";

    [JsonProperty("font_size")]
    public int FontSize { get; set; } = 18;

    [JsonProperty("color")]
    public string Color { get; set; } = "000000";

    [JsonProperty("bold")]
    public bool Bold { get; set; }

    [JsonProperty("italic")]
    public bool Italic { get; set; }

    [JsonProperty("underline")]
    public string Underline { get; set; } = "none";

    [JsonProperty("strikethrough")]
    public bool Strikethrough { get; set; }
}

// Image specification
public class ImageSpec
{
    [JsonProperty("path")]
    public string Path { get; set; } = "";

    [JsonProperty("alt_text")]
    public string AltText { get; set; } = "";

    [JsonProperty("crop")]
    public CropSpec? Crop { get; set; }

    [JsonProperty("effects")]
    public List<EffectSpec> Effects { get; set; } = new();
}

// Crop specification
public class CropSpec
{
    [JsonProperty("left")]
    public double Left { get; set; }

    [JsonProperty("top")]
    public double Top { get; set; }

    [JsonProperty("right")]
    public double Right { get; set; }

    [JsonProperty("bottom")]
    public double Bottom { get; set; }
}

// Table specification
public class TableSpec
{
    [JsonProperty("rows")]
    public int Rows { get; set; }

    [JsonProperty("columns")]
    public int Columns { get; set; }

    [JsonProperty("cells")]
    public List<List<TableCellSpec>> Cells { get; set; } = new();

    [JsonProperty("style")]
    public TableStyleSpec Style { get; set; } = new();
}

// Table cell specification
public class TableCellSpec
{
    [JsonProperty("text")]
    public string Text { get; set; } = "";

    [JsonProperty("span")]
    public CellSpan Span { get; set; } = new();

    [JsonProperty("fill")]
    public FillSpec Fill { get; set; } = new();

    [JsonProperty("text_format")]
    public RunSpec TextFormat { get; set; } = new();
}

// Cell span
public class CellSpan
{
    [JsonProperty("row_span")]
    public int RowSpan { get; set; } = 1;

    [JsonProperty("col_span")]
    public int ColSpan { get; set; } = 1;
}

// Table style specification
public class TableStyleSpec
{
    [JsonProperty("header_row")]
    public bool HeaderRow { get; set; }

    [JsonProperty("total_row")]
    public bool TotalRow { get; set; }

    [JsonProperty("banded_rows")]
    public bool BandedRows { get; set; }

    [JsonProperty("first_column")]
    public bool FirstColumn { get; set; }

    [JsonProperty("last_column")]
    public bool LastColumn { get; set; }

    [JsonProperty("banded_columns")]
    public bool BandedColumns { get; set; }
}

// Theme specification
public class ThemeSpec
{
    [JsonProperty("name")]
    public string Name { get; set; } = "";

    [JsonProperty("color_scheme")]
    public ColorScheme ColorScheme { get; set; } = new();

    [JsonProperty("font_scheme")]
    public FontScheme FontScheme { get; set; } = new();
}

// Color scheme
public class ColorScheme
{
    [JsonProperty("accent1")]
    public string Accent1 { get; set; } = "";

    [JsonProperty("accent2")]
    public string Accent2 { get; set; } = "";

    [JsonProperty("accent3")]
    public string Accent3 { get; set; } = "";

    [JsonProperty("accent4")]
    public string Accent4 { get; set; } = "";

    [JsonProperty("accent5")]
    public string Accent5 { get; set; } = "";

    [JsonProperty("accent6")]
    public string Accent6 { get; set; } = "";

    [JsonProperty("dark1")]
    public string Dark1 { get; set; } = "";

    [JsonProperty("dark2")]
    public string Dark2 { get; set; } = "";

    [JsonProperty("light1")]
    public string Light1 { get; set; } = "";

    [JsonProperty("light2")]
    public string Light2 { get; set; } = "";

    [JsonProperty("hyperlink")]
    public string Hyperlink { get; set; } = "";

    [JsonProperty("followed_hyperlink")]
    public string FollowedHyperlink { get; set; } = "";
}

// Font scheme
public class FontScheme
{
    [JsonProperty("major_font")]
    public string MajorFont { get; set; } = "";

    [JsonProperty("minor_font")]
    public string MinorFont { get; set; } = "";
}

// Font specification
public class FontSpec
{
    [JsonProperty("family")]
    public string Family { get; set; } = "";

    [JsonProperty("path")]
    public string Path { get; set; } = "";
}

// Test configuration
public class TestConfig
{
    [JsonProperty("render_dpi")]
    public int RenderDpi { get; set; } = 300;

    [JsonProperty("render_backend")]
    public string RenderBackend { get; set; } = "libreoffice";

    [JsonProperty("diff_threshold")]
    public double DiffThreshold { get; set; } = 0.01;

    [JsonProperty("diff_algorithm")]
    public string DiffAlgorithm { get; set; } = "perceptual";

    [JsonProperty("ignore_antialiasing")]
    public bool IgnoreAntialiasing { get; set; } = true;

    [JsonProperty("generator_timeout")]
    public int GeneratorTimeout { get; set; } = 60;

    [JsonProperty("render_timeout")]
    public int RenderTimeout { get; set; } = 120;

    [JsonProperty("generate_diff_image")]
    public bool GenerateDiffImage { get; set; } = true;

    [JsonProperty("save_intermediates")]
    public bool SaveIntermediates { get; set; } = false;

    [JsonProperty("retry_count")]
    public int RetryCount { get; set; } = 2;

    [JsonProperty("retry_delay")]
    public int RetryDelay { get; set; } = 5;
}
