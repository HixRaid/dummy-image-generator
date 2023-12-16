package generator

import (
	"errors"
	"fmt"
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

	var text string
	if info.Text != "" {
		text = info.Text
	} else {
		text = fmt.Sprintf("%d x %d", info.Size[0], info.Size[1])
	}

	fontSize := max(min(float64(info.Size[0])/float64(len(text))*1.15, float64(info.Size[1])*0.5), 5.0)
	if err := dc.LoadFontFace("fonts/arial.ttf", fontSize); err != nil {
		return nil, errors.New("can't load font")
	}

	dc.DrawStringAnchored(text, float64(info.Size[0])/2, float64(info.Size[1])/2, 0.5, 0.5)

	return dc.Image(), nil
}
