package main

// EMU (English Metric Units) position and size constants
const (
	// Title shape positioning
	titlePosX   = 914400  // Title X position
	titlePosY   = 365125  // Title Y position (top center)
	titleWidth  = 7315200 // Title width
	titleHeight = 914400  // Title height

	// Chart shape positioning
	chartPosX   = 1524000 // Chart X position (center)
	chartPosY   = 1828800 // Chart Y position
	chartWidth  = 6096000 // Chart width
	chartHeight = 4572000 // Chart height

	// Font sizes (in hundredths of a point)
	titleFontSize        = 2800 // 28pt
	chartPlaceholderSize = 1400 // 14pt

	// Chart specific
	doughnutHoleSize = 50 // 50% hole size for doughnut chart
)

// Excel cell reference constants for chart data
const (
	// Bar chart references (Quarterly data)
	barCategoryRange = "Sheet1!$A$2:$A$5"
	barNorthRange    = "Sheet1!$B$2:$B$5"
	barSouthRange    = "Sheet1!$C$2:$C$5"
	barEastRange     = "Sheet1!$D$2:$D$5"

	// Line chart references (Monthly data)
	lineCategoryRange = "Sheet1!$A$2:$A$13"
	lineNYRange       = "Sheet1!$B$2:$B$13"
	lineLARange       = "Sheet1!$C$2:$C$13"
	lineChicagoRange  = "Sheet1!$D$2:$D$13"

	// Pie chart references
	pieCategoryRange = "Sheet1!$A$2:$A$6"
	pieValuesRange   = "Sheet1!$B$2:$B$6"

	// Area chart references
	areaCategoryRange = "Sheet1!$A$2:$A$7"
	areaProductARange = "Sheet1!$B$2:$B$7"
	areaProductBRange = "Sheet1!$C$2:$C$7"
	areaProductCRange = "Sheet1!$D$2:$D$7"

	// Scatter chart references
	scatterXRange         = "Sheet1!$A$2:$A$11"
	scatterProduct1YRange = "Sheet1!$B$2:$B$11"
	scatterProduct2YRange = "Sheet1!$C$2:$C$11"

	// Doughnut chart references
	doughnutCategoryRange = "Sheet1!$A$2:$A$7"
	doughnutValuesRange   = "Sheet1!$B$2:$B$7"
)

// Axis ID constants
const (
	categoryAxisID = 1
	valueAxisID    = 2
)

// Error message constants
const (
	errFailedToAddSlide = "failed to add slide: %w"
)
