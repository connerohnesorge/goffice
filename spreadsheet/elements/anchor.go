package elements

// EditAs represents the anchor behavior type.
type EditAs string

// EditAs values.
const (
	EditAsTwoCell  EditAs = "twoCell"  // Drawing resizes with cells
	EditAsOneCell  EditAs = "oneCell"  // Drawing moves but doesn't resize
	EditAsAbsolute EditAs = "absolute" // Drawing doesn't move or resize
)
