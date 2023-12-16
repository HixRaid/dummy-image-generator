package middleware

import (
	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/utils"
)

func parseFormat(path []string) (format data.ImageFormat, ok bool) {
	if format, ok = utils.ParseFormat(path[len(path)-1]); !ok {
		format = data.PNG
	}

	return format, ok
}
