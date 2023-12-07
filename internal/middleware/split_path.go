package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func SplitPath(ctx *gin.Context) {
	path := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")
	ctx.Set("path", path)
}

func GetPath(ctx *gin.Context) ([]string, error) {
	v, ok := ctx.Get("path")
	if !ok {
		return nil, errors.New("not found path")
	}

	path, ok := v.([]string)
	if !ok {
		return nil, errors.New("invalid type path")
	}

	return path, nil
}
