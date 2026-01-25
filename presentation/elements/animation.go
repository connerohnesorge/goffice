//
//nolint:revive // This file contains many public types for OOXML animation elements.
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ===========================================================================
// Animation Enums
// ===========================================================================

// TimeNodeType represents the type of time node.
type TimeNodeType string

const (
	// TimeNodeTypeClickEffect is a click-triggered effect.
	TimeNodeTypeClickEffect TimeNodeType = "clickEffect"
	// TimeNodeTypeWithEffect is an effect that runs with the previous.
	TimeNodeTypeWithEffect TimeNodeType = "withEffect"
	// TimeNodeTypeAfterEffect is an effect that runs after the previous.
	TimeNodeTypeAfterEffect TimeNodeType = "afterEffect"
	// TimeNodeTypeMainSeq is the main sequence of animations.
	TimeNodeTypeMainSeq TimeNodeType = "mainSeq"
	// TimeNodeTypeInteractiveSeq is an interactive sequence.
	TimeNodeTypeInteractiveSeq TimeNodeType = "interactiveSeq"
)

// TimeNodeFill represents how a time node fills after playback.
type TimeNodeFill string

const (
	// TimeNodeFillRemove removes the effect after playback.
	TimeNodeFillRemove TimeNodeFill = "remove"
	// TimeNodeFillFreeze freezes the effect at end state.
	TimeNodeFillFreeze TimeNodeFill = "freeze"
	// TimeNodeFillHold holds the effect until next animation.
	TimeNodeFillHold TimeNodeFill = "hold"
	// TimeNodeFillTransition transitions to next effect.
	TimeNodeFillTransition TimeNodeFill = "transition"
)

// TimeNodeRestart represents when a time node can restart.
type TimeNodeRestart string

const (
	// TimeNodeRestartAlways can restart always.
	TimeNodeRestartAlways TimeNodeRestart = "always"
	// TimeNodeRestartWhenNotActive can restart when not active.
	TimeNodeRestartWhenNotActive TimeNodeRestart = "whenNotActive"
	// TimeNodeRestartNever cannot restart.
	TimeNodeRestartNever TimeNodeRestart = "never"
)

// AnimationBuildType represents how content is built in animation.
type AnimationBuildType string

const (
	// AnimationBuildByParagraph builds paragraph by paragraph.
	AnimationBuildByParagraph AnimationBuildType = "p"
	// AnimationBuildAllAtOnce builds all at once.
	AnimationBuildAllAtOnce AnimationBuildType = "allAtOnce"
)

// AnimateEffectTransition represents the effect transition type.
type AnimateEffectTransition string

const (
	// AnimateEffectTransitionIn is entrance effect.
	AnimateEffectTransitionIn AnimateEffectTransition = "in"
	// AnimateEffectTransitionOut is exit effect.
	AnimateEffectTransitionOut AnimateEffectTransition = "out"
	// AnimateEffectTransitionNone is no transition.
	AnimateEffectTransitionNone AnimateEffectTransition = "none"
)

// AnimateEffectFilter represents animation effect filter types.
type AnimateEffectFilter string

const (
	// FilterBlindsHorizontal is horizontal blinds effect.
	FilterBlindsHorizontal AnimateEffectFilter = "blinds(horizontal)"
	// FilterBlindsVertical is vertical blinds effect.
	FilterBlindsVertical AnimateEffectFilter = "blinds(vertical)"
	// FilterBoxIn is box in effect.
	FilterBoxIn AnimateEffectFilter = "box(in)"
	// FilterBoxOut is box out effect.
	FilterBoxOut AnimateEffectFilter = "box(out)"
	// FilterCheckerboardAcross is checkerboard across effect.
	FilterCheckerboardAcross AnimateEffectFilter = "checkerboard(across)"
	// FilterCheckerboardDown is checkerboard down effect.
	FilterCheckerboardDown AnimateEffectFilter = "checkerboard(down)"
	// FilterCircleIn is circle in effect.
	FilterCircleIn AnimateEffectFilter = "circle(in)"
	// FilterCircleOut is circle out effect.
	FilterCircleOut AnimateEffectFilter = "circle(out)"
	// FilterDiamondIn is diamond in effect.
	FilterDiamondIn AnimateEffectFilter = "diamond(in)"
	// FilterDiamondOut is diamond out effect.
	FilterDiamondOut AnimateEffectFilter = "diamond(out)"
	// FilterDissolve is dissolve effect.
	FilterDissolve AnimateEffectFilter = "dissolve"
	// FilterFade is fade effect.
	FilterFade AnimateEffectFilter = "fade"
	// FilterFlyFromBottom is fly from bottom effect.
	FilterFlyFromBottom AnimateEffectFilter = "fly(fromBottom)"
	// FilterFlyFromLeft is fly from left effect.
	FilterFlyFromLeft AnimateEffectFilter = "fly(fromLeft)"
	// FilterFlyFromRight is fly from right effect.
	FilterFlyFromRight AnimateEffectFilter = "fly(fromRight)"
	// FilterFlyFromTop is fly from top effect.
	FilterFlyFromTop AnimateEffectFilter = "fly(fromTop)"
	// FilterPeekFromBottom is peek from bottom effect.
	FilterPeekFromBottom AnimateEffectFilter = "peek(fromBottom)"
	// FilterPeekFromLeft is peek from left effect.
	FilterPeekFromLeft AnimateEffectFilter = "peek(fromLeft)"
	// FilterPeekFromRight is peek from right effect.
	FilterPeekFromRight AnimateEffectFilter = "peek(fromRight)"
	// FilterPeekFromTop is peek from top effect.
	FilterPeekFromTop AnimateEffectFilter = "peek(fromTop)"
	// FilterRandomBarsHorizontal is random bars horizontal effect.
	FilterRandomBarsHorizontal AnimateEffectFilter = "randomBars(horizontal)"
	// FilterRandomBarsVertical is random bars vertical effect.
	FilterRandomBarsVertical AnimateEffectFilter = "randomBars(vertical)"
	// FilterSplitHorizontalIn is split horizontal in effect.
	FilterSplitHorizontalIn AnimateEffectFilter = "split(horizontalIn)"
	// FilterSplitHorizontalOut is split horizontal out effect.
	FilterSplitHorizontalOut AnimateEffectFilter = "split(horizontalOut)"
	// FilterSplitVerticalIn is split vertical in effect.
	FilterSplitVerticalIn AnimateEffectFilter = "split(verticalIn)"
	// FilterSplitVerticalOut is split vertical out effect.
	FilterSplitVerticalOut AnimateEffectFilter = "split(verticalOut)"
	// FilterStripsDownLeft is strips down left effect.
	FilterStripsDownLeft AnimateEffectFilter = "strips(downLeft)"
	// FilterStripsDownRight is strips down right effect.
	FilterStripsDownRight AnimateEffectFilter = "strips(downRight)"
	// FilterStripsUpLeft is strips up left effect.
	FilterStripsUpLeft AnimateEffectFilter = "strips(upLeft)"
	// FilterStripsUpRight is strips up right effect.
	FilterStripsUpRight AnimateEffectFilter = "strips(upRight)"
	// FilterWedge is wedge effect.
	FilterWedge AnimateEffectFilter = "wedge"
	// FilterWheelClockwise1 is wheel clockwise 1 spoke effect.
	FilterWheelClockwise1 AnimateEffectFilter = "wheel(1)"
	// FilterWheelClockwise2 is wheel clockwise 2 spokes effect.
	FilterWheelClockwise2 AnimateEffectFilter = "wheel(2)"
	// FilterWheelClockwise3 is wheel clockwise 3 spokes effect.
	FilterWheelClockwise3 AnimateEffectFilter = "wheel(3)"
	// FilterWheelClockwise4 is wheel clockwise 4 spokes effect.
	FilterWheelClockwise4 AnimateEffectFilter = "wheel(4)"
	// FilterWheelClockwise8 is wheel clockwise 8 spokes effect.
	FilterWheelClockwise8 AnimateEffectFilter = "wheel(8)"
	// FilterWipeDown is wipe down effect.
	FilterWipeDown AnimateEffectFilter = "wipe(down)"
	// FilterWipeLeft is wipe left effect.
	FilterWipeLeft AnimateEffectFilter = "wipe(left)"
	// FilterWipeRight is wipe right effect.
	FilterWipeRight AnimateEffectFilter = "wipe(right)"
	// FilterWipeUp is wipe up effect.
	FilterWipeUp AnimateEffectFilter = "wipe(up)"
	// FilterZoomIn is zoom in effect.
	FilterZoomIn AnimateEffectFilter = "zoom(in)"
	// FilterZoomOut is zoom out effect.
	FilterZoomOut AnimateEffectFilter = "zoom(out)"
)

// ===========================================================================
// SlideTiming Extended Methods
// ===========================================================================

// TimeNodeList returns the time node list.
func (t *SlideTiming) TimeNodeList() *TimeNodeList {
	elem := t.GetElement(
		"tnLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if tnl, ok := elem.(*TimeNodeList); ok {
		return tnl
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &TimeNodeList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTimeNodeList returns or creates the time node list.
func (t *SlideTiming) GetOrCreateTimeNodeList() *TimeNodeList {
	tnl := t.TimeNodeList()
	if tnl != nil {
		return tnl
	}
	tnl = NewTimeNodeList()
	t.AppendChild(tnl)

	return tnl
}

// BuildList returns the build list.
func (t *SlideTiming) BuildList() *BuildList {
	elem := t.GetElement(
		"bldLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if bl, ok := elem.(*BuildList); ok {
		return bl
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &BuildList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateBuildList returns or creates the build list.
func (t *SlideTiming) GetOrCreateBuildList() *BuildList {
	bl := t.BuildList()
	if bl != nil {
		return bl
	}
	bl = NewBuildList()
	t.AppendChild(bl)

	return bl
}

// AddClickEffect adds a click-triggered animation to the slide.
// shapeID is the ID of the shape to animate.
// filter specifies the animation effect type.
// durationMs is the animation duration in milliseconds.
func (t *SlideTiming) AddClickEffect(
	shapeID int,
	filter AnimateEffectFilter,
	durationMs int,
) *CommonTimeNode {
	tnl := t.GetOrCreateTimeNodeList()

	// Get or create main sequence
	mainSeq := tnl.GetOrCreateMainSequence()

	// Create click effect sequence
	seq := NewSequenceTimeNode()
	seq.SetConcurrent(true)
	seq.SetNextAction("seek")

	// Create parallel time node for the effect
	par := NewParallelTimeNode()

	// Create common time node with click behavior
	ctn := NewCommonTimeNode()
	ctn.SetID(t.nextNodeID())
	ctn.SetPresetID(1)
	ctn.SetPresetClass("entr")
	ctn.SetPresetSubtype(0)
	ctn.SetFill(TimeNodeFillHold)
	ctn.SetNodeType(TimeNodeTypeClickEffect)
	ctn.SetDuration(durationMs)

	// Add start condition on click
	stCondLst := NewStartConditionList()
	cond := NewCondition()
	cond.SetEvent("onClick")
	cond.SetDelay(0)

	// Target reference
	tgtEl := NewTargetElement()
	spTgt := NewShapeTarget()
	spTgt.SetShapeID(shapeID)
	tgtEl.AppendChild(spTgt)
	cond.AppendChild(tgtEl)
	stCondLst.AppendChild(cond)
	ctn.AppendChild(stCondLst)

	// Add animation effect child node list
	childTnLst := NewChildTimeNodeList()
	effectPar := NewParallelTimeNode()
	effectCtn := NewCommonTimeNode()
	effectCtn.SetID(t.nextNodeID())
	effectCtn.SetDuration(durationMs)
	effectCtn.SetFill(TimeNodeFillHold)

	// Add effect itself
	effectChildTnLst := NewChildTimeNodeList()
	animEffect := NewAnimateEffect()
	animEffect.SetFilter(filter)
	animEffect.SetTransition(
		AnimateEffectTransitionIn,
	)
	animEffect.SetDuration(durationMs)
	animEffect.SetShapeTarget(shapeID)
	effectChildTnLst.AppendChild(animEffect)
	effectCtn.AppendChild(effectChildTnLst)

	effectPar.AppendChild(effectCtn)
	childTnLst.AppendChild(effectPar)
	ctn.AppendChild(childTnLst)

	par.AppendChild(ctn)
	seq.AppendChild(par)
	mainSeq.AppendChild(seq)

	return ctn
}

// AddFadeIn adds a fade in entrance animation.
func (t *SlideTiming) AddFadeIn(
	shapeID int,
	durationMs int,
) *CommonTimeNode {
	return t.AddClickEffect(
		shapeID,
		FilterFade,
		durationMs,
	)
}

// AddFlyIn adds a fly in entrance animation.
func (t *SlideTiming) AddFlyIn(
	shapeID int,
	direction string,
	durationMs int,
) *CommonTimeNode {
	var filter AnimateEffectFilter
	switch direction {
	case "left":
		filter = FilterFlyFromLeft
	case "right":
		filter = FilterFlyFromRight
	case "top":
		filter = FilterFlyFromTop
	case "bottom":
		filter = FilterFlyFromBottom
	default:
		filter = FilterFlyFromBottom

	}

	return t.AddClickEffect(
		shapeID,
		filter,
		durationMs,
	)
}

// AddWipeIn adds a wipe entrance animation.
func (t *SlideTiming) AddWipeIn(
	shapeID int,
	direction string,
	durationMs int,
) *CommonTimeNode {
	var filter AnimateEffectFilter
	switch direction {
	case "left":
		filter = FilterWipeLeft
	case "right":
		filter = FilterWipeRight
	case "up":
		filter = FilterWipeUp
	case "down":
		filter = FilterWipeDown
	default:
		filter = FilterWipeRight

	}

	return t.AddClickEffect(
		shapeID,
		filter,
		durationMs,
	)
}

// AddZoomIn adds a zoom in entrance animation.
func (t *SlideTiming) AddZoomIn(
	shapeID int,
	durationMs int,
) *CommonTimeNode {
	return t.AddClickEffect(
		shapeID,
		FilterZoomIn,
		durationMs,
	)
}

// AddDissolve adds a dissolve entrance animation.
func (t *SlideTiming) AddDissolve(
	shapeID int,
	durationMs int,
) *CommonTimeNode {
	return t.AddClickEffect(
		shapeID,
		FilterDissolve,
		durationMs,
	)
}

// nextNodeID generates the next unique node ID.
func (t *SlideTiming) nextNodeID() int {
	// Simple incrementing counter - in practice would track used IDs
	count := 1
	tnl := t.TimeNodeList()
	if tnl != nil {
		// Count existing nodes
		for range tnl.Children() {
			count++
		}
	}

	return count
}

// ===========================================================================
// TimeNodeList (p:tnLst)
// ===========================================================================

// TimeNodeList represents the time node list container (p:tnLst).
type TimeNodeList struct {
	*openxml.CompositeElementBase
}

// NewTimeNodeList creates a new TimeNodeList element.
func NewTimeNodeList() *TimeNodeList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"tnLst",
		PrefixP,
	)

	return &TimeNodeList{
		CompositeElementBase: elem,
	}
}

// GetOrCreateMainSequence returns or creates the main animation sequence.
func (tnl *TimeNodeList) GetOrCreateMainSequence() *ParallelTimeNode {
	// Look for existing main sequence
	for child := range tnl.Children() {
		if par, ok := child.(*ParallelTimeNode); ok {
			return par
		}
		if comp := wrapCompositeElement(child); comp != nil {
			if child.LocalName() == "par" {
				return &ParallelTimeNode{
					CompositeElementBase: comp,
				}
			}
		}
	}

	// Create main sequence structure
	par := NewParallelTimeNode()
	ctn := NewCommonTimeNode()
	ctn.SetID(1)
	ctn.SetDuration(0)
	ctn.SetRestart(TimeNodeRestartNever)
	ctn.SetNodeType(TimeNodeTypeMainSeq)
	par.AppendChild(ctn)
	tnl.AppendChild(par)

	return par
}

// Clone creates a deep copy.
func (tnl *TimeNodeList) Clone() openxml.Element {
	cloned := tnl.CompositeElementBase.Clone()

	return &TimeNodeList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// BuildList (p:bldLst)
// ===========================================================================

// BuildList represents the build list container (p:bldLst).
type BuildList struct {
	*openxml.CompositeElementBase
}

// NewBuildList creates a new BuildList element.
func NewBuildList() *BuildList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"bldLst",
		PrefixP,
	)

	return &BuildList{CompositeElementBase: elem}
}

// AddBuildParagraph adds a paragraph build effect.
func (bl *BuildList) AddBuildParagraph(
	shapeID int,
	buildType AnimationBuildType,
) *BuildParagraph {
	bp := NewBuildParagraph()
	bp.SetShapeID(shapeID)
	bp.SetBuildType(buildType)
	bl.AppendChild(bp)

	return bp
}

// Clone creates a deep copy.
func (bl *BuildList) Clone() openxml.Element {
	cloned := bl.CompositeElementBase.Clone()

	return &BuildList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// BuildParagraph (p:bldP)
// ===========================================================================

// BuildParagraph represents a paragraph build effect (p:bldP).
type BuildParagraph struct {
	*openxml.CompositeElementBase
}

// NewBuildParagraph creates a new BuildParagraph element.
func NewBuildParagraph() *BuildParagraph {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"bldP",
		PrefixP,
	)

	return &BuildParagraph{
		CompositeElementBase: elem,
	}
}

// ShapeID returns the target shape ID.
func (bp *BuildParagraph) ShapeID() int {
	attr, found := bp.GetAttribute("spId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetShapeID sets the target shape ID.
func (bp *BuildParagraph) SetShapeID(id int) {
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"spId",
			"",
			strconv.Itoa(id),
		),
	)
}

// BuildType returns the build type.
func (bp *BuildParagraph) BuildType() AnimationBuildType {
	attr, found := bp.GetAttribute("build", "")
	if !found {
		return AnimationBuildByParagraph
	}

	return AnimationBuildType(attr.Value())
}

// SetBuildType sets the build type.
func (bp *BuildParagraph) SetBuildType(
	buildType AnimationBuildType,
) {
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"build",
			"",
			string(buildType),
		),
	)
}

// Clone creates a deep copy.
func (bp *BuildParagraph) Clone() openxml.Element {
	cloned := bp.CompositeElementBase.Clone()

	return &BuildParagraph{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// CommonTimeNode (p:cTn)
// ===========================================================================

// CommonTimeNode represents common time node properties (p:cTn).
type CommonTimeNode struct {
	*openxml.CompositeElementBase
}

// NewCommonTimeNode creates a new CommonTimeNode element.
func NewCommonTimeNode() *CommonTimeNode {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cTn",
		PrefixP,
	)

	return &CommonTimeNode{
		CompositeElementBase: elem,
	}
}

// ID returns the node ID.
func (ctn *CommonTimeNode) ID() int {
	attr, found := ctn.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetID sets the node ID.
func (ctn *CommonTimeNode) SetID(id int) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"id",
			"",
			strconv.Itoa(id),
		),
	)
}

// Duration returns the duration in milliseconds. Returns -1 for indefinite.
func (ctn *CommonTimeNode) Duration() int {
	attr, found := ctn.GetAttribute("dur", "")
	if !found {
		return -1
	}
	if attr.Value() == "indefinite" {
		return -1
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return -1
	}

	return val
}

// SetDuration sets the duration in milliseconds. Use -1 for indefinite.
func (ctn *CommonTimeNode) SetDuration(
	durationMs int,
) {
	if durationMs < 0 {
		ctn.SetAttribute(
			openxml.NewAttribute(
				"",
				"dur",
				"",
				"indefinite",
			),
		)
	} else {
		ctn.SetAttribute(openxml.NewAttribute("", "dur", "", strconv.Itoa(durationMs)))
	}
}

// Fill returns the fill behavior.
func (ctn *CommonTimeNode) Fill() TimeNodeFill {
	attr, found := ctn.GetAttribute("fill", "")
	if !found {
		return TimeNodeFillRemove
	}

	return TimeNodeFill(attr.Value())
}

// SetFill sets the fill behavior.
func (ctn *CommonTimeNode) SetFill(
	fill TimeNodeFill,
) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"fill",
			"",
			string(fill),
		),
	)
}

// Restart returns the restart behavior.
func (ctn *CommonTimeNode) Restart() TimeNodeRestart {
	attr, found := ctn.GetAttribute("restart", "")
	if !found {
		return TimeNodeRestartAlways
	}

	return TimeNodeRestart(attr.Value())
}

// SetRestart sets the restart behavior.
func (ctn *CommonTimeNode) SetRestart(
	restart TimeNodeRestart,
) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"restart",
			"",
			string(restart),
		),
	)
}

// NodeType returns the node type.
func (ctn *CommonTimeNode) NodeType() TimeNodeType {
	attr, found := ctn.GetAttribute(
		"nodeType",
		"",
	)
	if !found {
		return ""
	}

	return TimeNodeType(attr.Value())
}

// SetNodeType sets the node type.
func (ctn *CommonTimeNode) SetNodeType(
	nodeType TimeNodeType,
) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"nodeType",
			"",
			string(nodeType),
		),
	)
}

// PresetID returns the preset effect ID.
func (ctn *CommonTimeNode) PresetID() int {
	attr, found := ctn.GetAttribute(
		"presetID",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPresetID sets the preset effect ID.
func (ctn *CommonTimeNode) SetPresetID(id int) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"presetID",
			"",
			strconv.Itoa(id),
		),
	)
}

// PresetClass returns the preset class (entr, exit, emph, path).
func (ctn *CommonTimeNode) PresetClass() string {
	attr, found := ctn.GetAttribute(
		"presetClass",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPresetClass sets the preset class.
func (ctn *CommonTimeNode) SetPresetClass(
	class string,
) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"presetClass",
			"",
			class,
		),
	)
}

// PresetSubtype returns the preset subtype.
func (ctn *CommonTimeNode) PresetSubtype() int {
	attr, found := ctn.GetAttribute(
		"presetSubtype",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPresetSubtype sets the preset subtype.
func (ctn *CommonTimeNode) SetPresetSubtype(
	subtype int,
) {
	ctn.SetAttribute(
		openxml.NewAttribute(
			"",
			"presetSubtype",
			"",
			strconv.Itoa(subtype),
		),
	)
}

// Clone creates a deep copy.
func (ctn *CommonTimeNode) Clone() openxml.Element {
	cloned := ctn.CompositeElementBase.Clone()

	return &CommonTimeNode{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// ParallelTimeNode (p:par)
// ===========================================================================

// ParallelTimeNode represents a parallel time node group (p:par).
type ParallelTimeNode struct {
	*openxml.CompositeElementBase
}

// NewParallelTimeNode creates a new ParallelTimeNode element.
func NewParallelTimeNode() *ParallelTimeNode {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"par",
		PrefixP,
	)

	return &ParallelTimeNode{
		CompositeElementBase: elem,
	}
}

// CommonTimeNode returns the common time node.
func (par *ParallelTimeNode) CommonTimeNode() *CommonTimeNode {
	elem := par.GetElement(
		"cTn",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if ctn, ok := elem.(*CommonTimeNode); ok {
		return ctn
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &CommonTimeNode{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy.
func (par *ParallelTimeNode) Clone() openxml.Element {
	cloned := par.CompositeElementBase.Clone()

	return &ParallelTimeNode{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SequenceTimeNode (p:seq)
// ===========================================================================

// SequenceTimeNode represents a sequence time node group (p:seq).
type SequenceTimeNode struct {
	*openxml.CompositeElementBase
}

// NewSequenceTimeNode creates a new SequenceTimeNode element.
func NewSequenceTimeNode() *SequenceTimeNode {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"seq",
		PrefixP,
	)

	return &SequenceTimeNode{
		CompositeElementBase: elem,
	}
}

// Concurrent returns whether the sequence runs concurrently.
func (seq *SequenceTimeNode) Concurrent() bool {
	attr, found := seq.GetAttribute(
		"concurrent",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == "true"
}

// SetConcurrent sets whether the sequence runs concurrently.
func (seq *SequenceTimeNode) SetConcurrent(
	concurrent bool,
) {
	if concurrent {
		seq.SetAttribute(
			openxml.NewAttribute(
				"",
				"concurrent",
				"",
				"1",
			),
		)
	} else {
		seq.RemoveAttribute("concurrent", "")
	}
}

// NextAction returns the next action.
func (seq *SequenceTimeNode) NextAction() string {
	attr, found := seq.GetAttribute("nextAc", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetNextAction sets the next action.
func (seq *SequenceTimeNode) SetNextAction(
	action string,
) {
	seq.SetAttribute(
		openxml.NewAttribute(
			"",
			"nextAc",
			"",
			action,
		),
	)
}

// CommonTimeNode returns the common time node.
func (seq *SequenceTimeNode) CommonTimeNode() *CommonTimeNode {
	elem := seq.GetElement(
		"cTn",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if ctn, ok := elem.(*CommonTimeNode); ok {
		return ctn
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &CommonTimeNode{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy.
func (seq *SequenceTimeNode) Clone() openxml.Element {
	cloned := seq.CompositeElementBase.Clone()

	return &SequenceTimeNode{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// AnimateEffect (p:animEffect)
// ===========================================================================

// AnimateEffect represents an animation effect (p:animEffect).
type AnimateEffect struct {
	*openxml.CompositeElementBase
}

// NewAnimateEffect creates a new AnimateEffect element.
func NewAnimateEffect() *AnimateEffect {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"animEffect",
		PrefixP,
	)

	return &AnimateEffect{
		CompositeElementBase: elem,
	}
}

// Filter returns the effect filter.
func (ae *AnimateEffect) Filter() AnimateEffectFilter {
	attr, found := ae.GetAttribute("filter", "")
	if !found {
		return ""
	}

	return AnimateEffectFilter(attr.Value())
}

// SetFilter sets the effect filter.
func (ae *AnimateEffect) SetFilter(
	filter AnimateEffectFilter,
) {
	ae.SetAttribute(
		openxml.NewAttribute(
			"",
			"filter",
			"",
			string(filter),
		),
	)
}

// Transition returns the transition type.
func (ae *AnimateEffect) Transition() AnimateEffectTransition {
	attr, found := ae.GetAttribute(
		"transition",
		"",
	)
	if !found {
		return AnimateEffectTransitionNone
	}

	return AnimateEffectTransition(attr.Value())
}

// SetTransition sets the transition type.
func (ae *AnimateEffect) SetTransition(
	transition AnimateEffectTransition,
) {
	ae.SetAttribute(
		openxml.NewAttribute(
			"",
			"transition",
			"",
			string(transition),
		),
	)
}

// SetDuration sets the duration using common behavior pattern.
func (ae *AnimateEffect) SetDuration(
	durationMs int,
) {
	// Duration is set via cBhvr child
	cBhvr := ae.GetElement(
		"cBhvr",
		NamespacePresentationML,
	)
	if cBhvr == nil {
		cBhvr = openxml.NewCompositeElement(
			NamespacePresentationML,
			"cBhvr",
			PrefixP,
		)
		ae.AppendChild(cBhvr)
	}
	if comp, ok := cBhvr.(*openxml.CompositeElementBase); ok {
		cTn := comp.GetElement(
			"cTn",
			NamespacePresentationML,
		)
		if cTn == nil {
			cTn = openxml.NewCompositeElement(
				NamespacePresentationML,
				"cTn",
				PrefixP,
			)
			comp.AppendChild(cTn)
		}
		if ctnComp, ok := cTn.(*openxml.CompositeElementBase); ok {
			ctnComp.SetAttribute(
				openxml.NewAttribute(
					"",
					"id",
					"",
					"1",
				),
			)
			ctnComp.SetAttribute(
				openxml.NewAttribute(
					"",
					"dur",
					"",
					strconv.Itoa(durationMs),
				),
			)
			ctnComp.SetAttribute(
				openxml.NewAttribute(
					"",
					"fill",
					"",
					"hold",
				),
			)
		}
	}
}

// SetShapeTarget sets the shape target.
func (ae *AnimateEffect) SetShapeTarget(
	shapeID int,
) {
	cBhvr := ae.GetElement(
		"cBhvr",
		NamespacePresentationML,
	)
	if cBhvr == nil {
		cBhvr = openxml.NewCompositeElement(
			NamespacePresentationML,
			"cBhvr",
			PrefixP,
		)
		ae.AppendChild(cBhvr)
	}
	if comp, ok := cBhvr.(*openxml.CompositeElementBase); ok {
		tgtEl := comp.GetElement(
			"tgtEl",
			NamespacePresentationML,
		)
		if tgtEl == nil {
			tgtEl = openxml.NewCompositeElement(
				NamespacePresentationML,
				"tgtEl",
				PrefixP,
			)
			comp.AppendChild(tgtEl)
		}
		if tgtComp, ok := tgtEl.(*openxml.CompositeElementBase); ok {
			spTgt := tgtComp.GetElement(
				"spTgt",
				NamespacePresentationML,
			)
			if spTgt == nil {
				spTgt = openxml.NewLeafElement(
					NamespacePresentationML,
					"spTgt",
					PrefixP,
				)
				tgtComp.AppendChild(spTgt)
			}
			if leaf, ok := spTgt.(*openxml.LeafElementBase); ok {
				leaf.SetAttribute(
					openxml.NewAttribute(
						"",
						"spid",
						"",
						strconv.Itoa(shapeID),
					),
				)
			}
		}
	}
}

// Clone creates a deep copy.
func (ae *AnimateEffect) Clone() openxml.Element {
	cloned := ae.CompositeElementBase.Clone()

	return &AnimateEffect{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// Animate (p:anim)
// ===========================================================================

// Animate represents a generic animation element (p:anim).
type Animate struct {
	*openxml.CompositeElementBase
}

// NewAnimate creates a new Animate element.
func NewAnimate() *Animate {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"anim",
		PrefixP,
	)

	return &Animate{CompositeElementBase: elem}
}

// CalcMode returns the calculation mode.
func (a *Animate) CalcMode() string {
	attr, found := a.GetAttribute("calcmode", "")
	if !found {
		return "discrete"
	}

	return attr.Value()
}

// SetCalcMode sets the calculation mode.
func (a *Animate) SetCalcMode(mode string) {
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"calcmode",
			"",
			mode,
		),
	)
}

// ValueType returns the value type.
func (a *Animate) ValueType() string {
	attr, found := a.GetAttribute("valueType", "")
	if !found {
		return "str"
	}

	return attr.Value()
}

// SetValueType sets the value type.
func (a *Animate) SetValueType(valType string) {
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"valueType",
			"",
			valType,
		),
	)
}

// Clone creates a deep copy.
func (a *Animate) Clone() openxml.Element {
	cloned := a.CompositeElementBase.Clone()

	return &Animate{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// AnimateColor (p:animClr)
// ===========================================================================

// AnimateColor represents a color animation element (p:animClr).
type AnimateColor struct {
	*openxml.CompositeElementBase
}

// NewAnimateColor creates a new AnimateColor element.
func NewAnimateColor() *AnimateColor {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"animClr",
		PrefixP,
	)

	return &AnimateColor{
		CompositeElementBase: elem,
	}
}

// ColorSpace returns the color space.
func (ac *AnimateColor) ColorSpace() string {
	attr, found := ac.GetAttribute("clrSpc", "")
	if !found {
		return "rgb"
	}

	return attr.Value()
}

// SetColorSpace sets the color space.
func (ac *AnimateColor) SetColorSpace(
	space string,
) {
	ac.SetAttribute(
		openxml.NewAttribute(
			"",
			"clrSpc",
			"",
			space,
		),
	)
}

// Direction returns the animation direction.
func (ac *AnimateColor) Direction() string {
	attr, found := ac.GetAttribute("dir", "")
	if !found {
		return "cw"
	}

	return attr.Value()
}

// SetDirection sets the animation direction.
func (ac *AnimateColor) SetDirection(dir string) {
	ac.SetAttribute(
		openxml.NewAttribute("", "dir", "", dir),
	)
}

// Clone creates a deep copy.
func (ac *AnimateColor) Clone() openxml.Element {
	cloned := ac.CompositeElementBase.Clone()

	return &AnimateColor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// AnimateMotion (p:animMotion)
// ===========================================================================

// AnimateMotion represents a motion path animation element (p:animMotion).
type AnimateMotion struct {
	*openxml.CompositeElementBase
}

// NewAnimateMotion creates a new AnimateMotion element.
func NewAnimateMotion() *AnimateMotion {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"animMotion",
		PrefixP,
	)

	return &AnimateMotion{
		CompositeElementBase: elem,
	}
}

// Origin returns the motion origin.
func (am *AnimateMotion) Origin() string {
	attr, found := am.GetAttribute("origin", "")
	if !found {
		return "layout"
	}

	return attr.Value()
}

// SetOrigin sets the motion origin.
func (am *AnimateMotion) SetOrigin(
	origin string,
) {
	am.SetAttribute(
		openxml.NewAttribute(
			"",
			"origin",
			"",
			origin,
		),
	)
}

// Path returns the motion path.
func (am *AnimateMotion) Path() string {
	attr, found := am.GetAttribute("path", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPath sets the motion path using SVG-like path notation.
func (am *AnimateMotion) SetPath(path string) {
	am.SetAttribute(
		openxml.NewAttribute(
			"",
			"path",
			"",
			path,
		),
	)
}

// PathEditMode returns the path edit mode.
func (am *AnimateMotion) PathEditMode() string {
	attr, found := am.GetAttribute(
		"pathEditMode",
		"",
	)
	if !found {
		return "relative"
	}

	return attr.Value()
}

// SetPathEditMode sets the path edit mode.
func (am *AnimateMotion) SetPathEditMode(
	mode string,
) {
	am.SetAttribute(
		openxml.NewAttribute(
			"",
			"pathEditMode",
			"",
			mode,
		),
	)
}

// Clone creates a deep copy.
func (am *AnimateMotion) Clone() openxml.Element {
	cloned := am.CompositeElementBase.Clone()

	return &AnimateMotion{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// AnimateRotation (p:animRot)
// ===========================================================================

// AnimateRotation represents a rotation animation element (p:animRot).
type AnimateRotation struct {
	*openxml.CompositeElementBase
}

// NewAnimateRotation creates a new AnimateRotation element.
func NewAnimateRotation() *AnimateRotation {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"animRot",
		PrefixP,
	)

	return &AnimateRotation{
		CompositeElementBase: elem,
	}
}

// By returns the rotation amount in 60000ths of a degree.
func (ar *AnimateRotation) By() int {
	attr, found := ar.GetAttribute("by", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetBy sets the rotation amount in 60000ths of a degree.
func (ar *AnimateRotation) SetBy(amount int) {
	ar.SetAttribute(
		openxml.NewAttribute(
			"",
			"by",
			"",
			strconv.Itoa(amount),
		),
	)
}

// From returns the starting rotation.
func (ar *AnimateRotation) From() int {
	attr, found := ar.GetAttribute("from", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFrom sets the starting rotation.
func (ar *AnimateRotation) SetFrom(degrees int) {
	ar.SetAttribute(
		openxml.NewAttribute(
			"",
			"from",
			"",
			strconv.Itoa(degrees),
		),
	)
}

// To returns the ending rotation.
func (ar *AnimateRotation) To() int {
	attr, found := ar.GetAttribute("to", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetTo sets the ending rotation.
func (ar *AnimateRotation) SetTo(degrees int) {
	ar.SetAttribute(
		openxml.NewAttribute(
			"",
			"to",
			"",
			strconv.Itoa(degrees),
		),
	)
}

// Clone creates a deep copy.
func (ar *AnimateRotation) Clone() openxml.Element {
	cloned := ar.CompositeElementBase.Clone()

	return &AnimateRotation{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// AnimateScale (p:animScale)
// ===========================================================================

// AnimateScale represents a scale animation element (p:animScale).
type AnimateScale struct {
	*openxml.CompositeElementBase
}

// NewAnimateScale creates a new AnimateScale element.
func NewAnimateScale() *AnimateScale {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"animScale",
		PrefixP,
	)

	return &AnimateScale{
		CompositeElementBase: elem,
	}
}

// ZoomContents returns whether to zoom contents.
func (as *AnimateScale) ZoomContents() bool {
	attr, found := as.GetAttribute(
		"zoomContents",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == "true"
}

// SetZoomContents sets whether to zoom contents.
func (as *AnimateScale) SetZoomContents(
	zoom bool,
) {
	if zoom {
		as.SetAttribute(
			openxml.NewAttribute(
				"",
				"zoomContents",
				"",
				"1",
			),
		)
	} else {
		as.RemoveAttribute("zoomContents", "")
	}
}

// Clone creates a deep copy.
func (as *AnimateScale) Clone() openxml.Element {
	cloned := as.CompositeElementBase.Clone()

	return &AnimateScale{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// Set (p:set)
// ===========================================================================

// Set represents an instant property set element (p:set).
type Set struct {
	*openxml.CompositeElementBase
}

// NewSet creates a new Set element.
func NewSet() *Set {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"set",
		PrefixP,
	)

	return &Set{CompositeElementBase: elem}
}

// Clone creates a deep copy.
func (s *Set) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &Set{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// Helper Elements for Animation Structure
// ===========================================================================

// StartConditionList represents the start condition list (p:stCondLst).
type StartConditionList struct {
	*openxml.CompositeElementBase
}

// NewStartConditionList creates a new StartConditionList.
func NewStartConditionList() *StartConditionList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"stCondLst",
		PrefixP,
	)

	return &StartConditionList{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy.
func (scl *StartConditionList) Clone() openxml.Element {
	cloned := scl.CompositeElementBase.Clone()

	return &StartConditionList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// EndConditionList represents the end condition list (p:endCondLst).
type EndConditionList struct {
	*openxml.CompositeElementBase
}

// NewEndConditionList creates a new EndConditionList.
func NewEndConditionList() *EndConditionList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"endCondLst",
		PrefixP,
	)

	return &EndConditionList{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy.
func (ecl *EndConditionList) Clone() openxml.Element {
	cloned := ecl.CompositeElementBase.Clone()

	return &EndConditionList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Condition represents a timing condition (p:cond).
type Condition struct {
	*openxml.CompositeElementBase
}

// NewCondition creates a new Condition.
func NewCondition() *Condition {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cond",
		PrefixP,
	)

	return &Condition{CompositeElementBase: elem}
}

// Event returns the event type.
func (c *Condition) Event() string {
	attr, found := c.GetAttribute("evt", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetEvent sets the event type (onClick, onBegin, onEnd, etc.).
func (c *Condition) SetEvent(event string) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"evt",
			"",
			event,
		),
	)
}

// Delay returns the delay in milliseconds.
func (c *Condition) Delay() int {
	attr, found := c.GetAttribute("delay", "")
	if !found {
		return 0
	}
	if attr.Value() == "indefinite" {
		return -1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDelay sets the delay in milliseconds.
func (c *Condition) SetDelay(delayMs int) {
	if delayMs < 0 {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"delay",
				"",
				"indefinite",
			),
		)
	} else {
		c.SetAttribute(openxml.NewAttribute("", "delay", "", strconv.Itoa(delayMs)))
	}
}

// Clone creates a deep copy.
func (c *Condition) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Condition{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TargetElement represents a target element container (p:tgtEl).
type TargetElement struct {
	*openxml.CompositeElementBase
}

// NewTargetElement creates a new TargetElement.
func NewTargetElement() *TargetElement {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"tgtEl",
		PrefixP,
	)

	return &TargetElement{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy.
func (te *TargetElement) Clone() openxml.Element {
	cloned := te.CompositeElementBase.Clone()

	return &TargetElement{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ShapeTarget represents a shape target reference (p:spTgt).
type ShapeTarget struct {
	*openxml.LeafElementBase
}

// NewShapeTarget creates a new ShapeTarget.
func NewShapeTarget() *ShapeTarget {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"spTgt",
		PrefixP,
	)

	return &ShapeTarget{LeafElementBase: elem}
}

// ShapeID returns the target shape ID.
func (st *ShapeTarget) ShapeID() int {
	attr, found := st.GetAttribute("spid", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetShapeID sets the target shape ID.
func (st *ShapeTarget) SetShapeID(id int) {
	st.SetAttribute(
		openxml.NewAttribute(
			"",
			"spid",
			"",
			strconv.Itoa(id),
		),
	)
}

// Clone creates a deep copy.
func (st *ShapeTarget) Clone() openxml.Element {
	cloned := st.LeafElementBase.Clone()

	return &ShapeTarget{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ChildTimeNodeList represents the child time node list (p:childTnLst).
type ChildTimeNodeList struct {
	*openxml.CompositeElementBase
}

// NewChildTimeNodeList creates a new ChildTimeNodeList.
func NewChildTimeNodeList() *ChildTimeNodeList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"childTnLst",
		PrefixP,
	)

	return &ChildTimeNodeList{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy.
func (ctl *ChildTimeNodeList) Clone() openxml.Element {
	cloned := ctl.CompositeElementBase.Clone()

	return &ChildTimeNodeList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
