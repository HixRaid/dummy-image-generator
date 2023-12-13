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
	imageInfoCtx = "image_info"
)

func ParseURL(ctx *gin.Context) {
	var imageInfo = &data.ImageInfo{}

	path := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")

	if len(path) > 4 {
		return
	}

	imageInfo.Text = ctx.Query("text")
	imageInfo.Size = parseSize(path)

	backgroundColor, textColor := parseImageColors(path)
	imageInfo.BackgroundColor = backgroundColor
	imageInfo.TextColor = textColor

	format, err := parseFormat(path)
	if len(path) == 4 && err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	imageInfo.Format = format

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
