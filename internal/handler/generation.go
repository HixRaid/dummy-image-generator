package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/middleware"
	"github.com/hixraid/dummy-image/internal/response"
	"github.com/hixraid/dummy-image/pkg/service"
)

func generateImage(ctx *gin.Context) {
	imageInfo, err := middleware.GetImageInfo(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	imageFormat, err := middleware.GetImageFormat(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Header("Content-type", fmt.Sprintf("image/%s", imageFormat))

	err = service.GenerateImage(ctx.Writer, imageFormat, imageInfo)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}
