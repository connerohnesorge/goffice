//go:build ignore
// +build ignore

//nolint:revive // Data file contains test values that should not be converted to constants
package main

// Data helper types and functions for ML results example
// These contain test data values that should not be converted to constants

type modelData struct {
	name   string
	params float64
	train  float64
	speed  float64
	top1   float64
	top5   float64
	mAP    float64
	memory float64
}

func getModelData() []modelData {
	return []modelData{
		{
			"ResNet-50",
			25.5,
			48.2,
			1250,
			76.1,
			92.9,
			37.2,
			2.1,
		},
		{
			"ResNet-101",
			44.5,
			72.5,
			890,
			77.8,
			93.7,
			39.1,
			3.2,
		},
		{
			"EfficientNet-B0",
			5.3,
			36.1,
			2150,
			77.1,
			93.5,
			38.4,
			0.9,
		},
		{
			"EfficientNet-B4",
			19.3,
			98.5,
			1050,
			82.9,
			96.3,
			42.8,
			2.8,
		},
		{
			"Vision Transformer-B",
			86.6,
			156.0,
			850,
			81.1,
			95.3,
			40.5,
			4.2,
		},
		{
			"YOLO-v8n",
			3.2,
			24.0,
			3500,
			72.5,
			90.1,
			35.6,
			0.6,
		},
		{
			"YOLO-v8s",
			11.2,
			42.0,
			2800,
			75.2,
			91.8,
			37.9,
			1.2,
		},
		{
			"YOLO-v8m",
			25.9,
			68.0,
			1950,
			79.3,
			94.1,
			40.2,
			2.3,
		},
		{
			"Swin-Transformer-T",
			28.3,
			120.0,
			720,
			81.5,
			95.5,
			41.1,
			3.5,
		},
		{
			"ConvNeXt-T",
			28.6,
			110.0,
			950,
			82.1,
			95.9,
			41.8,
			3.3,
		},
	}
}

type analysisData struct {
	name   string
	params float64
	acc    float64
}

func getAnalysisData() []analysisData {
	return []analysisData{
		{"ResNet-50", 25.5, 76.1},
		{"ResNet-101", 44.5, 77.8},
		{"EfficientNet-B0", 5.3, 77.1},
		{"EfficientNet-B4", 19.3, 82.9},
		{"Vision Transformer-B", 86.6, 81.1},
		{"YOLO-v8n", 3.2, 72.5},
		{"YOLO-v8s", 11.2, 75.2},
		{"YOLO-v8m", 25.9, 79.3},
		{"Swin-Transformer-T", 28.3, 81.5},
		{"ConvNeXt-T", 28.6, 82.1},
	}
}

type epochData struct {
	epoch   int
	trLoss  float64
	valLoss float64
	trAcc   float64
	valAcc  float64
	lr      float64
}

func getEpochData() []epochData {
	return []epochData{
		{1, 2.302, 2.298, 12.5, 15.2, 0.001},
		{2, 1.956, 1.845, 28.3, 35.1, 0.001},
		{3, 1.542, 1.423, 42.7, 48.9, 0.001},
		{5, 0.892, 0.934, 65.3, 62.1, 0.001},
		{10, 0.345, 0.412, 87.2, 84.5, 0.0005},
		{20, 0.156, 0.198, 94.2, 91.3, 0.0001},
		{30, 0.089, 0.127, 96.5, 93.8, 0.00005},
		{40, 0.052, 0.098, 97.8, 94.9, 0.00001},
		{50, 0.031, 0.084, 98.6, 95.5, 0.000001},
	}
}

type hyperparameterData struct {
	exp    int
	lr     float64
	batch  int
	wd     float64
	warmup int
	acc    float64
	time   float64
}

func getHyperparameterData() []hyperparameterData {
	return []hyperparameterData{
		{1, 0.001, 32, 0.0001, 500, 91.2, 48.0},
		{2, 0.001, 64, 0.0001, 500, 92.1, 45.0},
		{3, 0.001, 128, 0.0001, 500, 91.8, 42.0},
		{4, 0.002, 64, 0.0001, 1000, 93.5, 46.0},
		{5, 0.002, 64, 0.00005, 1000, 93.2, 47.0},
		{6, 0.0005, 64, 0.0001, 500, 92.8, 52.0},
		{7, 0.002, 128, 0.0001, 1000, 94.1, 44.0},
		{
			8,
			0.002,
			128,
			0.00005,
			1500,
			94.8,
			48.5,
		},
	}
}
