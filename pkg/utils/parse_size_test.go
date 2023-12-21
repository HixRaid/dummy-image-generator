package utils

import (
	"testing"

	"github.com/hixraid/dummy-image/pkg/data"
)

func TestParseSize(t *testing.T) {
	sizes := map[string]struct {
		Size  data.ImageResolution
		Error error
	}{
		"100":       {[2]int{100, 100}, nil},
		"100x100":   {[2]int{100, 100}, nil},
		"hd":        {[2]int{1280, 720}, nil},
		"-100":      {[2]int{0, 0}, ErrInvalidSize},
		"-100x-100": {[2]int{0, 0}, ErrInvalidSize},
		"abc":       {[2]int{0, 0}, ErrInvalidSize},
	}

	for k, v := range sizes {
		size, err := ParseSize(k)
		if size != v.Size || err != v.Error {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
