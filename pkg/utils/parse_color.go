package utils

import (
	"errors"
	"fmt"
	"image/color"
	"regexp"
	"strings"
)

const (
	hexColorRegexpPattern = "^(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	rgbColorRegexpPattern = "^(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5]),(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5]),(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])$"

	hexFactor = 17

	hexFormat      = "%02x%02x%02x"
	hexShortFormat = "%1x%1x%1x"
	rgbFormat      = "%d,%d,%d"
)

func ParseColor(s string) (color.Color, error) {
	if ok, _ := regexp.MatchString(hexColorRegexpPattern, s); ok {
		return parseHEX(s), nil
	}

	if ok, _ := regexp.MatchString(rgbColorRegexpPattern, s); ok {
		return parseRGB(s), nil
	}

	return nil, errors.New("invalid color")
}

func parseHEX(s string) color.Color {
	s = strings.ToLower(s)

	var r, g, b uint8

	if len(s) == 3 {
		fmt.Sscanf(s, hexShortFormat, &r, &g, &b)
		r *= hexFactor
		g *= hexFactor
		b *= hexFactor
	} else {
		fmt.Sscanf(s, hexFormat, &r, &g, &b)
	}

	return color.RGBA{r, g, b, 255}
}

func parseRGB(s string) color.Color {
	var r, g, b uint8

	fmt.Sscanf(s, rgbFormat, &r, &g, &b)

	return color.RGBA{r, g, b, 255}
}
