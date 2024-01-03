package middleware

import (
	"github.com/hixraid/dummy-image/internal/config"
	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/utils"
)

type imageSizeParser struct {
	minSize   data.ImageResolution
	maxSize   data.ImageResolution
	sizeClamp bool
}

func newImageSizeParser(cfg *config.SizeConfig) *imageSizeParser {
	minSize, err := utils.ParseSize(cfg.MinSize)
	if err != nil {
		minSize = [2]int{10, 10}
	}

	maxSize, err := utils.ParseSize(cfg.MaxSize)
	if err != nil {
		maxSize = [2]int{10, 10}
	}

	return &imageSizeParser{
		minSize,
		maxSize,
		cfg.SizeClamp,
	}
}

func (p *imageSizeParser) parseSize(s string) (data.ImageResolution, error) {
	size, err := utils.ParseSize(s)
	if err != nil {
		return [2]int{0, 0}, err
	}

	if !p.sizeClamp {
		return size, nil
	}

	if size[0] < p.minSize[0] {
		size[0] = p.minSize[0]
	} else if size[0] > p.maxSize[0] {
		size[0] = p.maxSize[0]
	}

	if size[1] < p.minSize[1] {
		size[1] = p.minSize[1]
	} else if size[1] > p.maxSize[1] {
		size[1] = p.maxSize[1]
	}

	return size, nil
}
