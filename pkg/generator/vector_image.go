package generator

import (
	"fmt"
	"image/color"

	_ "embed"

	svg "github.com/ajstarks/svgo"
	"github.com/hixraid/dummy-image-generator/pkg/data"
)

const (
	backgroundColorStyles = "fill: rgb(%d, %d, %d);"
	fontStyles            = "font-size: %dpx; fill: rgb(%d, %d, %d);"
)

//go:embed svg.css
var svgStyles string

func GenerateVectorImage(canvas *svg.SVG, info *data.ImageInfo) error {
	canvas.Start(info.Size[0], info.Size[1])
	canvas.Style("text/css", svgStyles)

	r, g, b := rgbFromColor(info.BackgroundColor)
	canvas.Rect(0, 0, info.Size[0], info.Size[1], fmt.Sprintf(backgroundColorStyles, r, g, b))

	var imageText = text(info)

	if imageText != "" {
		fontSize := textSize(float64(info.Size[0]), float64(info.Size[1]), float64(len(imageText)))

		r, g, b = rgbFromColor(info.TextColor)
		textStyles := fmt.Sprintf(
			fontStyles,
			int(fontSize),
			r, g, b,
		)
		canvas.Text(info.Size[0]/2, info.Size[1]/2, imageText, textStyles)
	}
	canvas.End()

	return nil
}

func rgbFromColor(clr color.Color) (r, g, b uint8) {
	if v, ok := clr.(color.RGBA); ok {
		return v.R, v.G, v.B
	}
	return
}
