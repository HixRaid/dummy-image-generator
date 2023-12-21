package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/response"
	"github.com/hixraid/dummy-image/pkg/data"
)

const (
	imageFormatCtx = "image_format"
	imageInfoCtx   = "image_info"
)

func ParseURL(ctx *gin.Context) {
	var imageInfo = data.ImageInfo{}

	path := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")

	if len(path) > 4 {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "too many params")
		return
	}

	imageInfo.Text = ctx.Query("text")
	size, err := parseSize(path[0])
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	imageInfo.Size = size

	backgroundColor, textColor := parseImageColors(path)
	imageInfo.BackgroundColor = backgroundColor
	imageInfo.TextColor = textColor

	format, ok := parseFormat(path)
	if len(path) == 4 && !ok {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid image_format")
		return
	}

	ctx.Set(imageFormatCtx, format)
	ctx.Set(imageInfoCtx, imageInfo)
}

func GetImageFormat(ctx *gin.Context) (data.ImageFormat, error) {
	v, ok := ctx.Get(imageFormatCtx)
	if !ok {
		return "", errors.New("not found image_format")
	}

	imageFormat, ok := v.(data.ImageFormat)
	if !ok {
		return "", errors.New("invalid type image_format")
	}

	return imageFormat, nil
}

func GetImageInfo(ctx *gin.Context) (data.ImageInfo, error) {
	v, ok := ctx.Get(imageInfoCtx)
	if !ok {
		return data.ImageInfo{}, errors.New("not found image_info")
	}

	imageInfo, ok := v.(data.ImageInfo)
	if !ok {
		return data.ImageInfo{}, errors.New("invalid type image_info")
	}

	return imageInfo, nil
}
