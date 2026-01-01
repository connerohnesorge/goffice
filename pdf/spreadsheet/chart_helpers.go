package spreadsheet

import (
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

func chartRelIDFromGraphicFrame(
	frame *elements.GraphicFrame,
) string {
	if frame == nil {
		return ""
	}
	graphic := frame.Graphic()
	if graphic == nil {
		return ""
	}
	graphicData := graphic.GraphicData()
	if graphicData == nil {
		return ""
	}

	for child := range graphicData.Children() {
		if child.LocalName() != "chart" ||
			child.NamespaceURI() != drawingml.NamespaceChart {
			continue
		}
		attr, found := child.GetAttribute(
			"id",
			openxml.NamespaceRelationships,
		)
		if !found {
			return ""
		}

		return attr.Value()
	}

	return ""
}
