package utils

import (
	"image/color"
	"testing"
)

func TestParseColor(t *testing.T) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	colors := map[string]color.Color{
		"black":       black,
		"white":       white,
		"000":         black,
		"fff":         white,
		"000000":      black,
		"ffffff":      white,
		"0,0,0":       black,
		"255,255,255": white,
		"nice":        nil,
		"ggg":         nil,
		"ffff":        nil,
		"-1,-1,-1":    nil,
		"256,256,256": nil,
	}

	for k, v := range colors {
		result, _ := ParseColor(k)
		if result != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
