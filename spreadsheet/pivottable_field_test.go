// This file contains tests for PivotField wrapper functionality.
package spreadsheet

import (
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// TestPivotFieldNameIndex tests the Name() and Index() methods.
func TestPivotFieldNameIndex(t *testing.T) {
	t.Run(
		"NameReturnsFieldName",
		func(t *testing.T) {
			pf := &PivotField{
				name:  "ProductName",
				index: 2,
			}

			if got := pf.Name(); got != "ProductName" {
				t.Errorf(
					"Name() = %q, want %q",
					got,
					"ProductName",
				)
			}
		},
	)

	t.Run(
		"IndexReturnsFieldIndex",
		func(t *testing.T) {
			pf := &PivotField{
				name:  "ProductName",
				index: 2,
			}

			if got := pf.Index(); got != 2 {
				t.Errorf(
					"Index() = %d, want %d",
					got,
					2,
				)
			}
		},
	)
}

// TestPivotFieldAxis tests the Axis() method for determining field axis.
func TestPivotFieldAxis(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *PivotField
		wantAxis AxisType
	}{
		{
			name: "RowField",
			setup: func() *PivotField {
				pt := &PivotTable{
					rowFields: []string{
						"Region",
						"Category",
					},
				}

				return &PivotField{
					pivot: pt,
					name:  "Region",
					index: 0,
				}
			},
			wantAxis: AxisTypeRow,
		},
		{
			name: "ColumnField",
			setup: func() *PivotField {
				pt := &PivotTable{
					colFields: []string{
						"Year",
						"Quarter",
					},
				}

				return &PivotField{
					pivot: pt,
					name:  "Year",
					index: 1,
				}
			},
			wantAxis: AxisTypeColumn,
		},
		{
			name: "PageField",
			setup: func() *PivotField {
				pt := &PivotTable{
					pageFields: []string{
						"Country",
					},
				}

				return &PivotField{
					pivot: pt,
					name:  "Country",
					index: 2,
				}
			},
			wantAxis: AxisTypePage,
		},
		{
			name: "DataField",
			setup: func() *PivotField {
				pt := &PivotTable{
					dataFields: []pivotDataField{
						{
							name:      "Sales",
							aggregate: AggregateSUM,
						},
					},
				}

				return &PivotField{
					pivot: pt,
					name:  "Sales",
					index: 3,
				}
			},
			wantAxis: AxisTypeData,
		},
		{
			name: "UnassignedField",
			setup: func() *PivotField {
				pt := &PivotTable{
					rowFields: []string{"Region"},
					colFields: []string{"Year"},
				}

				return &PivotField{
					pivot: pt,
					name:  "UnusedField",
					index: 4,
				}
			},
			wantAxis: AxisTypeNone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pf := tt.setup()
			if got := pf.Axis(); got != tt.wantAxis {
				t.Errorf(
					"Axis() = %v, want %v",
					got,
					tt.wantAxis,
				)
			}
		})
	}
}

// TestPivotFieldShowAll tests ShowAll getter and setter.
func TestPivotFieldShowAll(t *testing.T) {
	t.Run("GetDefaultValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		// Default should be true
		if got := pf.ShowAll(); !got {
			t.Errorf(
				"ShowAll() = %v, want %v",
				got,
				true,
			)
		}
	})

	t.Run("SetAndGetValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		pf.SetShowAll(false)
		if got := pf.ShowAll(); got {
			t.Errorf(
				"After SetShowAll(false), ShowAll() = %v, want %v",
				got,
				false,
			)
		}

		pf.SetShowAll(true)
		if got := pf.ShowAll(); !got {
			t.Errorf(
				"After SetShowAll(true), ShowAll() = %v, want %v",
				got,
				true,
			)
		}
	})

	t.Run(
		"NilElementReturnsDefault",
		func(t *testing.T) {
			pf := &PivotField{element: nil}

			if got := pf.ShowAll(); !got {
				t.Errorf(
					"ShowAll() with nil element = %v, want %v",
					got,
					true,
				)
			}
		},
	)
}

// TestPivotFieldCompact tests Compact getter and setter.
func TestPivotFieldCompact(t *testing.T) {
	t.Run("GetDefaultValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		// Default should be true
		if got := pf.Compact(); !got {
			t.Errorf(
				"Compact() = %v, want %v",
				got,
				true,
			)
		}
	})

	t.Run("SetAndGetValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		pf.SetCompact(false)
		if got := pf.Compact(); got {
			t.Errorf(
				"After SetCompact(false), Compact() = %v, want %v",
				got,
				false,
			)
		}

		pf.SetCompact(true)
		if got := pf.Compact(); !got {
			t.Errorf(
				"After SetCompact(true), Compact() = %v, want %v",
				got,
				true,
			)
		}
	})
}

// TestPivotFieldHideNewItems tests HideNewItems getter and setter.
func TestPivotFieldHideNewItems(t *testing.T) {
	t.Run("GetDefaultValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		// Default should be false
		if got := pf.HideNewItems(); got {
			t.Errorf(
				"HideNewItems() = %v, want %v",
				got,
				false,
			)
		}
	})

	t.Run("SetAndGetValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		pf.SetHideNewItems(true)
		if got := pf.HideNewItems(); !got {
			t.Errorf(
				"After SetHideNewItems(true), HideNewItems() = %v, want %v",
				got,
				true,
			)
		}

		pf.SetHideNewItems(false)
		if got := pf.HideNewItems(); got {
			t.Errorf(
				"After SetHideNewItems(false), HideNewItems() = %v, want %v",
				got,
				false,
			)
		}
	})
}

// TestPivotFieldShowDropDowns tests ShowDropDowns getter and setter.
func TestPivotFieldShowDropDowns(t *testing.T) {
	t.Run("GetDefaultValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		// Default should be true
		if got := pf.ShowDropDowns(); !got {
			t.Errorf(
				"ShowDropDowns() = %v, want %v",
				got,
				true,
			)
		}
	})

	t.Run("SetAndGetValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		pf.SetShowDropDowns(false)
		if got := pf.ShowDropDowns(); got {
			t.Errorf(
				"After SetShowDropDowns(false), ShowDropDowns() = %v, want %v",
				got,
				false,
			)
		}

		pf.SetShowDropDowns(true)
		if got := pf.ShowDropDowns(); !got {
			t.Errorf(
				"After SetShowDropDowns(true), ShowDropDowns() = %v, want %v",
				got,
				true,
			)
		}
	})
}

// TestPivotFieldDefaultSubtotal tests DefaultSubtotal getter and setter.
func TestPivotFieldDefaultSubtotal(t *testing.T) {
	t.Run("GetDefaultValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		// Default should be true
		if got := pf.DefaultSubtotal(); !got {
			t.Errorf(
				"DefaultSubtotal() = %v, want %v",
				got,
				true,
			)
		}
	})

	t.Run("SetAndGetValue", func(t *testing.T) {
		elem := elements.NewPivotField()
		pf := &PivotField{element: elem}

		pf.SetDefaultSubtotal(false)
		if got := pf.DefaultSubtotal(); got {
			t.Errorf(
				"After SetDefaultSubtotal(false), DefaultSubtotal() = %v, want %v",
				got,
				false,
			)
		}

		pf.SetDefaultSubtotal(true)
		if got := pf.DefaultSubtotal(); !got {
			t.Errorf(
				"After SetDefaultSubtotal(true), DefaultSubtotal() = %v, want %v",
				got,
				true,
			)
		}
	})
}

// TestPivotFieldDataType tests DataType inference.
func TestPivotFieldDataType(t *testing.T) {
	t.Run(
		"NoCacheReturnsString",
		func(t *testing.T) {
			pt := &PivotTable{}
			pf := &PivotField{
				pivot: pt,
				index: 0,
			}

			// Without cache, should default to string
			if got := pf.DataType(); got != FieldDataTypeString {
				t.Errorf(
					"DataType() without cache = %v, want %v",
					got,
					FieldDataTypeString,
				)
			}
		},
	)
}

// TestPivotFieldItemCount tests ItemCount counting.
func TestPivotFieldItemCount(t *testing.T) {
	t.Run(
		"NoCacheReturnsZero",
		func(t *testing.T) {
			pt := &PivotTable{}
			pf := &PivotField{
				pivot: pt,
				index: 0,
			}

			// Without cache, should return 0
			if got := pf.ItemCount(); got != 0 {
				t.Errorf(
					"ItemCount() without cache = %d, want %d",
					got,
					0,
				)
			}
		},
	)
}

// TestPivotTableFieldLookup tests Field() and Fields() methods on PivotTable.
func TestPivotTableFieldLookup(t *testing.T) {
	t.Run(
		"FieldReturnsNilForEmptyName",
		func(t *testing.T) {
			pt := &PivotTable{}

			if got := pt.Field(""); got != nil {
				t.Errorf(
					"Field(\"\") = %v, want nil",
					got,
				)
			}
		},
	)

	t.Run(
		"FieldReturnsNilWhenNoPivotDef",
		func(t *testing.T) {
			pt := &PivotTable{}

			if got := pt.Field("SomeField"); got != nil {
				t.Errorf(
					"Field(\"SomeField\") without pivot def = %v, want nil",
					got,
				)
			}
		},
	)

	t.Run(
		"FieldsReturnsNilWhenNoPivotDef",
		func(t *testing.T) {
			pt := &PivotTable{}

			if got := pt.Fields(); got != nil {
				t.Errorf(
					"Fields() without pivot def = %v, want nil",
					got,
				)
			}
		},
	)
}

// TestPivotFieldSettersWithNilElement tests that setters handle nil elements gracefully.
func TestPivotFieldSettersWithNilElement(
	_ *testing.T,
) {
	pf := &PivotField{element: nil}

	// These should not panic
	pf.SetShowAll(false)
	pf.SetCompact(false)
	pf.SetHideNewItems(true)
	pf.SetShowDropDowns(false)
	pf.SetDefaultSubtotal(false)
}
