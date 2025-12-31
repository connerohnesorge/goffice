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
            var scenario = deserializer.Deserialize<TestScenario>(yaml);

            // Normalize operation properties into the Properties dictionary
            foreach (var op in scenario.Operations)
            {
                NormalizeOperationProperties(op);
            }

            return scenario;
        }

        static void NormalizeOperationProperties(Operation op)
        {
            // Move all individual properties into the Properties dictionary
            if (op.Text != null)
                op.Properties["text"] = op.Text;
            if (op.DocType != null)
                op.Properties["doc_type"] = op.DocType;
            if (op.Rows.HasValue)
                op.Properties["rows"] = op.Rows.Value;
            if (op.Cols.HasValue)
                op.Properties["cols"] = op.Cols.Value;
            if (op.Style != null)
                op.Properties["style"] = op.Style;
            if (op.WidthPercent.HasValue)
                op.Properties["width_percent"] = op.WidthPercent.Value;
            if (op.Operations != null)
                op.Properties["operations"] = op.Operations;
            if (op.Cells != null)
                op.Properties["cells"] = op.Cells;
            if (op.PropertiesNested != null)
            {
                foreach (var kvp in op.PropertiesNested)
                {
                    op.Properties[kvp.Key] = kvp.Value;
                }
            }
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
                    throw new NotImplementedException("Spreadsheet bridge not yet implemented");
                case "presentation":
                    throw new NotImplementedException("Presentation bridge not yet implemented");
                default:
                    throw new ArgumentException($"Unknown document type: {scenario.DocumentType}");
            }
        }
    }
}
