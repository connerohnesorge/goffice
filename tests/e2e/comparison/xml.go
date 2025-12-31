package comparison

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// XMLDiffResult contains XML comparison results
type XMLDiffResult struct {
	Match               bool
	ElementMismatches   []ElementMismatch
	AttributeMismatches []AttributeMismatch
	ContentMismatches   []ContentMismatch
}

// ElementMismatch represents a missing or extra element
type ElementMismatch struct {
	Path     string // XPath-style path
	Expected string // Element name expected (from .NET)
	Actual   string // Element name actual (from Go)
}

// AttributeMismatch represents attribute value difference
type AttributeMismatch struct {
	Path      string // XPath to element
	Attribute string // Attribute name
	Expected  string // Value from .NET
	Actual    string // Value from Go
}

// ContentMismatch represents text content difference
type ContentMismatch struct {
	Path     string
	Expected string
	Actual   string
}

// XMLNode represents a simplified XML tree node
type XMLNode struct {
	Name       xml.Name
	Attributes map[string]string
	Content    string
	Children   []*XMLNode
}

// CompareXMLStructure compares XML structure of two DOCX files
func CompareXMLStructure(
	goPath, dotnetPath string,
	tolerance framework.ToleranceConfig,
) (*XMLDiffResult, error) {
	result := &XMLDiffResult{Match: true}

	// Open both ZIP archives
	goZip, err := zip.OpenReader(goPath)
	if err != nil {
		return nil, fmt.Errorf(
			"open go docx: %w",
			err,
		)
	}
	defer func() {
		_ = goZip.Close()
	}()

	dotnetZip, err := zip.OpenReader(dotnetPath)
	if err != nil {
		return nil, fmt.Errorf(
			"open dotnet docx: %w",
			err,
		)
	}
	defer func() {
		_ = dotnetZip.Close()
	}()

	// Compare document.xml (main document)
	goDocXML, err := extractXMLPart(
		&goZip.Reader,
		"word/document.xml",
	)
	if err != nil {
		return nil, fmt.Errorf(
			"extract go document.xml: %w",
			err,
		)
	}

	dotnetDocXML, err := extractXMLPart(
		&dotnetZip.Reader,
		"word/document.xml",
	)
	if err != nil {
		return nil, fmt.Errorf(
			"extract dotnet document.xml: %w",
			err,
		)
	}

	// Parse XML to trees
	goTree, err := parseXMLToTree(goDocXML)
	if err != nil {
		return nil, fmt.Errorf(
			"parse go XML: %w",
			err,
		)
	}

	dotnetTree, err := parseXMLToTree(
		dotnetDocXML,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"parse dotnet XML: %w",
			err,
		)
	}

	// Compare trees recursively
	compareNodes(
		goTree,
		dotnetTree,
		"",
		result,
		tolerance,
	)

	if len(result.ElementMismatches) > 0 ||
		len(result.AttributeMismatches) > 0 ||
		len(result.ContentMismatches) > 0 {
		result.Match = false
	}

	return result, nil
}

// extractXMLPart extracts an XML part from DOCX ZIP
func extractXMLPart(
	zipReader *zip.Reader,
	partName string,
) ([]byte, error) {
	for _, file := range zipReader.File {
		if file.Name != partName {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(rc)
		closeErr := rc.Close()
		if err != nil {
			return nil, err
		}
		if closeErr != nil {
			return nil, closeErr
		}

		return data, nil
	}

	return nil, fmt.Errorf(
		"part not found: %s",
		partName,
	)
}

// parseXMLToTree parses XML into a tree structure
func parseXMLToTree(
	data []byte,
) (*XMLNode, error) {
	decoder := xml.NewDecoder(
		strings.NewReader(string(data)),
	)

	var root *XMLNode
	var stack []*XMLNode

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			node := &XMLNode{
				Name:       t.Name,
				Attributes: make(map[string]string),
			}

			for _, attr := range t.Attr {
				node.Attributes[attr.Name.Local] = attr.Value
			}

			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, node)
			} else {
				root = node
			}

			stack = append(stack, node)

		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		case xml.CharData:
			if len(stack) > 0 {
				node := stack[len(stack)-1]
				content := strings.TrimSpace(string(t))
				if content != "" {
					node.Content += content
				}
			}
		}
	}

	return root, nil
}

// compareNodes recursively compares XML nodes
func compareNodes(
	goNode, dotnetNode *XMLNode,
	path string,
	result *XMLDiffResult,
	tolerance framework.ToleranceConfig,
) {
	if goNode == nil && dotnetNode == nil {
		return
	}

	currentPath := path + "/" + dotnetNode.Name.Local

	// Check element names match
	if goNode == nil {
		result.ElementMismatches = append(
			result.ElementMismatches,
			ElementMismatch{
				Path:     currentPath,
				Expected: dotnetNode.Name.Local,
				Actual:   "(missing)",
			},
		)

		return
	}

	if goNode.Name.Local != dotnetNode.Name.Local {
		result.ElementMismatches = append(
			result.ElementMismatches,
			ElementMismatch{
				Path:     currentPath,
				Expected: dotnetNode.Name.Local,
				Actual:   goNode.Name.Local,
			},
		)

		return
	}

	// Check attributes
	for attrName, dotnetValue := range dotnetNode.Attributes {
		goValue, exists := goNode.Attributes[attrName]
		if !exists {
			result.AttributeMismatches = append(
				result.AttributeMismatches,
				AttributeMismatch{
					Path:      currentPath,
					Attribute: attrName,
					Expected:  dotnetValue,
					Actual:    "(missing)",
				},
			)
		} else if goValue != dotnetValue {
			// Skip if attribute order doesn't matter and value is equivalent
			if tolerance.XMLAttributeOrderSensitive || goValue != dotnetValue {
				result.AttributeMismatches = append(result.AttributeMismatches, AttributeMismatch{
					Path:      currentPath,
					Attribute: attrName,
					Expected:  dotnetValue,
					Actual:    goValue,
				})
			}
		}
	}

	// Check for extra attributes in Go
	for attrName := range goNode.Attributes {
		if _, exists := dotnetNode.Attributes[attrName]; !exists {
			result.AttributeMismatches = append(
				result.AttributeMismatches,
				AttributeMismatch{
					Path:      currentPath,
					Attribute: attrName,
					Expected:  "(not present)",
					Actual:    goNode.Attributes[attrName],
				},
			)
		}
	}

	// Check text content
	if goNode.Content != dotnetNode.Content {
		result.ContentMismatches = append(
			result.ContentMismatches,
			ContentMismatch{
				Path:     currentPath,
				Expected: dotnetNode.Content,
				Actual:   goNode.Content,
			},
		)
	}

	// Compare children
	if len(
		goNode.Children,
	) != len(
		dotnetNode.Children,
	) {
		result.ElementMismatches = append(
			result.ElementMismatches,
			ElementMismatch{
				Path: currentPath,
				Expected: fmt.Sprintf(
					"%d children",
					len(dotnetNode.Children),
				),
				Actual: fmt.Sprintf(
					"%d children",
					len(goNode.Children),
				),
			},
		)

		return
	}

	for i := range dotnetNode.Children {
		compareNodes(
			goNode.Children[i],
			dotnetNode.Children[i],
			currentPath,
			result,
			tolerance,
		)
	}
}
