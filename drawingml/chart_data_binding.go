package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// ChartDataBinding defines the interface for data binding in charts.
type ChartDataBinding interface {
	// Formula returns the formula string for the data binding.
	Formula() string
	// SetFormula sets the formula string for the data binding.
	SetFormula(formula string)
}

// ChartNumericBinding defines the interface for numeric data binding.
type ChartNumericBinding interface {
	ChartDataBinding
	// SetCache sets cached numeric values for the data binding.
	SetCache(values []float64)
	// SetCacheWithFormat sets cached numeric values with a custom format code.
	SetCacheWithFormat(values []float64, formatCode string)
}

// ChartStringBinding defines the interface for string data binding.
type ChartStringBinding interface {
	ChartDataBinding
	// SetCache sets cached string values for the data binding.
	SetCache(values []string)
}

// DataLabelProvider defines the interface for elements that support data labels.
type DataLabelProvider interface {
	// DataLabels returns the data labels element.
	DataLabels() *DataLabels
	// SetDataLabels sets the data labels element.
	SetDataLabels(dl *DataLabels)
}

// ChartElement is a marker interface for any chart element.
type ChartElement interface {
	openxml.Element
}
