package handler

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/middleware"
	"github.com/hixraid/dummy-image/internal/response"
)

func generateImage(ctx *gin.Context) {
	imageInfo, err := middleware.GetImageInfo(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, imageInfo.Size[0], imageInfo.Size[1]))

	for x := 0; x < imageInfo.Size[0]; x++ {
		for y := 0; y < imageInfo.Size[1]; y++ {
			img.Set(x, y, color.Black)
		}
	}

	ctx.Stream(func(w io.Writer) bool {
		png.Encode(w, img)

		return false
	})
}
