# Wordprocessing Elements Specification (Delta)

## MODIFIED Requirements

### Requirement: Field Elements
The system SHALL provide elements for field codes AND a high-level API for form field manipulation.

#### Scenario: Simple field (unchanged)
- GIVEN a SimpleField element
- WHEN Instruction() is accessed
- THEN the field code (e.g., "PAGE", "DATE") is returned

#### Scenario: Complex field (unchanged)
- GIVEN FieldChar elements (Begin, Separate, End) with FieldCode
- WHEN processed
- THEN the complex field instruction is understood

#### Scenario: Create page number field (unchanged)
- GIVEN need for page number
- WHEN NewSimpleField("PAGE") is called
- THEN a page number field is created

#### Scenario: Form field API access
- GIVEN a SimpleField with FormFieldData child
- WHEN wrapped in FormField object
- THEN high-level get/set methods are available

## ADDED Requirements

### Requirement: Form Field Wrapper
The system SHALL provide a FormField wrapper for legacy form fields.

#### Scenario: FormField wraps SimpleField
- GIVEN a SimpleField with FormFieldData
- WHEN FormField wrapper is created
- THEN the wrapper provides access to field properties

#### Scenario: Field type detection
- GIVEN a SimpleField with instr="FORMTEXT"
- WHEN FormField.Type() is called
- THEN FormFieldTypeText is returned

#### Scenario: Field name access
- GIVEN a FormField with name "CustomerName"
- WHEN FormField.Name() is called
- THEN "CustomerName" is returned

#### Scenario: Field name modification
- GIVEN a FormField
- WHEN FormField.SetName("NewName") is called
- THEN the ffData/name element is updated

### Requirement: Text Field Manipulation
The system SHALL provide methods for text form field manipulation.

#### Scenario: Get text field value
- GIVEN a text FormField with value "John Doe"
- WHEN GetTextValue() is called
- THEN "John Doe" is returned

#### Scenario: Set text field value
- GIVEN a text FormField
- WHEN SetTextValue("Jane Smith") is called
- THEN both ffData/textInput/default AND run/text are updated to "Jane Smith"

#### Scenario: Set maximum length
- GIVEN a text FormField
- WHEN SetMaxLength(50) is called
- THEN ffData/textInput/maxLength is set to 50

#### Scenario: Get maximum length
- GIVEN a text FormField with maxLength=50
- WHEN GetMaxLength() is called
- THEN 50 is returned

#### Scenario: Set default text value
- GIVEN a text FormField
- WHEN SetDefaultTextValue("Default Name") is called
- THEN ffData/textInput/default is set to "Default Name"

#### Scenario: Text value synchronization
- GIVEN a text FormField
- WHEN SetTextValue("New Value") is called
- THEN the run text element matches "New Value" exactly

### Requirement: CheckBox Field Manipulation
The system SHALL provide methods for checkbox form field manipulation.

#### Scenario: Check checkbox state
- GIVEN a checkbox FormField that is checked
- WHEN IsChecked() is called
- THEN true is returned

#### Scenario: Toggle checkbox to checked
- GIVEN an unchecked checkbox FormField
- WHEN SetChecked(true) is called
- THEN ffData/checkBox/checked is set to 1 AND symbol character is updated to F052

#### Scenario: Toggle checkbox to unchecked
- GIVEN a checked checkbox FormField
- WHEN SetChecked(false) is called
- THEN ffData/checkBox/checked is set to 0 AND symbol character is updated to F06F

#### Scenario: Set checkbox size
- GIVEN a checkbox FormField
- WHEN SetCheckBoxSize(20) is called
- THEN ffData/checkBox/size is set to 20 (10 points)

#### Scenario: Get checkbox size
- GIVEN a checkbox FormField with size=20
- WHEN GetCheckBoxSize() is called
- THEN 20 is returned

#### Scenario: Auto checkbox size
- GIVEN a checkbox FormField
- WHEN SetAutoCheckBoxSize(true) is called
- THEN ffData/checkBox/sizeAuto is set to 1

#### Scenario: Checkbox symbol synchronization
- GIVEN a checkbox FormField
- WHEN SetChecked(true) is called
- THEN the run/sym element has char="F052" (checked box symbol)

#### Scenario: Checkbox symbol must be set via SetChecked
- GIVEN a checkbox FormField
- WHEN SetChecked() is used to change state
- THEN both ffData/checkBox/checked AND run/sym/char are synchronized
- NOTE: Symbols MUST only be set via SetChecked() to ensure proper synchronization
- NOTE: Manual manipulation of symbol elements may create invalid state

### Requirement: DropDown Field Manipulation
The system SHALL provide methods for dropdown form field manipulation.

#### Scenario: Get dropdown items
- GIVEN a dropdown FormField with items ["USA", "Canada", "UK"]
- WHEN GetDropDownItems() is called
- THEN ["USA", "Canada", "UK"] is returned

#### Scenario: Set dropdown items
- GIVEN a dropdown FormField
- WHEN SetDropDownItems(["Option A", "Option B", "Option C"]) is called
- THEN ffData/ddList contains three listEntry elements with those values

#### Scenario: Get selected index
- GIVEN a dropdown FormField with selectedIndex=1
- WHEN GetSelectedIndex() is called
- THEN 1 is returned

#### Scenario: Set selected index
- GIVEN a dropdown FormField with items ["A", "B", "C"]
- WHEN SetSelectedIndex(2) is called
- THEN ffData/ddList/result is 2 AND run/text is "C"

#### Scenario: Get selected value
- GIVEN a dropdown FormField with items ["USA", "Canada"] and selectedIndex=0
- WHEN GetSelectedValue() is called
- THEN "USA" is returned

#### Scenario: DropDown value synchronization
- GIVEN a dropdown FormField with items ["X", "Y", "Z"]
- WHEN SetSelectedIndex(1) is called
- THEN the run text element contains "Y" exactly

#### Scenario: Invalid selected index handling
- GIVEN a dropdown FormField with 3 items
- WHEN SetSelectedIndex(5) is called (out of range)
- THEN the operation is ignored or error is returned

#### Scenario: Get default dropdown index
- GIVEN a dropdown FormField with default=2
- WHEN GetDefaultDropDownIndex() is called
- THEN 2 is returned

#### Scenario: Set default dropdown index
- GIVEN a dropdown FormField with items ["A", "B", "C"]
- WHEN SetDefaultDropDownIndex(1) is called
- THEN ffData/ddList/default is set to 1

#### Scenario: DropDown max items validation
- GIVEN a Paragraph
- WHEN SetDropDownItems() is called with 26 items
- THEN the operation panics with "cannot have more than 25 items"

#### Scenario: DropDown insert max items validation
- GIVEN a Paragraph
- WHEN InsertDropDown() is called with 26 items
- THEN the operation panics with schema constraint error

### Requirement: Form Field Insertion in Paragraphs
The system SHALL provide methods to insert form fields into paragraphs.

#### Scenario: Insert text field
- GIVEN a Paragraph
- WHEN InsertTextField("CustomerName", "Enter name") is called
- THEN a SimpleField with FORMTEXT instruction, FormFieldData, and run is added

#### Scenario: Insert checkbox
- GIVEN a Paragraph
- WHEN InsertCheckBox("AgreeTerms", false) is called
- THEN a SimpleField with FORMCHECKBOX instruction, FormFieldData with checkBox, and symbol run is added

#### Scenario: Insert dropdown
- GIVEN a Paragraph
- WHEN InsertDropDown("Country", ["USA", "UK"], 0) is called
- THEN a SimpleField with FORMDROPDOWN instruction, FormFieldData with ddList, and run with selected text is added

#### Scenario: Multiple fields in one paragraph
- GIVEN a Paragraph
- WHEN InsertTextField(), InsertCheckBox(), and InsertDropDown() are called sequentially
- THEN all three SimpleField elements are added to the paragraph

#### Scenario: Inserted field is immediately accessible
- GIVEN a Paragraph
- WHEN field := InsertTextField("Test", "Value") is called
- THEN field.GetTextValue() returns "Value" immediately

### Requirement: Form Field Removal
The system SHALL support removing form fields from paragraphs.

#### Scenario: Remove field from paragraph
- GIVEN a FormField in a Paragraph
- WHEN Remove() is called
- THEN the SimpleField element is removed from the paragraph

#### Scenario: Remove and re-iterate
- GIVEN a document with 3 form fields
- WHEN one field is removed
- THEN GetFormFields() yields only 2 fields

### Requirement: Form Field Enabled State
The system SHALL support enabling and disabling form fields.

#### Scenario: Check if field is enabled
- GIVEN a FormField with enabled element present
- WHEN Enabled() is called
- THEN true is returned

#### Scenario: Disable field
- GIVEN a FormField
- WHEN SetEnabled(false) is called
- THEN the ffData/enabled element is removed

#### Scenario: Enable field
- GIVEN a FormField
- WHEN SetEnabled(true) is called
- THEN the ffData/enabled element is added (if not present)

#### Scenario: Disabled field persists on save
- GIVEN a disabled FormField
- WHEN the document is saved and reopened
- THEN the field is still disabled

### Requirement: Form Field XML Structure Compliance
The system SHALL generate form field XML structures compatible with Microsoft Word.

#### Scenario: Text field XML structure
- GIVEN a created text FormField
- WHEN serialized to XML
- THEN the structure matches: fldSimple[@instr="FORMTEXT"]/ffData/textInput/default

#### Scenario: CheckBox field XML structure
- GIVEN a created checkbox FormField
- WHEN serialized to XML
- THEN the structure matches: fldSimple[@instr="FORMCHECKBOX"]/ffData/checkBox/checked

#### Scenario: DropDown field XML structure
- GIVEN a created dropdown FormField
- WHEN serialized to XML
- THEN the structure matches: fldSimple[@instr="FORMDROPDOWN"]/ffData/ddList/listEntry

#### Scenario: Field name in ffData
- GIVEN any FormField with name "TestField"
- WHEN serialized to XML
- THEN ffData contains: name[@val="TestField"]

#### Scenario: Enabled attribute
- GIVEN an enabled FormField
- WHEN serialized to XML
- THEN ffData contains: enabled element (empty)

### Requirement: Form Field Roundtrip Compatibility
The system SHALL maintain form field integrity across save/load cycles.

#### Scenario: Text field value roundtrip
- GIVEN a text FormField with value "Test Value"
- WHEN saved, closed, and reopened
- THEN GetTextValue() returns "Test Value"

#### Scenario: CheckBox state roundtrip
- GIVEN a checked checkbox FormField
- WHEN saved, closed, and reopened
- THEN IsChecked() returns true

#### Scenario: DropDown selection roundtrip
- GIVEN a dropdown FormField with selectedIndex=2
- WHEN saved, closed, and reopened
- THEN GetSelectedIndex() returns 2

#### Scenario: Field name roundtrip
- GIVEN a FormField with name "UniqueField"
- WHEN saved, closed, and reopened
- THEN Name() returns "UniqueField"

#### Scenario: Multiple fields roundtrip
- GIVEN a document with text, checkbox, and dropdown fields
- WHEN saved, closed, and reopened
- THEN all three fields exist with correct types and values
