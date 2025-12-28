// Package elements provides PresentationML element types.
//
//nolint:revive // This file contains many public types for OOXML transition elements.
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ===========================================================================
// Transition Direction and Orientation Enums
// ===========================================================================

// TransitionDirection represents direction values for transitions.
type TransitionDirection string

const (
	// TransitionDirectionDown is downward direction.
	TransitionDirectionDown TransitionDirection = "d"
	// TransitionDirectionLeft is leftward direction.
	TransitionDirectionLeft TransitionDirection = "l"
	// TransitionDirectionRight is rightward direction.
	TransitionDirectionRight TransitionDirection = "r"
	// TransitionDirectionUp is upward direction.
	TransitionDirectionUp TransitionDirection = "u"
	// TransitionDirectionLeftUp is left-up diagonal direction.
	TransitionDirectionLeftUp TransitionDirection = "lu"
	// TransitionDirectionLeftDown is left-down diagonal direction.
	TransitionDirectionLeftDown TransitionDirection = "ld"
	// TransitionDirectionRightUp is right-up diagonal direction.
	TransitionDirectionRightUp TransitionDirection = "ru"
	// TransitionDirectionRightDown is right-down diagonal direction.
	TransitionDirectionRightDown TransitionDirection = "rd"
	// TransitionDirectionHorizontal is horizontal direction.
	TransitionDirectionHorizontal TransitionDirection = "h"
	// TransitionDirectionVertical is vertical direction.
	TransitionDirectionVertical TransitionDirection = "v"
	// TransitionDirectionIn is inward direction (for zoom).
	TransitionDirectionIn TransitionDirection = "in"
	// TransitionDirectionOut is outward direction (for zoom).
	TransitionDirectionOut TransitionDirection = "out"
)

// TransitionOrientation represents orientation values for split transitions.
type TransitionOrientation string

const (
	// TransitionOrientationHorizontal is horizontal orientation.
	TransitionOrientationHorizontal TransitionOrientation = "horz"
	// TransitionOrientationVertical is vertical orientation.
	TransitionOrientationVertical TransitionOrientation = "vert"
)

// ===========================================================================
// Extended SlideTransition Methods
// ===========================================================================

// AdvanceTime returns the advance time in milliseconds.
// Returns 0 if not set.
func (tr *SlideTransition) AdvanceTime() int {
	attr, found := tr.GetAttribute("advTm", "")
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}

	return val
}

// SetAdvanceTime sets the advance time in milliseconds.
// Set to 0 or negative to remove.
func (tr *SlideTransition) SetAdvanceTime(
	milliseconds int,
) {
	if milliseconds <= 0 {
		tr.RemoveAttribute("advTm", "")

		return
	}
	tr.SetAttribute(openxml.NewAttribute(
		"",
		"advTm",
		"",
		strconv.Itoa(milliseconds),
	))
}

// clearTransitionEffect removes any existing transition effect child.
func (tr *SlideTransition) clearTransitionEffect() {
	// List of all transition effect element names
	effectNames := []string{
		"blinds", "checker", "circle", "comb", "cover", "cut",
		"diamond", "dissolve", "fade", "newsflash", "plus", "pull",
		"push", "random", "randomBar", "split", "strips", "wedge",
		"wheel", "wipe", "zoom",
	}
	for _, name := range effectNames {
		if elem := tr.GetElement(name, NamespacePresentationML); elem != nil {
			tr.RemoveChild(elem)
		}
	}
}

// SetBlinds sets blinds transition with direction (h or v).
func (tr *SlideTransition) SetBlinds(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"blinds",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetChecker sets checker transition with direction (h or v).
func (tr *SlideTransition) SetChecker(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"checker",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetCircle sets circle transition.
func (tr *SlideTransition) SetCircle() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"circle",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetComb sets comb transition with direction (h or v).
func (tr *SlideTransition) SetComb(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"comb",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetCover sets cover transition with direction.
func (tr *SlideTransition) SetCover(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"cover",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetCut sets cut transition with optional through black.
func (tr *SlideTransition) SetCut(thruBlk bool) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"cut",
		PrefixP,
	)
	if thruBlk {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"thruBlk",
				"",
				"1",
			),
		)
	}
	tr.AppendChild(effect)
}

// SetDiamond sets diamond transition.
func (tr *SlideTransition) SetDiamond() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"diamond",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetDissolve sets dissolve transition.
func (tr *SlideTransition) SetDissolve() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"dissolve",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetFade sets fade transition with optional through black.
func (tr *SlideTransition) SetFade(thruBlk bool) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"fade",
		PrefixP,
	)
	if thruBlk {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"thruBlk",
				"",
				"1",
			),
		)
	}
	tr.AppendChild(effect)
}

// SetNewsflash sets newsflash transition.
func (tr *SlideTransition) SetNewsflash() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"newsflash",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetPlus sets plus transition.
func (tr *SlideTransition) SetPlus() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"plus",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetPull sets pull transition with direction.
func (tr *SlideTransition) SetPull(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"pull",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetPush sets push transition with direction.
func (tr *SlideTransition) SetPush(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"push",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetRandom sets random transition.
func (tr *SlideTransition) SetRandom() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"random",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetRandomBar sets random bar transition with direction (h or v).
func (tr *SlideTransition) SetRandomBar(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"randomBar",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetSplit sets split transition with orientation and direction.
func (tr *SlideTransition) SetSplit(
	orient TransitionOrientation,
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"split",
		PrefixP,
	)
	if orient != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"orient",
				"",
				string(orient),
			),
		)
	}
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetStrips sets strips transition with direction.
func (tr *SlideTransition) SetStrips(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"strips",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetWedge sets wedge transition.
func (tr *SlideTransition) SetWedge() {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"wedge",
		PrefixP,
	)
	tr.AppendChild(effect)
}

// SetWheel sets wheel transition with number of spokes.
func (tr *SlideTransition) SetWheel(spokes int) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"wheel",
		PrefixP,
	)
	if spokes > 0 {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"spokes",
				"",
				strconv.Itoa(spokes),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetWipe sets wipe transition with direction.
func (tr *SlideTransition) SetWipe(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"wipe",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// SetZoom sets zoom transition with direction (in or out).
func (tr *SlideTransition) SetZoom(
	dir TransitionDirection,
) {
	tr.clearTransitionEffect()
	effect := openxml.NewLeafElement(
		NamespacePresentationML,
		"zoom",
		PrefixP,
	)
	if dir != "" {
		effect.SetAttribute(
			openxml.NewAttribute(
				"",
				"dir",
				"",
				string(dir),
			),
		)
	}
	tr.AppendChild(effect)
}

// ===========================================================================
// Transition Effect Elements (p:blinds, p:checker, etc.)
// ===========================================================================

// TransitionBlinds represents the blinds transition effect (p:blinds).
type TransitionBlinds struct {
	*openxml.LeafElementBase
}

// NewTransitionBlinds creates a new blinds transition element.
func NewTransitionBlinds() *TransitionBlinds {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"blinds",
		PrefixP,
	)

	return &TransitionBlinds{
		LeafElementBase: elem,
	}
}

// Direction returns the blinds direction.
func (t *TransitionBlinds) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionHorizontal
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the blinds direction.
func (t *TransitionBlinds) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionBlinds) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionBlinds{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionChecker represents the checker transition effect (p:checker).
type TransitionChecker struct {
	*openxml.LeafElementBase
}

// NewTransitionChecker creates a new checker transition element.
func NewTransitionChecker() *TransitionChecker {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"checker",
		PrefixP,
	)

	return &TransitionChecker{
		LeafElementBase: elem,
	}
}

// Direction returns the checker direction.
func (t *TransitionChecker) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionHorizontal
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the checker direction.
func (t *TransitionChecker) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionChecker) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionChecker{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionCircle represents the circle transition effect (p:circle).
type TransitionCircle struct {
	*openxml.LeafElementBase
}

// NewTransitionCircle creates a new circle transition element.
func NewTransitionCircle() *TransitionCircle {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"circle",
		PrefixP,
	)

	return &TransitionCircle{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy.
func (t *TransitionCircle) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionCircle{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionComb represents the comb transition effect (p:comb).
type TransitionComb struct {
	*openxml.LeafElementBase
}

// NewTransitionComb creates a new comb transition element.
func NewTransitionComb() *TransitionComb {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"comb",
		PrefixP,
	)

	return &TransitionComb{LeafElementBase: elem}
}

// Direction returns the comb direction.
func (t *TransitionComb) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionHorizontal
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the comb direction.
func (t *TransitionComb) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionComb) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionComb{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionCover represents the cover transition effect (p:cover).
type TransitionCover struct {
	*openxml.LeafElementBase
}

// NewTransitionCover creates a new cover transition element.
func NewTransitionCover() *TransitionCover {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"cover",
		PrefixP,
	)

	return &TransitionCover{LeafElementBase: elem}
}

// Direction returns the cover direction.
func (t *TransitionCover) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionLeft
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the cover direction.
func (t *TransitionCover) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionCover) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionCover{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionCut represents the cut transition effect (p:cut).
type TransitionCut struct {
	*openxml.LeafElementBase
}

// NewTransitionCut creates a new cut transition element.
func NewTransitionCut() *TransitionCut {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"cut",
		PrefixP,
	)

	return &TransitionCut{LeafElementBase: elem}
}

// ThroughBlack returns whether the cut goes through black.
func (t *TransitionCut) ThroughBlack() bool {
	attr, found := t.GetAttribute("thruBlk", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrValueTrue
}

// SetThroughBlack sets whether the cut goes through black.
func (t *TransitionCut) SetThroughBlack(
	thruBlk bool,
) {
	if thruBlk {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"thruBlk",
				"",
				"1",
			),
		)
	} else {
		t.RemoveAttribute("thruBlk", "")
	}
}

// Clone creates a deep copy.
func (t *TransitionCut) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionCut{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionDiamond represents the diamond transition effect (p:diamond).
type TransitionDiamond struct {
	*openxml.LeafElementBase
}

// NewTransitionDiamond creates a new diamond transition element.
func NewTransitionDiamond() *TransitionDiamond {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"diamond",
		PrefixP,
	)

	return &TransitionDiamond{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy.
func (t *TransitionDiamond) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionDiamond{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionDissolve represents the dissolve transition effect (p:dissolve).
type TransitionDissolve struct {
	*openxml.LeafElementBase
}

// NewTransitionDissolve creates a new dissolve transition element.
func NewTransitionDissolve() *TransitionDissolve {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"dissolve",
		PrefixP,
	)

	return &TransitionDissolve{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy.
func (t *TransitionDissolve) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionDissolve{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionFade represents the fade transition effect (p:fade).
type TransitionFade struct {
	*openxml.LeafElementBase
}

// NewTransitionFade creates a new fade transition element.
func NewTransitionFade() *TransitionFade {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"fade",
		PrefixP,
	)

	return &TransitionFade{LeafElementBase: elem}
}

// ThroughBlack returns whether the fade goes through black.
func (t *TransitionFade) ThroughBlack() bool {
	attr, found := t.GetAttribute("thruBlk", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrValueTrue
}

// SetThroughBlack sets whether the fade goes through black.
func (t *TransitionFade) SetThroughBlack(
	thruBlk bool,
) {
	if thruBlk {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"thruBlk",
				"",
				"1",
			),
		)
	} else {
		t.RemoveAttribute("thruBlk", "")
	}
}

// Clone creates a deep copy.
func (t *TransitionFade) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionFade{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionNewsflash represents the newsflash transition effect (p:newsflash).
type TransitionNewsflash struct {
	*openxml.LeafElementBase
}

// NewTransitionNewsflash creates a new newsflash transition element.
func NewTransitionNewsflash() *TransitionNewsflash {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"newsflash",
		PrefixP,
	)

	return &TransitionNewsflash{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy.
func (t *TransitionNewsflash) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionNewsflash{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionPlus represents the plus transition effect (p:plus).
type TransitionPlus struct {
	*openxml.LeafElementBase
}

// NewTransitionPlus creates a new plus transition element.
func NewTransitionPlus() *TransitionPlus {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"plus",
		PrefixP,
	)

	return &TransitionPlus{LeafElementBase: elem}
}

// Clone creates a deep copy.
func (t *TransitionPlus) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionPlus{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionPull represents the pull transition effect (p:pull).
type TransitionPull struct {
	*openxml.LeafElementBase
}

// NewTransitionPull creates a new pull transition element.
func NewTransitionPull() *TransitionPull {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"pull",
		PrefixP,
	)

	return &TransitionPull{LeafElementBase: elem}
}

// Direction returns the pull direction.
func (t *TransitionPull) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionLeft
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the pull direction.
func (t *TransitionPull) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionPull) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionPull{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionPush represents the push transition effect (p:push).
type TransitionPush struct {
	*openxml.LeafElementBase
}

// NewTransitionPush creates a new push transition element.
func NewTransitionPush() *TransitionPush {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"push",
		PrefixP,
	)

	return &TransitionPush{LeafElementBase: elem}
}

// Direction returns the push direction.
func (t *TransitionPush) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionLeft
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the push direction.
func (t *TransitionPush) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionPush) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionPush{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionRandom represents the random transition effect (p:random).
type TransitionRandom struct {
	*openxml.LeafElementBase
}

// NewTransitionRandom creates a new random transition element.
func NewTransitionRandom() *TransitionRandom {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"random",
		PrefixP,
	)

	return &TransitionRandom{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy.
func (t *TransitionRandom) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionRandom{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionRandomBar represents the random bar transition effect (p:randomBar).
type TransitionRandomBar struct {
	*openxml.LeafElementBase
}

// NewTransitionRandomBar creates a new random bar transition element.
func NewTransitionRandomBar() *TransitionRandomBar {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"randomBar",
		PrefixP,
	)

	return &TransitionRandomBar{
		LeafElementBase: elem,
	}
}

// Direction returns the random bar direction.
func (t *TransitionRandomBar) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionHorizontal
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the random bar direction.
func (t *TransitionRandomBar) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionRandomBar) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionRandomBar{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionSplit represents the split transition effect (p:split).
type TransitionSplit struct {
	*openxml.LeafElementBase
}

// NewTransitionSplit creates a new split transition element.
func NewTransitionSplit() *TransitionSplit {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"split",
		PrefixP,
	)

	return &TransitionSplit{LeafElementBase: elem}
}

// Orientation returns the split orientation.
func (t *TransitionSplit) Orientation() TransitionOrientation {
	attr, found := t.GetAttribute("orient", "")
	if !found {
		return TransitionOrientationHorizontal
	}

	return TransitionOrientation(attr.Value())
}

// SetOrientation sets the split orientation.
func (t *TransitionSplit) SetOrientation(
	orient TransitionOrientation,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"orient",
			"",
			string(orient),
		),
	)
}

// Direction returns the split direction.
func (t *TransitionSplit) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionOut
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the split direction.
func (t *TransitionSplit) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionSplit) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionSplit{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionStrips represents the strips transition effect (p:strips).
type TransitionStrips struct {
	*openxml.LeafElementBase
}

// NewTransitionStrips creates a new strips transition element.
func NewTransitionStrips() *TransitionStrips {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"strips",
		PrefixP,
	)

	return &TransitionStrips{
		LeafElementBase: elem,
	}
}

// Direction returns the strips direction.
func (t *TransitionStrips) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionLeftDown
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the strips direction.
func (t *TransitionStrips) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionStrips) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionStrips{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionWedge represents the wedge transition effect (p:wedge).
type TransitionWedge struct {
	*openxml.LeafElementBase
}

// NewTransitionWedge creates a new wedge transition element.
func NewTransitionWedge() *TransitionWedge {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"wedge",
		PrefixP,
	)

	return &TransitionWedge{LeafElementBase: elem}
}

// Clone creates a deep copy.
func (t *TransitionWedge) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionWedge{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionWheel represents the wheel transition effect (p:wheel).
type TransitionWheel struct {
	*openxml.LeafElementBase
}

// NewTransitionWheel creates a new wheel transition element.
func NewTransitionWheel() *TransitionWheel {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"wheel",
		PrefixP,
	)

	return &TransitionWheel{LeafElementBase: elem}
}

// Spokes returns the number of spokes.
func (t *TransitionWheel) Spokes() int {
	attr, found := t.GetAttribute("spokes", "")
	if !found {
		return 4
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 4
	}

	return val
}

// SetSpokes sets the number of spokes.
func (t *TransitionWheel) SetSpokes(spokes int) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"spokes",
			"",
			strconv.Itoa(spokes),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionWheel) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionWheel{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionWipe represents the wipe transition effect (p:wipe).
type TransitionWipe struct {
	*openxml.LeafElementBase
}

// NewTransitionWipe creates a new wipe transition element.
func NewTransitionWipe() *TransitionWipe {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"wipe",
		PrefixP,
	)

	return &TransitionWipe{LeafElementBase: elem}
}

// Direction returns the wipe direction.
func (t *TransitionWipe) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionLeft
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the wipe direction.
func (t *TransitionWipe) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionWipe) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionWipe{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TransitionZoom represents the zoom transition effect (p:zoom).
type TransitionZoom struct {
	*openxml.LeafElementBase
}

// NewTransitionZoom creates a new zoom transition element.
func NewTransitionZoom() *TransitionZoom {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"zoom",
		PrefixP,
	)

	return &TransitionZoom{LeafElementBase: elem}
}

// Direction returns the zoom direction.
func (t *TransitionZoom) Direction() TransitionDirection {
	attr, found := t.GetAttribute("dir", "")
	if !found {
		return TransitionDirectionOut
	}

	return TransitionDirection(attr.Value())
}

// SetDirection sets the zoom direction.
func (t *TransitionZoom) SetDirection(
	dir TransitionDirection,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			string(dir),
		),
	)
}

// Clone creates a deep copy.
func (t *TransitionZoom) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TransitionZoom{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
