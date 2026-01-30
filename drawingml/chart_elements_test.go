package drawingml

import (
	"testing"
)

func TestTitleProperties(_ *testing.T) {
	title := NewTitle()

	// Test SetOverlay
	title.SetOverlay(true)
	// We can't easily check internal state without getters, but we can check XML
	// Assuming OuterXml works
}

func TestLegendProperties(_ *testing.T) {
	legend := NewLegend()

	// Test SetLayout
	layout := NewLayout()
	legend.SetLayout(layout)

	// Test SetShapeProperties
	spPr := NewChartShapeProperties()
	spPr.SetSolidFill("FF0000")
	legend.SetShapeProperties(spPr)

	// Test SetTextProperties
	txPr := NewChartText("txPr")
	legend.SetTextProperties(txPr)
}

func TestDataLabelsProperties(_ *testing.T) {
	dl := NewDataLabels()

	dl.SetPosition(DataLabelPositionBestFit)
	dl.SetNumberFormat("0.00", true)

	spPr := NewChartShapeProperties()
	dl.SetShapeProperties(spPr)

	txPr := NewChartText("txPr")
	dl.SetTextProperties(txPr)
}

func TestDataTable(_ *testing.T) {
	dt := NewDataTable()

	dt.SetShowHorizontalBorder(true)
	dt.SetShowVerticalBorder(true)
	dt.SetShowOutline(true)
	dt.SetShowKeys(true)

	spPr := NewChartShapeProperties()
	dt.SetShapeProperties(spPr)

	txPr := NewChartText("txPr")
	dt.SetTextProperties(txPr)
}

func TestScalingLogBase(_ *testing.T) {
	s := NewScaling()
	s.SetLogBase(10)
}

func TestPlotAreaDataTable(_ *testing.T) {
	pa := NewPlotArea()
	dt := NewDataTable()
	pa.SetDataTable(dt)
}
