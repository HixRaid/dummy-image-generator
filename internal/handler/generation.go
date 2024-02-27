package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image-generator/internal/middleware"
	"github.com/hixraid/dummy-image-generator/internal/response"
	"github.com/hixraid/dummy-image-generator/pkg/data"
	"github.com/hixraid/dummy-image-generator/pkg/service"
)

func generateImage(ctx *gin.Context) {
	imageInfo, err := middleware.GetImageInfo(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Header("Content-type", getContentType(imageInfo.Format))
	ctx.Header("Cache-Control", "public, max-age=31557600")

	buf, err := service.GenerateImage(imageInfo)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, "can't generate image")
		return
	}

	_, err = io.Copy(ctx.Writer, buf)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, "can't write response")
		return
	}
}

func getContentType(format data.ImageFormat) string {
	if format == data.SVG {
		return "image/svg+xml"
	}
	return fmt.Sprintf("image/%s", format)
}
