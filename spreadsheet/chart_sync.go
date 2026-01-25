//go:build ignore
// +build ignore

// This file implements chart data cache synchronization.
package spreadsheet

import (
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
)

// Synchronize updates the chart data cache from the sheet data.
func (c *Chart) Synchronize() error {
	if c.chartPart == nil {
		return nil
	}
	chartSpace := c.chartPart.ChartSpace()
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

	// Helper to extract numeric values
	getNumericValues := func(ref string) ([]float64, error) {
		rawValues, err := c.sheet.ExtractValues(ref)
		if err != nil {
			return nil, err
		}
		numValues := make([]float64, len(rawValues))
		for i, v := range rawValues {
			if v == nil {
				numValues[i] = 0 // Or NaN? Excel uses empty usually.
				continue
			}
			switch val := v.(type) {
			case float64:
				numValues[i] = val
			case int:
				numValues[i] = float64(val)
			case string:
				f, _ := strconv.ParseFloat(val, 64)
				numValues[i] = f
			case bool:
				if val {
					numValues[i] = 1
				} else {
					numValues[i] = 0
				}
			}
		}
		return numValues, nil
	}

	// Helper to extract string values
	getStringValues := func(ref string) ([]string, error) {
		rawValues, err := c.sheet.ExtractValues(ref)
		if err != nil {
			return nil, err
		}
		strValues := make([]string, len(rawValues))
		for i, v := range rawValues {
			strValues[i] = stringValue(v)
		}
		return strValues, nil
	}

	for _, s := range c.series {
		// Find series in plot area
		var targetSeries interface {
			SetValues(values *drawingml.Values)
			SetCategoryAxisData(data *drawingml.CategoryAxisData)
		}
		var xySeries interface {
			SetXValues(values *drawingml.XValues)
			SetYValues(values *drawingml.YValues)
		}
		var bubbleSeries interface {
			SetXValues(values *drawingml.XValues)
			SetYValues(values *drawingml.YValues)
			SetBubbleSize(values *drawingml.BubbleSize)
		}

		found := false
		for child := range plotArea.Children() {
			switch ch := child.(type) {
			case *drawingml.BarChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.LineChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.PieChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.AreaChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.ScatterChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						xySeries = ser
						found = true
					}
				}
			case *drawingml.BubbleChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						bubbleSeries = ser
						found = true
					}
				}
			case *drawingml.DoughnutChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.RadarChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.Bar3DChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.Line3DChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.Pie3DChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.Area3DChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.SurfaceChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.Surface3DChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.OfPieChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			case *drawingml.StockChart:
				for _, ser := range ch.Series() {
					if ser.Index() == s.index {
						targetSeries = ser
						found = true
					}
				}
			}
			if found {
				break
			}
		}

		if !found {
			continue
		}

		if targetSeries != nil {
			if s.valuesRange != "" {
				vals, err := getNumericValues(s.valuesRange)
				if err == nil {
					v := drawingml.NewValues()
					v.SetNumberReferenceWithCache(s.valuesRange, vals)
					targetSeries.SetValues(v)
				}
			}
			if s.catRange != "" {
				cats, err := getStringValues(s.catRange)
				if err == nil {
					c := drawingml.NewCategoryAxisData()
					c.SetStringReference(s.catRange)
					if len(cats) > 0 {
						// Create string cache if we have values
						strRef := drawingml.NewStringReferenceWithCache(s.catRange, cats)
						c.AppendChild(strRef)
					}
					targetSeries.SetCategoryAxisData(c)
				}
			}
		} else if bubbleSeries != nil {
			if s.valuesRange != "" { // Y values
				vals, err := getNumericValues(s.valuesRange)
				if err == nil {
					v := drawingml.NewYValues()
					v.SetNumberReferenceWithCache(s.valuesRange, vals)
					bubbleSeries.SetYValues(v)
				}
			}
			if s.catRange != "" { // X values
				vals, err := getNumericValues(s.catRange)
				if err == nil {
					v := drawingml.NewXValues()
					v.SetNumberReferenceWithCache(s.catRange, vals)
					bubbleSeries.SetXValues(v)
				}
			}
			if s.bubbleSizeRange != "" {
				vals, err := getNumericValues(s.bubbleSizeRange)
				if err == nil {
					v := drawingml.NewBubbleSize()
					v.SetNumberReferenceWithCache(s.bubbleSizeRange, vals)
					bubbleSeries.SetBubbleSize(v)
				}
			}
		} else if xySeries != nil {
			if s.valuesRange != "" { // Y values
				vals, err := getNumericValues(s.valuesRange)
				if err == nil {
					v := drawingml.NewYValues()
					v.SetNumberReferenceWithCache(s.valuesRange, vals)
					xySeries.SetYValues(v)
				}
			}
			if s.catRange != "" { // X values
				vals, err := getNumericValues(s.catRange)
				if err == nil {
					v := drawingml.NewXValues()
					v.SetNumberReferenceWithCache(s.catRange, vals)
					xySeries.SetXValues(v)
				}
			}
		}
	}

	return nil
}
