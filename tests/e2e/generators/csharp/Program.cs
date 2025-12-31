using System.CommandLine;
using Newtonsoft.Json;
using PptxGenerator.Models;
using PptxGenerator.Generators;
using DocumentFormat.OpenXml.Packaging;
using DocumentFormat.OpenXml.Presentation;
using DocumentFormat.OpenXml;
using A = DocumentFormat.OpenXml.Drawing;
using P = DocumentFormat.OpenXml.Presentation;

namespace PptxGenerator;

class Program
{
    static async Task<int> Main(string[] args)
    {
        var inputOption = new Option<FileInfo>(
            name: "--input",
            description: "Input JSON test case file")
        {
            IsRequired = true
        };

        var outputOption = new Option<FileInfo>(
            name: "--output",
            description: "Output PPTX file path")
        {
            IsRequired = true
        };

        var rootCommand = new RootCommand("PowerPoint presentation generator for E2E visual testing")
        {
            inputOption,
            outputOption
        };

        rootCommand.SetHandler(async (inputFile, outputFile) =>
        {
            await GeneratePresentationAsync(inputFile, outputFile);
        }, inputOption, outputOption);

        return await rootCommand.InvokeAsync(args);
    }

    static async Task GeneratePresentationAsync(FileInfo inputFile, FileInfo outputFile)
    {
        try
        {
            // Load test case from JSON
            var json = await File.ReadAllTextAsync(inputFile.FullName);
            var testCase = JsonConvert.DeserializeObject<TestCase>(json);

            if (testCase == null)
            {
                Console.Error.WriteLine("Failed to deserialize test case");
                Environment.Exit(1);
            }

            Console.WriteLine($"Generating presentation: {testCase.Name}");
            Console.WriteLine($"  ID: {testCase.Id}");
            Console.WriteLine($"  Category: {testCase.Category}");
            Console.WriteLine($"  Slides: {testCase.Spec.Slides.Count}");

            // Generate presentation
            GeneratePresentation(testCase, outputFile.FullName);

            Console.WriteLine($"✓ Presentation generated: {outputFile.FullName}");
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Error generating presentation: {ex.Message}");
            Console.Error.WriteLine(ex.StackTrace);
            Environment.Exit(1);
        }
    }

    static void GeneratePresentation(TestCase testCase, string outputPath)
    {
        // Create presentation document
        using var presentationDoc = PresentationDocument.Create(outputPath, PresentationDocumentType.Presentation);

        // Create presentation part
        var presentationPart = presentationDoc.AddPresentationPart();
        presentationPart.Presentation = new Presentation();

        // Create slide ID list
        var slideIdList = new SlideIdList();
        presentationPart.Presentation.SlideIdList = slideIdList;

        // Create slide size
        var slideSize = new P.SlideSize
        {
            Cx = (int)testCase.Spec.SlideSize.Width,
            Cy = (int)testCase.Spec.SlideSize.Height
        };
        presentationPart.Presentation.SlideSize = slideSize;

        uint slideId = 256;

        // Generate each slide
        foreach (var slideSpec in testCase.Spec.Slides)
        {
            var slidePart = presentationPart.AddNewPart<SlidePart>();
            GenerateSlide(slidePart, slideSpec);

            // Add slide to presentation
            var slideIdEntry = new SlideId
            {
                Id = slideId++,
                RelationshipId = presentationPart.GetIdOfPart(slidePart)
            };
            slideIdList.Append(slideIdEntry);
        }

        // Save presentation
        presentationPart.Presentation.Save();
    }

    static void GenerateSlide(SlidePart slidePart, SlideSpec slideSpec)
    {
        slidePart.Slide = new Slide(new CommonSlideData(new ShapeTree()));

        var shapeTree = slidePart.Slide.CommonSlideData.ShapeTree!;

        // Add non-visual group shape properties
        shapeTree.Append(new P.NonVisualGroupShapeProperties(
            new P.NonVisualDrawingProperties { Id = 1, Name = "" },
            new P.NonVisualGroupShapeDrawingProperties(),
            new ApplicationNonVisualDrawingProperties()));

        // Add group shape properties
        shapeTree.Append(new P.GroupShapeProperties(new A.TransformGroup()));

        // Background
        if (slideSpec.Background != null)
        {
            var background = new P.Background(new P.BackgroundProperties());
            var fillGenerator = new FillGenerator();
            var fill = fillGenerator.GenerateFill(slideSpec.Background.Fill);
            background.BackgroundProperties.Append(fill);
            slidePart.Slide.InsertBefore(background, slidePart.Slide.CommonSlideData);
        }

        uint shapeId = 2;

        // Generate elements
        foreach (var element in slideSpec.Elements)
        {
            switch (element.Type)
            {
                case "chart":
                    if (element.Chart != null)
                    {
                        var chartGen = new ChartGenerator();
                        var chartShape = chartGen.GenerateChart(
                            slidePart,
                            element.Chart,
                            element.Position,
                            element.Size,
                            shapeId++);
                        shapeTree.Append(chartShape);
                    }
                    break;

                case "shape":
                    if (element.Shape != null)
                    {
                        var shapeGen = new ShapeGenerator();
                        var shape = shapeGen.GenerateShape(
                            element.Shape,
                            element.Position,
                            element.Size,
                            shapeId++);
                        shapeTree.Append(shape);
                    }
                    break;

                case "text":
                    if (element.Text != null)
                    {
                        var textGen = new TextGenerator();
                        var textBox = textGen.GenerateTextBox(
                            element.Text,
                            element.Position,
                            element.Size,
                            shapeId++);
                        shapeTree.Append(textBox);
                    }
                    break;
            }
        }
    }
}
