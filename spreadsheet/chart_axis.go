// This file implements chart axis functionality.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/drawingml"
)

// ChartAxis represents an axis in a chart.
type ChartAxis struct {
	chart    *Chart
	axisType string
	title    string
	minVal   *float64
	maxVal   *float64
}

// newChartAxis creates a new ChartAxis.
func newChartAxis(
	chart *Chart,
	axisType string,
) *ChartAxis {
	return &ChartAxis{
		chart:    chart,
		axisType: axisType,
	}
}

// Title returns the axis title.
func (a *ChartAxis) Title() string {
	return a.title
}

// SetTitle sets the axis title.
func (a *ChartAxis) SetTitle(title string) {
	a.title = title

	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if title == "" {
		// Remove title
		switch ax := axis.(type) {
		case *drawingml.CategoryAxis:
			ax.SetTitle(nil)
		case *drawingml.ValueAxis:
			ax.SetTitle(nil)
		}
	} else {
		// Create and set title
		titleElem := drawingml.NewTitleWithText(title)
		switch ax := axis.(type) {
		case *drawingml.CategoryAxis:
			ax.SetTitle(titleElem)
		case *drawingml.ValueAxis:
			ax.SetTitle(titleElem)
		}
	}
}

// SetMinimum sets the minimum value for the axis.
func (a *ChartAxis) SetMinimum(val float64) {
	a.minVal = &val

	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	// Get the scaling element and set minimum
	switch ax := axis.(type) {
	case *drawingml.CategoryAxis:
		// Category axes don't typically have numeric min/max
	case *drawingml.ValueAxis:
		scaling := a.getOrCreateScaling(ax)
		if scaling != nil {
			scaling.SetMinimum(val)
		}
	}
}

// SetMaximum sets the maximum value for the axis.
func (a *ChartAxis) SetMaximum(val float64) {
	a.maxVal = &val

	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	// Get the scaling element and set maximum
	switch ax := axis.(type) {
	case *drawingml.CategoryAxis:
		// Category axes don't typically have numeric min/max
	case *drawingml.ValueAxis:
		scaling := a.getOrCreateScaling(ax)
		if scaling != nil {
			scaling.SetMaximum(val)
		}
	}
}

// SetMajorGridlines sets whether to show major gridlines.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *ChartAxis) SetMajorGridlines(show bool) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if ax, ok := axis.(*drawingml.ValueAxis); ok {
		ax.SetMajorGridlines(show)
	}
}

// SetMinorGridlines sets whether to show minor gridlines.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *ChartAxis) SetMinorGridlines(show bool) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if ax, ok := axis.(*drawingml.ValueAxis); ok {
		ax.SetMinorGridlines(show)
	}
}

// SetNumberFormat sets the number format for axis labels.
func (a *ChartAxis) SetNumberFormat(
	format string,
) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if ax, ok := axis.(*drawingml.ValueAxis); ok {
		ax.SetNumberFormat(format, false)
	}
}

// getUnderlyingAxis returns the underlying axis element.
func (a *ChartAxis) getUnderlyingAxis() any {
	chartSpace := a.chart.chartPart.ChartSpace()
	if chartSpace == nil {
		return nil
	}

	chart := chartSpace.Chart()
	if chart == nil {
		return nil
	}

	plotArea := chart.PlotArea()
	if plotArea == nil {
		return nil
	}

	// Find the appropriate axis in the plot area
	// This is simplified - a real implementation would need to track
	// which specific axis element corresponds to this ChartAxis
	children := plotArea.Children()
	for child := range children {
		switch ax := child.(type) {
		case *drawingml.CategoryAxis:
			if a.axisType == "category" {
				return ax
			}
		case *drawingml.ValueAxis:
			if a.axisType == "value" {
				return ax
			}
		}
	}

	return nil
}

// getOrCreateScaling gets or creates a scaling element for a value axis.
func (*ChartAxis) getOrCreateScaling(
	ax *drawingml.ValueAxis,
) *drawingml.Scaling {
	// The ValueAxis already has a scaling element created in NewValueAxis
	// We need to access it through the element's children
	children := ax.Children()
	for child := range children {
		if scaling, ok := child.(*drawingml.Scaling); ok {
			return scaling
		}
	}

	return nil
}

// SetLogScale sets logarithmic scale for the axis.
func (a *ChartAxis) SetLogScale(logBase float64) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if ax, ok := axis.(*drawingml.ValueAxis); ok {
		scaling := a.getOrCreateScaling(ax)
		if scaling != nil {
			scaling.SetLogBase(logBase)
		}
	}
}

// SetOrientation sets axis orientation.
func (a *ChartAxis) SetOrientation(orientation drawingml.OrientationValue) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if ax, ok := axis.(*drawingml.ValueAxis); ok {
		scaling := a.getOrCreateScaling(ax)
		if scaling != nil {
			scaling.SetOrientation(orientation)
		}
	}
}

// SetReverseOrder reverses the axis order.
func (a *ChartAxis) SetReverseOrder(reverse bool) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	if ax, ok := axis.(*drawingml.ValueAxis); ok {
		scaling := a.getOrCreateScaling(ax)
		if scaling != nil {
			if reverse {
				scaling.SetOrientation(drawingml.OrientationMaxMin)
			} else {
				scaling.SetOrientation(drawingml.OrientationMinMax)
			}
		}
	}
}

// SetAxisPosition sets the axis position.
func (a *ChartAxis) SetAxisPosition(pos drawingml.AxisPositionValue) {
	axis := a.getUnderlyingAxis()
	if axis == nil {
		return
	}

	switch ax := axis.(type) {
	case *drawingml.CategoryAxis:
		ax.SetAxisPosition(pos)
	case *drawingml.ValueAxis:
		ax.SetAxisPosition(pos)
	}
}
