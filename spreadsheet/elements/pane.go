package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PaneState represents the state of the pane.
type PaneState string

const (
	// PaneStateSplit indicates the pane is split.
	PaneStateSplit PaneState = "split"
	// PaneStateFrozen indicates the pane is frozen.
	PaneStateFrozen PaneState = "frozen"
	// PaneStateFrozenSplit indicates the pane is frozen and split.
	PaneStateFrozenSplit PaneState = "frozenSplit"
)

// PanePosition represents the active pane position.
type PanePosition string

const (
	// PanePositionBottomRight is the bottom-right pane.
	PanePositionBottomRight PanePosition = "bottomRight"
	// PanePositionTopRight is the top-right pane.
	PanePositionTopRight PanePosition = "topRight"
	// PanePositionBottomLeft is the bottom-left pane.
	PanePositionBottomLeft PanePosition = "bottomLeft"
	// PanePositionTopLeft is the top-left pane.
	PanePositionTopLeft PanePosition = "topLeft"
)

// Pane represents the pane element (x:pane) for frozen or split panes.
type Pane struct {
	*openxml.CompositeElementBase
}

// NewPane creates a new Pane element.
func NewPane() *Pane {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pane",
		PrefixDefault,
	)

	return &Pane{CompositeElementBase: elem}
}

// XSplit returns the horizontal split position.
// For frozen panes, this is the column count to the left of the frozen pane.
// For split panes, this is the position in twips.
func (p *Pane) XSplit() float64 {
	attr, found := p.GetAttribute("xSplit", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetXSplit sets the horizontal split position.
func (p *Pane) SetXSplit(value float64) {
	if value == 0 {
		p.RemoveAttribute("xSplit", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"xSplit",
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// YSplit returns the vertical split position.
// For frozen panes, this is the row count above the frozen pane.
// For split panes, this is the position in twips.
func (p *Pane) YSplit() float64 {
	attr, found := p.GetAttribute("ySplit", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetYSplit sets the vertical split position.
func (p *Pane) SetYSplit(value float64) {
	if value == 0 {
		p.RemoveAttribute("ySplit", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"ySplit",
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// TopLeftCell returns the top-left cell visible in the bottom-right pane.
func (p *Pane) TopLeftCell() string {
	attr, found := p.GetAttribute(
		"topLeftCell",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTopLeftCell sets the top-left cell visible in the bottom-right pane.
func (p *Pane) SetTopLeftCell(cell string) {
	if cell == "" {
		p.RemoveAttribute("topLeftCell", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"topLeftCell",
			"",
			cell,
		),
	)
}

// ActivePane returns the active pane position.
func (p *Pane) ActivePane() PanePosition {
	attr, found := p.GetAttribute(
		"activePane",
		"",
	)
	if !found {
		return PanePositionTopLeft
	}

	return PanePosition(attr.Value())
}

// SetActivePane sets the active pane position.
func (p *Pane) SetActivePane(pos PanePosition) {
	if pos == "" || pos == PanePositionTopLeft {
		p.RemoveAttribute("activePane", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"activePane",
			"",
			string(pos),
		),
	)
}

// State returns the pane state.
func (p *Pane) State() PaneState {
	attr, found := p.GetAttribute("state", "")
	if !found {
		return PaneStateSplit
	}

	return PaneState(attr.Value())
}

// SetState sets the pane state.
func (p *Pane) SetState(state PaneState) {
	if state == "" || state == PaneStateSplit {
		p.RemoveAttribute("state", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"state",
			"",
			string(state),
		),
	)
}

// FreezePanes sets up frozen panes at the specified row and column.
// Rows above frozenRow and columns left of frozenCol will be frozen.
func (p *Pane) FreezePanes(
	frozenRow, frozenCol int,
) {
	p.SetState(PaneStateFrozen)
	if frozenCol > 0 {
		p.SetXSplit(float64(frozenCol))
	}
	if frozenRow > 0 {
		p.SetYSplit(float64(frozenRow))
	}

	// Set the active pane based on what's frozen
	switch {
	case frozenRow > 0 && frozenCol > 0:
		p.SetActivePane(PanePositionBottomRight)
	case frozenRow > 0:
		p.SetActivePane(PanePositionBottomLeft)
	case frozenCol > 0:
		p.SetActivePane(PanePositionTopRight)
	}
}

// Clone creates a deep copy of this Pane element.
func (p *Pane) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &Pane{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Pane element.
func (p *Pane) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &Pane{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
