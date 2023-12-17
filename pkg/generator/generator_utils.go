package generator

import (
	"fmt"

	"github.com/hixraid/dummy-image/pkg/data"
)

func textSize(x, y, len float64) float64 {
	return max(min(x/len*1.15, y*0.5), 5.0)
}

func text(info data.ImageInfo) string {
	if info.Text != "" {
		return info.Text
	}

	return fmt.Sprintf("%d x %d", info.Size[0], info.Size[1])
}
