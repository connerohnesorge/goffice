// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"strings"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing"
)

// OoxmlMailMergeMetadata contains OOXML-specific mail merge metadata
// extracted from the document settings (w:mailMerge element in settings.xml).
//
// This metadata is created by Microsoft Word when you set up a mail merge in the UI.
// It describes the data source connection and field name mappings, enabling compatibility
// with Word-created templates.
//
// The metadata is automatically loaded during Execute() and ExecuteToDocuments() and
// applied transparently - you don't need to interact with it directly unless you're
// implementing advanced customizations.
//
// Field mappings allow document field names to differ from data source column names.
// For example, a MERGEFIELD named "FirstName" could map to a data source column "fname".
type OoxmlMailMergeMetadata struct {
	// DataType indicates the type of data source.
	// Common values: "textFile", "database", "spreadsheet", "addressBook"
	DataType string

	// ConnectString contains the connection information for the data source.
	// For files: file path
	// For databases: connection string
	ConnectString string

	// Query contains the query to execute against the data source.
	// For databases: SQL query
	// For files: file path or filter expression
	Query string

	// FieldMappings maps document field names to data source field names.
	// Key: document field name (what appears in MERGEFIELD)
	// Value: data source field name (column name in data source)
	// This supports scenarios where document fields have different names
	// than the actual data source columns.
	FieldMappings map[string]string
}

//
//nolint:revive // cognitive-complexity: metadata extraction has many validation steps
func loadMailMergeMetadata(
	doc *wordprocessing.Document,
) (*OoxmlMailMergeMetadata, error) {
	// Get the main document part
	mainPart := doc.MainPart()
	if mainPart == nil {
		return nil, nil
	}

	// Get the settings part
	settingsPart := mainPart.SettingsPart()
	if settingsPart == nil {
		// No settings part means no mail merge metadata
		return nil, nil
	}

	// Get the settings root element
	settings := settingsPart.Settings()
	if settings == nil {
		return nil, nil
	}

	// Get the mail merge element
	mailMerge := settings.MailMerge()
	if mailMerge == nil {
		// No mail merge configuration
		return nil, nil
	}

	// Create metadata structure
	metadata := &OoxmlMailMergeMetadata{
		FieldMappings: make(map[string]string),
	}

	// Extract DataType
	metadata.DataType = getChildElementValue(
		mailMerge,
		"dataType",
	)

	// Extract ConnectString
	metadata.ConnectString = getChildElementValue(
		mailMerge,
		"connectString",
	)

	// Extract Query
	metadata.Query = getChildElementValue(
		mailMerge,
		"query",
	)

	// Extract field mappings from ODSO (Office Data Source Object)
	extractFieldMappings(mailMerge, metadata)

	return metadata, nil
}

// extractFieldMappings extracts field mappings from the ODSO element.
func extractFieldMappings(
	mailMerge openxml.CompositeElement,
	metadata *OoxmlMailMergeMetadata,
) {
	odso := mailMerge.GetElement(
		"odso",
		openxml.NamespaceWordprocessingML,
	)
	if odso == nil {
		return
	}

	odsoComposite, ok := odso.(openxml.CompositeElement)
	if !ok {
		return
	}

	// Iterate through child elements looking for fieldMapData
	for child := range odsoComposite.Children() {
		if child.LocalName() != "fieldMapData" {
			continue
		}

		fmd, ok := child.(openxml.CompositeElement)
		if !ok {
			continue
		}

		// Extract name (document field) and mappedName (data source field)
		docFieldName := getChildElementValue(
			fmd,
			"name",
		)
		sourceFieldName := getChildElementValue(
			fmd,
			"mappedName",
		)

		if docFieldName != "" &&
			sourceFieldName != "" {
			// Store mapping (case-insensitive key for lookups)
			metadata.FieldMappings[strings.ToLower(docFieldName)] = sourceFieldName
		}
	}
}

// getChildElementValue extracts the w:val attribute value from a child element.
// This is a common pattern in OOXML where elements have a w:val attribute.
func getChildElementValue(
	parent openxml.CompositeElement,
	childName string,
) string {
	elem := parent.GetElement(
		childName,
		openxml.NamespaceWordprocessingML,
	)
	if elem == nil {
		return ""
	}

	attr, found := elem.GetAttribute(
		"val",
		openxml.NamespaceWordprocessingML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// applyFieldMapping applies field name mapping from OOXML metadata.
// If metadata is nil or no mapping exists, returns the original field name.
// Field name matching is case-insensitive.
func applyFieldMapping(
	fieldName string,
	metadata *OoxmlMailMergeMetadata,
) string {
	if metadata == nil {
		return fieldName
	}

	if len(metadata.FieldMappings) == 0 {
		return fieldName
	}

	// Look up mapping (case-insensitive)
	mappedName, found := metadata.FieldMappings[strings.ToLower(fieldName)]
	if !found {
		return fieldName
	}

	return mappedName
}
