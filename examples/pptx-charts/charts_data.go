package main

// Bar chart data - Quarterly sales for 3 regions
var (
	barCategories = []string{
		"Q1",
		"Q2",
		"Q3",
		"Q4",
	}
	barNorthData = []float64{120, 150, 180, 200}
	barSouthData = []float64{100, 130, 160, 190}
	barEastData  = []float64{80, 110, 140, 170}
)

// Line chart data - Monthly temperatures for 3 cities
var (
	lineCategories = []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}
	lineNYData = []float64{
		32,
		35,
		45,
		56,
		67,
		76,
		81,
		79,
		71,
		60,
		48,
		37,
	}
	lineLAData = []float64{
		58,
		60,
		63,
		66,
		69,
		73,
		77,
		78,
		76,
		71,
		64,
		59,
	}
	lineChicagoData = []float64{
		25,
		29,
		40,
		52,
		63,
		73,
		77,
		76,
		68,
		55,
		42,
		30,
	}
)

// Area chart data - Revenue growth for 3 product lines
var (
	areaCategories = []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
	}
	areaProductAData = []float64{
		100,
		120,
		140,
		160,
		180,
		200,
	}
	areaProductBData = []float64{
		80,
		95,
		110,
		125,
		140,
		155,
	}
	areaProductCData = []float64{
		60,
		70,
		80,
		90,
		100,
		110,
	}
)

// Pie chart data - Market share distribution
var (
	pieCategories = []string{
		"Company A",
		"Company B",
		"Company C",
		"Company D",
		"Company E",
	}
	pieValues = []float64{30, 25, 20, 15, 10}
)

// Scatter chart data - Price vs Demand correlation
var (
	scatterPrices = []float64{
		10,
		15,
		20,
		25,
		30,
		35,
		40,
		45,
		50,
		55,
	}
	scatterProduct1Demand = []float64{
		100,
		95,
		85,
		75,
		65,
		55,
		45,
		35,
		25,
		20,
	}
	scatterProduct2Demand = []float64{
		90,
		85,
		78,
		70,
		62,
		54,
		46,
		38,
		30,
		25,
	}
)

// Doughnut chart data - Budget allocation by department
var (
	doughnutCategories = []string{
		"Engineering",
		"Marketing",
		"Sales",
		"Operations",
		"HR",
		"R&D",
	}
	doughnutValues = []float64{
		35,
		20,
		18,
		12,
		8,
		7,
	}
)
