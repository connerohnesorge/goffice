# Wordprocessing Settings Specification

## Requirements

### Requirement: Settings Root Element
The system SHALL provide a Settings element as the root of the document settings part.

#### Scenario: Access settings
- GIVEN a Settings element
- WHEN individual setting elements are accessed
- THEN document settings are returned

#### Scenario: Document protection
- GIVEN a Settings element
- WHEN DocumentProtection() is called
- THEN document protection settings are returned

#### Scenario: Default tab stop
- GIVEN a Settings element
- WHEN DefaultTabStop() is called
- THEN the default tab stop width is returned

#### Scenario: Set default tab stop
- GIVEN a Settings element
- WHEN DefaultTabStop().Val.SetValue(720) is called
- THEN 0.5 inch default tabs are set

### Requirement: Zoom Settings
The system SHALL provide zoom settings.

#### Scenario: Get zoom
- GIVEN a Settings element
- WHEN Zoom() is called
- THEN zoom percentage is returned

#### Scenario: Set zoom
- GIVEN a Settings element
- WHEN Zoom().Percent.SetValue("100") is called
- THEN 100% zoom is set

#### Scenario: Zoom view type
- GIVEN a Settings element
- WHEN Zoom().Val is accessed
- THEN None, FullPage, BestFit, or TextFit is returned

### Requirement: Display Settings
The system SHALL provide display-related settings.

#### Scenario: Display background shape
- GIVEN a Settings element
- WHEN DisplayBackgroundShape() is called
- THEN whether background is displayed is returned

#### Scenario: Print fractional width
- GIVEN a Settings element
- WHEN PrintFractionalCharacterWidth() is called
- THEN fractional character width setting is returned

#### Scenario: Mirror margins
- GIVEN a Settings element
- WHEN MirrorMargins() is called
- THEN whether margins mirror for binding is returned

#### Scenario: Even and odd headers
- GIVEN a Settings element
- WHEN EvenAndOddHeaders() is called
- THEN whether different even/odd headers is returned

### Requirement: Proofing Settings
The system SHALL provide spell check and grammar settings.

#### Scenario: Proofing state
- GIVEN a Settings element
- WHEN ProofState() is called
- THEN spelling and grammar check state is returned

#### Scenario: Hide spelling errors
- GIVEN a Settings element
- WHEN HideSpellingErrors() is called
- THEN whether spelling marks are hidden is returned

#### Scenario: Hide grammar errors
- GIVEN a Settings element
- WHEN HideGrammaticalErrors() is called
- THEN whether grammar marks are hidden is returned

#### Scenario: Default language
- GIVEN a Settings element
- WHEN ThemeFontLang() is called
- THEN default language settings are returned

### Requirement: Compatibility Settings
The system SHALL provide compatibility settings for different Word versions.

#### Scenario: Access compatibility
- GIVEN a Settings element
- WHEN Compatibility() is called
- THEN compatibility settings element is returned

#### Scenario: Compatibility settings
- GIVEN a Compatibility element
- WHEN individual compatibility options are accessed
- THEN specific compatibility modes are returned

#### Scenario: Use Word 2003 table style rules
- GIVEN a Compatibility element
- WHEN UseWord2003TableStyleRules() is accessed
- THEN Word 2003 compatibility flag is returned

#### Scenario: Grow autofit
- GIVEN a Compatibility element
- WHEN GrowAutofit() is accessed
- THEN autofit behavior setting is returned

### Requirement: Track Changes Settings
The system SHALL provide revision tracking settings.

#### Scenario: Track revisions
- GIVEN a Settings element
- WHEN TrackRevisions() is called
- THEN whether revisions are tracked is returned

#### Scenario: Set track revisions
- GIVEN a Settings element
- WHEN SetTrackRevisions(true) is called
- THEN revision tracking is enabled

#### Scenario: Revision view
- GIVEN a Settings element
- WHEN RevisionView() is called
- THEN what revision marks are shown is returned

#### Scenario: Document protection for tracked changes
- GIVEN a Settings element
- WHEN DocumentProtection().Edit.SetValue("trackedChanges") is called
- THEN only tracked changes are allowed

### Requirement: Document Protection Settings
The system SHALL provide document protection configuration.

#### Scenario: Protection type
- GIVEN a DocumentProtection element
- WHEN Edit() is called
- THEN ReadOnly, Comments, TrackedChanges, or Forms is returned

#### Scenario: Set protection
- GIVEN a DocumentProtection element
- WHEN Edit().SetValue("readOnly") is called
- THEN document is read-only protected

#### Scenario: Password protection
- GIVEN a DocumentProtection element
- WHEN Hash and Salt attributes are set
- THEN password protection is applied

#### Scenario: Formatting restriction
- GIVEN a DocumentProtection element
- WHEN Formatting() is accessed
- THEN whether formatting is restricted is returned

#### Scenario: Enforcement
- GIVEN a DocumentProtection element
- WHEN Enforcement() is accessed
- THEN whether protection is enforced is returned

### Requirement: Mail Merge Settings
The system SHALL provide mail merge configuration.

#### Scenario: Mail merge type
- GIVEN a Settings element
- WHEN MailMerge() is called
- THEN mail merge settings are returned

#### Scenario: Main document type
- GIVEN a MailMerge element
- WHEN MainDocumentType() is called
- THEN Form Letters, Labels, Envelopes, etc. is returned

#### Scenario: Data source
- GIVEN a MailMerge element
- WHEN DataSource properties are accessed
- THEN data source connection info is returned

### Requirement: Writing Protection Settings
The system SHALL provide write protection.

#### Scenario: Write protection
- GIVEN a Settings element
- WHEN WriteProtection() is called
- THEN write protection settings are returned

#### Scenario: Recommended read-only
- GIVEN a WriteProtection element
- WHEN Recommended() is accessed
- THEN whether read-only is recommended is returned

### Requirement: Rsid Settings
The system SHALL provide revision save ID settings.

#### Scenario: Rsid root
- GIVEN a Settings element
- WHEN RsidRoot() is called
- THEN the original document rsid is returned

#### Scenario: Rsids collection
- GIVEN a Settings element
- WHEN Rsids() is called
- THEN all revision save IDs are returned

### Requirement: Document Variables
The system SHALL provide document variable storage.

#### Scenario: Access document variables
- GIVEN a Settings element
- WHEN DocumentVariables() is called
- THEN document variables are returned

#### Scenario: Get variable
- GIVEN a DocumentVariables element
- WHEN GetVariable(name) is called
- THEN the variable value is returned

#### Scenario: Set variable
- GIVEN a DocumentVariables element
- WHEN SetVariable(name, value) is called
- THEN the variable is added or updated

### Requirement: WebSettings Root Element
The system SHALL provide WebSettings for web-related configuration.

#### Scenario: Optimized for browser
- GIVEN a WebSettings element
- WHEN OptimizeForBrowser() is called
- THEN browser optimization setting is returned

#### Scenario: Allow PNG
- GIVEN a WebSettings element
- WHEN AllowPNG() is called
- THEN whether PNG images are allowed is returned

#### Scenario: Target screen size
- GIVEN a WebSettings element
- WHEN TargetScreenSize() is called
- THEN target screen resolution is returned

#### Scenario: Encoding
- GIVEN a WebSettings element
- WHEN Encoding() is called
- THEN the web page encoding is returned

### Requirement: Font Table
The system SHALL provide font table configuration.

#### Scenario: Fonts root element
- GIVEN a Fonts element (FontTablePart root)
- WHEN Elements[Font]() is called
- THEN all font definitions are returned

#### Scenario: Get font by name
- GIVEN a Fonts element
- WHEN GetFont(name) is called
- THEN the Font element with that name is returned

#### Scenario: Font properties
- GIVEN a Font element
- WHEN Name(), Charset(), Family(), Pitch() are accessed
- THEN font properties are returned

#### Scenario: Embed font
- GIVEN a Font element
- WHEN EmbedRegular(), EmbedBold(), etc. are accessed
- THEN embedded font data references are returned

