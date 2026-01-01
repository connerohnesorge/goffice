package layout

// ListLayout positions shapes in a linear flow
type ListLayout struct {
	Orientation Orientation
	Constraints Constraints
}

// NewListLayout creates a new list layout engine
func NewListLayout() *ListLayout {
	return &ListLayout{
		Orientation: OrientationVertical,
		Constraints: DefaultConstraints(),
	}
}

// Type returns the layout type this engine handles
func (*ListLayout) Type() LayoutType {
	return LayoutTypeList
}

// Layout positions shapes in a linear flow (vertical or horizontal)
func (l *ListLayout) Layout(shapes []LayoutShape, bounds Rectangle) error {
	if len(shapes) == 0 {
		return nil
	}

	c := l.Constraints
	availWidth := bounds.Width - 2*c.Padding
	availHeight := bounds.Height - 2*c.Padding

	if l.Orientation == OrientationVertical {
		return l.layoutVertical(shapes, bounds, availWidth, availHeight)
	}

	return l.layoutHorizontal(shapes, bounds, availWidth, availHeight)
}

func (l *ListLayout) layoutVertical(shapes []LayoutShape, bounds Rectangle, availWidth, availHeight int64) error {
	c := l.Constraints
	n := int64(len(shapes))

	// Calculate shape height with equal distribution
	totalSpacing := c.Spacing * (n - 1)
	shapeHeight := (availHeight - totalSpacing) / n
	if shapeHeight < c.MinHeight {
		shapeHeight = c.MinHeight
	}

	// Position each shape
	y := bounds.Y + c.Padding
	for i := range shapes {
		shapes[i].Width = availWidth
		shapes[i].Height = shapeHeight
		shapes[i].X = bounds.X + c.Padding
		shapes[i].Y = y
		y += shapeHeight + c.Spacing
	}

	return nil
}

func (l *ListLayout) layoutHorizontal(shapes []LayoutShape, bounds Rectangle, availWidth, availHeight int64) error {
	c := l.Constraints
	n := int64(len(shapes))

	// Calculate shape width with equal distribution
	totalSpacing := c.Spacing * (n - 1)
	shapeWidth := (availWidth - totalSpacing) / n
	if shapeWidth < c.MinWidth {
		shapeWidth = c.MinWidth
	}

	// Position each shape
	x := bounds.X + c.Padding
	for i := range shapes {
		shapes[i].Width = shapeWidth
		shapes[i].Height = availHeight
		shapes[i].X = x
		shapes[i].Y = bounds.Y + c.Padding
		x += shapeWidth + c.Spacing
	}

	return nil
}

func init() {
	DefaultRegistry().Register(NewListLayout())
}
