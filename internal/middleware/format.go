package middleware

import (
	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/utils"
)

func parseFormat(path []string) (format data.ImageFormat, err error) {
	if format, err = utils.ParseFormat(path[len(path)-1]); err != nil {
		return
	}

	return data.PNG, err
}
