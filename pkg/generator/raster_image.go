package generator

import (
	"errors"
	"image"

	"github.com/fogleman/gg"
	"github.com/hixraid/dummy-image/pkg/data"
)

func GenerateRasterImage(info data.ImageInfo) (image.Image, error) {
	dc := gg.NewContext(info.Size[0], info.Size[1])

	r, g, b, a := info.BackgroundColor.RGBA()
	dc.SetRGBA255(int(r), int(g), int(b), int(a))

	dc.Clear()

	r, g, b, a = info.TextColor.RGBA()
	dc.SetRGBA255(int(r), int(g), int(b), int(a))

	var imageText = text(info)

	fontSize := textSize(float64(info.Size[0]), float64(info.Size[1]), float64(len(imageText)))
	if err := dc.LoadFontFace("fonts/arial.ttf", fontSize); err != nil {
		return nil, errors.New("can't load font")
	}

	dc.DrawStringAnchored(imageText, float64(info.Size[0])/2, float64(info.Size[1])/2, 0.5, 0.5)

	return dc.Image(), nil
}
