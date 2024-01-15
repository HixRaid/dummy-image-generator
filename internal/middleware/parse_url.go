package middleware

import (
	"errors"
	"image/color"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image-generator/internal/config"
	"github.com/hixraid/dummy-image-generator/internal/response"
	"github.com/hixraid/dummy-image-generator/pkg/data"
	"github.com/hixraid/dummy-image-generator/pkg/utils"
)

const (
	imageFormatCtx = "image_format"
	imageInfoCtx   = "image_info"
)

type ImageURLParser struct {
	defaultFormat data.ImageFormat
	*imageSizeParser
	defaultBackgroundColor color.Color
	defaultTextColor       color.Color
}

func NewImageURLParser(cfg *config.ImageConfig) *ImageURLParser {
	defaultFormat, ok := utils.ParseFormat(cfg.DefaultFormat)
	if !ok {
		defaultFormat = data.PNG
	}

	defaultBackgroundColor, err := utils.ParseColor(cfg.Color.DefaultBackgroundColor)
	if err != nil {
		defaultBackgroundColor = data.Black
	}

	defaultTextColor, err := utils.ParseColor(cfg.Color.DefaultTextColor)
	if err != nil {
		defaultTextColor = data.White
	}

	return &ImageURLParser{
		defaultFormat,
		newImageSizeParser(cfg.Size),
		defaultBackgroundColor,
		defaultTextColor,
	}
}

func (p *ImageURLParser) ParseURL(ctx *gin.Context) {
	var imageInfo = &data.ImageInfo{}

	pathStr := strings.Trim(ctx.Request.URL.Path, "/")
	if len(pathStr) == 0 {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, "missing path")
		return
	}

	path := strings.Split(pathStr, "/")
	pathLen := len(path)

	size, err := p.imageSizeParser.parseSize(path[0])
	if err != nil {
		newInvalidPathResponse(ctx)
		return
	}
	imageInfo.Size = size

	pathIndex := 1

	imageInfo.BackgroundColor = p.defaultBackgroundColor
	imageInfo.TextColor = p.defaultTextColor

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

	imageFormat := p.defaultFormat

	if pathLen > pathIndex {
		format, ok := utils.ParseFormat(path[pathIndex])
		if pathLen == 4 && !ok {
			newInvalidPathResponse(ctx)
			return
		}

		if ok {
			pathIndex++
			imageFormat = format
		}
	}

	if pathLen > pathIndex {
		newInvalidPathResponse(ctx)
		return
	}

	if pathLen > 4 {
		newInvalidPathResponse(ctx)
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

func newInvalidPathResponse(ctx *gin.Context) {
	response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid path")
}
