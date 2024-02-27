package service

import (
	"bytes"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/chai2010/webp"
	"github.com/hixraid/dummy-image-generator/pkg/data"
	"github.com/hixraid/dummy-image-generator/pkg/generator"
)

func GenerateImage(info *data.ImageInfo) (io.Reader, error) {
	switch info.Format {
	case data.SVG:
		return writeVectorImage(info)
	default:
		return writeRasterImage(info)
	}
}

func writeRasterImage(info *data.ImageInfo) (io.Reader, error) {
	img, err := generator.GenerateRasterImage(info)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	switch info.Format {
	case data.PNG:
		return buf, png.Encode(buf, img)
	case data.JPEG:
		return buf, jpeg.Encode(buf, img, nil)
	case data.WEBP:
		return buf, webp.Encode(buf, img, nil)
	case data.GIF:
		return buf, gif.Encode(buf, img, nil)
	}

	return nil, errors.New("can't found image format")
}

func writeVectorImage(info *data.ImageInfo) (io.Reader, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	canvas := svg.New(buf)

	return buf, generator.GenerateVectorImage(canvas, info)
}
