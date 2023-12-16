package service

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/generator"
)

func GenerateImage(w io.Writer, format data.ImageFormat, info data.ImageInfo) error {
	switch format {
	case data.PNG, data.JPEG:
		img, err := generator.GenerateRasterImage(info)
		if err != nil {
			return err
		}
		return writeRasterImage(w, format, img)
	}

	return errors.New("can't found image format")
}

func writeRasterImage(w io.Writer, format data.ImageFormat, img image.Image) error {
	switch format {
	case data.PNG:
		return png.Encode(w, img)
	case data.JPEG:
		return jpeg.Encode(w, img, nil)
	}

	return errors.New("can't found raster image format")
}
