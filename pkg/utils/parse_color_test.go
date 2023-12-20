package utils

import (
	"image/color"
	"testing"
)

func TestParseColor(t *testing.T) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	colors := map[string]struct {
		Color color.Color
		Error error
	}{
		"black":       {black, nil},
		"white":       {white, nil},
		"000":         {black, nil},
		"fff":         {white, nil},
		"000000":      {black, nil},
		"ffffff":      {white, nil},
		"0,0,0":       {black, nil},
		"255,255,255": {white, nil},
		"nice":        {nil, ErrInvalidColor},
		"ggg":         {nil, ErrInvalidColor},
		"ffff":        {nil, ErrInvalidColor},
		"-1,-1,-1":    {nil, ErrInvalidColor},
		"256,256,256": {nil, ErrInvalidColor},
	}

	for k, v := range colors {
		result, err := ParseColor(k)
		if result != v.Color || err != v.Error {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
