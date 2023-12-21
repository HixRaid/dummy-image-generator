package middleware

import (
	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/utils"
)

const (
	minSizeWidth  = 10
	minSizeHeight = 10

	maxSizeWidth  = 2000
	maxSizeHeight = 2000
)

func parseSize(s string) (data.ImageResolution, error) {
	size, err := utils.ParseSize(s)
	if err != nil {
		return [2]int{0, 0}, err
	}

	if size[0] < minSizeWidth {
		size[0] = minSizeWidth
	}
	if size[1] < minSizeHeight {
		size[1] = minSizeHeight
	}

	if size[0] > maxSizeWidth {
		size[0] = maxSizeWidth
	}
	if size[1] > maxSizeHeight {
		size[1] = maxSizeHeight
	}

	return size, nil
}
