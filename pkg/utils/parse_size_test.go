package utils

import (
	"testing"
)

func TestParseSize(t *testing.T) {
	sizes := map[string][2]int{
		"100":       {100, 100},
		"100x100":   {100, 100},
		"hd":        {1280, 720},
		"-100":      {0, 0},
		"-100x-100": {0, 0},
		"abc":       {0, 0},
	}

	for k, v := range sizes {
		size := ParseSize(k)
		if size != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
