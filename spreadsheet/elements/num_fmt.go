package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Built-in number format IDs for Excel.
// These are the standard formats that Excel recognizes without requiring
// explicit format code definitions in the stylesheet.
const (
	// NumFmtGeneral is "General" format (0).
	NumFmtGeneral uint32 = 0
	// NumFmtInteger is "0" format (1).
	NumFmtInteger uint32 = 1
	// NumFmtDecimal2 is "0.00" format (2).
	NumFmtDecimal2 uint32 = 2
	// NumFmtThousands is "#,##0" format (3).
	NumFmtThousands uint32 = 3
	// NumFmtThousandsDecimal2 is "#,##0.00" format (4).
	NumFmtThousandsDecimal2 uint32 = 4
	// NumFmtPercent is "0%" format (9).
	NumFmtPercent uint32 = 9
	// NumFmtPercentDecimal2 is "0.00%" format (10).
	NumFmtPercentDecimal2 uint32 = 10
	// NumFmtScientific2 is "0.00E+00" format (11).
	NumFmtScientific2 uint32 = 11
	// NumFmtFraction1 is "# ?/?" format (12).
	NumFmtFraction1 uint32 = 12
	// NumFmtFraction2 is "# ??/??" format (13).
	NumFmtFraction2 uint32 = 13
	// NumFmtDateShort is "m/d/yy" format (14).
	NumFmtDateShort uint32 = 14
	// NumFmtDateLong is "d-mmm-yy" format (15).
	NumFmtDateLong uint32 = 15
	// NumFmtDateDayMonth is "d-mmm" format (16).
	NumFmtDateDayMonth uint32 = 16
	// NumFmtDateMonthYear is "mmm-yy" format (17).
	NumFmtDateMonthYear uint32 = 17
	// NumFmtTimeHourMin is "h:mm AM/PM" format (18).
	NumFmtTimeHourMin uint32 = 18
	// NumFmtTimeHourMinSec is "h:mm:ss AM/PM" format (19).
	NumFmtTimeHourMinSec uint32 = 19
	// NumFmtTime24HourMin is "h:mm" format (20).
	NumFmtTime24HourMin uint32 = 20
	// NumFmtTime24HourMinSec is "h:mm:ss" format (21).
	NumFmtTime24HourMinSec uint32 = 21
	// NumFmtDateTime is "m/d/yy h:mm" format (22).
	NumFmtDateTime uint32 = 22
	// NumFmtAccountingNoSymbol is "#,##0_);(#,##0)" format (37).
	NumFmtAccountingNoSymbol uint32 = 37
	// NumFmtAccountingNoSymbolRed is "#,##0_);[Red](#,##0)" format (38).
	NumFmtAccountingNoSymbolRed uint32 = 38
	// NumFmtAccountingNoSymbolDecimal2 is "#,##0.00_);(#,##0.00)" format (39).
	NumFmtAccountingNoSymbolDecimal2 uint32 = 39
	// NumFmtAccountingNoSymbolDecimal2Red is "#,##0.00_);[Red](#,##0.00)"
	// format (40).
	NumFmtAccountingNoSymbolDecimal2Red uint32 = 40
	// NumFmtTimeMinSec is "mm:ss" format (45).
	NumFmtTimeMinSec uint32 = 45
	// NumFmtTimeElapsed is "[h]:mm:ss" format (46).
	NumFmtTimeElapsed uint32 = 46
	// NumFmtTimeMinSecFraction is "mm:ss.0" format (47).
	NumFmtTimeMinSecFraction uint32 = 47
	// NumFmtScientific1 is "##0.0E+0" format (48).
	NumFmtScientific1 uint32 = 48
	// NumFmtText is "@" format (49).
	NumFmtText uint32 = 49
)

// BuiltInNumFmtCodes maps built-in number format IDs to their format codes.
var BuiltInNumFmtCodes = map[uint32]string{
	0:  "General",
	1:  "0",
	2:  "0.00",
	3:  "#,##0",
	4:  "#,##0.00",
	9:  "0%",
	10: "0.00%",
	11: "0.00E+00",
	12: "# ?/?",
	13: "# ??/??",
	14: "m/d/yy",
	15: "d-mmm-yy",
	16: "d-mmm",
	17: "mmm-yy",
	18: "h:mm AM/PM",
	19: "h:mm:ss AM/PM",
	20: "h:mm",
	21: "h:mm:ss",
	22: "m/d/yy h:mm",
	37: "#,##0_);(#,##0)",
	38: "#,##0_);[Red](#,##0)",
	39: "#,##0.00_);(#,##0.00)",
	40: "#,##0.00_);[Red](#,##0.00)",
	45: "mm:ss",
	46: "[h]:mm:ss",
	47: "mm:ss.0",
	48: "##0.0E+0",
	49: "@",
}

// IsBuiltInNumFmt returns true if the given format ID is a built-in format.
func IsBuiltInNumFmt(numFmtId uint32) bool {
	_, ok := BuiltInNumFmtCodes[numFmtId]

	return ok
}

// NumFmts represents the number formats container element (x:numFmts).
type NumFmts struct {
	*openxml.CompositeElementBase
}

// NewNumFmts creates a new NumFmts element.
func NewNumFmts() *NumFmts {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"numFmts",
		PrefixDefault,
	)

	return &NumFmts{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (nf *NumFmts) Count() uint32 {
	attr, found := nf.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetCount sets the count attribute.
func (nf *NumFmts) SetCount(count uint32) {
	nf.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// NumFmts returns an iterator over all NumFmt elements.
func (nf *NumFmts) NumFmts() iter.Seq[*NumFmt] {
	return func(yield func(*NumFmt) bool) {
		for child := range nf.Children() {
			if child.LocalName() != "numFmt" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var numFmt *NumFmt
			if n, ok := child.(*NumFmt); ok {
				numFmt = n
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				numFmt = &NumFmt{LeafElementBase: leaf}
			}
			if numFmt != nil && !yield(numFmt) {
				return
			}
		}
	}
}

// GetNumFmt returns the number format with the given ID, or nil if not found.
func (nf *NumFmts) GetNumFmt(
	numFmtId uint32,
) *NumFmt {
	for numFmt := range nf.NumFmts() {
		if numFmt.NumFmtId() == numFmtId {
			return numFmt
		}
	}

	return nil
}

// AddNumFmt adds a new number format with the given ID and format code.
func (nf *NumFmts) AddNumFmt(
	numFmtId uint32,
	formatCode string,
) *NumFmt {
	numFmt := NewNumFmt()
	numFmt.SetNumFmtId(numFmtId)
	numFmt.SetFormatCode(formatCode)
	nf.AppendChild(numFmt)
	nf.SetCount(nf.Count() + 1)

	return numFmt
}

// ItemCount returns the actual number of NumFmt children.
func (nf *NumFmts) ItemCount() int {
	count := 0
	for range nf.NumFmts() {
		count++
	}

	return count
}

// Clone creates a deep copy of this NumFmts element.
func (nf *NumFmts) Clone() openxml.Element {
	cloned := nf.CompositeElementBase.Clone()

	return &NumFmts{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this NumFmts element.
func (nf *NumFmts) CloneNode(
	deep bool,
) openxml.Element {
	cloned := nf.CompositeElementBase.CloneNode(
		deep,
	)

	return &NumFmts{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NumFmt represents a number format element (x:numFmt).
type NumFmt struct {
	*openxml.LeafElementBase
}

// NewNumFmt creates a new NumFmt element.
func NewNumFmt() *NumFmt {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"numFmt",
		PrefixDefault,
	)

	return &NumFmt{LeafElementBase: elem}
}

// NumFmtId returns the number format ID.
func (n *NumFmt) NumFmtId() uint32 {
	attr, found := n.GetAttribute("numFmtId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetNumFmtId sets the number format ID.
func (n *NumFmt) SetNumFmtId(id uint32) {
	n.SetAttribute(
		openxml.NewAttribute(
			"",
			"numFmtId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// FormatCode returns the format code string.
func (n *NumFmt) FormatCode() string {
	attr, found := n.GetAttribute(
		"formatCode",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFormatCode sets the format code string.
func (n *NumFmt) SetFormatCode(code string) {
	n.SetAttribute(
		openxml.NewAttribute(
			"",
			"formatCode",
			"",
			code,
		),
	)
}

// Clone creates a deep copy of this NumFmt element.
func (n *NumFmt) Clone() openxml.Element {
	cloned := n.LeafElementBase.Clone()

	return &NumFmt{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this NumFmt element.
func (n *NumFmt) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.LeafElementBase.CloneNode(deep)

	return &NumFmt{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
