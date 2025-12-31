module github.com/connerohnesorge/goffice/examples/pdf-rendering/word-images

go 1.25.0

replace github.com/connerohnesorge/goffice => ../../../

replace github.com/connerohnesorge/goffice-pdf => ../../../pdf

require (
	github.com/connerohnesorge/goffice v0.0.0-00010101000000-000000000000
	github.com/connerohnesorge/goffice-pdf v0.0.0-00010101000000-000000000000
)

require (
	github.com/hhrutter/lzw v1.0.0 // indirect
	github.com/hhrutter/tiff v1.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/pdfcpu/pdfcpu v0.9.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/image v0.21.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
