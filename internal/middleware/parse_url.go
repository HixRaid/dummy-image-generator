package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/response"
	"github.com/hixraid/dummy-image/pkg/data"
	"github.com/hixraid/dummy-image/pkg/utils"
)

const (
	imageFormatCtx = "image_format"
	imageInfoCtx   = "image_info"
)

func ParseURL(ctx *gin.Context) {
	var imageInfo = &data.ImageInfo{}

	pathStr := strings.Trim(ctx.Request.URL.Path, "/")
	if len(pathStr) == 0 {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "missing path")
		return
	}

	path := strings.Split(pathStr, "/")
	pathLen := len(path)

	size, err := parseSize(path[0])
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	imageInfo.Size = size

	pathIndex := 1

	imageInfo.BackgroundColor = data.Black
	imageInfo.TextColor = data.White

	if pathLen > pathIndex {
		clr, err := utils.ParseColor(path[pathIndex])
		if err == nil {
			imageInfo.BackgroundColor = clr
			pathIndex++

			if pathLen > pathIndex {
				clr, err := utils.ParseColor(path[pathIndex])
				if err == nil {
					imageInfo.TextColor = clr
					pathIndex++
				}
			}
		}
	}

	imageFormat := data.PNG

	if pathLen > pathIndex {
		format, ok := utils.ParseFormat(path[pathIndex])
		if pathLen == 4 && !ok {
			response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid image_format")
			return
		}

		if ok {
			pathIndex++
			imageFormat = format
		}
	}

	if pathLen > pathIndex {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid path")
		return
	}

	if pathLen > 4 {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "too many params")
		return
	}

	imageInfo.Text = ctx.Query("text")
	ctx.Set(imageFormatCtx, imageFormat)
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
