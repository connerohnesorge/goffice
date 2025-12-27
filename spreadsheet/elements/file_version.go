package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// FileVersion represents the file version element (x:fileVersion).
// It contains application version information about the workbook.
type FileVersion struct {
	*openxml.CompositeElementBase
}

// NewFileVersion creates a new FileVersion element.
func NewFileVersion() *FileVersion {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"fileVersion",
		PrefixDefault,
	)

	return &FileVersion{
		CompositeElementBase: elem,
	}
}

// AppName returns the application name that created the workbook.
func (fv *FileVersion) AppName() string {
	attr, found := fv.GetAttribute("appName", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAppName sets the application name.
func (fv *FileVersion) SetAppName(name string) {
	if name == "" {
		fv.RemoveAttribute("appName", "")

		return
	}
	fv.SetAttribute(
		openxml.NewAttribute(
			"",
			"appName",
			"",
			name,
		),
	)
}

// LastEdited returns the version of the application that last edited the
// workbook.
func (fv *FileVersion) LastEdited() string {
	attr, found := fv.GetAttribute(
		"lastEdited",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLastEdited sets the last edited version.
func (fv *FileVersion) SetLastEdited(
	version string,
) {
	if version == "" {
		fv.RemoveAttribute("lastEdited", "")

		return
	}
	fv.SetAttribute(
		openxml.NewAttribute(
			"",
			"lastEdited",
			"",
			version,
		),
	)
}

// LowestEdited returns the lowest version of the application that edited
// the workbook.
func (fv *FileVersion) LowestEdited() string {
	attr, found := fv.GetAttribute(
		"lowestEdited",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLowestEdited sets the lowest edited version.
func (fv *FileVersion) SetLowestEdited(
	version string,
) {
	if version == "" {
		fv.RemoveAttribute("lowestEdited", "")

		return
	}
	fv.SetAttribute(
		openxml.NewAttribute(
			"",
			"lowestEdited",
			"",
			version,
		),
	)
}

// RupBuild returns the incremental release of the application.
func (fv *FileVersion) RupBuild() string {
	attr, found := fv.GetAttribute("rupBuild", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRupBuild sets the incremental release build number.
func (fv *FileVersion) SetRupBuild(build string) {
	if build == "" {
		fv.RemoveAttribute("rupBuild", "")

		return
	}
	fv.SetAttribute(
		openxml.NewAttribute(
			"",
			"rupBuild",
			"",
			build,
		),
	)
}

// CodeName returns the code name of the application.
func (fv *FileVersion) CodeName() string {
	attr, found := fv.GetAttribute("codeName", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCodeName sets the code name.
func (fv *FileVersion) SetCodeName(name string) {
	if name == "" {
		fv.RemoveAttribute("codeName", "")

		return
	}
	fv.SetAttribute(
		openxml.NewAttribute(
			"",
			"codeName",
			"",
			name,
		),
	)
}

// SetVersionInfo is a convenience method to set all version attributes at
// once.
func (fv *FileVersion) SetVersionInfo(
	appName, lastEdited, lowestEdited, rupBuild string,
) {
	fv.SetAppName(appName)
	fv.SetLastEdited(lastEdited)
	fv.SetLowestEdited(lowestEdited)
	fv.SetRupBuild(rupBuild)
}

// SetExcelVersion sets version information for Microsoft Excel.
func (fv *FileVersion) SetExcelVersion(
	majorVersion, minorVersion int,
) {
	fv.SetAppName("xl")
	fv.SetLastEdited(strconv.Itoa(majorVersion))
	fv.SetLowestEdited(strconv.Itoa(minorVersion))
}

// Clone creates a deep copy of this FileVersion element.
func (fv *FileVersion) Clone() openxml.Element {
	cloned := fv.CompositeElementBase.Clone()

	return &FileVersion{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FileVersion element.
func (fv *FileVersion) CloneNode(
	deep bool,
) openxml.Element {
	cloned := fv.CompositeElementBase.CloneNode(
		deep,
	)

	return &FileVersion{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
