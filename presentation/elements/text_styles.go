package elements

import (
	"fmt"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// getParagraphProps returns the paragraph properties for the given level.
// level 0 = defPPr, 1-9 = lvl1pPr...lvl9pPr
func getParagraphProps(
	c *openxml.CompositeElementBase,
	level int,
) *drawingml.TextParagraphProperties {
	var name string
	if level == 0 {
		name = "defPPr"
	} else if level >= 1 && level <= 9 {
		name = fmt.Sprintf("lvl%dpPr", level)
	} else {
		return nil
	}

	elem := c.GetElement(name, openxml.NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if ppr, ok := elem.(*drawingml.TextParagraphProperties); ok {
		return ppr
	}
	// Fallback for wrapped elements
	// Note: We can't easily reconstruct strict types if they are generic,
	// but drawingml types are usually well-defined.
	return nil
}

// getOrCreateParagraphProps returns or creates the paragraph properties.
func getOrCreateParagraphProps(
	c *openxml.CompositeElementBase,
	level int,
) *drawingml.TextParagraphProperties {
	ppr := getParagraphProps(c, level)
	if ppr != nil {
		return ppr
	}

	var name string
	if level == 0 {
		name = "defPPr"
	} else if level >= 1 && level <= 9 {
		name = fmt.Sprintf("lvl%dpPr", level)
	} else {
		return nil
	}

	// Create new element
	// We use drawingml.NewTextParagraphProperties() but we need to set the tag name
	newPPr := drawingml.NewTextParagraphProperties()
	// Helper to set name/namespace on the underlying element
	// Note: TextParagraphProperties in drawingml usually has "pPr" as tag.
	// We need to override it to "defPPr" or "lvlXpPr" and namespace "a" (drawingml).
	// The CompositeElementBase inside newPPr needs to be updated.

	// However, we can't easily replace the inner CompositeElementBase of an existing struct
	// without reflection or if exported.
	// openxml.NewCompositeElement is used in NewTextParagraphProperties.

	// Better approach: Create a new CompositeElement with correct name,
	// and wrap it in TextParagraphProperties.

	elem := openxml.NewCompositeElement(openxml.NamespaceDrawingML, name, "a")
	newPPr.CompositeElementBase = elem

	c.AppendChild(newPPr)
	return newPPr
}

// ===========================================================================
// TitleStyle Methods
// ===========================================================================

func (ts *TitleStyle) DefaultParagraphProperties() *drawingml.TextParagraphProperties {
	return getParagraphProps(ts.CompositeElementBase, 0)
}

func (ts *TitleStyle) GetOrCreateDefaultParagraphProperties() *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(ts.CompositeElementBase, 0)
}

func (ts *TitleStyle) Level1ParagraphProperties() *drawingml.TextParagraphProperties {
	return getParagraphProps(ts.CompositeElementBase, 1)
}

func (ts *TitleStyle) GetOrCreateLevel1ParagraphProperties() *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(ts.CompositeElementBase, 1)
}

// ... We can add more levels as needed, but for now 1 is most critical.
// Adding generic accessor by level
func (ts *TitleStyle) LevelParagraphProperties(level int) *drawingml.TextParagraphProperties {
	return getParagraphProps(ts.CompositeElementBase, level)
}

func (ts *TitleStyle) GetOrCreateLevelParagraphProperties(
	level int,
) *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(ts.CompositeElementBase, level)
}

// ===========================================================================
// BodyStyle Methods
// ===========================================================================

func (bs *BodyStyle) DefaultParagraphProperties() *drawingml.TextParagraphProperties {
	return getParagraphProps(bs.CompositeElementBase, 0)
}

func (bs *BodyStyle) GetOrCreateDefaultParagraphProperties() *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(bs.CompositeElementBase, 0)
}

func (bs *BodyStyle) LevelParagraphProperties(level int) *drawingml.TextParagraphProperties {
	return getParagraphProps(bs.CompositeElementBase, level)
}

func (bs *BodyStyle) GetOrCreateLevelParagraphProperties(
	level int,
) *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(bs.CompositeElementBase, level)
}

// ===========================================================================
// OtherStyle Methods
// ===========================================================================

func (os *OtherStyle) DefaultParagraphProperties() *drawingml.TextParagraphProperties {
	return getParagraphProps(os.CompositeElementBase, 0)
}

func (os *OtherStyle) GetOrCreateDefaultParagraphProperties() *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(os.CompositeElementBase, 0)
}

func (os *OtherStyle) LevelParagraphProperties(level int) *drawingml.TextParagraphProperties {
	return getParagraphProps(os.CompositeElementBase, level)
}

func (os *OtherStyle) GetOrCreateLevelParagraphProperties(
	level int,
) *drawingml.TextParagraphProperties {
	return getOrCreateParagraphProps(os.CompositeElementBase, level)
}
