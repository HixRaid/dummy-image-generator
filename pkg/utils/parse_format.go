package utils

import (
	"slices"

	"github.com/hixraid/dummy-image-generator/pkg/data"
)

var formats = []data.ImageFormat{
	data.PNG,
	data.JPEG,
	data.WEBP,
	data.GIF,
	data.SVG,
}

func ParseFormat(s string) (data.ImageFormat, bool) {
	format := data.ImageFormat(s)
	if ok := slices.Contains(formats, format); ok {
		return format, true
	}

	return "", false
}
