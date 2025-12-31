package validation

// DetectMinimumVersion walks through a document and detects the minimum
// Office version required based on the elements and attributes used.
// It returns the highest version requirement found among all elements.
func DetectMinimumVersion(
	doc any,
) FileFormatVersions {
	detector := &versionDetector{
		minVersion: Office2016, // Start with base version
	}

	// Walk through all elements in the document
	detector.walkElement(doc)

	return detector.minVersion
}

// versionDetector is a helper for detecting version requirements.
type versionDetector struct {
	minVersion FileFormatVersions
}

// walkElement recursively walks through an element and its children
// to find the highest version requirement.
func (d *versionDetector) walkElement(
	element any,
) {
	if element == nil {
		return
	}

	// Check if this element has version information
	if ve, ok := element.(VersionedElement); ok {
		avail := ve.Availability()
		if avail != nil {
			// Update minimum version if this element requires a higher version
			if avail.IntroducedIn > d.minVersion {
				d.minVersion = avail.IntroducedIn
			}
		}
	}

	// Walk children
	children := getChildElements(element)
	for _, child := range children {
		d.walkElement(child)
	}
}

// DetectMinimumVersionForPackage walks through an entire package
// (all parts) and detects the minimum Office version required.
func DetectMinimumVersionForPackage(
	pkg any,
) FileFormatVersions {
	detector := &versionDetector{
		minVersion: Office2016, // Start with base version
	}

	// Get all parts from the package
	parts := getPackageParts(pkg)
	for _, part := range parts {
		// Get root element from part
		root := getPartRootElement(part)
		if root != nil {
			detector.walkElement(root)
		}
	}

	return detector.minVersion
}
