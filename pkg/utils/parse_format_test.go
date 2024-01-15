package utils

import (
	"testing"

	"github.com/hixraid/dummy-image-generator/pkg/data"
)

func TestParseFormat(t *testing.T) {
	formats := map[string]struct {
		Format data.ImageFormat
		Exist  bool
	}{
		"png":  {data.PNG, true},
		"jpeg": {data.JPEG, true},
		"webp": {data.WEBP, true},
		"gif":  {data.GIF, true},
		"svg":  {data.SVG, true},
		"abc":  {"", false},
		"def":  {"", false},
	}

	for k, v := range formats {
		format, exist := ParseFormat(k)
		if format != v.Format || exist != v.Exist {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
