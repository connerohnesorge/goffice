package elements

import (
	"encoding/xml"
)

// SpaceProcessingModeValues for xml:space attribute
type SpaceProcessingModeValues string

const (
	SpaceProcessingModeDefault  SpaceProcessingModeValues = "default"
	SpaceProcessingModePreserve SpaceProcessingModeValues = "preserve"
)

func (e SpaceProcessingModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SpaceProcessingModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SpaceProcessingModeValues(attr.Value)

	return nil
}
