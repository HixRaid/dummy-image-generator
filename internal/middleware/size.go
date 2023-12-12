package middleware

import "github.com/hixraid/dummy-image/pkg/utils"

const (
	minSizeWidth  = 10
	minSizeHeight = 10

	maxSizeWidth  = 2000
	maxSizeHeight = 2000
)

var defaultSize = [2]int{100, 100}

func parseSize(path []string) (size [2]int) {
	if len(path) < 1 {
		return defaultSize
	}

	size = utils.ParseSize(path[0])
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

	return
}
