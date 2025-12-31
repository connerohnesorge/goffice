package validation

import (
	"testing"
)

// Mock versioned element for testing
type mockVersionedElement struct {
	localName    string
	namespaceURI string
	attributes   map[string]string
	children     []any
	parent       any
	availability *VersionAvailability
}

func (e *mockVersionedElement) LocalName() string {
	return e.localName
}

func (e *mockVersionedElement) NamespaceURI() string {
	return e.namespaceURI
}

func (e *mockVersionedElement) Parent() any {
	return e.parent
}

func (e *mockVersionedElement) Attributes() []mockAttribute {
	return mapToAttrs(e.attributes)
}

func (e *mockVersionedElement) Children() []any {
	return e.children
}

func (e *mockVersionedElement) Availability() *VersionAvailability {
	return e.availability
}

func newMockVersionedElement(
	localName string,
	introducedIn FileFormatVersions,
) *mockVersionedElement {
	return &mockVersionedElement{
		localName:    localName,
		namespaceURI: "",
		attributes:   make(map[string]string),
		children:     make([]any, 0),
		availability: NewVersionAvailability(
			introducedIn,
		),
	}
}

func TestDetectMinimumVersion(t *testing.T) {
	t.Run(
		"detects Office2016 for base elements",
		func(t *testing.T) {
			// Create a document with only Office 2016 elements
			doc := newMockElement("document").
				withChild(newMockElement("body")).
				withChild(newMockElement("paragraph"))

			version := DetectMinimumVersion(doc)
			if version != Office2016 {
				t.Errorf(
					"expected Office2016, got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"detects Office2019 when element requires it",
		func(t *testing.T) {
			// Create a document with an Office 2019 element
			feature := newMockVersionedElement(
				"newFeature",
				Office2019,
			)
			doc := newMockElement("document").
				withChild(feature)

			version := DetectMinimumVersion(doc)
			if version != Office2019 {
				t.Errorf(
					"expected Office2019, got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"detects Office2021 when nested element requires it",
		func(t *testing.T) {
			// Create a document with nested Office 2021 element
			nestedElement := newMockVersionedElement(
				"advancedFeature",
				Office2021,
			)
			doc := newMockElement("document").
				withChild(
					newMockElement("body").
						withChild(nestedElement),
				)

			version := DetectMinimumVersion(doc)
			if version != Office2021 {
				t.Errorf(
					"expected Office2021, got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"detects highest version in mixed document",
		func(t *testing.T) {
			// Create a document with Office 2016, 2019, and 2021 elements
			doc := newMockElement("document").
				withChild(
					newMockVersionedElement(
						"feature2019",
						Office2019,
					),
				).
				withChild(
					newMockVersionedElement(
						"feature2021",
						Office2021,
					),
				).
				withChild(newMockElement("baseFeature"))

			version := DetectMinimumVersion(doc)
			if version != Office2021 {
				t.Errorf(
					"expected Office2021 (highest), got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"detects Microsoft365 for exclusive features",
		func(t *testing.T) {
			// Create a document with Microsoft 365 exclusive element
			doc := newMockVersionedElement(
				"m365Feature",
				Microsoft365,
			)

			version := DetectMinimumVersion(doc)
			if version != Microsoft365 {
				t.Errorf(
					"expected Microsoft365, got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"handles nil document",
		func(t *testing.T) {
			version := DetectMinimumVersion(nil)
			if version != Office2016 {
				t.Errorf(
					"expected Office2016 for nil doc, got %s",
					version.String(),
				)
			}
		},
	)
}

func TestVersionDetector_WalkElement(
	t *testing.T,
) {
	t.Run(
		"walks through complex tree structure",
		func(t *testing.T) {
			// Build a complex tree:
			// doc
			//   ├── body (Office 2016)
			//   │   ├── para1 (Office 2016)
			//   │   └── para2 (Office 2019)
			//   └── footer (Office 2021)
			doc := newMockElement("document").
				withChild(
					newMockElement("body").
						withChild(newMockElement("para1")).
						withChild(
							newMockVersionedElement(
								"para2",
								Office2019,
							),
						),
				).
				withChild(
					newMockVersionedElement(
						"footer",
						Office2021,
					),
				)

			detector := &versionDetector{
				minVersion: Office2016,
			}
			detector.walkElement(doc)

			if detector.minVersion != Office2021 {
				t.Errorf(
					"expected Office2021, got %s",
					detector.minVersion.String(),
				)
			}
		},
	)

	t.Run(
		"handles elements without version info",
		func(t *testing.T) {
			doc := newMockElement("document").
				withChild(newMockElement("body")).
				withChild(newMockElement("para"))

			detector := &versionDetector{
				minVersion: Office2016,
			}
			detector.walkElement(doc)

			if detector.minVersion != Office2016 {
				t.Errorf(
					"expected Office2016, got %s",
					detector.minVersion.String(),
				)
			}
		},
	)
}

func TestDetectMinimumVersionForPackage(
	t *testing.T,
) {
	t.Run(
		"detects version across multiple parts",
		func(t *testing.T) {
			// Create a mock package with multiple parts
			part1 := &mockPart{
				root: newMockVersionedElement(
					"doc1",
					Office2019,
				),
			}
			part2 := &mockPart{
				root: newMockVersionedElement(
					"doc2",
					Office2021,
				),
			}

			pkg := &mockPackage{
				parts: []any{part1, part2},
			}

			version := DetectMinimumVersionForPackage(
				pkg,
			)
			if version != Office2021 {
				t.Errorf(
					"expected Office2021 (highest across parts), got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"handles empty package",
		func(t *testing.T) {
			pkg := &mockPackage{
				parts: make([]any, 0),
			}

			version := DetectMinimumVersionForPackage(
				pkg,
			)
			if version != Office2016 {
				t.Errorf(
					"expected Office2016 for empty package, got %s",
					version.String(),
				)
			}
		},
	)

	t.Run(
		"handles nil package",
		func(t *testing.T) {
			version := DetectMinimumVersionForPackage(
				nil,
			)
			if version != Office2016 {
				t.Errorf(
					"expected Office2016 for nil package, got %s",
					version.String(),
				)
			}
		},
	)
}

// Mock package for testing
type mockPackage struct {
	parts []any
}

func (p *mockPackage) Parts() []any {
	return p.parts
}

// Mock part for testing
type mockPart struct {
	root any
}

func (p *mockPart) RootElement() any {
	return p.root
}
