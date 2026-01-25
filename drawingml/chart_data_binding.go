package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// ChartDataBinding defines the interface for data binding in charts.
type ChartDataBinding interface {
	Formula() string
	SetFormula(formula string)
}

// ChartNumericBinding defines the interface for numeric data binding.
type ChartNumericBinding interface {
	ChartDataBinding
	SetCache(values []float64)
	SetCacheWithFormat(values []float64, formatCode string)
}

// ChartStringBinding defines the interface for string data binding.
type ChartStringBinding interface {
	ChartDataBinding
	SetCache(values []string)
}

// DataLabelProvider defines the interface for elements that support data labels.
type DataLabelProvider interface {
	DataLabels() *DataLabels
	SetDataLabels(dl *DataLabels)
}

// ChartElement is a marker interface for any chart element.
type ChartElement interface {
	openxml.Element
}
