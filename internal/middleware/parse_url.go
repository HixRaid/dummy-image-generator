package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/data"
)

const (
	imageInfoCtx = "image_info"
)

func ParseURL(ctx *gin.Context) {
	var imageInfo = &data.ImageInfo{}

	path := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")

	imageInfo.Size = parseSize(path)
	imageInfo.Text = ctx.Query("text")

	backgroundColor, textColor := parseImageColors(path)
	imageInfo.BackgroundColor = backgroundColor
	imageInfo.TextColor = textColor

	ctx.Set(imageInfoCtx, imageInfo)
}

func GetImageInfo(ctx *gin.Context) (*data.ImageInfo, error) {
	v, ok := ctx.Get(imageInfoCtx)
	if !ok {
		return nil, errors.New("not found image_info")
	}

	imageInfo, ok := v.(*data.ImageInfo)
	if !ok {
		return nil, errors.New("invalid type image_info")
	}

	return imageInfo, nil
}
