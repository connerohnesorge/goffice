//nolint:revive // receiver names and flag-parameters are acceptable here
package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// LatentStyles represents latent style behavior settings (w:latentStyles).
type LatentStyles struct {
	*openxml.CompositeElementBase
}

// NewLatentStyles creates a new LatentStyles element.
func NewLatentStyles() *LatentStyles {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"latentStyles",
		PrefixW,
	)

	return &LatentStyles{
		CompositeElementBase: elem,
	}
}

// DefLockedState returns the default locked state for latent styles.
func (ls *LatentStyles) DefLockedState() bool {
	attr, found := ls.GetAttribute(
		"defLockedState",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetDefLockedState sets the default locked state for latent styles.
func (ls *LatentStyles) SetDefLockedState(
	b bool,
) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"defLockedState",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("defLockedState", NamespaceWML)
	}
}

// DefSemiHidden returns the default semi-hidden state for latent styles.
func (ls *LatentStyles) DefSemiHidden() bool {
	attr, found := ls.GetAttribute(
		"defSemiHidden",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetDefSemiHidden sets the default semi-hidden state for latent styles.
func (ls *LatentStyles) SetDefSemiHidden(b bool) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"defSemiHidden",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("defSemiHidden", NamespaceWML)
	}
}

// DefUnhideWhenUsed returns the default unhide-when-used state for latent styles.
func (ls *LatentStyles) DefUnhideWhenUsed() bool {
	attr, found := ls.GetAttribute(
		"defUnhideWhenUsed",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetDefUnhideWhenUsed sets the default unhide-when-used state for latent styles.
func (ls *LatentStyles) SetDefUnhideWhenUsed(
	b bool,
) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"defUnhideWhenUsed",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("defUnhideWhenUsed", NamespaceWML)
	}
}

// DefQFormat returns the default quick format state for latent styles.
func (ls *LatentStyles) DefQFormat() bool {
	attr, found := ls.GetAttribute(
		"defQFormat",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetDefQFormat sets the default quick format state for latent styles.
func (ls *LatentStyles) SetDefQFormat(b bool) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"defQFormat",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("defQFormat", NamespaceWML)
	}
}

// DefUIPriority returns the default UI priority for latent styles.
func (ls *LatentStyles) DefUIPriority() int {
	attr, found := ls.GetAttribute(
		"defUIPriority",
		NamespaceWML,
	)
	if !found {
		return 99
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 99
	}

	return val
}

// SetDefUIPriority sets the default UI priority for latent styles.
func (ls *LatentStyles) SetDefUIPriority(
	priority int,
) {
	ls.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"defUIPriority",
			PrefixW,
			strconv.Itoa(priority),
		),
	)
}

// Count returns the count of latent styles.
func (ls *LatentStyles) Count() int {
	attr, found := ls.GetAttribute(
		"count",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}

	return val
}

// SetCount sets the count of latent styles.
func (ls *LatentStyles) SetCount(count int) {
	ls.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"count",
			PrefixW,
			strconv.Itoa(count),
		),
	)
}

// LatentStyleExceptions returns an iterator over all latent style exceptions.
func (ls *LatentStyles) LatentStyleExceptions() iter.Seq[*LatentStyleException] {
	return func(yield func(*LatentStyleException) bool) {
		for child := range ls.Children() {
			if child.LocalName() == "lsdException" &&
				child.NamespaceURI() == NamespaceWML {
				var lse *LatentStyleException
				switch v := child.(type) {
				case *LatentStyleException:
					lse = v
				case *openxml.CompositeElementBase:
					lse = &LatentStyleException{CompositeElementBase: v}
				}
				if lse != nil && !yield(lse) {
					return
				}
			}
		}
	}
}

// GetException returns the exception for the specified style name, or nil if not found.
func (ls *LatentStyles) GetException(
	name string,
) *LatentStyleException {
	for ex := range ls.LatentStyleExceptions() {
		if ex.Name() == name {
			return ex
		}
	}

	return nil
}

// AddException adds a new latent style exception.
func (ls *LatentStyles) AddException(
	name string,
) *LatentStyleException {
	lse := NewLatentStyleException(name)
	ls.AppendChild(lse)

	return lse
}

// Clone creates a deep copy of this LatentStyles element.
func (ls *LatentStyles) Clone() openxml.Element {
	cloned := ls.CompositeElementBase.Clone()

	return &LatentStyles{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this LatentStyles element.
func (ls *LatentStyles) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ls.CompositeElementBase.CloneNode(
		deep,
	)

	return &LatentStyles{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// LatentStyleException represents an exception to default latent style settings (w:lsdException).
type LatentStyleException struct {
	*openxml.CompositeElementBase
}

// NewLatentStyleException creates a new LatentStyleException element.
func NewLatentStyleException(
	name string,
) *LatentStyleException {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"lsdException",
		PrefixW,
	)
	lse := &LatentStyleException{
		CompositeElementBase: elem,
	}
	lse.SetName(name)

	return lse
}

// Name returns the style name for this exception.
func (ls *LatentStyleException) Name() string {
	attr, found := ls.GetAttribute(
		"name",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the style name for this exception.
func (ls *LatentStyleException) SetName(
	name string,
) {
	ls.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"name",
			PrefixW,
			name,
		),
	)
}

// Locked returns whether this style is locked.
func (ls *LatentStyleException) Locked() bool {
	attr, found := ls.GetAttribute(
		"locked",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetLocked sets whether this style is locked.
func (ls *LatentStyleException) SetLocked(
	b bool,
) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"locked",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("locked", NamespaceWML)
	}
}

// SemiHidden returns whether this style is semi-hidden.
func (ls *LatentStyleException) SemiHidden() bool {
	attr, found := ls.GetAttribute(
		"semiHidden",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetSemiHidden sets whether this style is semi-hidden.
func (ls *LatentStyleException) SetSemiHidden(
	b bool,
) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"semiHidden",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("semiHidden", NamespaceWML)
	}
}

// UnhideWhenUsed returns whether this style becomes visible when used.
func (ls *LatentStyleException) UnhideWhenUsed() bool {
	attr, found := ls.GetAttribute(
		"unhideWhenUsed",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetUnhideWhenUsed sets whether this style becomes visible when used.
func (ls *LatentStyleException) SetUnhideWhenUsed(
	b bool,
) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"unhideWhenUsed",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("unhideWhenUsed", NamespaceWML)
	}
}

// QFormat returns whether this style appears in quick styles.
func (ls *LatentStyleException) QFormat() bool {
	attr, found := ls.GetAttribute(
		"qFormat",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == attrValueTrue ||
		val == "on"
}

// SetQFormat sets whether this style appears in quick styles.
func (ls *LatentStyleException) SetQFormat(
	b bool,
) {
	if b {
		ls.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"qFormat",
				PrefixW,
				"1",
			),
		)
	} else {
		ls.RemoveAttribute("qFormat", NamespaceWML)
	}
}

// UIPriority returns the UI priority for this style.
func (ls *LatentStyleException) UIPriority() int {
	attr, found := ls.GetAttribute(
		"uiPriority",
		NamespaceWML,
	)
	if !found {
		return -1 // indicates not set
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return -1
	}

	return val
}

// SetUIPriority sets the UI priority for this style.
func (ls *LatentStyleException) SetUIPriority(
	priority int,
) {
	if priority < 0 {
		ls.RemoveAttribute(
			"uiPriority",
			NamespaceWML,
		)

		return
	}
	ls.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"uiPriority",
			PrefixW,
			strconv.Itoa(priority),
		),
	)
}

// Clone creates a deep copy of this LatentStyleException element.
func (ls *LatentStyleException) Clone() openxml.Element {
	cloned := ls.CompositeElementBase.Clone()

	return &LatentStyleException{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
