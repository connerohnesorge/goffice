using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Drawing;
using DocumentFormat.OpenXml.Drawing.Charts;
using DocumentFormat.OpenXml.Packaging;
using DocumentFormat.OpenXml.Presentation;
using PptxGenerator.Models;
using A = DocumentFormat.OpenXml.Drawing;
using C = DocumentFormat.OpenXml.Drawing.Charts;
using P = DocumentFormat.OpenXml.Presentation;

namespace PptxGenerator.Generators;

public class ChartGenerator
{
    public P.GraphicFrame GenerateChart(
        SlidePart slidePart,
        ChartSpec chartSpec,
        Models.Position position,
        Models.Size size,
        uint shapeId)
    {
        // Create chart part
        var chartPart = slidePart.AddNewPart<ChartPart>();
        chartPart.ChartSpace = new ChartSpace();
        chartPart.ChartSpace.Append(new EditingLanguage { Val = "en-US" });

        var chart = new C.Chart();
        chartPart.ChartSpace.Append(chart);

        // Create plot area
        var plotArea = new PlotArea();
        plotArea.Append(new Layout());
        chart.Append(plotArea);

        // Generate chart based on type
        switch (chartSpec.Type)
        {
            case "bar_clustered":
            case "bar_stacked":
            case "bar_percent_stacked":
            case "column":
                GenerateBarChart(plotArea, chartSpec);
                break;
            case "line":
            case "line_smooth":
            case "line_markers":
                GenerateLineChart(plotArea, chartSpec);
                break;
            case "pie":
                GeneratePieChart(plotArea, chartSpec);
                break;
            case "doughnut":
                GenerateDoughnutChart(plotArea, chartSpec);
                break;
            case "area":
            case "area_stacked":
                GenerateAreaChart(plotArea, chartSpec);
                break;
            case "scatter":
            case "scatter_smooth":
            case "scatter_line":
                GenerateScatterChart(plotArea, chartSpec);
                break;
            case "bubble":
                GenerateBubbleChart(plotArea, chartSpec);
                break;
            default:
                GenerateBarChart(plotArea, chartSpec); // Fallback
                break;
        }

        // Add chart axes (if not pie/doughnut)
        if (!chartSpec.Type.StartsWith("pie") && !chartSpec.Type.StartsWith("doughnut"))
        {
            AddChartAxes(plotArea, chartSpec);
        }

        // Add legend
        if (chartSpec.Legend?.ShowLegend == true)
        {
            AddLegend(chart, chartSpec.Legend);
        }

        // Add title
        if (!string.IsNullOrEmpty(chartSpec.Title))
        {
            AddChartTitle(chart, chartSpec.Title);
        }

        // Create graphic frame
        var graphicFrame = new P.GraphicFrame();
        graphicFrame.Append(new P.NonVisualGraphicFrameProperties(
            new P.NonVisualDrawingProperties { Id = shapeId, Name = $"Chart {shapeId}" },
            new P.NonVisualGraphicFrameDrawingProperties(),
            new ApplicationNonVisualDrawingProperties()));

        graphicFrame.Append(new P.Transform(
            new A.Offset { X = position.X, Y = position.Y },
            new A.Extents { Cx = size.Width, Cy = size.Height }));

        var graphic = new A.Graphic();
        var graphicData = new A.GraphicData { Uri = "http://schemas.openxmlformats.org/drawingml/2006/chart" };
        graphicData.Append(new C.ChartReference { Id = slidePart.GetIdOfPart(chartPart) });
        graphic.Append(graphicData);
        graphicFrame.Append(graphic);

        return graphicFrame;
    }

    private void GenerateBarChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var barChart = new BarChart();

        // Set bar direction and grouping
        barChart.Append(new BarDirection { Val = BarDirectionValues.Column });

        var grouping = chartSpec.Type switch
        {
            "bar_stacked" => BarGroupingValues.Stacked,
            "bar_percent_stacked" => BarGroupingValues.PercentStacked,
            _ => BarGroupingValues.Clustered
        };
        barChart.Append(new BarGrouping { Val = grouping });

        // Add series
        for (uint i = 0; i < chartSpec.Data.Series.Count; i++)
        {
            var series = chartSpec.Data.Series[(int)i];
            var barChartSeries = new BarChartSeries();
            barChartSeries.Append(new C.Index { Val = i });
            barChartSeries.Append(new Order { Val = i });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            barChartSeries.Append(seriesText);

            // Category axis data
            barChartSeries.Append(CreateCategoryAxisData(chartSpec.Data.Categories));

            // Values
            barChartSeries.Append(CreateValueData(series.Values));

            // Color
            if (!string.IsNullOrEmpty(series.Color))
            {
                barChartSeries.Append(new ChartShapeProperties(
                    new A.SolidFill(new A.RgbColorModelHex { Val = series.Color })));
            }

            barChart.Append(barChartSeries);
        }

        // Axis IDs
        barChart.Append(new AxisId { Val = 48650112U });
        barChart.Append(new AxisId { Val = 48672768U });

        plotArea.Append(barChart);
    }

    private void GenerateLineChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var lineChart = new LineChart();
        lineChart.Append(new Grouping { Val = GroupingValues.Standard });

        // Add series
        for (uint i = 0; i < chartSpec.Data.Series.Count; i++)
        {
            var series = chartSpec.Data.Series[(int)i];
            var lineChartSeries = new LineChartSeries();
            lineChartSeries.Append(new C.Index { Val = i });
            lineChartSeries.Append(new Order { Val = i });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            lineChartSeries.Append(seriesText);

            // Category axis data
            lineChartSeries.Append(CreateCategoryAxisData(chartSpec.Data.Categories));

            // Values
            lineChartSeries.Append(CreateValueData(series.Values));

            // Smooth line
            if (chartSpec.Type == "line_smooth" || series.Smooth)
            {
                lineChartSeries.Append(new C.Smooth { Val = true });
            }

            // Markers
            if (chartSpec.Type == "line_markers")
            {
                lineChartSeries.Append(new C.Marker(new C.Symbol { Val = MarkerStyleValues.Circle }));
            }
            else
            {
                lineChartSeries.Append(new C.Marker(new C.Symbol { Val = MarkerStyleValues.None }));
            }

            // Color
            if (!string.IsNullOrEmpty(series.Color))
            {
                lineChartSeries.Append(new ChartShapeProperties(
                    new A.Outline(new A.SolidFill(new A.RgbColorModelHex { Val = series.Color }))));
            }

            lineChart.Append(lineChartSeries);
        }

        // Axis IDs
        lineChart.Append(new AxisId { Val = 48650112U });
        lineChart.Append(new AxisId { Val = 48672768U });

        plotArea.Append(lineChart);
    }

    private void GeneratePieChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var pieChart = new PieChart();

        if (chartSpec.Data.Series.Count > 0)
        {
            var series = chartSpec.Data.Series[0];
            var pieChartSeries = new PieChartSeries();
            pieChartSeries.Append(new C.Index { Val = 0U });
            pieChartSeries.Append(new Order { Val = 0U });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            pieChartSeries.Append(seriesText);

            // Category axis data (pie slice labels)
            pieChartSeries.Append(CreateCategoryAxisData(chartSpec.Data.Categories));

            // Values
            pieChartSeries.Append(CreateValueData(series.Values));

            pieChart.Append(pieChartSeries);
        }

        plotArea.Append(pieChart);
    }

    private void GenerateDoughnutChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var doughnutChart = new DoughnutChart();

        if (chartSpec.Data.Series.Count > 0)
        {
            var series = chartSpec.Data.Series[0];
            var pieChartSeries = new PieChartSeries();
            pieChartSeries.Append(new C.Index { Val = 0U });
            pieChartSeries.Append(new Order { Val = 0U });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            pieChartSeries.Append(seriesText);

            // Category axis data
            pieChartSeries.Append(CreateCategoryAxisData(chartSpec.Data.Categories));

            // Values
            pieChartSeries.Append(CreateValueData(series.Values));

            doughnutChart.Append(pieChartSeries);
        }

        // Hole size (50% for doughnut)
        doughnutChart.Append(new HoleSize { Val = (byte)50 });

        plotArea.Append(doughnutChart);
    }

    private void GenerateAreaChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var areaChart = new AreaChart();

        var grouping = chartSpec.Type == "area_stacked"
            ? GroupingValues.Stacked
            : GroupingValues.Standard;
        areaChart.Append(new Grouping { Val = grouping });

        // Add series
        for (uint i = 0; i < chartSpec.Data.Series.Count; i++)
        {
            var series = chartSpec.Data.Series[(int)i];
            var areaChartSeries = new AreaChartSeries();
            areaChartSeries.Append(new C.Index { Val = i });
            areaChartSeries.Append(new Order { Val = i });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            areaChartSeries.Append(seriesText);

            // Category axis data
            areaChartSeries.Append(CreateCategoryAxisData(chartSpec.Data.Categories));

            // Values
            areaChartSeries.Append(CreateValueData(series.Values));

            // Color
            if (!string.IsNullOrEmpty(series.Color))
            {
                areaChartSeries.Append(new ChartShapeProperties(
                    new A.SolidFill(new A.RgbColorModelHex { Val = series.Color })));
            }

            areaChart.Append(areaChartSeries);
        }

        // Axis IDs
        areaChart.Append(new AxisId { Val = 48650112U });
        areaChart.Append(new AxisId { Val = 48672768U });

        plotArea.Append(areaChart);
    }

    private void GenerateScatterChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var scatterChart = new ScatterChart();
        scatterChart.Append(new ScatterStyle { Val = ScatterStyleValues.LineMarker });

        // Add series
        for (uint i = 0; i < chartSpec.Data.Series.Count; i++)
        {
            var series = chartSpec.Data.Series[(int)i];
            var scatterChartSeries = new ScatterChartSeries();
            scatterChartSeries.Append(new C.Index { Val = i });
            scatterChartSeries.Append(new Order { Val = i });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            scatterChartSeries.Append(seriesText);

            // X values
            if (series.XValues != null && series.XValues.Count > 0)
            {
                scatterChartSeries.Append(CreateXValueData(series.XValues));
            }
            else
            {
                // Use index as X values
                var xVals = Enumerable.Range(0, series.Values.Count).Select(x => (double)x).ToList();
                scatterChartSeries.Append(CreateXValueData(xVals));
            }

            // Y values
            scatterChartSeries.Append(CreateYValueData(series.Values));

            // Smooth line
            if (chartSpec.Type == "scatter_smooth" || series.Smooth)
            {
                scatterChartSeries.Append(new C.Smooth { Val = true });
            }

            // Color
            if (!string.IsNullOrEmpty(series.Color))
            {
                scatterChartSeries.Append(new ChartShapeProperties(
                    new A.Outline(new A.SolidFill(new A.RgbColorModelHex { Val = series.Color }))));
            }

            scatterChart.Append(scatterChartSeries);
        }

        // Axis IDs
        scatterChart.Append(new AxisId { Val = 48650112U });
        scatterChart.Append(new AxisId { Val = 48672768U });

        plotArea.Append(scatterChart);
    }

    private void GenerateBubbleChart(PlotArea plotArea, ChartSpec chartSpec)
    {
        var bubbleChart = new BubbleChart();

        // Add series
        for (uint i = 0; i < chartSpec.Data.Series.Count; i++)
        {
            var series = chartSpec.Data.Series[(int)i];
            var bubbleChartSeries = new BubbleChartSeries();
            bubbleChartSeries.Append(new C.Index { Val = i });
            bubbleChartSeries.Append(new Order { Val = i });

            // Series text
            var seriesText = new SeriesText();
            seriesText.Append(new StringReference(
                new Formula { Text = "" },
                new StringCache(
                    new PointCount { Val = 1U },
                    new StringPoint { Index = 0U, NumericValue = new NumericValue { Text = series.Name } })));
            bubbleChartSeries.Append(seriesText);

            // X values
            if (series.XValues != null && series.XValues.Count > 0)
            {
                bubbleChartSeries.Append(CreateXValueData(series.XValues));
            }
            else
            {
                var xVals = Enumerable.Range(0, series.Values.Count).Select(x => (double)x).ToList();
                bubbleChartSeries.Append(CreateXValueData(xVals));
            }

            // Y values
            bubbleChartSeries.Append(CreateYValueData(series.Values));

            // Bubble size (use values as size)
            bubbleChartSeries.Append(CreateBubbleSizeData(series.Values));

            // Color
            if (!string.IsNullOrEmpty(series.Color))
            {
                bubbleChartSeries.Append(new ChartShapeProperties(
                    new A.SolidFill(new A.RgbColorModelHex { Val = series.Color })));
            }

            bubbleChart.Append(bubbleChartSeries);
        }

        // Axis IDs
        bubbleChart.Append(new AxisId { Val = 48650112U });
        bubbleChart.Append(new AxisId { Val = 48672768U });

        plotArea.Append(bubbleChart);
    }

    private CategoryAxisData CreateCategoryAxisData(List<string> categories)
    {
        var catAx = new CategoryAxisData();
        var strRef = new StringReference();
        strRef.Append(new Formula { Text = "" });

        var strCache = new StringCache();
        strCache.Append(new PointCount { Val = (uint)categories.Count });

        for (uint i = 0; i < categories.Count; i++)
        {
            strCache.Append(new StringPoint
            {
                Index = i,
                NumericValue = new NumericValue { Text = categories[(int)i] }
            });
        }

        strRef.Append(strCache);
        catAx.Append(strRef);
        return catAx;
    }

    private C.Values CreateValueData(List<double> values)
    {
        var vals = new C.Values();
        var numRef = new NumberReference();
        numRef.Append(new Formula { Text = "" });

        var numCache = new NumberingCache();
        numCache.Append(new FormatCode { Text = "General" });
        numCache.Append(new PointCount { Val = (uint)values.Count });

        for (uint i = 0; i < values.Count; i++)
        {
            numCache.Append(new NumericPoint
            {
                Index = i,
                NumericValue = new NumericValue { Text = values[(int)i].ToString("G") }
            });
        }

        numRef.Append(numCache);
        vals.Append(numRef);
        return vals;
    }

    private XValues CreateXValueData(List<double> values)
    {
        var xVals = new XValues();
        var numRef = new NumberReference();
        numRef.Append(new Formula { Text = "" });

        var numCache = new NumberingCache();
        numCache.Append(new FormatCode { Text = "General" });
        numCache.Append(new PointCount { Val = (uint)values.Count });

        for (uint i = 0; i < values.Count; i++)
        {
            numCache.Append(new NumericPoint
            {
                Index = i,
                NumericValue = new NumericValue { Text = values[(int)i].ToString("G") }
            });
        }

        numRef.Append(numCache);
        xVals.Append(numRef);
        return xVals;
    }

    private YValues CreateYValueData(List<double> values)
    {
        var yVals = new YValues();
        var numRef = new NumberReference();
        numRef.Append(new Formula { Text = "" });

        var numCache = new NumberingCache();
        numCache.Append(new FormatCode { Text = "General" });
        numCache.Append(new PointCount { Val = (uint)values.Count });

        for (uint i = 0; i < values.Count; i++)
        {
            numCache.Append(new NumericPoint
            {
                Index = i,
                NumericValue = new NumericValue { Text = values[(int)i].ToString("G") }
            });
        }

        numRef.Append(numCache);
        yVals.Append(numRef);
        return yVals;
    }

    private BubbleSize CreateBubbleSizeData(List<double> values)
    {
        var bubbleSize = new BubbleSize();
        var numRef = new NumberReference();
        numRef.Append(new Formula { Text = "" });

        var numCache = new NumberingCache();
        numCache.Append(new FormatCode { Text = "General" });
        numCache.Append(new PointCount { Val = (uint)values.Count });

        for (uint i = 0; i < values.Count; i++)
        {
            numCache.Append(new NumericPoint
            {
                Index = i,
                NumericValue = new NumericValue { Text = values[(int)i].ToString("G") }
            });
        }

        numRef.Append(numCache);
        bubbleSize.Append(numRef);
        return bubbleSize;
    }

    private void AddChartAxes(PlotArea plotArea, ChartSpec chartSpec)
    {
        // Category axis
        var catAx = new CategoryAxis();
        catAx.Append(new AxisId { Val = 48650112U });
        catAx.Append(new Scaling(new Orientation { Val = OrientationValues.MinMax }));
        catAx.Append(new AxisPosition { Val = AxisPositionValues.Bottom });
        catAx.Append(new TickLabelPosition { Val = TickLabelPositionValues.NextTo });
        catAx.Append(new CrossingAxis { Val = 48672768U });
        catAx.Append(new Crosses { Val = CrossesValues.AutoZero });

        if (chartSpec.Axes?.Category?.Title != null && !string.IsNullOrEmpty(chartSpec.Axes.Category.Title))
        {
            var title = new C.Title();
            var chartText = new ChartText();
            var richText = new C.RichText();
            richText.Append(new A.BodyProperties());
            richText.Append(new A.ListStyle());
            var para = new A.Paragraph();
            var run = new A.Run();
            run.Append(new A.Text { Text = chartSpec.Axes.Category.Title });
            para.Append(run);
            richText.Append(para);
            chartText.Append(richText);
            title.Append(chartText);
            catAx.Append(title);
        }

        plotArea.Append(catAx);

        // Value axis
        var valAx = new ValueAxis();
        valAx.Append(new AxisId { Val = 48672768U });
        valAx.Append(new Scaling(new Orientation { Val = OrientationValues.MinMax }));
        valAx.Append(new AxisPosition { Val = AxisPositionValues.Left });
        valAx.Append(new MajorGridlines());
        valAx.Append(new TickLabelPosition { Val = TickLabelPositionValues.NextTo });
        valAx.Append(new CrossingAxis { Val = 48650112U });
        valAx.Append(new Crosses { Val = CrossesValues.AutoZero });
        valAx.Append(new CrossBetween { Val = CrossBetweenValues.Between });

        if (chartSpec.Axes?.Value?.Title != null && !string.IsNullOrEmpty(chartSpec.Axes.Value.Title))
        {
            var title = new C.Title();
            var chartText = new ChartText();
            var richText = new C.RichText();
            richText.Append(new A.BodyProperties());
            richText.Append(new A.ListStyle());
            var para = new A.Paragraph();
            var run = new A.Run();
            run.Append(new A.Text { Text = chartSpec.Axes.Value.Title });
            para.Append(run);
            richText.Append(para);
            chartText.Append(richText);
            title.Append(chartText);
            valAx.Append(title);
        }

        plotArea.Append(valAx);
    }

    private void AddLegend(C.Chart chart, LegendSpec legendSpec)
    {
        var legend = new Legend();

        var position = legendSpec.Position switch
        {
            "top" => LegendPositionValues.Top,
            "bottom" => LegendPositionValues.Bottom,
            "left" => LegendPositionValues.Left,
            "right" => LegendPositionValues.Right,
            "top_right" => LegendPositionValues.TopRight,
            _ => LegendPositionValues.Right
        };

        legend.Append(new LegendPosition { Val = position });
        legend.Append(new Layout());

        chart.Append(legend);
    }

    private void AddChartTitle(C.Chart chart, string title)
    {
        var chartTitle = new C.Title();
        var chartText = new ChartText();
        var richText = new C.RichText();
        richText.Append(new A.BodyProperties());
        richText.Append(new A.ListStyle());

        var para = new A.Paragraph();
        var run = new A.Run();
        run.Append(new A.Text { Text = title });
        para.Append(run);
        richText.Append(para);

        chartText.Append(richText);
        chartTitle.Append(chartText);
        chartTitle.Append(new Overlay { Val = false });

        chart.InsertBefore(chartTitle, chart.PlotArea);
    }
}
