package core

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// RegisterImage registers an image XObject on the given page and returns
// the resource name used for drawing (e.g., "Im1").
func (d *Document) RegisterImage(
	page *Page,
	img ImageResource,
) (string, error) {
	if page == nil {
		return "", fmt.Errorf("page is nil")
	}
	if img == nil {
		return "", fmt.Errorf("image is nil")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	page.mu.Lock()
	defer page.mu.Unlock()

	imageRef, err := d.registerImageStreamLocked(img)
	if err != nil {
		return "", err
	}

	name := nextImageNameLocked(page)

	if page.resources == nil {
		page.resources = types.NewDict()
	}

	xObjectDict, ok := page.resources["XObject"].(types.Dict)
	if !ok {
		xObjectDict = types.NewDict()
	}

	xObjectDict[name] = *imageRef
	page.resources["XObject"] = xObjectDict

	return name, nil
}

func nextImageNameLocked(page *Page) string {
	if page.images == nil {
		page.images = make(map[string]bool)
	}

	index := len(page.images) + 1
	name := fmt.Sprintf("Im%d", index)
	for page.images[name] {
		index++
		name = fmt.Sprintf("Im%d", index)
	}

	page.images[name] = true

	return name
}

func (d *Document) registerImageStreamLocked(
	img ImageResource,
) (*types.IndirectRef, error) {
	if img == nil {
		return nil, fmt.Errorf("image is nil")
	}

	dict := types.NewDict()
	dict.Insert("Type", types.Name("XObject"))
	dict.Insert("Subtype", types.Name("Image"))
	dict.Insert("Width", types.Integer(img.Width()))
	dict.Insert("Height", types.Integer(img.Height()))
	dict.Insert(
		"BitsPerComponent",
		types.Integer(img.ImageBitsPerComponent()),
	)

	colorSpace := img.ImageColorSpace()
	if colorSpace != "" {
		dict.Insert("ColorSpace", types.Name(colorSpace))
	}

	filter := img.ImageFilter()
	if filter != "" {
		dict.Insert("Filter", types.Name(filter))
	}

	if decode := img.ImageDecode(); len(decode) > 0 {
		floatVals := make([]float64, 0, len(decode))
		for _, val := range decode {
			floatVals = append(floatVals, val)
		}
		dict.Insert("Decode", types.NewNumberArray(floatVals...))
	}

	if decodeParms := img.ImageDecodeParms(); len(decodeParms) > 0 {
		dict.Insert("DecodeParms", imageDecodeParmsToDict(decodeParms))
	}

	if img.ImageInterpolate() {
		dict.Insert("Interpolate", types.Boolean(true))
	}

	if smask := img.ImageSMask(); smask != nil {
		smaskRef, err := d.registerImageStreamLocked(smask)
		if err != nil {
			return nil, err
		}
		dict.Insert("SMask", *smaskRef)
	}

	streamDict := &types.StreamDict{
		Dict:    dict,
		Content: img.Data(),
	}

	if err := streamDict.Encode(); err != nil {
		return nil, fmt.Errorf("encode image stream: %w", err)
	}

	return d.ctx.XRefTable.IndRefForNewObject(*streamDict)
}

func imageDecodeParmsToDict(
	decodeParms map[string]interface{},
) types.Dict {
	dict := types.NewDict()
	for key, val := range decodeParms {
		switch typed := val.(type) {
		case int:
			dict.Insert(key, types.Integer(typed))
		case int64:
			dict.Insert(key, types.Integer(typed))
		case float64:
			dict.Insert(key, types.Float(typed))
		case bool:
			dict.Insert(key, types.Boolean(typed))
		case string:
			dict.Insert(key, types.Name(typed))
		}
	}

	return dict
}
