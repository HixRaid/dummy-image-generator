package generator

import (
	"fmt"

	"github.com/hixraid/dummy-image/pkg/data"
)

const noTextKey = "no_text"

func textSize(x, y, len float64) float64 {
	return max(min(x/len*1.15, y*0.5), 5.0)
}

func text(info *data.ImageInfo) string {
	var s string

	switch info.Text {
	case noTextKey:
		s = ""
	case "":
		s = fmt.Sprintf("%d x %d", info.Size[0], info.Size[1])
	default:
		s = info.Text
	}

	return s
}
