package service

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/chai2010/webp"
	"github.com/hixraid/dummy-image-generator/pkg/data"
	"github.com/hixraid/dummy-image-generator/pkg/generator"
)

func GenerateImage(w io.Writer, format data.ImageFormat, info *data.ImageInfo) error {
	switch format {
	case data.PNG, data.JPEG, data.WEBP, data.GIF:
		img, err := generator.GenerateRasterImage(info)
		if err != nil {
			return err
		}
		return writeRasterImage(w, format, img)
	case data.SVG:
		canvas := svg.New(w)
		return generator.GenerateVectorImage(canvas, info)
	}
	return errors.New("can't found image format")
}

func writeRasterImage(w io.Writer, format data.ImageFormat, img image.Image) error {
	switch format {
	case data.PNG:
		return png.Encode(w, img)
	case data.JPEG:
		return jpeg.Encode(w, img, nil)
	case data.WEBP:
		return webp.Encode(w, img, nil)
	case data.GIF:
		return gif.Encode(w, img, nil)
	}

	return errors.New("can't found raster image format")
}
