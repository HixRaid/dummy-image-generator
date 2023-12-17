package middleware

import (
	"image/color"

	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/utils"
)

var (
	defaultBackgroundColor = data.Black
	defaultTextColor       = data.White
)

func parseImageColors(path []string) (backgroundColor, textColor color.Color) {
	pathLen := len(path)
	if pathLen < 2 {
		return defaultBackgroundColor, defaultTextColor
	}

	backgroundColor, err := utils.ParseColor(path[1])
	if err != nil {
		return defaultBackgroundColor, defaultTextColor
	}

	if pathLen < 3 {
		return backgroundColor, defaultTextColor
	}

	textColor, err = utils.ParseColor(path[2])
	if err != nil {
		return backgroundColor, defaultTextColor
	}

	return
}
