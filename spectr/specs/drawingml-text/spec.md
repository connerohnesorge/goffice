# Drawingml Text Specification

## Requirements

### Requirement: Text Body Element

The system SHALL provide a `TextBody` element (a:txBody) as the container for shape text.

#### Scenario: TextBody structure
- GIVEN a TextBody element
- WHEN children are accessed
- THEN BodyProperties (a:bodyPr), ListStyle (a:lstStyle), and Paragraph (a:p) elements are available

#### Scenario: Multiple paragraphs
- GIVEN a TextBody element
- WHEN paragraphs are accessed
- THEN multiple Paragraph (a:p) elements are supported
- AND paragraphs are rendered in document order

### Requirement: Body Properties Element

The system SHALL provide a `BodyProperties` element (a:bodyPr) for text container settings.

#### Scenario: Text wrapping
- GIVEN a BodyProperties element
- WHEN wrap attribute is accessed
- THEN values none, square are available
- AND none means text does not wrap

#### Scenario: Text anchor
- GIVEN a BodyProperties element
- WHEN anchor attribute is accessed
- THEN values t (top), ctr (center), b (bottom), just (justified), dist (distributed) are available

#### Scenario: Text insets
- GIVEN a BodyProperties element
- WHEN inset attributes are accessed
- THEN lIns, tIns, rIns, bIns (in EMUs) define text margins from shape edge

#### Scenario: Text rotation
- GIVEN a BodyProperties element
- WHEN rot attribute is accessed
- THEN text rotation in 1/60000 degrees is available (independent of shape rotation)

#### Scenario: Vertical text
- GIVEN a BodyProperties element
- WHEN vert attribute is accessed
- THEN values horz (horizontal), vert (vertical), vert270 (270 degree), wordArtVert, eaVert, mongolianVert, wordArtVertRtl are available

#### Scenario: Auto-fit options
- GIVEN a BodyProperties element
- WHEN auto-fit is accessed
- THEN NoAutoFit (a:noAutofit), NormalAutoFit (a:normAutofit with fontScale, lnSpcReduction), ShapeAutoFit (a:spAutoFit) are available

#### Scenario: Text columns
- GIVEN a BodyProperties element
- WHEN numCol attribute is accessed
- THEN the number of text columns (1 or more) is available
- AND spcCol defines space between columns in EMUs

#### Scenario: Overflow behavior
- GIVEN a BodyProperties element
- WHEN horzOverflow and vertOverflow are accessed
- THEN values overflow, clip are available

### Requirement: List Style Element

The system SHALL provide a `ListStyle` element (a:lstStyle) for paragraph level defaults.

#### Scenario: Default paragraph properties
- GIVEN a ListStyle element
- WHEN level properties are accessed
- THEN DefParagraphProperties (a:defPPr) provides document-wide defaults
- AND Level1ParagraphProperties through Level9ParagraphProperties (a:lvl1pPr - a:lvl9pPr) provide level-specific defaults

### Requirement: Paragraph Element

The system SHALL provide a `Paragraph` element (a:p) for text paragraphs.

#### Scenario: Paragraph structure
- GIVEN a Paragraph element
- WHEN children are accessed
- THEN ParagraphProperties (a:pPr), Run (a:r), LineBreak (a:br), Field (a:fld), EndParagraphRunProperties (a:endParaRPr) are available

#### Scenario: Paragraph content ordering
- GIVEN a Paragraph element
- WHEN content is accessed
- THEN Run, LineBreak, and Field elements appear in document order
- AND EndParagraphRunProperties applies to the paragraph end mark

### Requirement: Paragraph Properties Element

The system SHALL provide a `ParagraphProperties` element (a:pPr) for paragraph formatting.

#### Scenario: Text alignment
- GIVEN a ParagraphProperties element
- WHEN algn attribute is accessed
- THEN values l (left), ctr (center), r (right), just (justified), justLow, dist, thaiDist are available

#### Scenario: Indentation
- GIVEN a ParagraphProperties element
- WHEN indentation is accessed
- THEN marL (left margin), marR (right margin), indent (first line), lvl (outline level 0-8) are available
- AND values are in EMUs

#### Scenario: Line spacing
- GIVEN a ParagraphProperties element
- WHEN LineSpacing (a:lnSpc) is accessed
- THEN SpacingPercent (a:spcPct with val in 1/1000 percent) or SpacingPoints (a:spcPts with val in 1/100 points) is available

#### Scenario: Space before/after
- GIVEN a ParagraphProperties element
- WHEN SpaceBefore (a:spcBef) and SpaceAfter (a:spcAft) are accessed
- THEN SpacingPercent or SpacingPoints is available for each

#### Scenario: Bullet properties
- GIVEN a ParagraphProperties element
- WHEN bullet properties are accessed
- THEN BulletFont (a:buFont), BulletColor (a:buClr), BulletSizePercent (a:buSzPct), BulletSizePoints (a:buSzPts), BulletCharacter (a:buChar), BulletAutoNumber (a:buAutoNum), BulletBlip (a:buBlip), BulletNone (a:buNone) are available

#### Scenario: Auto-numbering bullets
- GIVEN a BulletAutoNumber element
- WHEN type attribute is accessed
- THEN values alphaLcParenBoth, alphaLcParenR, alphaLcPeriod, alphaUcParenBoth, alphaUcParenR, alphaUcPeriod, arabic1Minus, arabic2Minus, arabicDbPeriod, arabicDbPlain, arabicParenBoth, arabicParenR, arabicPeriod, arabicPlain, circleNumDbPlain, circleNumWdBlackPlain, circleNumWdWhitePlain, ea1ChsPeriod, ea1ChsPlain, ea1JpnChsDbPeriod, ea1JpnKorPeriod, ea1JpnKorPlain, hebrew2Minus, hindiAlpha1Period, hindiAlphaPeriod, hindiNumParenR, hindiNumPeriod, romanLcParenBoth, romanLcParenR, romanLcPeriod, romanUcParenBoth, romanUcParenR, romanUcPeriod, thaiAlphaParenBoth, thaiAlphaParenR, thaiAlphaPeriod, thaiNumParenBoth, thaiNumParenR, thaiNumPeriod are available
- AND startAt attribute defines starting number

#### Scenario: Tab stops
- GIVEN a ParagraphProperties element
- WHEN TabList (a:tabLst) is accessed
- THEN Tab (a:tab) elements with pos (position in EMUs) and algn (l, ctr, r, dec) are available

#### Scenario: Right-to-left
- GIVEN a ParagraphProperties element
- WHEN rtl attribute is accessed
- THEN true/false indicates right-to-left paragraph direction

### Requirement: Run Element

The system SHALL provide a `Run` element (a:r) for formatted text spans.

#### Scenario: Run structure
- GIVEN a Run element
- WHEN children are accessed
- THEN RunProperties (a:rPr) and Text (a:t) are available

#### Scenario: Text content
- GIVEN a Text element (a:t)
- WHEN text is accessed
- THEN the actual text string is available
- AND whitespace is preserved according to xml:space attribute

### Requirement: Run Properties Element

The system SHALL provide a `RunProperties` element (a:rPr) for character formatting.

#### Scenario: Font properties
- GIVEN a RunProperties element
- WHEN font is accessed
- THEN Latin (a:latin), EastAsian (a:ea), ComplexScript (a:cs), Symbol (a:sym) fonts are available
- AND each has typeface, panose, pitchFamily, charset attributes

#### Scenario: Font size
- GIVEN a RunProperties element
- WHEN sz attribute is accessed
- THEN font size in 1/100 points is available (e.g., 1200 = 12pt)

#### Scenario: Bold and italic
- GIVEN a RunProperties element
- WHEN b and i attributes are accessed
- THEN true/false indicates bold and italic respectively

#### Scenario: Underline
- GIVEN a RunProperties element
- WHEN u attribute is accessed
- THEN values none, words, sng (single), dbl (double), heavy, dotted, dottedHeavy, dash, dashHeavy, dashLong, dashLongHeavy, dotDash, dotDashHeavy, dotDotDash, dotDotDashHeavy, wavy, wavyHeavy, wavyDbl are available

#### Scenario: Strike
- GIVEN a RunProperties element
- WHEN strike attribute is accessed
- THEN values noStrike, sngStrike (single), dblStrike (double) are available

#### Scenario: Baseline shift
- GIVEN a RunProperties element
- WHEN baseline attribute is accessed
- THEN percentage value for superscript (positive) or subscript (negative) is available
- AND typically +/-25000 (25%)

#### Scenario: Character spacing
- GIVEN a RunProperties element
- WHEN spc attribute is accessed
- THEN spacing in 1/100 points is available (can be negative for condensed)

#### Scenario: Kerning
- GIVEN a RunProperties element
- WHEN kern attribute is accessed
- THEN minimum font size in 1/100 points for kerning to apply is available

#### Scenario: Text color
- GIVEN a RunProperties element
- WHEN SolidFill (a:solidFill) is accessed
- THEN text color with any color type and transforms is available

#### Scenario: Text outline
- GIVEN a RunProperties element
- WHEN Outline (a:ln) is accessed
- THEN text outline with line properties is available

#### Scenario: Text effects
- GIVEN a RunProperties element
- WHEN EffectList (a:effectLst) or EffectDag (a:effectDag) is accessed
- THEN text effects (shadow, glow, etc.) are available

#### Scenario: Highlight
- GIVEN a RunProperties element
- WHEN Highlight (a:highlight) is accessed
- THEN highlight color is available

#### Scenario: Hyperlink
- GIVEN a RunProperties element
- WHEN HyperlinkClick (a:hlinkClick) or HyperlinkMouseOver (a:hlinkMouseOver) is accessed
- THEN r:id references relationship to target
- AND action, tgtFrame, tooltip, history, invalidUrl, endSnd are available

#### Scenario: Language
- GIVEN a RunProperties element
- WHEN lang and altLang attributes are accessed
- THEN BCP 47 language tags (e.g., "en-US") are available

### Requirement: Line Break Element

The system SHALL provide a `LineBreak` element (a:br) for explicit line breaks.

#### Scenario: Line break properties
- GIVEN a LineBreak element
- WHEN RunProperties (a:rPr) is accessed
- THEN formatting for the line break character is available

### Requirement: Field Element

The system SHALL provide a `Field` element (a:fld) for dynamic content.

#### Scenario: Field structure
- GIVEN a Field element
- WHEN properties are accessed
- THEN type (field type like slidenum, datetime, etc.), id (GUID) are available
- AND RunProperties and Text provide current display value

### Requirement: End Paragraph Run Properties Element

The system SHALL provide an `EndParagraphRunProperties` element (a:endParaRPr) for paragraph end formatting.

#### Scenario: End paragraph formatting
- GIVEN an EndParagraphRunProperties element
- WHEN properties are accessed
- THEN same properties as RunProperties are available
- AND these apply to the paragraph end mark (carriage return)

### Requirement: Text Hyperlinks

The system SHALL support hyperlinks within text.

#### Scenario: Hyperlink click
- GIVEN a HyperlinkClick element (a:hlinkClick)
- WHEN properties are accessed
- THEN r:id references relationship to target URL or document location
- AND tooltip provides hover text
- AND action can trigger custom actions

#### Scenario: Mouse over link
- GIVEN a HyperlinkMouseOver element (a:hlinkMouseOver)
- WHEN properties are accessed
- THEN same structure as HyperlinkClick is available
- AND activates on hover instead of click

