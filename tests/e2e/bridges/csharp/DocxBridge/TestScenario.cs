using System.Collections.Generic;
using YamlDotNet.Serialization;

namespace DocxBridge
{
    public class TestScenario
    {
        public string Name { get; set; } = string.Empty;
        public string Description { get; set; } = string.Empty;
        public string DocumentType { get; set; } = string.Empty;
        public string Category { get; set; } = string.Empty;
        public List<string> Tags { get; set; } = new();
        public ToleranceConfig Tolerance { get; set; } = new();
        public List<Operation> Operations { get; set; } = new();
    }

    public class ToleranceConfig
    {
        public bool XmlAttributeOrderSensitive { get; set; }
        public double VisualPixelTolerance { get; set; }
        public double VisualDiffThreshold { get; set; }
    }

    public class Operation
    {
        public string Action { get; set; } = string.Empty;

        // All other properties get collected here via YamlMember attribute
        [YamlIgnore]
        public Dictionary<object, object> Properties { get; set; } = new();

        // Individual properties that might appear in YAML
        // These are collected and moved to Properties dictionary after deserialization
        public string? Text { get; set; }
        public string? DocType { get; set; }
        public int? Rows { get; set; }
        public int? Cols { get; set; }
        public string? Style { get; set; }
        public int? WidthPercent { get; set; }
        public List<object>? Operations { get; set; }
        public List<object>? Cells { get; set; }

        // Handle nested 'properties' key in YAML (for paragraph/run properties)
        [YamlMember(Alias = "properties")]
        public Dictionary<object, object>? PropertiesNested { get; set; }
    }
}
